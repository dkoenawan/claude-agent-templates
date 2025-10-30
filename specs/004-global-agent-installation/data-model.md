# Data Model: Global Agent Installation

**Feature**: 004-global-agent-installation
**Date**: 2025-10-30
**Purpose**: Define data structures and their relationships

## Overview

This document defines the data entities required for global agent installation. The system uses file-based storage with JSON for structured data and filesystem paths for content.

---

## Entity 1: EmbeddedFileSet

**Purpose**: Represents the collection of source files embedded within the Go binary

**Attributes**:
- `rootFS`: Virtual filesystem containing all embedded files
- `agentsFS`: Sub-filesystem for agent definitions (agents/)
- `specifyFS`: Sub-filesystem for spec-kit files (.specify/)
- `version`: Version string embedded at build time
- `buildTimestamp`: Unix timestamp when binary was built
- `checksums`: Map of file paths to SHA256 checksums

**Relationships**:
- Used by `InstallationProcess` when repository files are unavailable
- Provides fallback content for `SourceFileProvider`

**Validation Rules**:
- All file paths must be relative (no absolute paths)
- Checksums must be valid SHA256 hashes (64 hex characters)
- Version must follow semantic versioning (vX.Y.Z format)
- Total size must not exceed 15MB compressed

**State Transitions**:
- Immutable after binary build (embedded at compile time)
- Accessed read-only at runtime

**Storage Format**:
```go
// Embedded via //go:embed directive
type EmbeddedFileSet struct {
    Files   embed.FS
    Version string
    BuildTime time.Time
    Checksums map[string]string // path -> SHA256
}
```

---

## Entity 2: DownloadPackage

**Purpose**: Represents an update package downloaded from GitHub Releases

**Attributes**:
- `version`: Version string (e.g., "v2.1.0")
- `downloadURL`: Full URL to archive file
- `checksumURL`: Full URL to checksum file
- `archiveFormat`: Format type ("tar.gz" or "zip")
- `expectedChecksum`: SHA256 checksum from manifest or checksum file
- `localPath`: Path to downloaded archive file
- `size`: File size in bytes
- `downloadedAt`: Timestamp of download completion

**Relationships**:
- Created by `DownloadManager` based on `VersionManifest`
- Used by `InstallationProcess` for updates
- Validated against `SourceChecksum` before extraction

**Validation Rules**:
- Version must be valid semver (validated by Masterminds/semver)
- Checksum must match file content (validated after download)
- Archive must be valid tar.gz or zip format
- Size must match expected size from manifest (±1KB tolerance)

**State Transitions**:
1. **Created**: URLs resolved, ready for download
2. **Downloading**: HTTP transfer in progress
3. **Downloaded**: File on disk, checksum not yet verified
4. **Verified**: Checksum matched, ready for extraction
5. **Extracted**: Contents available as SourceFileProvider
6. **Failed**: Download error, checksum mismatch, or extraction error

**Storage Format**:
```json
{
  "version": "v2.1.0",
  "downloadURL": "https://github.com/.../spec-kit-agents-sources-v2.1.0.tar.gz",
  "checksumURL": "https://github.com/.../spec-kit-agents-sources-v2.1.0.tar.gz.sha256",
  "archiveFormat": "tar.gz",
  "expectedChecksum": "a3f5...",
  "localPath": "/tmp/spec-kit-agents-download-abc123/sources.tar.gz",
  "size": 1534821,
  "downloadedAt": "2025-10-30T14:32:15Z"
}
```

---

## Entity 3: GlobalInstallation

**Purpose**: Represents agents and templates installed in ~/.claude/ for cross-repository access

**Attributes**:
- `installationType`: Fixed value "global"
- `installationID`: UUID uniquely identifying this installation
- `claudeDir`: Path to ~/.claude/ directory
- `agentsDir`: Path to ~/.claude/agents/
- `commandsDir`: Path to ~/.claude/commands/
- `versionLockPath`: Path to version lock file
- `installedAt`: Timestamp of installation
- `sourceType`: Origin of files ("embedded", "repository", or "downloaded")

**Relationships**:
- Creates one `VersionLock` with installationType="global"
- Contains multiple `AgentFile` instances in agentsDir
- Contains multiple `CommandFile` instances in commandsDir
- May coexist with `RepositoryInstallation` instances

**Validation Rules**:
- claudeDir must be writable
- Must have at least 100MB free disk space
- Cannot overwrite existing global installation without --force flag
- Version lock must be valid JSON

**State Transitions**:
1. **Pre-Install**: Directories may not exist
2. **Installing**: Copying files, creating version lock
3. **Installed**: Version lock exists, agents accessible
4. **Updating**: Backup created, new files being copied
5. **Backed Up**: Previous version saved before update
6. **Rolled Back**: Previous version restored after update failure

**Storage Format**:
- Filesystem-based (directories and files)
- Metadata in version lock (see VersionLock entity)

**Precedence Rules**:
- Repository-local installations take precedence over global
- Claude Code checks current directory first, then ~/.claude/
- Agents with same name: repository-local shadows global

---

## Entity 4: RepositoryInstallation

**Purpose**: Represents agents and templates installed within a specific repository

**Attributes**:
- `installationType`: Fixed value "repository-local"
- `installationID`: UUID uniquely identifying this installation
- `repositoryRoot`: Path to repository root (contains .git/)
- `prefix`: Installation prefix ("." or "spec-kit-agents")
- `specifyDir`: Path to {prefix}/.specify/
- `versionLockPath`: Path to {prefix}/.version-lock.json
- `installedAt`: Timestamp of installation
- `sourceType`: Fixed value "repository" (uses local .specify/ and agents/)

**Relationships**:
- Creates one `VersionLock` with installationType="repository-local"
- Coexists with `GlobalInstallation` (repository-local takes precedence)
- Managed independently of global installation

**Validation Rules**:
- repositoryRoot must contain .git/ directory
- Source files must exist (.specify/, agents/)
- Prefix must be valid directory name
- Version lock must be valid JSON

**State Transitions**:
- Same as GlobalInstallation but scoped to repository

**Storage Format**:
- Filesystem-based within repository
- Metadata in {prefix}/.version-lock.json

---

## Entity 5: VersionLock (Extended)

**Purpose**: Tracks installation metadata, version history, and enables rollback

**Existing Attributes** (from current implementation):
- `installationID`: UUID
- `installedAt`: ISO 8601 timestamp
- `lastVerified`: ISO 8601 timestamp
- `components`: Array of ComponentVersion

**New Attributes**:
- `installationType`: "global" or "repository-local"
- `installationPath`: Absolute path to installation root
- `sourceType`: "embedded", "repository", or "downloaded"
- `backups`: Array of BackupMetadata

**ComponentVersion** (existing):
- `name`: Component name ("spec-kit-agents", "spec-kit")
- `version`: Semver string
- `installedAt`: ISO 8601 timestamp

**BackupMetadata** (new):
- `backupID`: UUID
- `createdAt`: ISO 8601 timestamp
- `previousVersion`: Semver string
- `backupPath`: Absolute path to backup directory
- `reason`: "pre-update", "manual", or "scheduled"

**Relationships**:
- One per installation (global or repository-local)
- Referenced by `InstallationProcess` for version tracking
- Used by rollback mechanism

**Validation Rules**:
- installationType must be "global" or "repository-local"
- installationPath must be absolute and exist
- sourceType must be "embedded", "repository", or "downloaded"
- All timestamps must be valid ISO 8601
- Component versions must be valid semver

**State Transitions**:
- Created on initial installation
- Updated on each modification (install, update, verify)
- History array grows with each operation (capped at 50 entries)

**Storage Format**:
```json
{
  "installationID": "550e8400-e29b-41d4-a716-446655440000",
  "installedAt": "2025-10-30T14:00:00Z",
  "lastVerified": "2025-10-30T14:00:00Z",
  "installationType": "global",
  "installationPath": "/home/user/.claude",
  "sourceType": "embedded",
  "components": [
    {
      "name": "spec-kit-agents",
      "version": "v2.1.0",
      "installedAt": "2025-10-30T14:00:00Z"
    },
    {
      "name": "spec-kit",
      "version": "v0.0.72",
      "installedAt": "2025-10-30T14:00:00Z"
    }
  ],
  "backups": [
    {
      "backupID": "650e8400-e29b-41d4-a716-446655440001",
      "createdAt": "2025-10-30T13:55:00Z",
      "previousVersion": "v2.0.1",
      "backupPath": "/home/user/.claude/.backups/backup-650e8400",
      "reason": "pre-update"
    }
  ],
  "history": [
    {
      "timestamp": "2025-10-30T14:00:00Z",
      "action": "install",
      "version": "v2.1.0",
      "success": true
    }
  ]
}
```

---

## Entity 6: SourceFileProvider

**Purpose**: Abstract interface for accessing source files from different origins

**Attributes**:
- `providerType`: "embedded", "repository", or "downloaded"
- `rootPath`: Base path for file access
- `filesystem`: Virtual or real filesystem interface

**Methods** (conceptual interface):
- `GetAgentFiles()`: Returns list of agent markdown files
- `GetCommandFiles()`: Returns list of command markdown files
- `GetSpecifyFiles()`: Returns list of .specify/ directory contents
- `ReadFile(path)`: Returns file contents as bytes
- `ValidateChecksum(path)`: Verifies file integrity

**Relationships**:
- Implemented by `EmbeddedFileSet`, repository filesystem, or `DownloadPackage`
- Used by `InstallationProcess` to copy files

**Validation Rules**:
- All file paths must be valid and accessible
- Checksums must match if validation is enabled
- Must provide complete set of required files (agents/, .specify/)

**Implementation Strategy**:
- Use Go's io/fs.FS interface for filesystem abstraction
- Repository and download sources use os.DirFS
- Embedded source uses embed.FS
- Consistent API regardless of source

---

## Entity 7: InstallationProcess

**Purpose**: Orchestrates the installation workflow with source fallback logic

**Attributes**:
- `installationType`: Target type ("global" or "repository-local")
- `sourceProvider`: SourceFileProvider instance
- `targetPaths`: InstallationPaths for destination
- `options`: InstallationOptions (force, dryRun, quiet)
- `state`: Current state ("validating", "copying", "verifying", "complete")
- `progress`: Percentage complete (0-100)

**Workflow Steps**:
1. **Source Selection**:
   - Check for repository source files (.specify/, agents/)
   - If not found, fall back to embedded files
   - If updating, use downloaded package

2. **Pre-Installation Validation**:
   - Verify disk space (>100MB free)
   - Check target directory permissions
   - Validate source files (checksums)
   - Detect conflicts with existing installations

3. **Backup Creation** (if updating):
   - Copy existing installation to backup directory
   - Update version lock with backup metadata
   - Verify backup integrity

4. **File Copying**:
   - Copy agents to ~/.claude/agents/ (with "cat-" prefix)
   - Copy commands to ~/.claude/commands/ (with "speckit." prefix)
   - Copy .specify/ directory if needed
   - Update version manifest

5. **Version Lock Creation/Update**:
   - Generate or update version lock file
   - Record installation metadata
   - Add history entry

6. **Post-Installation Verification**:
   - Verify all files copied successfully
   - Validate checksums
   - Test that at least one agent is invokable

7. **Cleanup**:
   - Remove temporary files
   - Log installation success

**Rollback Procedure** (on failure):
1. Stop installation process
2. Remove partially copied files
3. Restore from backup (if updating)
4. Update version lock with rollback history
5. Return error to user

**Relationships**:
- Uses `SourceFileProvider` for source files
- Creates or updates `VersionLock`
- Creates `GlobalInstallation` or `RepositoryInstallation`

---

## Data Flow Diagrams

### Fresh Global Installation Flow

```
User runs: spec-kit-agents install --global
                    ↓
        Check for repository source
                    ↓
            Not found → Use EmbeddedFileSet
                    ↓
        Validate disk space & permissions
                    ↓
        Create ~/.claude/ directories
                    ↓
        Copy files via SourceFileProvider
        (agents → ~/.claude/agents/cat-*.md)
        (commands → ~/.claude/commands/speckit.*.md)
                    ↓
        Create VersionLock (type: global)
                    ↓
        Verify installation
                    ↓
        Report success to user
```

### Update Flow

```
User runs: spec-kit-agents update
                    ↓
        Load existing VersionLock
                    ↓
        Check for newer version (GitHub API)
                    ↓
        Download DownloadPackage
                    ↓
        Validate checksum
                    ↓
        Create backup (BackupMetadata)
                    ↓
        Extract downloaded package
                    ↓
        Install via InstallationProcess
        (source: DownloadPackage)
                    ↓
        Update VersionLock
        (add backup entry, update components)
                    ↓
        Verify new installation
                    ↓
        Clean up old backup (keep last 3)
                    ↓
        Report success to user
```

### Repository-Local with Fallback Flow

```
User runs: spec-kit-agents install (in repo)
                    ↓
        Check for .specify/ and agents/
                    ↓
            Found → Use repository source
            Not found → Fall back to EmbeddedFileSet
                    ↓
        Determine prefix (. or spec-kit-agents)
                    ↓
        Install to {prefix}/.specify/
                    ↓
        Setup Claude integration (global)
        (agents → ~/.claude/agents/)
        (commands → ~/.claude/commands/)
                    ↓
        Create VersionLock (type: repository-local)
                    ↓
        Report installation complete
```

---

## Conflict Resolution

### Scenario 1: Global + Repository-Local Coexistence

**Rule**: Repository-local takes precedence within that repository

**Implementation**:
- Both installations can exist simultaneously
- VersionLock tracks each installation independently
- Claude Code naturally checks current directory first
- No manual conflict resolution needed

**Example**:
```
Repository A:
  - Has repository-local installation
  - Uses local agents (custom modifications)

Repository B:
  - No local installation
  - Uses global agents from ~/.claude/

~/.claude/:
  - Global installation
  - Serves all repositories without local installation
```

### Scenario 2: Agent Name Collision

**Rule**: Repository-local agent shadows global agent with same name

**Implementation**:
- Claude Code checks .claude/ in current directory before ~/.claude/
- No warning needed (expected behavior)
- Version lock tracks source of each agent

**Example**:
```
~/.claude/agents/cat-solution-architect-python.md (global)
./repo/.claude/agents/cat-solution-architect-python.md (local)

When in ./repo/, local version is used
When in other directories, global version is used
```

### Scenario 3: Update During Active Use

**Rule**: Backup before update, rollback on failure

**Implementation**:
- Backup created before any file modifications
- Atomic operations where possible (rename, not copy)
- If update fails, restore from backup automatically
- User notified of rollback

---

## Size and Performance Estimates

### Embedded File Set Size
- ~50 agents × 8KB avg = 400KB
- ~10 commands × 5KB avg = 50KB
- .specify/ directory = ~500KB
- Total uncompressed: ~950KB
- Compressed (gzip): ~300-400KB
- **Actual binary impact: ~1.5MB** (includes Go runtime overhead)

### Download Package Size
- Source files (tar.gz): ~300-400KB
- Checksum file: ~100 bytes
- **Total download: ~400KB** (< 1 second on typical connection)

### Installation Time
- Copy 60 files: ~100ms
- Verify checksums: ~50ms
- Update version lock: ~10ms
- **Total installation: ~200ms** (well under 2-minute goal)

### Update Time
- Download package: ~1 second
- Create backup: ~200ms
- Extract + install: ~300ms
- Verify: ~50ms
- **Total update: ~2 seconds** (excluding download time)

---

## Schema Versioning

**Current Version**: 1.0 (matches existing version lock format)

**Proposed Version**: 2.0 (adds global installation fields)

**Migration Strategy**:
- Version 1.0 locks continue to work (repository-local)
- Version 2.0 adds optional fields (backward compatible)
- Missing fields default to version 1.0 behavior
- No forced migration required

**Version 2.0 Changes**:
- Add `installationType` field (default: "repository-local")
- Add `installationPath` field (computed if missing)
- Add `sourceType` field (default: "repository")
- Add `backups` array (default: empty)

**Backward Compatibility**:
```go
// Reading version 1.0 lock
func loadVersionLock(path string) (*VersionLock, error) {
    lock := &VersionLock{}
    // ... read JSON ...

    // Set defaults for v1.0 locks
    if lock.InstallationType == "" {
        lock.InstallationType = "repository-local"
    }
    if lock.SourceType == "" {
        lock.SourceType = "repository"
    }
    if lock.Backups == nil {
        lock.Backups = []BackupMetadata{}
    }

    return lock, nil
}
```

---

## Data Model Summary

| Entity | Purpose | Storage | Size | Mutable |
|--------|---------|---------|------|---------|
| EmbeddedFileSet | Bundled source files | Binary (embed.FS) | ~1.5MB | No |
| DownloadPackage | Update archives | Temp filesystem | ~400KB | No |
| GlobalInstallation | Cross-repo agents | ~/.claude/ | ~1MB | Yes |
| RepositoryInstallation | Per-repo agents | {repo}/.claude/ | ~1MB | Yes |
| VersionLock | Installation metadata | JSON file | ~2-5KB | Yes |
| SourceFileProvider | Abstract file access | Interface | N/A | No |
| InstallationProcess | Workflow state | Memory | N/A | N/A |

**Total Storage Impact**:
- Binary: +1.5MB (embedded files)
- Global install: ~1MB in ~/.claude/
- Repository install: ~1MB per repository
- Version lock: ~5KB per installation
- Backups: ~1MB per backup (kept for 3 most recent)

---

**Data Model Status**: ✅ COMPLETE
**Entities Defined**: 7
**Ready for Contract Generation**: YES

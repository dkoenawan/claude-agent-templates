# Research: Spec-Kit Lockstep Installation

**Date**: 2025-10-22
**Feature**: 003-spec-kit-lockstep-install
**Purpose**: Resolve technical unknowns from plan.md Technical Context

## Research Questions & Findings

### Q1: How is spec-kit distributed?

**Decision**: Spec-kit is distributed as a Git repository with versioned releases

**Rationale**:
- Current project structure shows `.specify/` directory containing spec-kit templates and scripts
- GitHub's spec-driven development toolkit is open source at https://github.com/github/spec-kit
- The project is currently using spec-kit scripts (`.specify/scripts/bash/`) suggesting git-based installation
- No package.json or requirements.txt found, indicating direct git usage rather than package manager

**Alternatives Considered**:
1. **npm package** - Would require Node.js dependency, adds complexity
2. **curl install script** - Less control over versioning, harder to pin specific versions
3. **Git submodule** - Considered but more complex for end users to manage
4. **Direct git clone** (CHOSEN) - Simple, version-controllable, widely understood

**Implementation Approach**:
- Use `git clone --depth 1 --branch vX.Y.Z` to fetch specific version
- Clone to `~/.claude-agent-templates/.specify/` user directory
- Alternatively, include spec-kit version directly in this repository
- Recommended: **Vendor spec-kit into repository** - copy spec-kit files into `.specify/` and version control them directly

**Trade-offs**:
- Vendoring pros: No external dependency, guaranteed availability, simple installation
- Vendoring cons: Need to manually update when spec-kit releases
- Git clone pros: Always get official version, easier to track upstream
- Git clone cons: Network dependency, more complex error handling

**Final Decision**: **Vendor spec-kit files into repository**
- Eliminates network dependency during installation
- Ensures version stability
- Simplifies installation to just cloning claude-agent-templates
- Already using this approach (`.specify/` directory exists in repo)

---

### Q2: What testing framework should be used for bash scripts?

**Decision**: Use bats-core (Bash Automated Testing System)

**Rationale**:
- Industry standard for bash script testing
- TAP-compliant output integrates with CI/CD
- Active maintenance and community support
- Simple syntax familiar to developers
- Already supported in GitHub Actions

**Alternatives Considered**:
1. **shunit2** - Older framework, less active development
2. **Manual testing** - Not repeatable, no CI/CD integration
3. **pytest with subprocess** - Overcomplicates bash testing
4. **bats-core** (CHOSEN) - Modern, well-maintained, TAP-compliant

**Implementation Details**:
```bash
# Install bats-core as dev dependency
git submodule add https://github.com/bats-core/bats-core.git tests/bats
git submodule add https://github.com/bats-core/bats-support.git tests/test_helper/bats-support
git submodule add https://github.com/bats-core/bats-assert.git tests/test_helper/bats-assert

# Test structure
tests/
├── integration/
│   ├── fresh_install.bats
│   ├── upgrade.bats
│   └── version_conflict.bats
└── unit/
    └── version_checker.bats
```

**Test Coverage Goals**:
- Unit tests for version comparison logic
- Integration tests for full install/upgrade flows
- Edge case tests for conflicts and failures
- Cross-platform tests (Linux, macOS, Windows/Git Bash)

---

### Q3: Version manifest format and schema

**Decision**: Use JSON format with semantic versioning constraints

**Rationale**:
- JSON is universally parsable (bash, python, jq)
- Human-readable and git-diffable
- Standard for dependency manifests (package.json, composer.json)
- Supports semantic version constraints

**Schema Design**:
```json
{
  "$schema": "https://json-schema.org/draft-07/schema#",
  "version": "1.0",
  "name": "claude-agent-templates",
  "dependencies": {
    "spec-kit": {
      "version": "0.0.72",
      "source": "vendored",
      "install_path": ".specify",
      "integrity": "sha256-[hash]",
      "compatibility": {
        "min_version": "0.0.70",
        "max_version": "0.1.0",
        "breaking_versions": []
      }
    }
  },
  "update_policy": "manual",
  "last_updated": "2025-10-22"
}
```

**Alternatives Considered**:
1. **YAML** - Less bash-friendly parsing
2. **TOML** - Requires external parser
3. **Shell script** - Not data format, harder to validate
4. **JSON** (CHOSEN) - Universal, parsable, validatable

**Location**: `.specify/version-manifest.json` (part of repository)

---

### Q4: Version lock format (user installation state)

**Decision**: JSON format tracking installed versions and installation metadata

**Rationale**:
- Mirrors version manifest structure for consistency
- Stores installation history for debugging
- Enables rollback capability
- Machine-readable for automation

**Schema Design**:
```json
{
  "version": "1.0",
  "installation_id": "[uuid]",
  "installed_at": "2025-10-22T12:00:00Z",
  "last_verified": "2025-10-22T12:00:00Z",
  "components": {
    "claude-agent-templates": {
      "version": "1.0.0",
      "installed_from": "git",
      "commit": "abc123",
      "install_path": "~/.claude-agent-templates"
    },
    "spec-kit": {
      "version": "0.0.72",
      "installed_from": "vendored",
      "install_path": "~/.claude-agent-templates/.specify"
    }
  },
  "history": [
    {
      "timestamp": "2025-10-22T12:00:00Z",
      "action": "install",
      "component": "claude-agent-templates",
      "version": "1.0.0"
    }
  ]
}
```

**Location**: `~/.claude-agent-templates/.version-lock.json` (user directory)

---

### Q5: Cross-platform compatibility strategy

**Decision**: Support Linux, macOS, Windows (Git Bash/WSL) with POSIX-compliant bash

**Rationale**:
- Claude Code runs on all three platforms
- Git Bash provides bash environment on Windows
- POSIX compliance ensures maximum compatibility
- Existing scripts already use bash

**Implementation Guidelines**:
- Use `#!/usr/bin/env bash` shebang (portable)
- Avoid GNU-specific extensions (use POSIX alternatives)
- Test path handling for Windows (Git Bash converts paths)
- Use `mktemp` for temp files (cross-platform)
- Avoid `realpath` (not on macOS), use `cd "$(dirname "$0")" && pwd`

**Platform-Specific Considerations**:

| Aspect | Linux | macOS | Windows (Git Bash) |
|--------|-------|-------|-------------------|
| Bash version | 4.0+ | 3.2+ (default) or 5.x (homebrew) | 4.4+ (Git Bash) |
| Home directory | `$HOME` | `$HOME` | `$HOME` (Git Bash converts) |
| Path separator | `/` | `/` | `/` (Git Bash) or `\` (cmd) |
| Case sensitivity | Yes | No (default) | No |
| Symlinks | Yes | Yes | Limited (Git Bash) |

**Testing Strategy**:
- GitHub Actions matrix: ubuntu-latest, macos-latest, windows-latest
- Use `git-bash` shell on Windows
- Test file operations, path handling, version detection

---

### Q6: Installation error handling and rollback

**Decision**: Transactional installation with automatic rollback on failure

**Rationale**:
- Prevents partial installations that leave system in broken state
- User experience requirement: installation either succeeds completely or reverts
- Enables retry without manual cleanup

**Implementation Pattern**:
```bash
#!/usr/bin/env bash
set -e  # Exit on error

# Backup current state
backup_installation() {
    if [[ -d "$INSTALL_DIR" ]]; then
        mv "$INSTALL_DIR" "$INSTALL_DIR.backup.$$"
    fi
}

# Restore from backup on failure
rollback_installation() {
    local exit_code=$?
    if [[ $exit_code -ne 0 ]]; then
        echo "ERROR: Installation failed, rolling back..."
        rm -rf "$INSTALL_DIR"
        if [[ -d "$INSTALL_DIR.backup.$$" ]]; then
            mv "$INSTALL_DIR.backup.$$" "$INSTALL_DIR"
        fi
    fi
}

trap rollback_installation EXIT

# Main installation
backup_installation
install_claude_agent_templates
install_spec_kit
verify_installation

# Success - cleanup backup
rm -rf "$INSTALL_DIR.backup.$$"
```

**Error Categories**:
1. **Network errors** - Graceful degradation, offline mode support
2. **Permission errors** - Clear error messages with resolution steps
3. **Version conflicts** - Detect and report with recommended actions
4. **Partial failures** - Automatic rollback to previous state

---

### Q7: Version checking and comparison logic

**Decision**: Semantic versioning comparison using bash arithmetic

**Rationale**:
- Spec-kit uses semver (observed: 0.0.72)
- Need to compare versions for compatibility checks
- Bash-native solution avoids external dependencies

**Implementation**:
```bash
# Compare semantic versions
# Returns: 0 if v1 == v2, 1 if v1 > v2, 2 if v1 < v2
compare_versions() {
    local v1=$1
    local v2=$2

    # Split versions into components
    IFS='.' read -ra V1 <<< "$v1"
    IFS='.' read -ra V2 <<< "$v2"

    # Compare major, minor, patch
    for i in 0 1 2; do
        local num1=${V1[$i]:-0}
        local num2=${V2[$i]:-0}

        if ((num1 > num2)); then
            return 1  # v1 > v2
        elif ((num1 < num2)); then
            return 2  # v1 < v2
        fi
    done

    return 0  # v1 == v2
}

# Check if version is within range
version_in_range() {
    local version=$1
    local min_version=$2
    local max_version=$3

    compare_versions "$version" "$min_version"
    if [[ $? -eq 2 ]]; then
        return 1  # version < min
    fi

    compare_versions "$version" "$max_version"
    if [[ $? -eq 1 ]]; then
        return 1  # version > max
    fi

    return 0  # version in range
}
```

**Testing**: Comprehensive unit tests for version comparison edge cases

---

### Q8: User communication and documentation

**Decision**: Progressive disclosure in documentation, verbose CLI output during installation

**Rationale**:
- SC-006 requires 90% self-service success rate
- Users need clear feedback during installation
- Documentation must cover common issues

**Documentation Structure**:
1. **Quickstart** (Phase 1) - 5-minute getting started guide
2. **Installation Guide** - Detailed installation instructions with screenshots
3. **Troubleshooting** - Common errors and resolution steps
4. **FAQ** - Frequently asked questions about versioning

**CLI Output Strategy**:
```bash
#!/usr/bin/env bash

# Verbose output with progress indicators
echo "🔍 Checking for existing installation..."
echo "✓ No conflicts detected"
echo ""
echo "📦 Installing claude-agent-templates v1.0.0..."
echo "  ✓ Cloning repository"
echo "  ✓ Installing to ~/.claude-agent-templates"
echo ""
echo "📦 Installing spec-kit v0.0.72 (vendored)..."
echo "  ✓ Copying spec-kit files"
echo "  ✓ Verifying installation"
echo ""
echo "✅ Installation complete!"
echo ""
echo "📝 Version Information:"
echo "  claude-agent-templates: v1.0.0"
echo "  spec-kit: v0.0.72"
echo ""
echo "Run 'claude-agent-templates --version' to verify installation"
```

**Error Message Format**:
```bash
# Clear, actionable error messages
echo "❌ ERROR: Version conflict detected"
echo ""
echo "  Current spec-kit version: 0.0.68"
echo "  Required version range: 0.0.70 - 0.1.0"
echo ""
echo "Resolution steps:"
echo "  1. Backup your current installation:"
echo "     mv ~/.claude-agent-templates ~/.claude-agent-templates.backup"
echo "  2. Re-run installation script"
echo "  3. If issue persists, see:"
echo "     https://github.com/dkoenawan/claude-agent-templates/docs/troubleshooting"
```

---

## Technology Stack Summary

Based on research findings:

| Component | Technology | Justification |
|-----------|-----------|---------------|
| **Distribution** | Git repository (vendored spec-kit) | Simple, version-controlled, no external dependencies |
| **Version Manifest** | JSON (`.specify/version-manifest.json`) | Universal format, parsable, validatable |
| **Version Lock** | JSON (`~/.claude-agent-templates/.version-lock.json`) | Consistent with manifest, machine-readable |
| **Installation Script** | Bash 4.0+ (`scripts/install.sh`) | Cross-platform, no runtime dependencies |
| **Version Checker** | Bash 4.0+ (`scripts/check-version.sh`) | Native implementation, no external tools |
| **Testing Framework** | bats-core | Industry standard, CI/CD compatible |
| **Error Handling** | Transactional with rollback | User-friendly, prevents broken states |
| **Documentation** | Markdown + CLI help text | Accessible, version-controlled |

---

## Implementation Recommendations

### Phase 1 Priorities

1. **Create version manifest** - Define current spec-kit version (0.0.72 based on CHANGELOG reference)
2. **Implement version checker** - Bash script with semver comparison
3. **Create installation script** - Orchestrate installation with rollback
4. **Write quickstart guide** - Enable 90% self-service success

### Phase 2 Priorities (Implementation)

1. **Write tests first** (TDD) - bats-core test suite
2. **Implement installation logic** - Following test specifications
3. **Cross-platform testing** - GitHub Actions matrix
4. **Documentation completion** - Troubleshooting guide, FAQ

### Risk Mitigation

| Risk | Mitigation |
|------|------------|
| Spec-kit breaking changes | Min/max version constraints in manifest |
| Network failures | Vendored spec-kit eliminates network dependency |
| Platform incompatibilities | POSIX-compliant bash, cross-platform testing |
| Partial installation failures | Transactional installation with automatic rollback |
| User errors | Clear error messages, comprehensive documentation |

---

## Next Steps

1. Proceed to Phase 1: Design
2. Generate `data-model.md` with entity specifications
3. Create `contracts/` with script interfaces
4. Write `quickstart.md` for users
5. Update agent context with technology stack

---

## Q9: Installation Scope and Claude Integration

**CRITICAL UPDATE**: User clarification revealed fundamental architecture questions about installation scope and Claude integration.

### User Scenarios Analysis

#### Scenario 1A: Global Installation for Multiple Projects

**User Need**: Install claude-agent-templates globally to be available across all projects

**Current Challenge**: Claude Code looks for:
- Agents in `~/.claude/agents/`
- Commands (slash commands) in project-specific `.claude/commands/`
- Skills in `.claude/skills/`

**Decision**: **Hybrid approach - Global agents + Per-project spec-kit**

**Rationale**:
- Agents are templates that don't change per-project → Install globally to `~/.claude/agents/`
- Spec-kit slash commands are project-specific → Install in each project's `.claude/commands/`
- This matches Claude Code's architecture expectations

**Implementation**:
```bash
# Global installation (agents only)
mkdir -p ~/.claude/agents
cp -r agents/**/*.md ~/.claude/agents/

# Per-project installation (spec-kit with pinned version)
# Each project gets its own .specify/ with vendored spec-kit
git clone https://github.com/dkoenawan/claude-agent-templates.git myproject
cd myproject
# .specify/ already contains vendored spec-kit at correct version
```

**Benefits**:
- Agents available globally across all projects
- Each project controls its own spec-kit version
- No version conflicts between projects

---

#### Scenario 1B: Local Installation for Single Project

**User Need**: Install claude-agent-templates for one specific project only

**Decision**: **Project-scoped installation with local .claude directory**

**Implementation**:
```bash
# Clone into project
cd /path/to/myproject
git clone https://github.com/dkoenawan/claude-agent-templates.git .claude-agent-templates

# Symlink or copy agents to .claude/agents/
mkdir -p .claude/agents
ln -s $(pwd)/.claude-agent-templates/agents/**/*.md .claude/agents/

# .specify/ is already present with vendored spec-kit
```

**Benefits**:
- Isolated per-project
- Different projects can use different claude-agent-templates versions
- No global installation required

---

#### Scenario 2A: User Already Uses Spec-Kit

**User Need**: Install claude-agent-templates when user already has spec-kit installed

**Challenge**: Potential version conflicts between:
- User's existing spec-kit (in their own `.specify/`)
- claude-agent-templates' vendored spec-kit

**Decision**: **Coexist with user's spec-kit - use different directory**

**Implementation**:
```bash
# User's existing setup
myproject/
├── .specify/                    # User's own spec-kit installation
└── src/

# After installing claude-agent-templates
myproject/
├── .specify/                    # User's own spec-kit (unchanged)
├── .claude-agent-templates/     # NEW: Our installation
│   ├── .specify/                # Our vendored spec-kit (isolated)
│   ├── agents/
│   └── scripts/
└── src/

# Install script detects existing .specify/
./scripts/install.sh --prefix .claude-agent-templates
```

**Conflict Resolution**:
1. **Detect existing .specify/** during installation
2. **Warn user** about potential conflicts
3. **Offer options**:
   - Option A: Install to `.claude-agent-templates/` subdirectory (default)
   - Option B: Use user's existing spec-kit (risky - version compatibility)
   - Option C: Abort installation (user decides manually)

**Benefits**:
- No conflicts with user's existing spec-kit
- Both can coexist
- User chooses integration level

---

#### Scenario 2B: Running Custom Commands from Claude

**User Need**: When using Claude Code, user wants to run custom slash commands and skills

**Challenge**: Claude Code looks for:
- Commands in `.claude/commands/`
- Skills in `.claude/skills/`
- But our spec-kit commands are in `.claude-agent-templates/.specify/templates/commands/`

**Decision**: **COPY (not symlink) spec-kit commands to .claude/commands/ with namespace prefixes**

**CRITICAL ISSUE IDENTIFIED**: Symlinks would overwrite or conflict with user's own custom agents, skills, and commands!

**Revised Approach**:
```bash
# Installation uses COPY with namespaced filenames
myproject/
├── .claude/
│   ├── commands/
│   │   ├── my-custom-command.md         # User's existing commands (preserved!)
│   │   ├── speckit.specify.md           # Our commands (namespaced)
│   │   ├── speckit.plan.md
│   │   └── speckit.tasks.md
│   ├── skills/
│   │   └── [user's existing skills preserved]
│   └── agents/
│       ├── my-custom-agent.md           # User's agents (preserved!)
│       ├── cat-requirements-analyst.md   # Our agents (prefixed)
│       ├── cat-solution-architect-python.md
│       └── cat-software-engineer-python.md
└── .claude-agent-templates/
    ├── .specify/
    │   └── templates/
    │       └── commands/
    └── agents/
```

**Installation Script Update**:
```bash
#!/usr/bin/env bash
# scripts/install.sh

# Detect installation mode
if [[ -d ".specify" ]]; then
    echo "⚠️  Existing .specify/ detected"
    echo "Installing to .claude-agent-templates/ to avoid conflicts"
    INSTALL_DIR=".claude-agent-templates"
else
    echo "No existing .specify/ detected"
    echo "Installing to project root"
    INSTALL_DIR="."
fi

# Create .claude structure for Claude Code
mkdir -p .claude/{commands,skills,agents}

# COPY (not symlink) slash commands with namespace prefix
# Commands from spec-kit templates
cp "$INSTALL_DIR/.specify/templates/commands/specify.md" .claude/commands/speckit.specify.md
cp "$INSTALL_DIR/.specify/templates/commands/plan.md" .claude/commands/speckit.plan.md
cp "$INSTALL_DIR/.specify/templates/commands/tasks.md" .claude/commands/speckit.tasks.md

# COPY (not symlink) agents with "cat-" prefix (claude-agent-templates)
# This avoids conflicts with user's custom agents
for agent in "$INSTALL_DIR/agents/"**/*.md; do
    filename=$(basename "$agent")
    cp "$agent" ".claude/agents/cat-$filename"
done

echo "✅ Claude Code integration complete"
echo "   Commands available: /speckit.specify, /speckit.plan, /speckit.tasks"
echo "   Agents available in Claude Code UI (prefix: cat-)"
echo ""
echo "⚠️  NOTE: Files were copied (not symlinked) to preserve your custom agents/commands"
echo "   To update: Run install script again or manually copy updated files"
```

**Benefits**:
- ✅ No conflicts with user's existing commands/agents/skills
- ✅ User's custom files are preserved
- ✅ Clear namespace (`speckit.*` for commands, `cat-*` for agents)
- ✅ Works across all platforms (no symlink issues on Windows)

**Trade-offs**:
- ❌ Updates require re-running install script (not automatic via symlinks)
- ✅ But safer - no risk of overwriting user's work
- ✅ User can modify copied files without affecting source

---

### Updated Architecture Decision

**Final Installation Model**:

1. **Repository Structure** (what we ship):
   ```
   claude-agent-templates/
   ├── .specify/              # Vendored spec-kit (pinned version)
   │   ├── version-manifest.json
   │   ├── scripts/
   │   └── templates/
   │       └── commands/      # Slash commands
   ├── agents/                # Agent templates
   │   ├── core/
   │   ├── python/
   │   └── ...
   └── scripts/
       ├── install.sh         # Smart installer
       └── check-version.sh   # Version checker
   ```

2. **After Installation** (user's project):
   ```
   myproject/
   ├── .claude/                        # Claude Code directory
   │   ├── commands/
   │   │   ├── speckit.specify -> ...  # Symlink
   │   │   ├── speckit.plan -> ...
   │   │   └── speckit.tasks -> ...
   │   └── agents/
   │       └── [agent symlinks]
   ├── .claude-agent-templates/        # Our installation (if .specify exists)
   │   ├── .specify/                   # Our vendored spec-kit
   │   ├── agents/
   │   └── scripts/
   ├── .specify/                       # User's existing spec-kit (if any)
   └── specs/                          # Created by spec-kit commands
   ```

3. **Version Management**:
   - **Version manifest** stays in our `.specify/version-manifest.json`
   - **Version lock** goes in `.claude-agent-templates/.version-lock.json`
   - Each project can have different claude-agent-templates version
   - No global version lock needed

---

### Implications for Implementation

#### Changes to Data Model:

1. **Version Lock Location** (updated):
   - Global install: `~/.claude/agents/.claude-agent-templates-version.json`
   - Project install: `.claude-agent-templates/.version-lock.json`

2. **Installation Paths** (updated):
   - Must support custom `--prefix` flag
   - Default prefix: `.claude-agent-templates` (if .specify exists) or `.` (if not)

#### Changes to Scripts:

1. **install.sh must**:
   - Detect existing `.specify/` directory
   - Choose appropriate installation prefix
   - Create `.claude/` directory structure
   - Create symlinks for commands and agents
   - Handle both global (`~/.claude/agents/`) and local installs

2. **check-version.sh must**:
   - Look for version lock in multiple locations
   - Support `--prefix` flag
   - Handle symlinked commands

#### Changes to Contracts:

1. **Install script contract** needs:
   - `--prefix <directory>` option
   - `--global` flag for global installation
   - `--integrate-claude` flag to create .claude/ structure
   - Symlink creation logic

2. **Version manifest** needs:
   - Support for multiple installation locations

---

### Risk Mitigation

| Risk | Mitigation |
|------|------------|
| Symlinks break on Windows | Use git-bash compatible symlinks OR copy files instead |
| User's .specify conflicts | Install to separate directory, detect and warn |
| Commands not found by Claude | Document .claude/ structure requirement |
| Version drift across projects | Per-project version lock, clear documentation |

---

**REVISED DECISION**: Change from single global installation to **flexible installation model**:
- **Option 1**: Global agents (`~/.claude/agents/`) + per-project spec-kit
- **Option 2**: Project-local installation (`.claude-agent-templates/`)
- **Option 3**: Coexist with user's spec-kit (separate directory)

All options create `.claude/` structure with symlinks for Claude Code integration.

---

---

## Q10: Cross-Platform Installation Language

**CRITICAL DECISION**: Eliminate Bash/PowerShell maintenance burden with single-language solution

### Problem Statement

Original design used Bash scripts, which creates maintenance issues:
- ❌ Need separate PowerShell scripts for Windows
- ❌ Two languages to maintain (Bash + PowerShell)
- ❌ Platform-specific bugs and quirks
- ❌ Testing complexity (need Windows + Unix environments)

### User Requirements

From user clarification:
- ✅ **Prefer minimal dependencies** - Important but not critical
- ✅ **One-liner installation** - Must match GitHub spec-kit's `uv` approach
- ✅ **Zero runtime dependencies preferred** - Go compiled binaries ideal

### Decision: Go with Compiled Binaries

**Rationale**:
- Single language (Go) for all platforms
- Zero runtime dependencies (compiled to native binary)
- Fast execution
- Excellent cross-platform file operations
- Native JSON handling (no `jq` dependency)
- Easy distribution via GitHub Releases

**Alternatives Considered**:

| Option | Pros | Cons | Verdict |
|--------|------|------|---------|
| **Bash/PowerShell** | No runtime deps | 2 languages to maintain | ❌ Rejected |
| **Python + uv** | Aligns with spec-kit, already in project | Requires Python runtime | ❌ Not zero deps |
| **Node.js + npm** | Ubiquitous, can publish to npm | Requires Node runtime | ❌ Not zero deps |
| **Go binary** (CHOSEN) | Zero deps, single language, fast | ~5-10MB binary size | ✅ **SELECTED** |

### Implementation Design

#### 1. Go Project Structure

```
spec-kit-agents/
├── cmd/
│   └── spec-kit-agents/
│       └── main.go              # CLI entry point (cobra)
├── internal/
│   ├── install/
│   │   ├── install.go           # Installation orchestration
│   │   ├── detect.go            # Detect existing .specify/
│   │   ├── claude.go            # .claude/ directory integration
│   │   └── copy.go              # File copy operations
│   ├── version/
│   │   ├── manifest.go          # Version manifest handling
│   │   ├── lock.go              # Version lock management
│   │   ├── check.go             # Version compatibility checking
│   │   └── compare.go           # Semver comparison
│   └── config/
│       └── paths.go             # Cross-platform path handling
├── pkg/
│   └── models/
│       ├── manifest.go          # Version manifest struct
│       └── lock.go              # Version lock struct
├── scripts/
│   └── install.sh               # One-liner wrapper (downloads Go binary)
└── go.mod
```

#### 2. One-liner Installation

**Installation URL**: `https://install.spec-kit-agents.dev` (or raw GitHub URL)

```bash
# One-liner install
curl -fsSL https://raw.githubusercontent.com/yourusername/spec-kit-agents/main/scripts/install.sh | bash

# With explicit version
curl -fsSL https://raw.githubusercontent.com/yourusername/spec-kit-agents/main/scripts/install.sh | bash -s -- v1.0.0

# Windows (PowerShell) - uses same install.sh via Git Bash
# Or separate one-liner:
irm https://install.spec-kit-agents.dev/install.ps1 | iex
```

**What install.sh does**:
1. Detects OS and architecture (Linux/macOS/Windows, amd64/arm64)
2. Downloads pre-compiled Go binary from GitHub Releases
3. Installs to `~/.local/bin/spec-kit-agents`
4. Adds to PATH if needed
5. Runs `spec-kit-agents install` in current directory (optional)

#### 3. Go Binary Commands

```bash
# Install in current project
spec-kit-agents install [OPTIONS]

# Options:
#   --prefix <dir>      Installation prefix (default: auto-detect)
#   --global            Install agents globally to ~/.claude/agents/
#   --force             Force installation even if conflicts detected
#   --quiet             Suppress non-error output

# Check version compatibility
spec-kit-agents check [OPTIONS]

# Options:
#   --json              Output in JSON format
#   --fix               Auto-fix version conflicts

# Update installation
spec-kit-agents update

# Show version
spec-kit-agents version

# Show installation status
spec-kit-agents status
```

#### 4. Distribution via GitHub Releases

**Build matrix** (GitHub Actions):
- Linux: amd64, arm64
- macOS: amd64, arm64 (Intel + Apple Silicon)
- Windows: amd64, arm64

**Release artifacts**:
```
spec-kit-agents-v1.0.0-linux-amd64
spec-kit-agents-v1.0.0-linux-arm64
spec-kit-agents-v1.0.0-darwin-amd64
spec-kit-agents-v1.0.0-darwin-arm64
spec-kit-agents-v1.0.0-windows-amd64.exe
spec-kit-agents-v1.0.0-windows-arm64.exe
```

#### 5. Key Go Implementation Details

**Cross-platform file operations**:
```go
// internal/install/copy.go
func copyWithPrefix(src, destDir, prefix string) error {
    entries, _ := os.ReadDir(src)
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        srcPath := filepath.Join(src, entry.Name())
        destPath := filepath.Join(destDir, prefix+entry.Name())

        // Cross-platform copy
        input, _ := os.ReadFile(srcPath)
        os.WriteFile(destPath, input, 0644)
    }
    return nil
}
```

**Version manifest parsing**:
```go
// pkg/models/manifest.go
type Manifest struct {
    Version      string                 `json:"version"`
    Name         string                 `json:"name"`
    Dependencies map[string]Dependency  `json:"dependencies"`
    UpdatePolicy string                 `json:"update_policy"`
    LastUpdated  string                 `json:"last_updated"`
}

type Dependency struct {
    Version       string         `json:"version"`
    Source        string         `json:"source"`
    InstallPath   string         `json:"install_path"`
    Integrity     string         `json:"integrity"`
    Compatibility Compatibility  `json:"compatibility"`
}

// internal/version/manifest.go
func LoadManifest(path string) (*models.Manifest, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    var manifest models.Manifest
    if err := json.Unmarshal(data, &manifest); err != nil {
        return nil, err
    }

    return &manifest, nil
}
```

**Semantic version comparison**:
```go
// internal/version/compare.go
import "github.com/Masterminds/semver/v3"

func CompareVersions(v1, v2 string) (int, error) {
    ver1, err := semver.NewVersion(v1)
    if err != nil {
        return 0, err
    }

    ver2, err := semver.NewVersion(v2)
    if err != nil {
        return 0, err
    }

    return ver1.Compare(ver2), nil
}

func InRange(version, minVer, maxVer string) (bool, error) {
    constraint := fmt.Sprintf(">=%s, <=%s", minVer, maxVer)
    c, err := semver.NewConstraint(constraint)
    if err != nil {
        return false, err
    }

    v, err := semver.NewVersion(version)
    if err != nil {
        return false, err
    }

    return c.Check(v), nil
}
```

**Installation workflow**:
```go
// internal/install/install.go
func Run(opts Options) error {
    // 1. Detect existing installation
    hasSpecKit := pathExists(".specify")
    prefix := detectPrefix(hasSpecKit, opts)

    log.Info("Installation prefix: %s", prefix)

    // 2. Create installation directory
    if err := os.MkdirAll(prefix, 0755); err != nil {
        return fmt.Errorf("failed to create install dir: %w", err)
    }

    // 3. Copy spec-kit files (vendored)
    if err := copySpecKit(prefix); err != nil {
        return fmt.Errorf("failed to copy spec-kit: %w", err)
    }

    // 4. Set up .claude/ directory structure
    if err := setupClaudeDir(prefix); err != nil {
        return fmt.Errorf("failed to setup .claude/: %w", err)
    }

    // 5. Create version lock
    lock := createVersionLock(prefix)
    if err := saveVersionLock(filepath.Join(prefix, ".version-lock.json"), lock); err != nil {
        return fmt.Errorf("failed to save version lock: %w", err)
    }

    // 6. Verify installation
    if err := verifyInstallation(prefix); err != nil {
        return fmt.Errorf("installation verification failed: %w", err)
    }

    log.Success("Installation complete!")
    log.Info("Commands available: /speckit.specify, /speckit.plan, /speckit.tasks")
    log.Info("Agents available with 'cat-' prefix")

    return nil
}
```

### Benefits of Go Approach

| Benefit | Impact |
|---------|--------|
| **Single language** | No Bash/PowerShell split, easier maintenance |
| **Zero dependencies** | Users don't need Python/Node/jq/etc |
| **Native performance** | Fast file operations, instant startup |
| **Cross-platform** | Same code runs on Linux/macOS/Windows |
| **Strong typing** | Compile-time safety for JSON parsing |
| **Easy distribution** | GitHub Releases, no package registry needed |
| **Small binary** | ~5-10MB (acceptable for developer tool) |
| **Testing** | Go's excellent testing framework |

### Testing Strategy

**Unit tests** (Go testing):
```go
// internal/version/compare_test.go
func TestCompareVersions(t *testing.T) {
    tests := []struct {
        v1, v2 string
        want   int
    }{
        {"0.0.72", "0.0.70", 1},   // v1 > v2
        {"0.0.68", "0.0.72", -1},  // v1 < v2
        {"0.0.72", "0.0.72", 0},   // v1 == v2
    }

    for _, tt := range tests {
        got, err := CompareVersions(tt.v1, tt.v2)
        if err != nil {
            t.Fatalf("CompareVersions error: %v", err)
        }
        if got != tt.want {
            t.Errorf("CompareVersions(%q, %q) = %d, want %d",
                tt.v1, tt.v2, got, tt.want)
        }
    }
}
```

**Integration tests**:
```bash
# tests/integration/test_install.sh
#!/bin/bash
set -e

# Test fresh installation
cd $(mktemp -d)
spec-kit-agents install
[ -d ".claude/commands" ] || exit 1
[ -f ".claude/commands/speckit.specify.md" ] || exit 1
[ -f ".claude/agents/cat-requirements-analyst.md" ] || exit 1

# Test with existing .specify/
mkdir .specify
spec-kit-agents install
[ -d ".claude-agent-templates" ] || exit 1

echo "✅ Integration tests passed"
```

### Dependencies (Go Modules)

```go
// go.mod
module github.com/yourusername/spec-kit-agents

go 1.22

require (
    github.com/spf13/cobra v1.8.0           // CLI framework
    github.com/Masterminds/semver/v3 v3.2.1 // Semantic versioning
    github.com/google/uuid v1.6.0           // UUID generation
)
```

### Migration from Bash

**Phase 1**: Create Go implementation alongside existing Bash scripts
**Phase 2**: Test Go implementation thoroughly
**Phase 3**: Deprecate Bash scripts (keep for one release cycle)
**Phase 4**: Remove Bash scripts completely

**Backward compatibility**: Keep `install.sh` as thin wrapper that downloads Go binary

---

**Status**: ✅ All Technical Context clarifications resolved (including installation scope, Claude integration, and cross-platform language decision)

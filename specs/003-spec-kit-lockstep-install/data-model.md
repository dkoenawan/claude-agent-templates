# Data Model: Spec-Kit Lockstep Installation

**Feature**: 003-spec-kit-lockstep-install
**Date**: 2025-10-22
**Purpose**: Define data structures for version management and installation tracking

## Entity Relationship Overview

```
┌─────────────────────┐
│  Version Manifest   │ (Repository)
│  ───────────────    │
│  - spec-kit version │────┐
│  - compatibility    │    │
│  - integrity hash   │    │
└─────────────────────┘    │
                           │ references
                           │
                           ↓
┌─────────────────────┐  ┌──────────────────┐
│   Version Lock      │  │  Installation    │
│   ────────────      │  │  Config          │
│  - installed vers   │  │  ──────────      │
│  - install history  │  │  - install path  │
│  - verification     │  │  - user prefs    │
└─────────────────────┘  └──────────────────┘
         │
         │ creates
         ↓
┌─────────────────────┐
│  Installation Log   │
│  ───────────────    │
│  - timestamps       │
│  - actions          │
│  - errors           │
└─────────────────────┘
```

## Core Entities

### 1. Version Manifest

**Purpose**: Specifies pinned spec-kit version and compatibility constraints (stored in repository)

**Location**: `.specify/version-manifest.json`

**Schema**:
```json
{
  "$schema": "https://json-schema.org/draft-07/schema#",
  "type": "object",
  "required": ["version", "name", "dependencies"],
  "properties": {
    "version": {
      "type": "string",
      "pattern": "^[0-9]+\\.[0-9]+$",
      "description": "Manifest schema version"
    },
    "name": {
      "type": "string",
      "const": "claude-agent-templates"
    },
    "dependencies": {
      "type": "object",
      "required": ["spec-kit"],
      "properties": {
        "spec-kit": {
          "type": "object",
          "required": ["version", "source", "install_path"],
          "properties": {
            "version": {
              "type": "string",
              "pattern": "^[0-9]+\\.[0-9]+\\.[0-9]+$",
              "description": "Pinned spec-kit version (semver)"
            },
            "source": {
              "type": "string",
              "enum": ["vendored", "git", "npm"],
              "description": "Distribution method"
            },
            "install_path": {
              "type": "string",
              "description": "Relative path where spec-kit is installed"
            },
            "integrity": {
              "type": "string",
              "pattern": "^sha256-[a-f0-9]{64}$",
              "description": "SHA-256 hash for verification"
            },
            "compatibility": {
              "type": "object",
              "properties": {
                "min_version": {
                  "type": "string",
                  "pattern": "^[0-9]+\\.[0-9]+\\.[0-9]+$"
                },
                "max_version": {
                  "type": "string",
                  "pattern": "^[0-9]+\\.[0-9]+\\.[0-9]+$"
                },
                "breaking_versions": {
                  "type": "array",
                  "items": {
                    "type": "string",
                    "pattern": "^[0-9]+\\.[0-9]+\\.[0-9]+$"
                  }
                }
              }
            }
          }
        }
      }
    },
    "update_policy": {
      "type": "string",
      "enum": ["manual", "patch", "minor"],
      "description": "Automatic update strategy"
    },
    "last_updated": {
      "type": "string",
      "format": "date",
      "description": "ISO 8601 date of last manifest update"
    }
  }
}
```

**Example**:
```json
{
  "version": "1.0",
  "name": "claude-agent-templates",
  "dependencies": {
    "spec-kit": {
      "version": "0.0.72",
      "source": "vendored",
      "install_path": ".specify",
      "integrity": "sha256-abc123...",
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

**Validation Rules**:
- `version` must follow semver pattern
- `dependencies.spec-kit.version` must be within compatibility range
- `integrity` hash must match actual spec-kit files
- `last_updated` must be valid ISO 8601 date

**State Transitions**: Updated manually by maintainers when spec-kit version changes

---

### 2. Version Lock

**Purpose**: Records installed component versions and installation history (user-specific)

**Location**: `~/.claude-agent-templates/.version-lock.json`

**Schema**:
```json
{
  "$schema": "https://json-schema.org/draft-07/schema#",
  "type": "object",
  "required": ["version", "installation_id", "installed_at", "components"],
  "properties": {
    "version": {
      "type": "string",
      "pattern": "^[0-9]+\\.[0-9]+$"
    },
    "installation_id": {
      "type": "string",
      "format": "uuid"
    },
    "installed_at": {
      "type": "string",
      "format": "date-time"
    },
    "last_verified": {
      "type": "string",
      "format": "date-time"
    },
    "components": {
      "type": "object",
      "required": ["claude-agent-templates", "spec-kit"],
      "properties": {
        "claude-agent-templates": {
          "type": "object",
          "required": ["version", "installed_from", "install_path"],
          "properties": {
            "version": {
              "type": "string",
              "pattern": "^[0-9]+\\.[0-9]+\\.[0-9]+$"
            },
            "installed_from": {
              "type": "string",
              "enum": ["git", "archive", "manual"]
            },
            "commit": {
              "type": "string",
              "pattern": "^[a-f0-9]{7,40}$"
            },
            "install_path": {
              "type": "string"
            }
          }
        },
        "spec-kit": {
          "type": "object",
          "required": ["version", "installed_from", "install_path"],
          "properties": {
            "version": {
              "type": "string",
              "pattern": "^[0-9]+\\.[0-9]+\\.[0-9]+$"
            },
            "installed_from": {
              "type": "string",
              "enum": ["vendored", "git", "npm"]
            },
            "install_path": {
              "type": "string"
            }
          }
        }
      }
    },
    "history": {
      "type": "array",
      "items": {
        "type": "object",
        "required": ["timestamp", "action", "component"],
        "properties": {
          "timestamp": {
            "type": "string",
            "format": "date-time"
          },
          "action": {
            "type": "string",
            "enum": ["install", "upgrade", "verify", "rollback"]
          },
          "component": {
            "type": "string",
            "enum": ["claude-agent-templates", "spec-kit", "all"]
          },
          "version": {
            "type": "string",
            "pattern": "^[0-9]+\\.[0-9]+\\.[0-9]+$"
          },
          "status": {
            "type": "string",
            "enum": ["success", "failure", "partial"]
          },
          "error": {
            "type": "string"
          }
        }
      }
    }
  }
}
```

**Example**:
```json
{
  "version": "1.0",
  "installation_id": "550e8400-e29b-41d4-a716-446655440000",
  "installed_at": "2025-10-22T12:00:00Z",
  "last_verified": "2025-10-22T12:00:00Z",
  "components": {
    "claude-agent-templates": {
      "version": "1.0.0",
      "installed_from": "git",
      "commit": "abc1234",
      "install_path": "/home/user/.claude-agent-templates"
    },
    "spec-kit": {
      "version": "0.0.72",
      "installed_from": "vendored",
      "install_path": "/home/user/.claude-agent-templates/.specify"
    }
  },
  "history": [
    {
      "timestamp": "2025-10-22T12:00:00Z",
      "action": "install",
      "component": "all",
      "version": "1.0.0",
      "status": "success"
    }
  ]
}
```

**Validation Rules**:
- `installation_id` must be unique UUID v4
- `components.spec-kit.version` must match version manifest
- History timestamps must be chronological
- `last_verified` must be updated on each verification

**State Transitions**:
- **Created** → Install operation creates initial lock file
- **Updated** → Upgrade operations append to history
- **Verified** → Version checks update `last_verified`
- **Rolled back** → Restore operation records rollback event

---

### 3. Installation Configuration

**Purpose**: User preferences and installation settings (optional, user-specific)

**Location**: `~/.claude-agent-templates/config.json`

**Schema**:
```json
{
  "$schema": "https://json-schema.org/draft-07/schema#",
  "type": "object",
  "properties": {
    "install_path": {
      "type": "string",
      "description": "Custom installation directory"
    },
    "auto_update": {
      "type": "boolean",
      "description": "Enable automatic patch updates"
    },
    "verify_on_use": {
      "type": "boolean",
      "description": "Verify versions before each use"
    },
    "backup_on_upgrade": {
      "type": "boolean",
      "default": true
    },
    "log_level": {
      "type": "string",
      "enum": ["quiet", "normal", "verbose", "debug"]
    }
  }
}
```

**Default Values**:
- `install_path`: `~/.claude-agent-templates`
- `auto_update`: `false`
- `verify_on_use`: `false`
- `backup_on_upgrade`: `true`
- `log_level`: `normal`

**Validation Rules**:
- `install_path` must be absolute path
- Boolean fields must be `true` or `false`
- `log_level` must be one of enumerated values

---

### 4. Installation Log

**Purpose**: Detailed installation/upgrade activity log for debugging

**Location**: `~/.claude-agent-templates/.install-log.txt`

**Format**: Plain text with structured log entries

**Schema** (line format):
```
[TIMESTAMP] [LEVEL] [COMPONENT] MESSAGE
```

**Example**:
```
[2025-10-22T12:00:00Z] [INFO] [installer] Starting installation
[2025-10-22T12:00:01Z] [INFO] [installer] Checking for existing installation
[2025-10-22T12:00:02Z] [INFO] [spec-kit] Installing spec-kit v0.0.72 (vendored)
[2025-10-22T12:00:03Z] [INFO] [spec-kit] Verifying integrity: sha256-abc123...
[2025-10-22T12:00:04Z] [INFO] [installer] Creating version lock
[2025-10-22T12:00:05Z] [INFO] [installer] Installation complete
```

**Log Levels**:
- `DEBUG`: Detailed diagnostic information
- `INFO`: General informational messages
- `WARN`: Warning conditions (recoverable)
- `ERROR`: Error conditions (recoverable with user intervention)
- `FATAL`: Critical errors requiring installation abort

**Rotation**: Log files rotated when size exceeds 10MB (keep last 3 files)

**Validation Rules**:
- Timestamps must be ISO 8601 format
- Levels must be one of defined values
- Component names should be lowercase-hyphenated

---

## Data Relationships

### Version Manifest → Version Lock

**Relationship**: One-to-many (one manifest version, many user installations)

**Constraint**: Version lock `components.spec-kit.version` MUST match or be compatible with manifest `dependencies.spec-kit.version`

**Validation**:
```bash
# Pseudo-code validation
manifest_version=$(jq -r '.dependencies["spec-kit"].version' .specify/version-manifest.json)
lock_version=$(jq -r '.components["spec-kit"].version' ~/.claude-agent-templates/.version-lock.json)

if [[ "$manifest_version" != "$lock_version" ]]; then
    check_compatibility "$lock_version" "$manifest_version"
fi
```

---

### Version Lock → Installation History

**Relationship**: One-to-many (one lock file, many history entries)

**Constraint**: History must be append-only, chronologically ordered

**Validation**:
- New history entries must have timestamp > last entry timestamp
- History entries cannot be deleted (only file rotation)
- Each entry must include required fields: timestamp, action, component, status

---

### Version Manifest → Integrity Hash

**Relationship**: One-to-one (each manifest has one integrity hash for spec-kit)

**Purpose**: Cryptographic verification of spec-kit files

**Validation**:
```bash
# Verify integrity hash
expected_hash=$(jq -r '.dependencies["spec-kit"].integrity' .specify/version-manifest.json)
actual_hash=$(tar -cf - .specify/ | sha256sum | awk '{print "sha256-"$1}')

if [[ "$expected_hash" != "$actual_hash" ]]; then
    echo "ERROR: Integrity check failed"
    exit 1
fi
```

---

## Storage Locations

| Entity | Location | Scope | Versioned |
|--------|----------|-------|-----------|
| **Version Manifest** | `.specify/version-manifest.json` | Repository | Yes (git) |
| **Version Lock** | `~/.claude-agent-templates/.version-lock.json` | User | No |
| **Installation Config** | `~/.claude-agent-templates/config.json` | User | No |
| **Installation Log** | `~/.claude-agent-templates/.install-log.txt` | User | No |

---

## Data Access Patterns

### Pattern 1: Fresh Installation

1. Read version manifest from repository
2. Create installation directory (`~/.claude-agent-templates/`)
3. Install components according to manifest
4. Verify integrity against manifest hash
5. Create version lock with initial state
6. Initialize installation log
7. (Optional) Create default configuration

### Pattern 2: Version Check

1. Read version lock from user directory
2. Read version manifest from repository
3. Compare `components.spec-kit.version` with `dependencies.spec-kit.version`
4. Check compatibility constraints (min/max version)
5. Return compatibility status

### Pattern 3: Upgrade

1. Read current version lock
2. Backup current installation (transactional)
3. Read new version manifest
4. Install new spec-kit version
5. Verify integrity
6. Update version lock with new versions
7. Append upgrade event to history
8. Remove backup on success, restore on failure

### Pattern 4: Conflict Detection

1. Read version lock (user's installed version)
2. Read version manifest (required version)
3. Compare versions using semver logic
4. Check against `breaking_versions` list
5. Check compatibility range (`min_version` to `max_version`)
6. Generate conflict report with resolution steps

---

## Edge Cases & Error Handling

### Missing Version Lock

**Scenario**: User deletes `.version-lock.json`

**Handling**:
1. Detect missing lock file
2. Scan installation directory for installed versions
3. Prompt user: "Reinstall or recreate lock file?"
4. If recreate: generate new lock from current state
5. If reinstall: run fresh installation

### Corrupted Manifest

**Scenario**: `version-manifest.json` is invalid JSON or missing required fields

**Handling**:
1. Detect JSON parse error or schema validation failure
2. Abort installation with error message
3. Direct user to report issue (likely repository corruption)
4. Suggest re-cloning repository

### Version Lock Conflict

**Scenario**: Installed spec-kit version incompatible with manifest

**Handling**:
1. Detect version mismatch during compatibility check
2. Display clear error message with versions
3. Offer options:
   - Upgrade to required version (automatic)
   - Keep current version (requires manifest update)
   - Abort operation
4. Log conflict details to installation log

### Partial Installation Failure

**Scenario**: Installation fails midway (network error, disk full, etc.)

**Handling**:
1. Trap EXIT signal during installation
2. On non-zero exit code, trigger rollback
3. Restore backup installation
4. Clean up partial files
5. Log failure details with error message
6. Exit with informative error code

---

## Performance Considerations

### File Size Limits

| File | Max Size | Rotation Strategy |
|------|----------|-------------------|
| Version Manifest | <100 KB | N/A (version controlled) |
| Version Lock | <1 MB | N/A (bounded by history length) |
| Installation Log | 10 MB | Rotate to `.install-log.txt.1` (keep 3) |
| Installation Config | <10 KB | N/A |

### Read/Write Patterns

- **Version Manifest**: Read-heavy (every installation/check), write-rare (maintainer updates)
- **Version Lock**: Read-heavy (every check), write on install/upgrade
- **Installation Log**: Write-heavy (every operation), read on error/debug
- **Installation Config**: Read on every operation, write-rare (user changes)

---

## Security Considerations

### Integrity Verification

- SHA-256 hash in manifest ensures spec-kit files unchanged
- Hash verification mandatory before installation
- Failed verification aborts installation with error

### File Permissions

```bash
# Recommended permissions
~/.claude-agent-templates/          # 755 (rwxr-xr-x)
├── .version-lock.json             # 644 (rw-r--r--)
├── config.json                     # 644 (rw-r--r--)
└── .install-log.txt               # 644 (rw-r--r--)
```

### Data Validation

- All JSON files validated against JSON Schema before use
- Semantic version strings validated with regex
- File paths validated for directory traversal attacks
- UUIDs validated for correct format

---

## Testing Strategy

### Unit Tests

- JSON schema validation
- Version comparison logic
- Integrity hash verification
- Path handling (cross-platform)

### Integration Tests

- Fresh installation workflow
- Upgrade workflow with version changes
- Conflict detection and resolution
- Rollback on failure

### Property-based Tests

- Version comparison transitivity (if A > B and B > C, then A > C)
- History chronological ordering
- Integrity hash determinism

---

## Migration Path

For existing installations (no version lock):

1. Detect missing version lock on first run
2. Create version lock from current installation
3. Set `installed_from` to "manual"
4. Leave `commit` field empty (unknown)
5. Add migration event to history
6. Continue normal operation

---

**Status**: ✅ Data model complete and ready for contract generation

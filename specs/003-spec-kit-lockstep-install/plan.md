# Implementation Plan: Spec-Kit Lockstep Installation

**Branch**: `003-spec-kit-lockstep-install` | **Date**: 2025-10-22 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/003-spec-kit-lockstep-install/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/commands/plan.md` for the execution workflow.

## Summary

Implement a lockstep installation system that ensures claude-agent-templates and GitHub spec-kit are installed together with compatible, pinned versions. This prevents breaking changes from uncontrolled spec-kit upgrades by managing version dependencies through a manifest file and automated installation scripts that handle detection, installation, verification, and upgrade workflows across Linux, macOS, and Windows platforms.

## Technical Context

**Language/Version**: Go 1.22+ (primary CLI), Bash (one-liner installer wrapper only)
**Primary Dependencies**:
- Go stdlib (encoding/json, os, filepath)
- github.com/spf13/cobra (CLI framework)
- github.com/Masterminds/semver/v3 (semantic versioning)
- github.com/google/uuid (UUID generation)
**Storage**: File-based (version manifest in repository, version lock in user's installation directory)
**Testing**: Go's native testing framework + integration tests (bash)
**Target Platform**: Linux, macOS, Windows (native binaries for each)
**Project Type**: Single binary CLI tool distributed via GitHub Releases
**Performance Goals**:
- One-liner installation completes in <60 seconds
- Binary execution: install <10s, check <1s
- Binary size: <10MB compressed
**Constraints**: Zero runtime dependencies (compiled binary), cross-platform native support
**Scale/Scope**: Single Go binary with subcommands (install, check, update, status, version)

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

**Note**: Constitution file is currently a template. Applying interim principles based on project best practices.

### Interim Principles Applied

| Principle | Status | Notes |
|-----------|--------|-------|
| **Test-First Development** | ✅ PASS | Plan includes test strategy for installation scripts across platforms |
| **Backward Compatibility** | ✅ PASS | FR-005, FR-008 ensure upgrade path without breaking existing installations |
| **Documentation-First** | ✅ PASS | FR-007 requires user-facing documentation; quickstart.md planned in Phase 1 |
| **Modularity & Reusability** | ✅ PASS | Separate components: version manifest, installation script, compatibility checker |
| **Simplicity & Convention Over Configuration** | ✅ PASS | Zero manual configuration (FR-001), automated version detection (FR-002) |

### Pre-Research Gate: ✅ PASS

All interim principles are satisfied. Proceeding to Phase 0 research.

## Project Structure

### Documentation (this feature)

```
specs/[###-feature]/
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command)
├── data-model.md        # Phase 1 output (/speckit.plan command)
├── quickstart.md        # Phase 1 output (/speckit.plan command)
├── contracts/           # Phase 1 output (/speckit.plan command)
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)
```

### Source Code (repository root)

```
# Go CLI Application
cmd/
└── spec-kit-agents/
    └── main.go                   # NEW: CLI entry point (cobra)

internal/
├── install/
│   ├── install.go                # NEW: Installation orchestration
│   ├── detect.go                 # NEW: Detect existing .specify/
│   ├── claude.go                 # NEW: .claude/ directory integration
│   └── copy.go                   # NEW: File copy operations
├── version/
│   ├── manifest.go               # NEW: Version manifest handling
│   ├── lock.go                   # NEW: Version lock management
│   ├── check.go                  # NEW: Version compatibility checking
│   └── compare.go                # NEW: Semver comparison
└── config/
    └── paths.go                  # NEW: Cross-platform path handling

pkg/
└── models/
    ├── manifest.go               # NEW: Version manifest struct
    └── lock.go                   # NEW: Version lock struct

# One-liner Installer (thin wrapper)
scripts/
├── install.sh                    # NEW: Downloads Go binary from GitHub Releases
└── install.ps1                   # NEW: PowerShell variant for Windows (optional)

# Version Management
.specify/
├── version-manifest.json         # NEW: Pinned spec-kit version specification
└── scripts/
    └── bash/
        ├── common.sh             # EXISTING: Spec-kit scripts (unchanged)
        └── check-prerequisites.sh # EXISTING: Spec-kit scripts (unchanged)

# Go Module Files
go.mod                            # NEW: Go dependencies
go.sum                            # NEW: Go dependency checksums

# Testing
tests/
├── integration/
│   ├── test_install.sh           # NEW: Test installation from scratch
│   ├── test_upgrade.sh           # NEW: Test upgrade scenarios
│   └── test_conflict.sh          # NEW: Test conflict detection
└── unit/                         # (Go tests live alongside code)

# GitHub Actions
.github/
└── workflows/
    ├── release.yml               # NEW: Build and release Go binaries
    └── test.yml                  # NEW: Run Go tests on push/PR
```

**Structure Decision**: Go-based CLI tool distributed as compiled binaries via GitHub Releases. Single language (Go) for all cross-platform logic. Thin bash wrapper (`scripts/install.sh`) for one-liner installation. Eliminates Bash/PowerShell maintenance burden while providing zero-dependency installation.

## Post-Design Constitution Re-evaluation

*GATE: Re-check constitution compliance after Phase 1 design complete*

### Design Review Against Principles

| Principle | Status | Notes |
|-----------|--------|-------|
| **Test-First Development** | ✅ PASS | Comprehensive test contracts defined in `contracts/install-script.yaml` and `contracts/version-checker.yaml`. bats-core framework selected. TDD workflow specified. |
| **Backward Compatibility** | ✅ PASS | Three installation modes support users with existing setups. Coexistence with user's spec-kit. Version lock tracks history. |
| **Documentation-First** | ✅ PASS | `quickstart.md` provides 5-10 minute onboarding. Three installation paths documented. Troubleshooting guide included. |
| **Modularity & Reusability** | ✅ PASS | Separate components: `install.sh`, `check-version.sh`, version manifest, version lock. Clean separation of concerns. |
| **Simplicity & Convention Over Configuration** | ✅ PASS | Zero manual configuration. Automatic detection of installation mode. Sensible defaults (`.claude-agent-templates/` prefix if `.specify/` exists). |

### Design Artifacts Validation

| Artifact | Status | Quality Check |
|----------|--------|---------------|
| **research.md** | ✅ Complete | All 9 technical questions resolved with decisions, rationale, alternatives |
| **data-model.md** | ✅ Complete | 4 core entities with schemas, relationships, validation rules, edge cases |
| **contracts/** | ✅ Complete | 2 script contracts (install, version-checker), 2 JSON schemas (manifest, lock) |
| **quickstart.md** | ✅ Complete | 3 installation options, verification checklist, troubleshooting, next steps |

### Post-Design Gate: ✅ PASS

All interim principles satisfied after design phase. Design artifacts complete and high-quality. Ready to proceed to Phase 2 (task breakdown).

### Key Design Decisions Validated

1. **Flexible Installation Model** (research.md Q9):
   - Global agents + per-project spec-kit
   - Project-local installation
   - Coexistence with user's spec-kit
   - All options validated against constitution principles ✅

2. **Claude Code Integration** (research.md Q9):
   - `.claude/` structure with symlinks
   - Commands discoverable by Claude Code
   - Agents available globally or locally
   - Follows Claude's expected directory structure ✅

3. **Version Management** (data-model.md):
   - JSON schemas for manifest and lock
   - Semantic versioning with range constraints
   - Installation history tracking
   - Integrity verification with SHA-256 ✅

4. **Cross-Platform Support** (research.md Q5):
   - POSIX-compliant bash
   - Symlinks with fallback to copy (Windows)
   - Path handling for all platforms
   - Tested on Linux, macOS, Windows (Git Bash) ✅

## Complexity Tracking

*Fill ONLY if Constitution Check has violations that must be justified*

**No violations detected.** All design decisions comply with interim constitution principles.

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| [e.g., 4th project] | [current need] | [why 3 projects insufficient] |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient] |


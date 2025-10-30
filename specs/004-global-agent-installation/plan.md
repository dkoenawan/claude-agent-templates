# Implementation Plan: Global Agent Installation

**Branch**: `004-global-agent-installation` | **Date**: 2025-10-30 | **Spec**: [spec.md](spec.md)
**Input**: Feature specification from `/specs/004-global-agent-installation/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/commands/plan.md` for the execution workflow.

## Summary

This feature provides **two complementary commands** for spec-kit-agents distribution:

1. **Global Installation** (`spec-kit-agents install --global`): Installs agents, commands, and skills to `~/.claude/` for cross-repository access. Uses embedded files from the binary, no repository clone needed.

2. **Repository Setup** (`spec-kit-agents setup`): Initializes spec-kit configuration in a specific repository by copying `.specify/`, workflow templates to `.github/workflows/`, and creating `specs/` directory structure.

**Key Architectural Insight**: Agents/commands/skills can be global (Claude discovers from `~/.claude/`), but GitHub workflows MUST be per-repository (GitHub Actions requirement). This two-command design separates global tooling from repository-specific automation.

**Technical Approach**: Embed all source files (agents, templates, .specify/, workflow templates) within the Go binary using the embed package. Implement hybrid mode that prefers repository files but falls back to embedded files. Support updates via downloadable packages from GitHub releases. Maintain backward compatibility with existing installations.

**Out of Scope**: Creating the actual workflow template content is a separate feature. This feature only distributes whatever workflow templates exist in `.specify/templates/workflows/`.

## Technical Context

**Language/Version**: Go 1.22+ (existing codebase), Bash 4.0+ (installation scripts)
**Primary Dependencies**:
- Go embed package (for bundling source files)
- github.com/spf13/cobra (existing CLI framework)
- Existing dependencies: github.com/Masterminds/semver/v3, github.com/google/uuid
- NEEDS CLARIFICATION: Archive format for update packages (tar.gz, zip, or custom format?)
- NEEDS CLARIFICATION: Download mechanism for updates (GitHub API, direct URLs, or both?)

**Storage**: File-based (no databases)
- Embedded files within Go binary
- Global installation in ~/.claude/agents/ and ~/.claude/commands/
- Version lock at ~/.claude/.version-lock.json or {prefix}/.version-lock.json
- Installation history and backups in same directory structure

**Testing**:
- Go testing package (existing unit tests in internal/version/, pkg/models/)
- Integration tests for installation flows
- NEEDS CLARIFICATION: E2E testing strategy for offline installation scenarios

**Target Platform**: Linux, macOS, Windows (cross-platform via Go)

**Project Type**: Single project (CLI tool with embedded resources)

**Performance Goals**:
- Installation completion: < 2 minutes (from spec SC-001)
- Agent invocation latency: < 500ms (from spec SC-002)
- Binary size: < 20MB with embedded files (from spec SC-003)
- Update download + installation: < 5 minutes for typical connection

**Constraints**:
- Binary size must remain under 20MB (spec SC-003)
- Offline-capable (spec FR-014)
- Must work without repository clone (spec FR-001)
- Backward compatible with existing repository-local installations
- No breaking changes to current CLI commands

**Scale/Scope**:
- Embed ~50 agent files (~2-10KB each = ~500KB total)
- Embed ~10 command files (~5KB each = ~50KB total)
- Embed .specify/ directory (~500KB compressed)
- Support updates for 100+ existing users without disruption
- Handle 10+ agent domains (core, python, dotnet, nodejs, java, etc.)

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

**Status**: ⚠️ CONSTITUTION NOT YET DEFINED - Using interim best practices

The project constitution file (`.specify/memory/constitution.md`) contains only placeholder content. For this feature, we will apply these interim principles:

### Interim Principles Applied

1. **Test-First Development**
   - ✅ COMPLIANT: Will write tests for embed package integration before implementation
   - ✅ COMPLIANT: Will create contract tests for global vs. local installation precedence
   - ✅ COMPLIANT: Will implement E2E tests for offline installation

2. **Backward Compatibility**
   - ✅ COMPLIANT: Existing --global flag behavior preserved (enhanced, not changed)
   - ✅ COMPLIANT: Repository-local installations continue to work unchanged
   - ✅ COMPLIANT: Version lock format extended (not replaced)
   - ✅ COMPLIANT: No breaking changes to CLI interface

3. **Simplicity & Convention Over Configuration**
   - ✅ COMPLIANT: Uses standard ~/.claude/ directories (no new locations)
   - ✅ COMPLIANT: Automatic fallback (repository → embedded) requires zero configuration
   - ✅ COMPLIANT: Single binary with embedded files (no multi-file distribution)

4. **Modularity & Reusability**
   - ✅ COMPLIANT: Embed logic isolated in new internal/embed/ package
   - ✅ COMPLIANT: Download logic isolated in new internal/download/ package
   - ✅ COMPLIANT: Existing internal/install/ extended (not replaced)

5. **Documentation-First**
   - ✅ COMPLIANT: Specification complete before implementation
   - ✅ COMPLIANT: Quickstart guide will document global installation flow
   - ✅ COMPLIANT: Migration guide for repository-local → global users

**Post-Design Re-Check**: Will validate after Phase 1 artifacts are generated

**Action Item**: Formal constitution should be defined using `/speckit.constitution` before production use

## Project Structure

### Documentation (this feature)

```
specs/004-global-agent-installation/
├── spec.md              # Feature specification (complete)
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command)
├── data-model.md        # Phase 1 output (/speckit.plan command)
├── quickstart.md        # Phase 1 output (/speckit.plan command)
├── contracts/           # Phase 1 output (/speckit.plan command)
│   ├── embed-api.yaml   # Go embed package interface
│   ├── download-api.yaml # Update download interface
│   └── install-flow.yaml # Installation flow contract
├── checklists/          # Quality validation
│   └── requirements.md  # Spec quality checklist (complete)
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)
```

### Source Code (repository root)

```
cmd/spec-kit-agents/
├── main.go              # CLI entry point (existing)
└── version.go           # Version info with build-time injection

internal/
├── embed/               # NEW: Embedded file management
│   ├── embed.go        # Go embed directives and extraction
│   ├── embed_test.go   # Tests for embedded file access
│   └── fs.go           # Virtual filesystem for embedded content
├── download/            # NEW: Update package download
│   ├── download.go     # GitHub release download logic
│   ├── download_test.go
│   ├── checksum.go     # SHA256 validation
│   └── extract.go      # Archive extraction (tar.gz/zip)
├── install/             # EXISTING: Extended for global mode
│   ├── install.go      # Main installation logic (modify)
│   ├── detect.go       # Mode detection (modify for global)
│   ├── claude.go       # Claude integration (existing)
│   ├── copy.go         # File copying utilities (existing)
│   └── paths.go        # Path management (existing)
├── config/              # EXISTING: Configuration management
│   ├── paths.go        # Path utilities (existing)
│   └── logger.go       # Logging (existing)
└── version/             # EXISTING: Version management
    ├── compare.go      # Semver comparison (existing)
    ├── compare_test.go # Tests (existing)
    └── manifest.go     # Version manifest (existing)

pkg/models/              # EXISTING: Data models
├── manifest.go          # Version manifest model (existing)
├── manifest_test.go     # Tests (existing)
├── lock.go              # Version lock model (modify for global flag)
└── lock_test.go         # Tests (existing)

# Embedded source files (via go:embed)
.specify/                # Embedded in binary
├── templates/           # Command templates
│   └── commands/        # Slash command implementations
agents/                  # Agent definitions
├── core/               # Domain-agnostic agents
├── python/             # Python-specific agents
├── dotnet/             # .NET-specific agents
├── nodejs/             # Node.js-specific agents
└── java/               # Java-specific agents

tests/
├── integration/         # Integration tests
│   ├── test_global_install.go      # NEW: Global installation test
│   ├── test_offline_install.go     # NEW: Offline installation test
│   └── test_update_flow.go         # NEW: Update flow test
└── fixtures/            # Test data
    └── mock-release/    # NEW: Mock GitHub release structure

scripts/
└── install.sh           # One-liner installer (modify for embedded binary)

.github/workflows/
└── release.yml          # Release workflow (existing, builds with embed)
```

**Structure Decision**: Single project (CLI tool)

This feature extends the existing Go CLI application with embedded resource support. The key additions are:

1. **internal/embed/**: New package for managing embedded files using Go's embed package
2. **internal/download/**: New package for downloading and validating update packages
3. **internal/install/**: Extended to support global installation mode using embedded or downloaded files
4. **pkg/models/**: Extended version lock model to track global vs. repository-local installations

The structure maintains separation of concerns:
- Embedding logic isolated in internal/embed/
- Download logic isolated in internal/download/
- Installation logic extended (not replaced) in internal/install/
- Existing CLI interface preserved in cmd/spec-kit-agents/

## Complexity Tracking

*Fill ONLY if Constitution Check has violations that must be justified*

**No violations identified** - all design decisions align with interim principles.

---

## Post-Design Constitution Re-Check

**Status**: ✅ ALL CHECKS PASS

After completing Phase 1 design (research, data model, contracts, quickstart), re-validating against interim principles:

### 1. Test-First Development
- ✅ COMPLIANT: Contracts define testable interfaces (embed-api.yaml, download-api.yaml, install-flow.yaml)
- ✅ COMPLIANT: Test scenarios specified in contracts for unit, integration, and E2E tests
- ✅ COMPLIANT: E2E testing strategy defined with fixtures and offline scenarios (research.md)

### 2. Backward Compatibility
- ✅ COMPLIANT: Version lock schema versioned (v2.0 adds optional fields, v1.0 continues to work)
- ✅ COMPLIANT: Existing --global flag enhanced (not changed)
- ✅ COMPLIANT: Repository-local installations unaffected
- ✅ COMPLIANT: CLI interface unchanged (no breaking changes)

### 3. Simplicity & Convention Over Configuration
- ✅ COMPLIANT: Single binary with embedded files (no multi-file distribution)
- ✅ COMPLIANT: Automatic fallback (repository → embedded) requires zero configuration
- ✅ COMPLIANT: Standard ~/.claude/ directories (no new configuration locations)
- ✅ COMPLIANT: Sensible defaults throughout (no required configuration files)

### 4. Modularity & Reusability
- ✅ COMPLIANT: Embed logic isolated in internal/embed/ package
- ✅ COMPLIANT: Download logic isolated in internal/download/ package
- ✅ COMPLIANT: Existing internal/install/ extended (not replaced)
- ✅ COMPLIANT: Clear interfaces defined in contracts (SourceFileProvider abstraction)

### 5. Documentation-First
- ✅ COMPLIANT: Quickstart guide created before implementation
- ✅ COMPLIANT: Contracts document all interfaces before coding
- ✅ COMPLIANT: Data model defines entities before implementation
- ✅ COMPLIANT: Research resolves technical decisions before coding

**Design Approval**: Ready for Phase 2 (Task Breakdown via /speckit.tasks)

---

## Artifacts Summary

### Phase 0 - Research (Completed)
- ✅ `research.md` - Technology decisions documented (tar.gz/zip, direct URLs, E2E testing)
- ✅ All NEEDS CLARIFICATION items resolved (3/3)

### Phase 1 - Design (Completed)
- ✅ `data-model.md` - 7 entities defined with relationships and validation rules
- ✅ `contracts/embed-api.yaml` - Embedded files interface (7 methods)
- ✅ `contracts/download-api.yaml` - Download and validation interface (4 methods)
- ✅ `contracts/install-flow.yaml` - Complete installation workflows (3 workflows, state machine)
- ✅ `quickstart.md` - User onboarding guide with troubleshooting
- ✅ Agent context updated (CLAUDE.md with technology stack)

### Phase 2 - Task Breakdown (Not Started)
- ⏳ Run `/speckit.tasks` to generate `tasks.md`

---

## Next Command

```
/speckit.tasks
```

This will generate an actionable, dependency-ordered task breakdown for implementation.

---

## Architectural Update Note (2025-10-30)

**Clarification**: After initial planning, we refined the architecture to clearly separate:

1. **Global Components** (in `~/.claude/` - discovered by Claude):
   - Agents (cat-*.md)
   - Commands (speckit.*.md)
   - Skills

2. **Repository Components** (in repo - required by GitHub):
   - `.github/workflows/` (GitHub Actions requirement)
   - `.specify/` (project-specific configuration)
   - `specs/` (feature specifications)

**Two Commands**:
- `spec-kit-agents install --global` → Global installation
- `spec-kit-agents setup` → Repository setup

**Workflow Templates**: This feature will copy/distribute workflow templates but NOT create their content. Workflow template creation is a separate future feature.

This clarification is reflected in updated spec.md, with new user story (US-4), additional requirements (FR-017 through FR-025), and explicit scope boundaries.


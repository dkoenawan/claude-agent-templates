# Feature Specification: Global Agent Installation + Repository Setup

**Feature Branch**: `004-global-agent-installation`
**Created**: 2025-10-30
**Updated**: 2025-10-30 (architecture clarification)
**Status**: Draft
**Input**: User description: "Create global installation to allow user to install agent in root and access it in any repository."

## Architecture Clarification

This feature provides **two complementary commands**:

1. **`spec-kit-agents install --global`** (one-time)
   - Installs agents, commands, and skills to `~/.claude/`
   - Accessible from all repositories
   - Embedded in binary (no repository needed)

2. **`spec-kit-agents setup`** (per-repository)
   - Copies `.specify/` directory to current repository
   - Copies workflow templates to `.github/workflows/`
   - Creates initial `specs/` directory
   - Repository-specific configuration

**Key Insight**: Agents/commands/skills can be global (Claude discovers them), but GitHub workflows MUST be per-repository (GitHub requirement).

## User Scenarios & Testing *(mandatory)*

### User Story 1 - One-Time Global Installation (Priority: P1)

As a developer, I want to install claude-agent-templates once on my system so that I can access agents and commands from any repository without needing to clone the source code or repeat installation steps.

**Why this priority**: This is the core functionality that enables all other use cases. Without global installation, users must clone the repository and install from source for every project, which is the primary pain point being addressed.

**Independent Test**: Can be fully tested by running the installer from any directory (without the repository cloned), verifying that agents appear in `~/.claude/agents/` and commands work in any repository, and delivers immediate usability without repository dependencies.

**Acceptance Scenarios**:

1. **Given** I have downloaded the spec-kit-agents binary, **When** I run `spec-kit-agents install --global` from any directory, **Then** the installation completes successfully without requiring repository source files
2. **Given** the global installation is complete, **When** I navigate to any repository and invoke a Claude agent, **Then** the agent is available and functional
3. **Given** I want to verify my installation, **When** I run `spec-kit-agents status`, **Then** I see confirmation that agents are globally installed in `~/.claude/agents/`

---

### User Story 2 - Update Global Installation (Priority: P2)

As a developer with an existing global installation, I want to update to the latest version of agents and templates so that I can benefit from new features and bug fixes without reinstalling or managing source repositories.

**Why this priority**: Updates are essential for ongoing maintenance but depend on the initial global installation being functional (P1). This ensures users can stay current without complexity.

**Independent Test**: Can be fully tested by running `spec-kit-agents update` after a global installation, verifying that new versions are downloaded and installed, agents are updated, and the version lock reflects the new versions.

**Acceptance Scenarios**:

1. **Given** I have a global installation, **When** I run `spec-kit-agents update`, **Then** the system downloads the latest version and updates my agents without requiring manual intervention
2. **Given** an update is available, **When** I check for updates with `spec-kit-agents check`, **Then** I see information about the new version and what would be updated
3. **Given** an update fails mid-process, **When** the error occurs, **Then** the system automatically rolls back to my previous working version

---

### User Story 3 - Repository-Independent Usage (Priority: P1)

As a developer working across multiple repositories, I want to use the same agents and commands in any project directory so that I have consistent tooling regardless of which codebase I'm working in.

**Why this priority**: This is the key value proposition of global installation - enabling cross-repository usage. It's tied with initial installation (P1) as the primary user benefit.

**Independent Test**: Can be fully tested by creating multiple test repositories, invoking agents from each directory, and verifying that all agents function identically without per-repository installation steps.

**Acceptance Scenarios**:

1. **Given** I have a global installation, **When** I create a new repository and run a slash command like `/speckit.specify`, **Then** the command executes successfully using globally installed templates
2. **Given** I'm working in any repository, **When** I invoke a domain-specific agent (e.g., `cat-solution-architect-python`), **Then** the agent is available with its full context and capabilities
3. **Given** I move between different project directories, **When** I run `spec-kit-agents status`, **Then** I see the same installation information showing my global installation is active

---

### User Story 4 - Repository Setup for GitHub Automation (Priority: P1)

As a developer using a specific repository, I want to set up spec-kit configuration and GitHub workflow templates so that I can use GitHub Issues-based automation and project-specific customization.

**Why this priority**: This is equally critical as global installation. While agents/commands work globally, GitHub workflows MUST be in each repository (GitHub requirement). Without this, users cannot leverage GitHub Issues automation.

**Independent Test**: Can be fully tested by running `spec-kit-agents setup` in a new repository, verifying that `.specify/`, `.github/workflows/`, and `specs/` directories are created with appropriate content, and confirming that workflows can be triggered by GitHub events.

**Acceptance Scenarios**:

1. **Given** I have a git repository, **When** I run `spec-kit-agents setup` from the repository root, **Then** the `.specify/` directory is created with templates, scripts, and memory configuration
2. **Given** I run the setup command, **When** the setup completes, **Then** `.github/workflows/` contains workflow templates ready for GitHub automation
3. **Given** I want to customize my project, **When** I check the created files, **Then** I see `.specify/memory/constitution.md` where I can define project-specific principles

---

### User Story 5 - Offline Installation Support (Priority: P3)

As a developer working in environments with limited internet connectivity, I want to install agents from a downloaded package so that I can set up my tooling even when offline.

**Why this priority**: While important for some users (security-restricted environments, air-gapped systems), most users will install with internet access. This is a nice-to-have that extends reach to specialized environments.

**Independent Test**: Can be fully tested by downloading the spec-kit-agents binary and source package, disabling network access, running the installation, and verifying all components install successfully from local files.

**Acceptance Scenarios**:

1. **Given** I have downloaded the installation package, **When** I run the installer without internet access, **Then** the installation succeeds using bundled source files
2. **Given** I'm offline, **When** I run `spec-kit-agents status`, **Then** I see my installation details without requiring network connectivity
3. **Given** I need to update offline, **When** I download a newer installation package and run the updater, **Then** the system uses the local package instead of attempting to download

---

### Edge Cases

**Global Installation**:
- What happens when a user tries to install globally but `~/.claude/` is not writable?
- What happens when the installation binary is updated but older versions of agents are already installed globally?
- How does the system manage disk space when keeping backups of previous installations?
- How does the system handle partial installations if the process is interrupted (network failure, disk full, killed process)?

**Repository Setup**:
- What happens if a user runs `spec-kit-agents setup` in a directory that's not a git repository?
- How does the system handle conflicts if `.specify/` or `.github/workflows/` already exist in the repository?
- What happens if a user runs setup multiple times in the same repository?
- How does the system handle different versions of workflow templates when updating?
- What happens if the repository already has GitHub workflows with conflicting names?

**Interaction Between Global and Repository**:
- How does the system handle conflicts if a repository has local agents with the same names as global agents?
- What happens if a user runs `spec-kit-agents install` from within a repository that has its own `.specify/` directory?

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST support installation without requiring the repository to be cloned locally
- **FR-002**: System MUST embed all required source files (agents, templates, spec-kit files) within the installation binary
- **FR-003**: System MUST install agents to `~/.claude/agents/` for global access across all repositories
- **FR-004**: System MUST install slash commands to `~/.claude/commands/` for global access
- **FR-005**: System MUST provide a `--global` flag that enables installation from embedded source files
- **FR-006**: System MUST fall back to embedded files when repository source files are not available
- **FR-007**: System MUST support updates that download new versions from GitHub releases
- **FR-008**: System MUST maintain version compatibility checking between global installations and project-specific configurations
- **FR-009**: System MUST create automatic backups before updates with rollback capability
- **FR-010**: System MUST track installation location in the version lock file (global vs. repository-local)
- **FR-011**: System MUST prevent conflicts when both global and local installations exist by documenting precedence rules
- **FR-012**: System MUST validate checksums of downloaded files during updates
- **FR-013**: Users MUST be able to check their current installation status and location using `spec-kit-agents status`
- **FR-014**: System MUST support offline installation from pre-downloaded packages
- **FR-015**: System MUST provide clear error messages when installation requirements are not met (permissions, disk space, etc.)
- **FR-016**: System MUST allow uninstallation that cleans up global files while preserving project-specific configurations

**Repository Setup Requirements**:
- **FR-017**: System MUST provide a `setup` command that initializes spec-kit configuration in the current repository
- **FR-018**: System MUST copy `.specify/` directory (templates, scripts, memory) to the repository root during setup
- **FR-019**: System MUST copy workflow templates from embedded files to `.github/workflows/` during setup
- **FR-020**: System MUST create an empty `specs/` directory structure during setup
- **FR-021**: System MUST detect if setup is run outside a git repository and provide appropriate error message
- **FR-022**: System MUST detect existing `.specify/` or `.github/workflows/` and offer options (skip, overwrite, merge)
- **FR-023**: System MUST allow repository setup to run multiple times with idempotent behavior
- **FR-024**: System MUST NOT require global installation to run repository setup (setup can work standalone)
- **FR-025**: System MUST preserve existing GitHub workflows that don't conflict with spec-kit workflows

### Key Entities

- **Global Installation**: Represents agents and templates installed in `~/.claude/` that are accessible from any repository. Includes installation metadata, version information, and installation timestamp.
- **Repository Setup**: Represents spec-kit configuration initialized in a specific repository. Includes `.specify/` directory, workflow templates in `.github/workflows/`, and `specs/` directory structure.
- **Embedded Source Files**: Agent definitions, templates, spec-kit files, and workflow templates bundled within the binary for repository-independent installation. Includes integrity checksums and version markers.
- **Version Lock**: Tracks installation details including whether installation is global or repository-local, installed component versions, and installation history for rollback support.
- **Installation Package**: Downloadable archive containing agents, templates, and spec-kit files for offline installation or updates. Includes manifest with checksums and compatibility information.
- **Workflow Templates**: GitHub Actions workflow YAML files for spec-kit automation (issue orchestration, execution phases, validation). Copied to repositories during setup but NOT created by this feature (out of scope).

## Success Criteria *(mandatory)*

### Measurable Outcomes

**Global Installation**:
- **SC-001**: Users can complete global installation in under 2 minutes without cloning the repository
- **SC-002**: Globally installed agents are accessible from any directory with less than 500ms invocation latency
- **SC-003**: Installation binary size remains under 20MB including all embedded source files
- **SC-004**: 100% of agents and commands function identically whether installed globally or repository-locally
- **SC-005**: Updates complete successfully with automatic rollback in case of failure 99% of the time
- **SC-006**: Users can verify their installation status in under 5 seconds using status commands
- **SC-007**: System correctly resolves conflicts between global and local installations without user intervention 100% of the time
- **SC-008**: Offline installations succeed at the same rate as online installations (>99%)
- **SC-009**: Installation process consumes less than 100MB of disk space including backups

**Repository Setup**:
- **SC-010**: Users can complete repository setup in under 1 minute
- **SC-011**: Setup command successfully creates `.specify/`, `.github/workflows/`, and `specs/` in 100% of valid repositories
- **SC-012**: Setup detects and handles existing files without data loss 100% of the time
- **SC-013**: Repository setup works independently of global installation (can be used standalone)

## Assumptions

- Users have write permissions to `~/.claude/` directory (standard Claude Code installation requirement)
- The installation binary is obtained through GitHub releases or built from source
- Internet connectivity is available for updates (offline installation supported as alternative)
- Go's embed package can efficiently bundle required source files without significant binary bloat
- Claude Code reads agents from `~/.claude/agents/` and commands from `~/.claude/commands/` on every invocation
- Repository-local installations will take precedence over global installations when both exist
- Users running the installer have basic command-line familiarity

## Dependencies

- Go embed package for bundling source files in the binary
- GitHub releases infrastructure for hosting installation packages and update artifacts
- Existing version lock and manifest system for tracking installations
- Claude Code's agent and command discovery mechanism (filesystem-based)
- Cross-platform file system APIs for handling paths on Linux, macOS, and Windows

## Scope

### In Scope

**Global Installation**:
- Implementing Go embed package integration to bundle source files in the binary
- Updating the `--global` flag to use embedded files instead of requiring repository
- Creating a hybrid installation mode that prefers repository files but falls back to embedded files
- Implementing download and installation of update packages from GitHub releases
- Adding installation location tracking to the version lock (global vs. repository-local)
- Documenting precedence rules when both global and local installations exist
- Creating comprehensive error messages for installation failures
- Supporting offline installation from pre-downloaded packages
- Implementing uninstall functionality for global installations

**Repository Setup**:
- Creating a `setup` command to initialize spec-kit configuration in repositories
- Copying `.specify/` directory structure to repository root
- Copying workflow templates (whatever exists) to `.github/workflows/`
- Creating initial `specs/` directory structure
- Detecting git repositories and existing files
- Handling conflicts and idempotent re-runs
- Making setup work independently of global installation

### Out of Scope

- **Creating workflow template content** (workflows will be created in a separate feature; this feature only distributes whatever templates exist)
- **Modifying existing GitHub workflows** in the claude-agent-templates repository (those are for building/testing this project)
- Creating a web-based installer or GUI (command-line only)
- Supporting installation of custom or third-party agents beyond the core template set
- Implementing package management features beyond install/update/uninstall (no dependency resolution, plugins, etc.)
- Auto-update functionality that installs without user confirmation
- Migration tools for moving from repository-local to global installations (manual reinstall acceptable)
- Supporting installation paths other than `~/.claude/` (standard location only)
- Per-project agent customization or overrides (global agents are read-only)
- Integration with package managers (Homebrew, apt, chocolatey) - direct binary distribution only


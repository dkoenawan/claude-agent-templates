# Tasks: Global Agent Installation + Repository Setup

**Input**: Design documents from `/specs/004-global-agent-installation/`
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/
**Feature Branch**: `004-global-agent-installation`

**Tests**: Not explicitly requested in spec - implementation only

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions
- Go project structure: `cmd/`, `internal/`, `pkg/` at repository root
- Tests: `tests/integration/` for integration tests
- Embedded files: `agents/`, `.specify/` (embedded via go:embed)

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure for embedding support

- [ ] T001 Review existing Go project structure and dependencies in go.mod
- [ ] T002 [P] Add build scripts for embedding files in scripts/build-with-embed.sh
- [ ] T003 [P] Create version injection script in scripts/inject-version.sh
- [ ] T004 Update .gitignore to exclude temporary build artifacts

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [ ] T005 Create internal/embed/ package structure with embed.go
- [ ] T006 [P] Implement Go embed directives for agents/, .specify/ in internal/embed/embed.go
- [ ] T007 [P] Implement EmbeddedFiles interface: GetAgents() in internal/embed/embed.go
- [ ] T008 [P] Implement EmbeddedFiles interface: GetSpecify() in internal/embed/embed.go
- [ ] T009 [P] Implement EmbeddedFiles interface: GetCommands() in internal/embed/embed.go
- [ ] T010 Implement GetVersion() and GetBuildTime() with build-time injection in internal/embed/embed.go
- [ ] T011 Create checksum generation pre-build script in scripts/generate-checksums.sh
- [ ] T012 Implement GetChecksum() and ListFiles() methods in internal/embed/embed.go
- [ ] T013 Create SourceFileProvider interface abstraction in internal/install/provider.go
- [ ] T014 Implement EmbeddedSourceProvider implementing SourceFileProvider in internal/embed/provider.go
- [ ] T015 Implement RepositorySourceProvider implementing SourceFileProvider in internal/install/repository_provider.go
- [ ] T016 Extend VersionLock model with new fields (installationType, installationPath, sourceType, backups) in pkg/models/lock.go
- [ ] T017 Implement version lock schema migration from v1.0 to v2.0 in pkg/models/lock.go
- [ ] T018 Create GlobalInstallation entity and methods in internal/install/global.go
- [ ] T019 [P] Create DownloadPackage entity in pkg/models/download.go
- [ ] T020 [P] Update cross-platform path handling for ~/.claude/ in internal/config/paths.go

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - One-Time Global Installation (Priority: P1) 🎯 MVP

**Goal**: Enable global installation of agents from embedded files without repository clone

**Independent Test**: Run `spec-kit-agents install --global` from any directory (without repository cloned), verify agents in `~/.claude/agents/` and confirm slash commands work in any repository

### Implementation for User Story 1

- [ ] T021 [P] [US1] Implement source selection logic with embedded fallback in internal/install/detect.go
- [ ] T022 [P] [US1] Implement pre-installation validation (disk space, permissions) in internal/install/validate.go
- [ ] T023 [US1] Implement directory creation for ~/.claude/ structure in internal/install/setup.go
- [ ] T024 [US1] Implement agent file copying with cat- prefix in internal/install/copy_agents.go
- [ ] T025 [US1] Implement command file copying with speckit. prefix in internal/install/copy_commands.go
- [ ] T026 [US1] Implement global version lock creation in internal/install/version_lock.go
- [ ] T027 [US1] Implement post-installation verification in internal/install/verify.go
- [ ] T028 [US1] Add `install` command with --global flag to CLI in cmd/spec-kit-agents/install.go
- [ ] T029 [US1] Implement installation success reporting in cmd/spec-kit-agents/install.go
- [ ] T030 [US1] Add `status` command to show installation details in cmd/spec-kit-agents/status.go

**Checkpoint**: At this point, User Story 1 should be fully functional - global installation from embedded files works

---

## Phase 4: User Story 3 - Repository-Independent Usage (Priority: P1)

**Goal**: Enable consistent agent access across all repositories using global installation

**Independent Test**: Create multiple test directories, invoke agents from each, verify identical functionality without per-repository setup

### Implementation for User Story 3

- [ ] T031 [US3] Implement repository root detection in internal/install/detect.go
- [ ] T032 [US3] Implement repository source file detection (.specify/, agents/) in internal/install/detect.go
- [ ] T033 [US3] Implement source priority logic (repository → embedded) in internal/install/provider.go
- [ ] T034 [US3] Add repository context handling to `status` command in cmd/spec-kit-agents/status.go
- [ ] T035 [US3] Implement installation location precedence display in cmd/spec-kit-agents/status.go
- [ ] T036 [US3] Update version lock to track source type (embedded vs repository) in pkg/models/lock.go

**Checkpoint**: Global agents work consistently across all repositories

---

## Phase 5: User Story 4 - Repository Setup for GitHub Automation (Priority: P1)

**Goal**: Initialize spec-kit configuration in repositories for GitHub workflow automation

**Independent Test**: Run `spec-kit-agents setup` in new repo, verify `.specify/`, `.github/workflows/`, `specs/` created, confirm workflows can be triggered

### Implementation for User Story 4

- [ ] T037 [P] [US4] Create setup command in CLI in cmd/spec-kit-agents/setup.go
- [ ] T038 [P] [US4] Implement git repository detection in internal/install/git.go
- [ ] T039 [US4] Implement existing file detection (.specify/, .github/workflows/) in internal/install/setup.go
- [ ] T040 [US4] Implement conflict handling (skip, overwrite, merge options) in internal/install/setup.go
- [ ] T041 [US4] Implement .specify/ directory copying from embedded files in internal/install/setup.go
- [ ] T042 [US4] Implement workflow template copying to .github/workflows/ in internal/install/setup.go
- [ ] T043 [US4] Implement specs/ directory structure creation in internal/install/setup.go
- [ ] T044 [US4] Implement repository version lock creation in internal/install/setup.go
- [ ] T045 [US4] Add idempotent re-run support for setup command in cmd/spec-kit-agents/setup.go
- [ ] T046 [US4] Implement setup success reporting with directory details in cmd/spec-kit-agents/setup.go

**Checkpoint**: Repository setup works independently, creates all necessary structures

---

## Phase 6: User Story 2 - Update Global Installation (Priority: P2)

**Goal**: Enable updates to latest version with automatic backup and rollback

**Independent Test**: Run `spec-kit-agents update` after global installation, verify new version downloaded, agents updated, version lock reflects changes, rollback works on failure

### Implementation for User Story 2

- [ ] T047 [P] [US2] Create internal/download/ package structure
- [ ] T048 [P] [US2] Implement Downloader interface in internal/download/downloader.go
- [ ] T049 [P] [US2] Implement GetLatestVersion() using GitHub API in internal/download/github.go
- [ ] T050 [P] [US2] Implement platform detection (Linux/macOS/Windows) in internal/download/platform.go
- [ ] T051 [P] [US2] Implement direct URL construction for GitHub releases in internal/download/url.go
- [ ] T052 [US2] Implement DownloadRelease() with retry logic in internal/download/downloader.go
- [ ] T053 [US2] Implement checksum download and validation in internal/download/checksum.go
- [ ] T054 [US2] Implement progress reporting callback in internal/download/progress.go
- [ ] T055 [US2] Implement archive extraction (tar.gz and zip) in internal/download/extract.go
- [ ] T056 [US2] Create DownloadedSourceProvider implementing SourceFileProvider in internal/download/provider.go
- [ ] T057 [US2] Implement backup creation logic in internal/install/backup.go
- [ ] T058 [US2] Extend version lock with BackupMetadata in pkg/models/lock.go
- [ ] T059 [US2] Implement update workflow orchestration in internal/install/update.go
- [ ] T060 [US2] Implement automatic rollback on failure in internal/install/rollback.go
- [ ] T061 [US2] Implement old backup cleanup (keep last 3) in internal/install/cleanup.go
- [ ] T062 [US2] Add `update` command to CLI in cmd/spec-kit-agents/update.go
- [ ] T063 [US2] Add `check` command for update availability in cmd/spec-kit-agents/check.go
- [ ] T064 [US2] Add `rollback` command for manual rollback in cmd/spec-kit-agents/rollback.go

**Checkpoint**: Update flow works with automatic backup and rollback

---

## Phase 7: User Story 5 - Offline Installation Support (Priority: P3)

**Goal**: Enable installation in environments with limited internet connectivity

**Independent Test**: Download binary and source package, disable network, run installation, verify all components install from local files

### Implementation for User Story 5

- [ ] T065 [P] [US5] Implement offline detection in internal/download/offline.go
- [ ] T066 [P] [US5] Implement local package detection in internal/download/local.go
- [ ] T067 [US5] Add --offline flag to install command in cmd/spec-kit-agents/install.go
- [ ] T068 [US5] Add --package flag for local package path in cmd/spec-kit-agents/install.go
- [ ] T069 [US5] Implement offline update from local package in internal/install/update.go
- [ ] T070 [US5] Update status command to work offline in cmd/spec-kit-agents/status.go

**Checkpoint**: Offline installation and updates work from local packages

---

## Phase 8: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [ ] T071 [P] Add comprehensive error messages for all failure scenarios in internal/install/errors.go
- [ ] T072 [P] Implement detailed logging throughout installation flows in internal/config/logger.go
- [ ] T073 [P] Add --dry-run flag to install, update, and setup commands
- [ ] T074 [P] Add --force flag for overwriting existing installations
- [ ] T075 [P] Add --quiet and --verbose flags for output control
- [ ] T076 Update README.md with global installation instructions
- [ ] T077 Update quickstart.md with tested examples
- [ ] T078 [P] Update GitHub release workflow to create source archives in .github/workflows/release.yml
- [ ] T079 [P] Add checksum generation to release workflow in .github/workflows/release.yml
- [ ] T080 Create integration test for global installation in tests/integration/test_global_install.go
- [ ] T081 Create integration test for offline installation in tests/integration/test_offline_install.go
- [ ] T082 Create integration test for update flow in tests/integration/test_update_flow.go
- [ ] T083 Create integration test for repository setup in tests/integration/test_repository_setup.go
- [ ] T084 Add binary size check to CI workflow (< 20MB requirement)
- [ ] T085 Add performance benchmarks for installation time
- [ ] T086 Update CLAUDE.md with global installation context
- [ ] T087 Create migration guide for repository-local → global users in docs/migration-guide.md
- [ ] T088 Add uninstall command in cmd/spec-kit-agents/uninstall.go
- [ ] T089 Run full quickstart.md validation with real installation

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Story 1 (Phase 3)**: Depends on Foundational - Core MVP functionality
- **User Story 3 (Phase 4)**: Depends on US1 completion - Extends global installation
- **User Story 4 (Phase 5)**: Depends on Foundational - Can run in parallel with US1/US3
- **User Story 2 (Phase 6)**: Depends on US1 and US3 completion - Extends with updates
- **User Story 5 (Phase 7)**: Depends on US2 completion - Extends with offline support
- **Polish (Phase 8)**: Depends on all user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2) - No dependencies on other stories ✅ MVP
- **User Story 3 (P1)**: Depends on User Story 1 - Extends global installation to work across repos
- **User Story 4 (P1)**: Can start after Foundational (Phase 2) - Independent from US1/US3 (different functionality)
- **User Story 2 (P2)**: Depends on User Stories 1 and 3 - Adds update capability to global installation
- **User Story 5 (P3)**: Depends on User Story 2 - Extends update with offline support

### Within Each User Story

- Foundational infrastructure before specific features
- Models and data structures before business logic
- Core logic before CLI commands
- Implementation before integration tests
- Story complete before moving to next priority

### Parallel Opportunities

- **Phase 1 (Setup)**: T002, T003 can run in parallel (different files)
- **Phase 2 (Foundational)**: T006-T009, T019-T020 can run in parallel (different packages)
- **Phase 3 (US1)**: T021-T022 can run in parallel (different files)
- **Phase 5 (US4)**: T037-T038 can run in parallel (different files)
- **Phase 6 (US2)**: T047-T051 can run in parallel (different files in download package)
- **Phase 7 (US5)**: T065-T066 can run in parallel (different files)
- **Phase 8 (Polish)**: T071-T075, T078-T079, T080-T083 can run in parallel (different concerns)

---

## Parallel Example: User Story 1 (Global Installation)

```bash
# Launch foundational tasks in parallel:
Task: "Implement Go embed directives for agents/, .specify/ in internal/embed/embed.go"
Task: "Implement EmbeddedFiles interface: GetAgents() in internal/embed/embed.go"
Task: "Implement EmbeddedFiles interface: GetSpecify() in internal/embed/embed.go"
Task: "Implement EmbeddedFiles interface: GetCommands() in internal/embed/embed.go"

# Then launch US1 implementation tasks in parallel:
Task: "Implement source selection logic with embedded fallback in internal/install/detect.go"
Task: "Implement pre-installation validation (disk space, permissions) in internal/install/validate.go"
```

---

## Implementation Strategy

### MVP First (User Stories 1, 3, 4 - All P1)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL - blocks all stories)
3. Complete Phase 3: User Story 1 (Global Installation) → **TEST INDEPENDENTLY**
4. Complete Phase 4: User Story 3 (Repository-Independent Usage) → **TEST INDEPENDENTLY**
5. Complete Phase 5: User Story 4 (Repository Setup) → **TEST INDEPENDENTLY**
6. **STOP and VALIDATE**: Test all P1 user stories work together
7. Deploy/demo if ready

**Rationale**: All three P1 stories are critical for the complete feature. US1 enables global installation, US3 validates cross-repository access, US4 enables GitHub automation. Together they form the complete MVP.

### Incremental Delivery

1. **Phase 1-2** → Foundation ready
2. **Phase 3 (US1)** → Test independently → Deploy (basic global install works)
3. **Phase 4 (US3)** → Test independently → Deploy (cross-repo usage validated)
4. **Phase 5 (US4)** → Test independently → Deploy (repository setup works)
5. **Phase 6 (US2)** → Test independently → Deploy (updates work)
6. **Phase 7 (US5)** → Test independently → Deploy (offline support)
7. **Phase 8** → Final polish and documentation

Each phase adds value without breaking previous functionality.

### Parallel Team Strategy

With multiple developers:

1. **All team members** complete Phase 1-2 together (foundational)
2. Once Foundational is done:
   - **Developer A**: User Story 1 (T021-T030)
   - **Developer B**: User Story 4 (T037-T046) - Can run in parallel with US1
   - **Developer C**: Prepare User Story 3 prerequisites
3. After US1 complete:
   - **Developer A**: User Story 3 (T031-T036) - Extends US1
   - **Developer B**: Continue US4 or start US2 prep
4. After US3 complete:
   - **Developer A**: User Story 2 (T047-T064)
   - **Developer B**: User Story 5 or Polish
5. Stories integrate naturally through shared interfaces

---

## Technical Notes

### Go Embed Package Usage

```go
// internal/embed/embed.go
package embed

import "embed"

//go:embed agents/* .specify/*
var embeddedFiles embed.FS

// Access via fs.FS interface
func GetAgents() (fs.FS, error) {
    return fs.Sub(embeddedFiles, "agents")
}
```

### Build-Time Version Injection

```bash
# scripts/build-with-embed.sh
VERSION=$(git describe --tags --always)
BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)

go build -ldflags="-X main.Version=$VERSION -X main.BuildTime=$BUILD_TIME" \
    -o bin/spec-kit-agents ./cmd/spec-kit-agents/
```

### Installation Paths

- **Global**: `~/.claude/agents/cat-*.md`, `~/.claude/commands/speckit.*.md`
- **Repository**: `{repo}/.specify/`, `{repo}/.github/workflows/`, `{repo}/specs/`
- **Version Locks**: `~/.claude/.version-lock.json` (global), `{repo}/.version-lock.json` (local)

### Archive Format Detection

- **Linux/macOS**: Download `spec-kit-agents-sources-v{version}.tar.gz`
- **Windows**: Download `spec-kit-agents-sources-v{version}.zip`
- Auto-detect based on `runtime.GOOS`

---

## Success Metrics (from spec.md)

- **SC-001**: Global installation completes in < 2 minutes (validate with T080)
- **SC-002**: Agent invocation latency < 500ms (validate with T085)
- **SC-003**: Binary size < 20MB with embedded files (validate with T084)
- **SC-004**: 100% of agents function identically (validate with T080-T083)
- **SC-005**: Updates succeed with 99% rollback success (validate with T082)
- **SC-010**: Repository setup completes in < 1 minute (validate with T083)
- **SC-011**: Setup creates all directories 100% of the time (validate with T083)

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- Each user story should be independently completable and testable
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
- All paths use forward slashes (Go convention, works on all platforms)
- Binary embedding increases size by ~1.5MB (well within 20MB budget)
- Tests are NOT included in this implementation (not requested in spec)

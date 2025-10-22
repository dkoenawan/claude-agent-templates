# Tasks: Spec-Kit Lockstep Installation

**Input**: Design documents from `/specs/003-spec-kit-lockstep-install/`
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/

**Tests**: Included (TDD approach per constitution)

**Organization**: Tasks grouped by user story to enable independent implementation and testing.

## Format: `[ID] [P?] [Story] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Implementation Strategy

**MVP Scope**: User Story 1 (Single Command Installation)
- Delivers core value: one-liner installation with lockstep version management
- Fully testable independently
- Foundation for US2 and US3

**Incremental Delivery**:
1. US1 (P1): Single command installation → Deploy to early adopters
2. US2 (P2): Version compatibility management → Enable maintainers
3. US3 (P3): Upgrade path management → Complete feature set

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Initialize Go project structure and tooling

- [X] T001 Initialize Go module in repository root with `go mod init github.com/yourusername/spec-kit-agents`
- [X] T002 Create directory structure per plan.md: cmd/spec-kit-agents/, internal/{install,version,config}/, pkg/models/, scripts/
- [X] T003 [P] Install Go dependencies: cobra, semver, uuid per go.mod
- [X] T004 [P] Create .gitignore for Go project (bin/, *.exe, go.sum initially)
- [X] T005 [P] Set up GitHub Actions workflow file .github/workflows/test.yml for running tests on push/PR

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core data models and utilities that ALL user stories depend on

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T006 [P] Create Version Manifest model in pkg/models/manifest.go with JSON schema validation
- [X] T007 [P] Create Version Lock model in pkg/models/lock.go with JSON schema validation
- [X] T008 [P] Implement semantic version comparison in internal/version/compare.go using semver library
- [X] T009 Create version manifest JSON file in .specify/version-manifest.json with spec-kit v0.0.72 pinned
- [X] T010 [P] Implement cross-platform path handling utilities in internal/config/paths.go
- [X] T011 [P] Create logging utility in internal/config/logger.go with levels (DEBUG, INFO, WARN, ERROR)

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Single Command Installation (Priority: P1) 🎯 MVP

**Goal**: User can install spec-kit-agents with correct spec-kit version in a single command

**Independent Test**: Run installation on fresh system, verify both components installed and functional

### Tests for User Story 1 (TDD)

**NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [X] T012 [P] [US1] Unit test for version comparison in internal/version/compare_test.go
- [X] T013 [P] [US1] Unit test for manifest parsing in pkg/models/manifest_test.go
- [X] T014 [P] [US1] Unit test for prefix detection in internal/install/detect_test.go
- [X] T015 [US1] Integration test for fresh install in tests/integration/test_install.sh

### Implementation for User Story 1

#### Step 1: Core Installation Logic

- [X] T016 [P] [US1] Implement .specify/ detection in internal/install/detect.go
- [X] T017 [P] [US1] Implement prefix determination logic in internal/install/detect.go (returns "." or ".claude-agent-templates")
- [X] T018 [US1] Implement file copy operations in internal/install/copy.go with prefix support
- [X] T019 [US1] Implement .claude/ directory setup in internal/install/claude.go (creates commands/ and agents/ subdirs)

#### Step 2: Version Management

- [X] T020 [US1] Implement version manifest loading in internal/version/manifest.go
- [X] T021 [US1] Implement version lock creation in internal/version/lock.go with UUID generation
- [X] T022 [US1] Implement version lock saving in internal/version/lock.go

#### Step 3: CLI Command

- [X] T023 [US1] Create cobra CLI main.go in cmd/spec-kit-agents/main.go with install command
- [X] T024 [US1] Implement install command handler in internal/install/install.go orchestrating all steps
- [X] T025 [US1] Add --prefix, --global, --force, --quiet flags to install command
- [X] T026 [US1] Implement installation verification in internal/install/install.go (checks files exist)

#### Step 4: Error Handling

- [X] T027 [US1] Implement error types in internal/install/errors.go (ErrVersionConflict, ErrPermission, etc.)
- [X] T028 [US1] Add clear error messages with resolution steps per contracts/go-cli.yaml
- [X] T029 [US1] Implement installation logging to ~/.claude-agent-templates/.install-log.txt

#### Step 5: One-liner Installer

- [X] T030 [US1] Create bash installer script in scripts/install.sh that detects OS/arch and downloads binary
- [X] T031 [US1] Add binary download logic from GitHub Releases to scripts/install.sh
- [X] T032 [P] [US1] Create PowerShell installer script in scripts/install.ps1 for Windows (optional)

**Checkpoint**: User Story 1 complete - users can run one-liner install and get working setup

---

## Phase 4: User Story 2 - Version Compatibility Management (Priority: P2)

**Goal**: Maintainers can update pinned spec-kit version and users get tested combination

**Independent Test**: Update version pin, run new install, verify new version used

### Tests for User Story 2 (TDD)

- [X] T033 [P] [US2] Unit test for version range checking in internal/version/compare_test.go
- [X] T034 [P] [US2] Unit test for compatibility validation in internal/version/check_test.go
- [X] T035 [US2] Integration test for version check in tests/integration/test_check.sh

### Implementation for User Story 2

#### Step 1: Version Checking Logic

- [X] T036 [P] [US2] Implement version compatibility checking in internal/version/check.go
- [X] T037 [P] [US2] Implement version range validation (min_version, max_version) in internal/version/check.go
- [X] T038 [P] [US2] Implement breaking version detection in internal/version/check.go

#### Step 2: CLI Command

- [X] T039 [US2] Add check command to cobra CLI in cmd/spec-kit-agents/main.go
- [ ] T040 [US2] Implement check command handler in internal/version/check.go
- [ ] T041 [US2] Add --json, --fix, --detailed flags to check command
- [ ] T042 [US2] Implement text output formatter for compatibility report in internal/version/check.go
- [ ] T043 [P] [US2] Implement JSON output formatter for compatibility report in internal/version/check.go

#### Step 3: Status Reporting

- [ ] T044 [US2] Add status command to cobra CLI in cmd/spec-kit-agents/main.go
- [ ] T045 [US2] Implement status command showing installation info and version history
- [ ] T046 [US2] Add version command showing binary version, build time, commit hash

**Checkpoint**: User Story 2 complete - maintainers can manage versions, users can check compatibility

---

## Phase 5: User Story 3 - Upgrade Path Management (Priority: P3)

**Goal**: Existing users can upgrade seamlessly without breaking workflows

**Independent Test**: Simulate upgrade from older version, verify smooth transition and backward compatibility

### Tests for User Story 3 (TDD)

- [ ] T047 [P] [US3] Unit test for backup creation in internal/install/backup_test.go
- [ ] T048 [P] [US3] Unit test for rollback logic in internal/install/rollback_test.go
- [ ] T049 [US3] Integration test for upgrade in tests/integration/test_upgrade.sh
- [ ] T050 [US3] Integration test for rollback on failure in tests/integration/test_rollback.sh

### Implementation for User Story 3

#### Step 1: Backup & Rollback

- [ ] T051 [P] [US3] Implement installation backup in internal/install/backup.go
- [ ] T052 [P] [US3] Implement rollback logic in internal/install/rollback.go
- [ ] T053 [US3] Add EXIT trap for automatic rollback on failure in internal/install/install.go

#### Step 2: Update Command

- [ ] T054 [US3] Add update command to cobra CLI in cmd/spec-kit-agents/main.go
- [ ] T055 [US3] Implement update command handler in internal/install/update.go
- [ ] T056 [US3] Add --version, --backup flags to update command
- [ ] T057 [US3] Implement version lock history appending in internal/version/lock.go

#### Step 3: Conflict Resolution

- [ ] T058 [US3] Implement version conflict detection in update workflow
- [ ] T059 [US3] Add user prompts for conflict resolution (upgrade, keep, abort)
- [ ] T060 [US3] Implement --fix flag for automatic conflict resolution in check command

**Checkpoint**: User Story 3 complete - full upgrade lifecycle with safety guarantees

---

## Phase 6: GitHub Release Automation

**Purpose**: Automate binary building and distribution

- [ ] T061 Create GitHub Actions release workflow in .github/workflows/release.yml
- [ ] T062 [P] Add build matrix for 6 platforms (linux/darwin/windows × amd64/arm64) to release.yml
- [ ] T063 [P] Configure GoReleaser for binary distribution (optional, or use manual build)
- [ ] T064 Add binary checksums generation to release workflow
- [ ] T065 Create release documentation template in .github/RELEASE_TEMPLATE.md

---

## Phase 7: Cross-Platform Testing

**Purpose**: Ensure installation works on all target platforms

- [ ] T066 [P] Add ubuntu-latest runner to GitHub Actions test workflow
- [ ] T067 [P] Add macos-latest runner to GitHub Actions test workflow
- [ ] T068 [P] Add windows-latest runner with Git Bash to GitHub Actions test workflow
- [ ] T069 Create integration test suite in tests/integration/test_all_platforms.sh
- [ ] T070 Test one-liner installer on fresh VM for each platform

---

## Phase 8: Documentation & Polish

**Purpose**: User-facing documentation and final touches

- [ ] T071 [P] Update main README.md with installation instructions and one-liner
- [ ] T072 [P] Update quickstart.md with actual CLI commands and examples
- [ ] T073 [P] Create CONTRIBUTING.md with development setup instructions
- [ ] T074 [P] Add inline help text to all CLI commands (--help output)
- [ ] T075 Create example .specify/version-manifest.json with comments
- [ ] T076 [P] Add shell completion scripts for bash/zsh (optional)
- [ ] T077 Validate all error messages are clear and actionable per contracts
- [ ] T078 Run final integration test suite across all platforms

---

## Dependencies & Execution Order

### User Story Dependencies

```
Setup (Phase 1) → Foundational (Phase 2) → ┬→ US1 (P1) ✅ MVP
                                            ├→ US2 (P2) (can start in parallel with US1)
                                            └→ US3 (P3) (depends on US1 install logic)

GitHub Release (Phase 6) can start after US1 core is stable
Cross-Platform Testing (Phase 7) runs throughout
Documentation (Phase 8) runs at the end
```

### Task-Level Dependencies

**Foundational** (T006-T011): All parallelizable except T009 (needs T006)

**US1 Tasks**:
- Tests (T012-T015): Fully parallel
- Installation logic (T016-T019): Fully parallel
- Version management (T020-T022): T022 depends on T020, T021
- CLI (T023-T026): T024 depends on T023, others depend on T024
- Error handling (T027-T029): Parallel
- Installer script (T030-T032): T031 depends on T030, T032 parallel

**US2 Tasks**:
- Tests (T033-T035): Fully parallel
- Version checking (T036-T038): Fully parallel
- CLI (T039-T043): T040-T043 depend on T039
- Status (T044-T046): Parallel

**US3 Tasks**:
- Tests (T047-T050): Fully parallel
- Backup/Rollback (T051-T053): T053 depends on T051, T052
- Update command (T054-T057): T055-T057 depend on T054
- Conflict resolution (T058-T060): Parallel

---

## Parallel Execution Examples

### Phase 2 (Foundational) - Can run in parallel

```bash
# Terminal 1: Data models
go test ./pkg/models/...

# Terminal 2: Version utilities
go test ./internal/version/...

# Terminal 3: Config utilities
go test ./internal/config/...

# Terminal 4: Version manifest file
# Manual creation of .specify/version-manifest.json
```

### User Story 1 - Tests phase

```bash
# All test files can be written in parallel by different developers
# Terminal 1
vim internal/version/compare_test.go

# Terminal 2
vim pkg/models/manifest_test.go

# Terminal 3
vim internal/install/detect_test.go

# Terminal 4
vim tests/integration/test_install.sh
```

### User Story 1 - Implementation phase

```bash
# Different components can be implemented in parallel
# Terminal 1: Detection logic
vim internal/install/detect.go

# Terminal 2: File operations
vim internal/install/copy.go

# Terminal 3: Claude integration
vim internal/install/claude.go

# Terminal 4: Version management
vim internal/version/manifest.go
```

---

## Testing Strategy

### Test Execution Order

1. **Unit tests** (run with `go test ./...`):
   - pkg/models/manifest_test.go
   - pkg/models/lock_test.go
   - internal/version/compare_test.go
   - internal/install/detect_test.go
   - All other *_test.go files

2. **Integration tests** (run with bash):
   - tests/integration/test_install.sh (fresh install)
   - tests/integration/test_check.sh (version checking)
   - tests/integration/test_upgrade.sh (upgrade workflow)
   - tests/integration/test_rollback.sh (rollback on failure)

3. **Cross-platform tests** (GitHub Actions):
   - Run all tests on Linux, macOS, Windows
   - Verify one-liner installer works on each platform

### Coverage Goals

- Unit tests: >80% code coverage
- Integration tests: All user stories covered
- Cross-platform: All platforms pass all tests

---

## Summary

**Total Tasks**: 78
**MVP Tasks** (US1): 33 tasks (Setup + Foundational + US1)
**Parallel Opportunities**: 42 tasks marked [P]

**Task Breakdown by Phase**:
- Setup: 5 tasks
- Foundational: 6 tasks
- US1 (P1 - MVP): 22 tasks
- US2 (P2): 14 tasks
- US3 (P3): 13 tasks
- GitHub Release: 5 tasks
- Cross-Platform Testing: 5 tasks
- Documentation: 8 tasks

**Suggested Delivery**:
1. **Week 1**: Setup + Foundational + US1 → MVP Release
2. **Week 2**: US2 + GitHub Release → Version 1.0
3. **Week 3**: US3 + Cross-Platform Testing → Version 1.1
4. **Week 4**: Documentation + Polish → Version 1.2 (production-ready)

**Independent Test Criteria**:
- ✅ **US1**: Fresh install command works, both components present
- ✅ **US2**: Version check detects conflicts, status shows history
- ✅ **US3**: Upgrade updates version, rollback restores on failure

Each user story delivers value independently and can be deployed separately.

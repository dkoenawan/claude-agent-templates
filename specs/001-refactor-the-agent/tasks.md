# Tasks: Agent Refactoring for Spec-Driven Development with GitHub Issues Interface

**Input**: Design documents from `/specs/001-refactor-the-agent/`
**Prerequisites**: plan.md (required), research.md, data-model.md, contracts/

## Execution Flow (main)
```
1. Load plan.md from feature directory ✓
   → Tech stack: Markdown/YAML for agent definitions, Bash for automation scripts
   → Structure: Single project - Claude agent template repository
2. Load design documents ✓:
   → data-model.md: Agent Specification, Domain, GitHub Issue Context, Validation Rules
   → contracts/: agent-specification-schema.yaml, github-workflow-api.yaml
   → research.md: S.P.E.C framework, domain specialization, GitHub Actions integration
3. Generate tasks by category ✓
4. Apply task rules ✓: [P] for parallel tasks (different files)
5. Number tasks sequentially ✓
6. Generate dependency graph ✓
7. Create parallel execution examples ✓
8. Validate task completeness ✓
```

## Format: `[ID] [P?] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- Include exact file paths in descriptions

## Path Conventions
- **Single project**: `agents/`, `scripts/`, `.github/workflows/` at repository root
- Agent specifications in domain-specific subdirectories
- Automation scripts in `.specify/scripts/` and `Taskfile.yml`

## Phase 3.1: Setup & Infrastructure

- [ ] **T001** Create domain-specific agent directories: `agents/dotnet/`, `agents/nodejs/`, `agents/java/`
- [ ] **T002** [P] Create agent specification validation script: `scripts/validate-agent-spec.sh`
- [ ] **T003** [P] Create GitHub Actions workflow for issue orchestration: `.github/workflows/issue-agent-orchestration.yml`
- [ ] **T004** [P] Create GitHub Actions workflow for execution phase: `.github/workflows/execute-phase.yml`
- [ ] **T005** [P] Create GitHub Actions workflow for agent validation: `.github/workflows/validate-agents.yml`

## Phase 3.2: Tests First (TDD) ⚠️ MUST COMPLETE BEFORE 3.3
**CRITICAL: These tests MUST be written and MUST FAIL before ANY implementation**

- [ ] **T006** [P] Contract test for Claude agent format validation: `tests/contract/test_claude_agent_format.py`
- [ ] **T007** [P] Contract test for GitHub Actions workflow specifications: `tests/contract/test_github_workflows.py`
- [ ] **T008** [P] Integration test for Python agent specialization: `tests/integration/test_python_agent_workflow.py`
- [ ] **T009** [P] Integration test for .NET agent specialization: `tests/integration/test_dotnet_agent_workflow.py`
- [ ] **T010** [P] Integration test for domain-agnostic agent workflow: `tests/integration/test_agnostic_agent_workflow.py`
- [ ] **T011** [P] Integration test for GitHub issue automation: `tests/integration/test_github_automation.py`
- [ ] **T012** [P] Integration test for agent specification validation: `tests/integration/test_agent_validation.py`

## Phase 3.3: Core Implementation (ONLY after tests are failing)

### Domain-Specific Agent Creation
- [ ] **T013** [P] Create solution-architect-dotnet agent: `agents/dotnet/solution-architect-dotnet.md`
- [ ] **T014** [P] Create software-engineer-dotnet agent: `agents/dotnet/software-engineer-dotnet.md`
- [ ] **T015** [P] Create test-engineer-dotnet agent: `agents/dotnet/test-engineer-dotnet.md`
- [ ] **T016** [P] Create solution-architect-nodejs agent: `agents/nodejs/solution-architect-nodejs.md`
- [ ] **T017** [P] Create software-engineer-nodejs agent: `agents/nodejs/software-engineer-nodejs.md`
- [ ] **T018** [P] Create test-engineer-nodejs agent: `agents/nodejs/test-engineer-nodejs.md`
- [ ] **T019** [P] Create solution-architect-java agent: `agents/java/solution-architect-java.md`
- [ ] **T020** [P] Create software-engineer-java agent: `agents/java/software-engineer-java.md`
- [ ] **T021** [P] Create test-engineer-java agent: `agents/java/test-engineer-java.md`

### Enhanced Specification Support
- [ ] **T022** Extend existing requirements-analyst agent with enhanced specifications: `agents/core/requirements-analyst.md`
- [ ] **T023** Extend existing solution-architect agent with enhanced specifications: `agents/core/solution-architect.md`
- [ ] **T024** Extend existing software-engineer-python agent with enhanced specifications: `agents/python/software-engineer-python.md`
- [ ] **T025** Extend existing test-engineer-python agent with enhanced specifications: `agents/python/test-engineer-python.md`
- [ ] **T026** Extend existing documentation agent with enhanced specifications: `agents/core/documentation.md`

### Validation and Schema Implementation
- [ ] **T027** [P] Implement Claude agent format validator: `scripts/validate-claude-agent.py`
- [ ] **T028** [P] Implement domain classifier logic: `scripts/classify-domain.py`
- [ ] **T029** [P] Implement workflow state tracker: `scripts/track-workflow.py`

## Phase 3.4: Integration & Automation

- [ ] **T030** Create agent selection logic for GitHub Issues: `.github/scripts/select-agent.py`
- [ ] **T031** Implement GitHub issue parsing for agent inputs: `.github/scripts/parse-issue.py`
- [ ] **T032** Create workflow metrics collection: `.github/scripts/collect-metrics.py`
- [ ] **T033** Update Taskfile with new validation commands: `Taskfile.yml`
- [ ] **T034** Create agent installation automation: `scripts/install-agents.sh`
- [ ] **T035** Create workflow installation automation: `scripts/install-workflows.sh`

## Phase 3.5: Documentation & Polish

- [ ] **T036** [P] Create agent specification documentation: `docs/agent-specifications.md`
- [ ] **T037** [P] Create domain specialization guide: `docs/domain-specialization.md`
- [ ] **T038** [P] Create GitHub Actions integration guide: `docs/github-actions-setup.md`
- [ ] **T039** [P] Create troubleshooting guide: `docs/troubleshooting.md`
- [ ] **T040** [P] Update main README with refactoring information: `README.md`
- [ ] **T041** [P] Create migration guide for existing users: `docs/migration-guide.md`
- [ ] **T042** [P] Unit tests for validation logic: `tests/unit/test_validation.py`
- [ ] **T043** [P] Unit tests for domain classification: `tests/unit/test_domain_classification.py`
- [ ] **T044** [P] Performance tests for agent processing (<2 seconds): `tests/performance/test_agent_performance.py`

## Dependencies

### Critical Path Dependencies
- **Setup First**: T001 blocks all agent creation tasks (T013-T021)
- **Tests Before Implementation**: T006-T012 must complete and FAIL before T013-T035
- **Agent Extensions**: T022-T026 can only begin after existing agents are analyzed
- **Schema Validation**: T027 depends on T006-T007 test implementation
- **Integration Logic**: T030-T032 depend on T027-T029 validation implementation
- **Automation**: T033-T035 depend on T030-T032 integration logic
- **Documentation**: T036-T041 can run in parallel after implementation is complete

### Parallel Execution Groups
1. **Setup Phase**: T002-T005 (different workflow files)
2. **Test Creation**: T006-T012 (different test files)
3. **Agent Creation**: T013-T021 (different agent files per domain)
4. **Schema Implementation**: T027-T029 (different script files)
5. **Documentation**: T036-T041 (different documentation files)

## Parallel Example
```bash
# Launch T006-T012 together (Test Creation Phase):
Task: "Contract test agent-specification-schema.yaml validation in tests/contract/test_agent_spec_schema.py"
Task: "Contract test github-workflow-api.yaml validation in tests/contract/test_workflow_api_schema.py"
Task: "Integration test Python agent workflow in tests/integration/test_python_agent_workflow.py"
Task: "Integration test .NET agent workflow in tests/integration/test_dotnet_agent_workflow.py"
Task: "Integration test domain-agnostic workflow in tests/integration/test_agnostic_agent_workflow.py"
Task: "Integration test GitHub automation in tests/integration/test_github_automation.py"
Task: "Integration test agent validation in tests/integration/test_agent_validation.py"

# Launch T013-T021 together (Agent Creation Phase):
Task: "Create solution-architect-dotnet agent in agents/dotnet/solution-architect-dotnet.md"
Task: "Create software-engineer-dotnet agent in agents/dotnet/software-engineer-dotnet.md"
Task: "Create test-engineer-dotnet agent in agents/dotnet/test-engineer-dotnet.md"
# ... etc for all domain-specific agents

# Launch T036-T041 together (Documentation Phase):
Task: "Create agent specification documentation in docs/agent-specifications.md"
Task: "Create domain specialization guide in docs/domain-specialization.md"
Task: "Create GitHub Actions setup guide in docs/github-actions-setup.md"
Task: "Create troubleshooting guide in docs/troubleshooting.md"
Task: "Update main README with refactoring info in README.md"
Task: "Create migration guide in docs/migration-guide.md"
```

## Notes
- [P] tasks = different files, no dependencies between them
- Verify all tests fail before implementing (TDD critical)
- Commit after each task completion
- Each agent specification must include all schema elements from contracts/
- Maintain backward compatibility throughout implementation
- Focus on MVP - avoid feature creep

## Task Generation Rules Applied

1. **From Contracts**:
   - `agent-specification-schema.yaml` → T006 contract test + agent spec tasks (T013-T026)
   - `github-workflow-specifications.yaml` → T007 contract test + workflow automation (T003-T005, T030-T032)

2. **From Data Model**:
   - Agent Specification entity → T013-T026 (agent creation tasks)
   - Domain entity → T013-T021 (domain-specific variants)
   - Validation Rules entity → T027, T042-T043 (validation implementation)
   - GitHub Issue Context → T008-T011, T030-T031 (issue processing)

3. **From Research Decisions**:
   - S.P.E.C Framework → T002, T027-T029 (specification validation)
   - Domain Specialization → T013-T021 (technology-specific agents)
   - GitHub Actions Integration → T003-T005, T030-T035 (automation)
   - Backward Compatibility → T022-T026 (existing agent enhancement)

4. **From Quickstart Scenarios**:
   - Python development workflow → T008
   - .NET development workflow → T009
   - Cross-domain documentation → T010
   - GitHub automation → T011

## Validation Checklist
*GATE: Checked before task execution*

- [x] All contracts have corresponding tests (T006-T007)
- [x] All entities have implementation tasks (T013-T029)
- [x] All tests come before implementation (T006-T012 before T013+)
- [x] Parallel tasks truly independent (different files/directories)
- [x] Each task specifies exact file path
- [x] No task modifies same file as another [P] task
- [x] Research decisions translated to implementation tasks
- [x] Quickstart scenarios covered by integration tests
- [x] MVP scope maintained throughout task list
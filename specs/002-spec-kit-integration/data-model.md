# Data Model: Spec-Kit Integration

**Feature**: 002-spec-kit-integration
**Date**: 2025-10-20
**Purpose**: Define data structures for specifications, agents, skills, and workflows

## Overview

This feature introduces several key entities that need clear data models to ensure consistency across the system. All entities are stored as files (markdown with YAML frontmatter) following the principle of simplicity and version control compatibility.

---

## Entity 1: Feature Specification

**Purpose**: Represents a single feature or enhancement with requirements, user scenarios, and success criteria

**Storage**: `specs/[###-feature-name]/spec.md`

**Structure**:
```yaml
---
feature_number: "002"          # Sequential number
feature_name: "short-name"     # Kebab-case identifier
created_date: "2025-10-20"     # ISO 8601 date
status: "draft"                # draft | in-planning | in-progress | completed
branch: "002-short-name"       # Git branch name
---
```

**Markdown Sections** (mandatory):
- Title and metadata header
- User Scenarios & Testing (with prioritized user stories P1, P2, P3)
- Requirements (functional requirements with FR-XXX identifiers)
- Success Criteria (measurable outcomes with SC-XXX identifiers)
- Assumptions
- Dependencies
- Scope (in-scope, out-of-scope)

**Key Entities** (optional section within spec):
- Listed as bullet points describing domain entities

**Validation Rules**:
- Feature number must be unique and sequential
- Status must be one of defined values
- All mandatory sections must be present
- Functional requirements must be testable
- Success criteria must be measurable and technology-agnostic
- No [NEEDS CLARIFICATION] markers in completed specs

**Relationships**:
- One Feature Specification → One Implementation Plan
- One Feature Specification → Many Tasks (via tasks.md)
- One Feature Specification → One Git Branch

---

## Entity 2: Implementation Plan

**Purpose**: Defines technical approach, architecture decisions, and phase-by-phase design

**Storage**: `specs/[###-feature-name]/plan.md`

**Structure**:
```yaml
---
feature_number: "002"
feature_name: "short-name"
plan_date: "2025-10-20"
spec_file: "spec.md"
phase: "0"                     # 0 (research) | 1 (design) | 2 (ready)
---
```

**Markdown Sections** (mandatory):
- Summary (extracted from spec + technical approach)
- Technical Context (language, dependencies, storage, testing, platform, constraints)
- Constitution Check (compliance validation)
- Project Structure (documentation + source code layout)
- Complexity Tracking (only if constitution violations exist)

**Phase-Specific Sections**:
- Phase 0: Research findings in separate research.md
- Phase 1: Data models, contracts, quickstart guide
- Phase 2: Task breakdown (external tasks.md file)

**Validation Rules**:
- Must reference valid spec.md
- Technical Context must have no unresolved NEEDS CLARIFICATION before Phase 1
- Constitution Check must pass (or violations justified)
- Project structure must match repository conventions

**Relationships**:
- One Implementation Plan → One Feature Specification
- One Implementation Plan → One Research Document
- One Implementation Plan → One Data Model Document
- One Implementation Plan → Many Contracts
- One Implementation Plan → One Quickstart Guide

---

## Entity 3: Domain Agent

**Purpose**: Technology-specific agent with specialized knowledge conforming to spec-kit format

**Storage**: `agents/[domain]/[role]-[domain].md`

**Structure**:
```yaml
---
name: "solution-architect-python"  # Unique identifier
description: "Brief description"
domain: "python"                   # python | dotnet | nodejs | java | core
role: "architect"                  # analyst | architect | engineer | test-engineer | documentation
spec_version: "1.0"               # Agent spec format version
tools: ["Bash", "Edit", "Read"]   # Available tools
model: "inherit"                   # Model selection
color: "blue"                      # UI color coding
workflow_position: 4               # 1-9 in workflow
github_integration:
  triggers: ["plan-approved"]      # Label triggers
  outputs: ["implementation-ready"] # Output labels
  permissions: ["contents:read"]   # GitHub permissions
---
```

**Markdown Sections** (mandatory):
- Role and responsibilities
- Domain-specific expertise
- Inputs (what agent receives)
- Outputs (what agent produces)
- Validation criteria (how to verify agent's work)
- Examples (with context, input, output)

**Validation Rules**:
- Name must match file name pattern
- Domain must be valid (python | dotnet | nodejs | java | core)
- Role must be valid (analyst | architect | engineer | test-engineer | documentation)
- Spec version must match current standard
- Must have at least one example
- Workflow position must be 1-9
- GitHub integration triggers/outputs must be valid labels

**Relationships**:
- Many Agents → One Domain
- Many Agents → Many GitHub Issue Labels (via triggers/outputs)
- One Agent → One Role in Workflow

**State Transitions**:
None - agents are immutable once defined (versioned updates create new versions)

---

## Entity 4: Claude Skill

**Purpose**: Reusable workflow capability that can be invoked to perform specialized tasks

**Storage**: `.claude/skills/[skill-name]/skill.md`

**Structure**:
```yaml
---
name: "domain-classifier"          # Unique identifier
version: "1.0.0"                   # Semantic version
description: "Brief description"
parameters:
  - name: "feature_description"
    type: "string"
    required: true
    description: "Feature text to classify"
  - name: "confidence_threshold"
    type: "number"
    required: false
    default: 0.7
    description: "Minimum confidence for classification"
outputs:
  - name: "domain"
    type: "string"
    values: ["python", "dotnet", "nodejs", "java", "core"]
  - name: "confidence"
    type: "number"
tools: ["Read", "Grep", "Bash"]    # Required Claude Code tools
dependencies: []                    # Other skills this depends on
---
```

**Markdown Sections** (mandatory):
- Purpose and use cases
- Parameter descriptions
- Output specifications
- Implementation logic
- Examples (input → output)
- Error handling

**Validation Rules**:
- Name must be kebab-case, unique within skills directory
- Version must follow semantic versioning (MAJOR.MINOR.PATCH)
- Required parameters must be listed
- Output types must be specified
- Must have at least one example
- Tools must be valid Claude Code tool names

**Relationships**:
- Many Skills → Many Slash Commands (skills can be invoked by commands)
- One Skill → Many Skills (via dependencies)
- One Skill → Many Claude Code Tools

**State Transitions**:
- Version increments on changes:
  - MAJOR: Breaking changes (parameter signature changes)
  - MINOR: New features (new parameters with defaults)
  - PATCH: Bug fixes (no interface changes)

---

## Entity 5: Slash Command

**Purpose**: User-facing command for executing spec-kit workflows

**Storage**: `.claude/commands/speckit.[command].md`

**Structure**:
```yaml
---
command: "/speckit.specify"        # Command invocation
description: "Brief description"
category: "spec-kit"               # Grouping category
parameters:
  - name: "feature_description"
    type: "string"
    required: true
    description: "Natural language feature description"
skills_used: ["domain-classifier"] # Skills invoked by this command
phase: "requirements"              # requirements | planning | implementation | testing | documentation
---
```

**Markdown Sections** (mandatory):
- Command overview
- Input requirements
- Execution workflow (step-by-step)
- Output artifacts
- Examples (command → results)
- Error handling and troubleshooting

**Validation Rules**:
- Command must start with `/speckit.`
- Must reference valid skills if skills_used is populated
- Phase must be valid workflow phase
- Must have at least one example showing usage

**Relationships**:
- One Slash Command → Many Skills (via skills_used)
- Many Slash Commands → One Workflow Phase
- One Slash Command → Many Output Files (artifacts created)

**State Transitions**:
None - commands are stateless, invoke skills which may have state

---

## Entity 6: Workflow Phase

**Purpose**: Stage in the spec-driven development lifecycle

**Representation**: Enum/constant (not a stored file, but a shared concept)

**Values**:
```yaml
phases:
  - id: 1
    name: "requirements"
    label: "agent-ready"
    agent_role: "analyst"
    outputs: ["spec.md", "checklists/requirements.md"]

  - id: 2
    name: "planning"
    label: "requirements-ready"
    agent_role: "architect"
    outputs: ["plan.md", "research.md", "data-model.md", "contracts/", "quickstart.md"]

  - id: 3
    name: "test-planning"
    label: "plan-approved"
    agent_role: "test-engineer"
    outputs: ["test-strategy.md", "checklists/testing.md"]

  - id: 4
    name: "implementation-ready"
    label: "tests-planned"
    outputs: ["tasks.md"]

  - id: 5
    name: "implementation"
    label: "implementation-ready"
    agent_role: "engineer"
    outputs: ["source code", "tests", "PR"]

  - id: 6
    name: "review"
    label: "implementation-complete"
    outputs: ["review comments", "approval"]

  - id: 7
    name: "user-acceptance"
    label: "user-accepted"
    outputs: ["acceptance confirmation"]

  - id: 8
    name: "documentation"
    label: "documentation-complete"
    agent_role: "documentation"
    outputs: ["updated docs", "CHANGELOG.md"]
```

**Validation Rules**:
- Phase IDs must be sequential 1-9
- Phase names must be unique
- GitHub labels must be unique
- Outputs must be valid file/artifact types

**Relationships**:
- Many Agents → Many Workflow Phases (via workflow_position)
- Many Slash Commands → Many Workflow Phases (via phase)
- One Feature Specification → One Current Phase at any time

**State Transitions**:
- Linear progression: phase N → phase N+1
- No skipping phases
- Can regress (e.g., implementation → planning if issues found)

---

## Entity 7: GitHub Issue Integration

**Purpose**: Connection between GitHub issues and spec-driven development workflow

**Storage**: Issue metadata (GitHub API), `.workflow-state.json` (local tracking)

**Structure** (in .workflow-state.json):
```json
{
  "issue_number": 123,
  "feature_number": "002",
  "current_phase": "planning",
  "current_label": "requirements-ready",
  "domain": "python",
  "assigned_agents": [
    "requirements-analyst",
    "solution-architect-python"
  ],
  "created_at": "2025-10-20T08:00:00Z",
  "updated_at": "2025-10-20T10:30:00Z",
  "artifacts": {
    "spec": "specs/002-spec-kit-integration/spec.md",
    "plan": "specs/002-spec-kit-integration/plan.md"
  }
}
```

**Validation Rules**:
- Issue number must be valid GitHub issue
- Feature number must match existing feature
- Current phase must be valid workflow phase
- Current label must match phase label
- Domain must be valid if assigned
- Assigned agents must exist in agents/ directory

**Relationships**:
- One GitHub Issue → One Feature Specification
- One GitHub Issue → One Current Workflow Phase
- One GitHub Issue → Many Agents (assigned over lifecycle)
- One GitHub Issue → Many Artifacts (created during workflow)

**State Transitions**:
Follows Workflow Phase transitions, triggered by:
- GitHub label changes
- Agent completion signals
- User approval actions

---

## Data Flow Diagram

```
[User Creates Issue]
       ↓
[Issue Classified] → [Domain Determined] → [Agents Assigned]
       ↓
[Requirements Phase] → [spec.md created]
       ↓
[Planning Phase] → [plan.md, research.md, data-model.md, contracts/, quickstart.md created]
       ↓
[Test Planning Phase] → [test-strategy.md created]
       ↓
[Task Breakdown] → [tasks.md created]
       ↓
[Implementation Phase] → [Source code, tests, PR created]
       ↓
[Review Phase] → [Code review, approval]
       ↓
[User Acceptance] → [User validates against spec.md]
       ↓
[Documentation Phase] → [Docs updated, CHANGELOG updated]
       ↓
[Complete] → [Issue closed, branch merged]
```

---

## Storage Summary

| Entity | Storage Location | Format | Version Control |
|--------|------------------|--------|-----------------|
| Feature Specification | `specs/[###-name]/spec.md` | Markdown + YAML | Git |
| Implementation Plan | `specs/[###-name]/plan.md` | Markdown + YAML | Git |
| Research Document | `specs/[###-name]/research.md` | Markdown | Git |
| Data Model | `specs/[###-name]/data-model.md` | Markdown | Git |
| Quickstart Guide | `specs/[###-name]/quickstart.md` | Markdown | Git |
| Tasks | `specs/[###-name]/tasks.md` | Markdown | Git |
| Contracts | `specs/[###-name]/contracts/*.yaml` | YAML/JSON | Git |
| Domain Agent | `agents/[domain]/[role]-[domain].md` | Markdown + YAML | Git |
| Claude Skill | `.claude/skills/[skill]/skill.md` | Markdown + YAML | Git |
| Slash Command | `.claude/commands/speckit.[cmd].md` | Markdown + YAML | Git |
| Workflow State | `.workflow-state.json` | JSON | Git (gitignored for local state) |

---

## Validation & Integrity

### Referential Integrity

- Feature specs reference valid branches (validated by git)
- Plans reference valid feature specs (validated by file existence)
- Agents reference valid domains (validated by enum)
- Skills reference valid tools (validated against Claude Code tool list)
- Slash commands reference valid skills (validated by file existence)
- Workflow states reference valid phases (validated by enum)

### Schema Validation

- Python script `scripts/validate-claude-agent.py` validates agent format
- YAML frontmatter parsed and validated on command execution
- Missing required fields cause validation errors
- Invalid enum values rejected

### Consistency Checks

- Feature numbers sequential (enforced by creation script)
- Branch names match feature names (enforced by creation script)
- Workflow phase transitions valid (enforced by state machine)
- Skill versions follow semantic versioning (validated by pattern)

---

## Migration from Existing Format

### Current Agent Format → Spec-Kit Agent Format

**Changes Required**:
- Add `spec_version: "1.0"` to frontmatter
- Add `workflow_position` field
- Add `github_integration` section
- Ensure `examples` section has proper structure

**Migration Skill**: `agent-migrator` skill handles automatic conversion

### Current Feature Structure → Spec-Kit Feature Structure

**Changes Required**:
- Move feature specs to `specs/[###-name]/` directory
- Add YAML frontmatter with metadata
- Ensure all mandatory sections present
- Generate missing artifacts (plan.md, tasks.md, etc.)

**Migration Command**: `/speckit.migrate-feature` (future enhancement)

---

**Data Model Status**: ✅ COMPLETE
**Entities Defined**: 7
**Ready for Contracts Generation**: YES

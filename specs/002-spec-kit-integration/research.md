# Research & Technology Decisions: Spec-Kit Integration

**Feature**: 002-spec-kit-integration
**Date**: 2025-10-20
**Purpose**: Resolve technical clarifications and document technology choices

## Research Questions

This document addresses the NEEDS CLARIFICATION items identified in the Technical Context:

1. Claude API integration approach (direct API vs Claude Code CLI integration)
2. Skill packaging format and distribution mechanism
3. Claude skills distribution and versioning strategy
4. Expected number of Claude skills (start with 3-5 but plan for growth)

---

## Decision 1: Claude Integration Approach

**Decision**: Use Claude Code CLI slash commands (not direct Claude API integration)

**Rationale**:
- Claude Code provides a standardized CLI interface for executing slash commands
- Slash commands in `.claude/commands/` are automatically discovered and executed
- No need for API key management or direct API calls
- Better integration with Claude Code's existing workflow and context management
- Slash commands can be version-controlled alongside the project
- Users already using Claude Code don't need additional setup

**Alternatives Considered**:
- **Direct Claude API**: Would require API key management, custom HTTP client implementation, and more complex error handling. Rejected because it adds unnecessary complexity and doesn't integrate with Claude Code's native workflow.
- **GitHub Actions with Claude API**: Would only work for GitHub-hosted workflows, not local development. Rejected because offline/local development is a core requirement (FR-008).

**Implementation Impact**:
- Create markdown files in `.claude/commands/` for each spec-kit workflow command
- Each command file contains the full prompt/workflow description
- Claude Code CLI automatically exposes these as `/speckit.*` commands
- No additional dependencies or authentication required

---

## Decision 2: Claude Skills Packaging Format

**Decision**: Use Claude Code's native skill format (markdown files with YAML frontmatter in `.claude/skills/` directory)

**Rationale**:
- Claude Code has built-in support for skills as discoverable capabilities
- Skills are defined using markdown files with structured frontmatter
- Skills can invoke tools, execute code, and chain workflows
- Version-controlled alongside the project for easy distribution
- Compatible with Claude Code's agent framework

**Alternatives Considered**:
- **Custom JSON/YAML format**: Would require custom parser and loader. Rejected because it duplicates Claude Code's native capability.
- **Python modules**: Would require Python runtime and complex packaging. Rejected because it adds language-specific dependencies.
- **Shell scripts**: Would lack structured metadata and discoverability. Rejected because it doesn't integrate with Claude Code's skill system.

**Skill Structure**:
```markdown
---
name: skill-name
description: Brief description
parameters:
  - name: param1
    type: string
    required: true
---

# Skill Implementation

[Skill workflow and logic here]
```

**Implementation Impact**:
- Create `.claude/skills/` directory structure
- Define 3-5 exemplar skills for common workflows
- Skills can be invoked from slash commands or other skills
- Skills inherit Claude Code's tool access and permissions

---

## Decision 3: Skills Distribution Strategy

**Decision**: Git-based distribution with semantic versioning in skill metadata

**Rationale**:
- Skills are version-controlled as part of the repository
- Users clone/fork the repository to get skills
- Updates distributed via git pull/fetch
- Semantic versioning in skill YAML frontmatter for compatibility tracking
- No separate package registry or distribution infrastructure needed
- Aligns with project goal of simplicity and file-based approach

**Alternatives Considered**:
- **NPM package**: Would require Node.js ecosystem. Rejected because it adds unnecessary dependency.
- **PyPI package**: Would require Python packaging. Rejected because skills are markdown, not Python modules.
- **Custom registry**: Would require infrastructure and maintenance. Rejected because it violates simplicity principle.
- **GitHub Releases**: Would require separate download/installation step. Rejected because git-based is more seamless.

**Versioning Strategy**:
- Each skill declares `version: "1.2.3"` in YAML frontmatter
- Breaking changes increment major version
- New features increment minor version
- Bug fixes increment patch version
- Project includes `SKILL_COMPATIBILITY.md` documenting version requirements

**Implementation Impact**:
- Skills distributed as part of repository
- Users get latest skills with `git pull`
- Version compatibility checked by validation scripts
- Migration guides provided when skills have breaking changes

---

## Decision 4: Initial Skills Scope

**Decision**: Start with 5 exemplar skills, plan for 20+ over time

**Initial 5 Skills** (for demonstration and core workflows):

1. **`domain-classifier`**: Analyze feature description and classify technology domain (Python/Node.js/.NET/Java)
2. **`agent-migrator`**: Migrate existing agent specifications to spec-kit format
3. **`contract-generator`**: Generate OpenAPI/GraphQL contracts from functional requirements
4. **`test-scaffold`**: Generate test structure and boilerplate based on test strategy
5. **`infra-initializer`**: Scaffold infrastructure-as-code (Terraform) and monitoring (Datadog) setup

**Growth Plan** (additional skills for production use):

- **SDLC Workflow Skills**:
  - `ci-pipeline-generator`: Generate GitHub Actions/GitLab CI configs
  - `pr-template-generator`: Create PR templates based on feature type
  - `changelog-updater`: Update CHANGELOG.md based on commits

- **Development Best Practices**:
  - `code-reviewer`: Automated code review against project standards
  - `security-scanner`: Check for common security issues
  - `performance-analyzer`: Identify performance bottlenecks

- **Documentation Skills**:
  - `api-doc-generator`: Generate API documentation from contracts
  - `architecture-diagram`: Create architecture diagrams from specs
  - `onboarding-guide`: Generate onboarding docs for new contributors

**Rationale for 5 Initial Skills**:
- Demonstrates all major skill categories (analysis, migration, generation, scaffolding)
- Covers critical path for spec-kit adoption (domain classification, agent migration)
- Showcases integration with SDLC best practices (contracts, tests, infrastructure)
- Small enough to implement and test thoroughly
- Large enough to validate skill framework design

**Implementation Impact**:
- Phase 1 implementation focuses on 5 core skills
- Skill framework designed to support 20+ skills without major refactoring
- Documentation includes skill development guide for community contributions
- Each skill has comprehensive tests and examples

---

## Best Practices Research

### GitHub Spec-Kit Integration

**Research Finding**: GitHub's spec-kit uses `.specify/` directory convention

**Key Patterns**:
- Feature specs in `specs/[feature-name]/spec.md`
- Templates in `.specify/templates/`
- Scripts in `.specify/scripts/`
- Memory/knowledge base in `.specify/memory/`

**Adoption Strategy**:
- Mirror GitHub's directory structure exactly for compatibility
- Extend with additional opinionated templates (infrastructure, monitoring)
- Add validation scripts for project-specific requirements
- Maintain separation between framework (`.specify/`) and features (`specs/`)

### Slash Command Best Practices

**Research Finding**: Effective slash commands follow clear naming and nesting conventions

**Key Patterns**:
- Namespace-based naming: `/namespace.command`
- Clear verb-based actions: specify, plan, implement, analyze
- Consistent parameter handling: `--option value` format
- JSON output support for automation: `--json` flag

**Adoption Strategy**:
- All spec-kit commands use `/speckit.*` namespace
- Each command has clear single responsibility
- Commands are composable (output of one feeds into another)
- Support both interactive and automated (CI/CD) usage

### Multi-Domain Agent Framework

**Research Finding**: Agent specifications benefit from standardized format with domain-specific extensions

**Key Patterns**:
- YAML frontmatter for structured metadata
- Markdown body for detailed instructions
- Clear separation between core and domain-specific knowledge
- Examples section for concrete guidance

**Adoption Strategy**:
- Migrate existing agents to spec-kit format while preserving content
- Maintain domain-specific directories (python/, dotnet/, nodejs/, java/)
- Core agents in `core/` directory for cross-domain capabilities
- Validation ensures format consistency while allowing domain flexibility

---

## Technology Stack Summary

| Component | Technology | Version | Rationale |
|-----------|-----------|---------|-----------|
| Command Interface | Claude Code Slash Commands | Latest | Native integration, no additional setup |
| Skills Framework | Claude Code Skills | Latest | Built-in support, markdown-based |
| Scripting | Bash | 4.0+ | Cross-platform, no dependencies |
| Validation | Python | 3.8+ | Rich testing ecosystem, existing scripts |
| Spec Format | Markdown + YAML | CommonMark | Human-readable, version-controllable |
| Distribution | Git | 2.x+ | Standard version control |
| CI/CD | GitHub Actions | Latest | Already in use, tight integration |
| Documentation | Markdown | CommonMark | Simple, portable, GitHub-rendered |

---

## Risk Assessment

### Technical Risks

1. **Claude Code API Changes**
   - **Risk**: Claude Code slash command/skill format changes in future versions
   - **Mitigation**: Version pin in documentation, follow semantic versioning, provide migration guides
   - **Severity**: Medium
   - **Likelihood**: Low (Claude Code appears stable)

2. **GitHub Spec-Kit Evolution**
   - **Risk**: GitHub's spec-kit introduces breaking changes
   - **Mitigation**: Monitor spec-kit repo, maintain loose coupling, override/extend where needed
   - **Severity**: Medium
   - **Likelihood**: Medium (early-stage project)

3. **Backward Compatibility**
   - **Risk**: Breaking existing agent specs or workflows during migration
   - **Mitigation**: Comprehensive test suite, gradual migration path, dual format support during transition
   - **Severity**: High
   - **Likelihood**: Low (with proper testing)

### Adoption Risks

1. **Learning Curve for Non-Technical Users**
   - **Risk**: Users unfamiliar with Git/CLI may struggle
   - **Mitigation**: Comprehensive quickstart guides, video tutorials, clear error messages, sensible defaults
   - **Severity**: High (primary target audience)
   - **Likelihood**: Medium

2. **Skill Discovery and Reuse**
   - **Risk**: Users may not discover available skills or understand when to use them
   - **Mitigation**: Skill catalog in documentation, clear naming conventions, examples in README
   - **Severity**: Medium
   - **Likelihood**: Medium

---

## Next Steps

With these technology decisions made, we can proceed to:

1. **Phase 1 - Design**:
   - Create data model for feature specifications, agents, skills
   - Define API contracts for CLI tools (if needed)
   - Generate quickstart guide for users

2. **Phase 2 - Implementation Planning**:
   - Break down into concrete tasks (`/speckit.tasks`)
   - Prioritize based on dependencies and risk
   - Estimate effort for each task

3. **Validation**:
   - Update agent context with technology stack
   - Re-check constitution compliance
   - Review with stakeholders

---

**Research Status**: ✅ COMPLETE
**Clarifications Resolved**: 4/4
**Ready for Phase 1**: YES

<!--
Sync Impact Report:
- Version change: 2.1.1 → 2.2.0 (MINOR: New principle structure + comprehensive governance)
- Modified principles: All principles restructured from template to concrete implementation
- Added sections: Agent Specialization, Quality Assurance, Development Workflow
- Removed sections: Template placeholders and example comments
- Templates requiring updates:
  ✅ .specify/templates/plan-template.md (reviewed - Constitution Check aligns)
  ✅ .specify/templates/spec-template.md (reviewed - requirements alignment maintained)
  ✅ .specify/templates/tasks-template.md (reviewed - TDD approach consistent)
- Follow-up TODOs: None - all placeholders resolved
-->

# Claude Agent Templates Constitution

## Core Principles

### I. Spec-Driven Development (NON-NEGOTIABLE)
All agent communication MUST occur through detailed specifications. Specifications serve as the primary protocol between agents, eliminating miscommunication. Every feature begins with a complete specification before any implementation. Agents MUST validate against specifications throughout the development lifecycle.

**Rationale**: Solves the critical problem of agent miscommunication by establishing specifications as the universal language between agents, ensuring predictable and reliable software delivery.

### II. Domain Specialization
Agents MUST be specialized by technology domain (Python, .NET, Node.js, Java) and workflow role (analyst, architect, engineer, test-engineer, documentation). Each agent MUST maintain deep expertise in their domain's frameworks, patterns, and best practices. Cross-domain knowledge is encouraged but secondary to domain mastery.

**Rationale**: Domain expertise ensures high-quality, idiomatic code that follows established patterns and leverages appropriate frameworks within each technology ecosystem.

### III. Test-First Development (NON-NEGOTIABLE)
TDD is mandatory: Tests written → User approved → Tests fail → Then implement. Red-Green-Refactor cycle strictly enforced. All implementations MUST achieve >80% code coverage. Contract tests required for all API specifications. Integration tests required for workflow transitions.

**Rationale**: Ensures code quality, prevents regressions, and validates that implementations meet specifications exactly as defined.

### IV. Hexagonal Architecture
All implementations MUST follow hexagonal/clean architecture patterns with clear separation of concerns. Domain logic isolated from infrastructure. Dependency injection required for all external dependencies. Port-adapter pattern enforced for external integrations.

**Rationale**: Maintains code maintainability, testability, and enables long-term evolution of the codebase across multiple technology domains.

### V. GitHub Integration
All workflows MUST integrate with GitHub Issues and Actions. Issue classification and agent assignment MUST be automated. Workflow state tracking through GitHub labels is required. Progress visibility through issue comments and status updates is mandatory.

**Rationale**: Provides complete traceability and automation of the development lifecycle while leveraging GitHub's native collaboration and tracking capabilities.

## Agent Specialization

Agent roles MUST follow the 9-step workflow pattern:
1. **Requirements Analyst** - Business requirement extraction and clarification
2-4. **Solution Architect (Domain)** - Architecture design with domain expertise
5. **Test Engineer (Domain)** - Comprehensive test strategy with domain frameworks
6. **Software Engineer (Domain)** - Implementation following domain patterns
9. **Documentation** - Final documentation and repository cleanup

All domain-specific agents MUST maintain expertise in their technology stack's frameworks, testing tools, and architectural patterns.

## Quality Assurance

All code MUST pass comprehensive validation:
- Agent specification format validation
- Domain-specific linting and formatting
- Test coverage >80% with meaningful tests
- Integration test validation for workflow steps
- Performance benchmarks where applicable
- Security best practices enforcement

No code reaches production without passing all quality gates.

## Development Workflow

Trunk-based development with short-lived feature branches. Each agent modification requires:
1. Specification update if behavior changes
2. Comprehensive testing including integration tests
3. Validation against existing workflows
4. Documentation updates
5. Real-world testing in sample projects

All changes MUST maintain backward compatibility unless explicitly versioned as breaking changes.

## Governance

This constitution supersedes all other development practices. All pull requests and code reviews MUST verify constitutional compliance. Complexity that violates principles MUST be justified with documented rationale and simpler alternatives considered.

Constitutional amendments require:
1. Documented rationale for the change
2. Impact assessment on existing agents and workflows
3. Migration plan for affected components
4. Approval from project maintainers
5. Version increment following semantic versioning

Use CLAUDE.md for runtime development guidance and day-to-day development practices.

**Version**: 2.2.0 | **Ratified**: 2025-09-22 | **Last Amended**: 2025-09-22
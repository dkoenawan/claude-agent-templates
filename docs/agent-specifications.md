# Agent Specifications Guide

This guide provides comprehensive documentation for Claude Agent specifications, including format requirements, validation rules, and best practices.

## Table of Contents

- [Specification Format](#specification-format)
- [Required Fields](#required-fields)
- [Domain Configuration](#domain-configuration)
- [Examples and Best Practices](#examples-and-best-practices)
- [Validation Rules](#validation-rules)

## Specification Format

All Claude agents must follow the standardized specification format with YAML frontmatter:

```yaml
---
name: agent-name-domain
description: Brief description of agent purpose and capabilities
domain: python|dotnet|nodejs|java|core
role: analyst|architect|engineer|test-engineer|documentation
spec_version: "1.0"
tools: Bash, Edit, MultiEdit, Write, Read, Glob, Grep, LS, WebFetch, WebSearch, NotebookEdit, TodoWrite, BashOutput, KillBash
model: inherit
color: blue|green|red|purple|orange|yellow
inputs:
  - List of expected inputs
  - GitHub issues with specific labels
  - Architectural plans or requirements
outputs:
  - List of expected outputs
  - Generated artifacts
  - Updated issue labels
validation:
  - Validation criteria
  - Quality requirements
  - Completeness checks
dependencies:
  - Required dependencies
  - Runtime environments
  - Tool requirements
workflow_position: 1-9
github_integration:
  triggers: ["label-names"]
  outputs: ["output-labels"]
  permissions: ["permission-list"]
examples:
  - context: "Example context description"
    input: "Example input description"
    output: "Expected output description"
---

# Agent Name

Agent description and detailed documentation...
```

## Required Fields

### Core Identity Fields

**name** (string, required)
- Unique identifier for the agent
- Format: `{role}-{domain}` or `{role}` for core agents
- Examples: `solution-architect-python`, `requirements-analyst`

**description** (string, required)
- Clear, concise description of agent purpose
- Should explain what the agent does and when to use it
- Length: 1-2 sentences recommended

**domain** (string, required)
- Technology domain classification
- Valid values: `python`, `dotnet`, `nodejs`, `java`, `core`
- Determines which technology stack the agent specializes in

**role** (string, required)
- Workflow role classification
- Valid values: `analyst`, `architect`, `engineer`, `test-engineer`, `documentation`
- Indicates the agent's position in the development workflow

### Technical Configuration

**spec_version** (string, required)
- Current specification version
- Format: `"X.Y"` (quoted string)
- Current version: `"1.0"`

**tools** (string, required)
- Comma-separated list of available tools
- Standard tools: `Bash, Edit, MultiEdit, Write, Read, Glob, Grep, LS, WebFetch, WebSearch, NotebookEdit, TodoWrite, BashOutput, KillBash`

**model** (string, required)
- AI model configuration
- Standard value: `inherit`

**color** (string, required)
- UI color coding for agent identification
- Valid values: `blue`, `green`, `red`, `purple`, `orange`, `yellow`
- Convention: architects=purple, engineers=blue, test-engineers=green

### Workflow Configuration

**workflow_position** (integer, required)
- Position in 9-step development workflow
- Range: 1-9
- Examples: 2=requirements analysis, 4=architecture, 6=implementation

**inputs** (array, required)
- List of expected input types
- Should be specific and actionable
- Examples: "GitHub issues with requirements-ready label", "Architectural plans"

**outputs** (array, required)
- List of expected output types
- Should describe deliverables
- Examples: "Implementation plan", "Updated GitHub issues", "Pull requests"

**validation** (array, required)
- Validation criteria and quality requirements
- Should be measurable and specific
- Examples: "Code coverage >80%", "All tests pass"

**dependencies** (array, required)
- Required dependencies and prerequisites
- Include runtime environments, tools, access requirements
- Examples: "Python 3.11+ runtime", "GitHub CLI access"

### GitHub Integration

**github_integration** (object, required)
- GitHub-specific configuration for automation

**triggers** (array, required)
- GitHub issue labels that trigger this agent
- Examples: `["requirements-ready"]`, `["plan-approved"]`

**outputs** (array, required)
- GitHub labels this agent adds to issues
- Examples: `["plan-approved"]`, `["implementation-complete"]`

**permissions** (array, required)
- Required GitHub permissions
- Examples: `["issues:write", "contents:write", "pull_requests:write"]`

### Documentation

**examples** (array, required)
- At least one example demonstrating agent usage
- Each example must have: `context`, `input`, `output`

## Domain Configuration

### Core Domain
- **Purpose**: Language-agnostic workflow agents
- **Agents**: requirements-analyst, solution-architect, documentation
- **Naming**: No domain suffix (e.g., `requirements-analyst`)

### Python Domain
- **Purpose**: Python ecosystem development
- **Frameworks**: FastAPI, Django, Flask, SQLAlchemy
- **Testing**: pytest, unittest, coverage
- **Quality**: black, ruff, mypy

### .NET Domain
- **Purpose**: .NET ecosystem development
- **Frameworks**: ASP.NET Core, Entity Framework Core
- **Testing**: xUnit, NUnit, Moq
- **Quality**: Static analysis, code formatting

### Node.js Domain
- **Purpose**: JavaScript/TypeScript development
- **Frameworks**: Express.js, Fastify, TypeScript
- **Testing**: Jest, Supertest, Cypress
- **Quality**: ESLint, Prettier

### Java Domain
- **Purpose**: Java ecosystem development
- **Frameworks**: Spring Boot, Spring Framework, JPA
- **Testing**: JUnit 5, Mockito, TestContainers
- **Build**: Maven, Gradle

## Examples and Best Practices

### Example 1: Solution Architect Agent

```yaml
---
name: solution-architect-python
description: Expert Python solution architect for designing scalable applications using hexagonal architecture
domain: python
role: architect
spec_version: "1.0"
tools: Bash, Glob, Grep, LS, Read, WebFetch, TodoWrite, WebSearch, BashOutput, KillBash
model: inherit
color: purple
inputs:
  - GitHub issues with requirements-ready label
  - Business requirements and acceptance criteria
  - Technical constraints and existing architecture
outputs:
  - Detailed implementation plan with hexagonal architecture
  - Work unit breakdown with dependencies
  - GitHub issue updates with plan-approved label
validation:
  - Requirements completeness validation
  - Technical feasibility assessment
  - Architecture pattern compliance
dependencies:
  - Python 3.11+ runtime environment
  - Understanding of hexagonal architecture
  - Knowledge of Python frameworks
workflow_position: 4
github_integration:
  triggers: ["requirements-ready"]
  outputs: ["plan-approved"]
  permissions: ["issues:write", "labels:write"]
examples:
  - context: User has complex API requirement needing architectural planning
    input: "Design scalable REST API for e-commerce with user management and order processing"
    output: "Create hexagonal architecture plan with FastAPI, SQLAlchemy, and domain separation"
---
```

### Example 2: Software Engineer Agent

```yaml
---
name: software-engineer-dotnet
description: Expert .NET software engineer implementing clean architecture solutions
domain: dotnet
role: engineer
spec_version: "1.0"
tools: Bash, Edit, MultiEdit, Write, Read, Glob, Grep, LS, WebFetch, WebSearch, NotebookEdit, TodoWrite, BashOutput, KillBash
model: inherit
color: blue
inputs:
  - GitHub issues with plan-approved label
  - Architectural plans with clean architecture
  - Test specifications from test engineers
outputs:
  - Implemented .NET solution following clean architecture
  - Comprehensive unit and integration tests
  - GitHub pull requests with implementation
validation:
  - Code quality and C# best practices compliance
  - Test coverage >80%
  - Clean architecture adherence
dependencies:
  - .NET 8+ SDK
  - Entity Framework Core tools
  - xUnit testing framework
workflow_position: 6
github_integration:
  triggers: ["plan-approved", "tests-planned"]
  outputs: ["implementation-complete", "ready-for-review"]
  permissions: ["contents:write", "pull_requests:write"]
examples:
  - context: Implementing approved architectural plan
    input: "Plan-approved issue for user authentication system implementation"
    output: "Complete .NET implementation with clean architecture, tests, and pull request"
---
```

## Validation Rules

### Format Validation
- YAML frontmatter must be valid
- All required fields must be present
- Field types must match specifications
- Enumerated values must be from valid sets

### Naming Conventions
- Agent names should follow `{role}-{domain}` pattern for domain-specific agents
- Core agents use role name only
- File names should match agent names: `{name}.md`

### Content Validation
- Examples must include context, input, and output
- Tools list should include commonly used tools
- Workflow position should match agent role
- Dependencies should be realistic and specific

### Cross-Agent Consistency
- No duplicate agent names
- Consistent color coding by role
- Proper workflow position assignments
- Domain coverage completeness

### GitHub Integration
- Trigger labels should align with workflow
- Output labels should indicate completion
- Permissions should match required actions
- Examples should demonstrate real usage

## Common Mistakes

1. **Missing Required Fields**: Ensure all required fields are present
2. **Invalid Domain Values**: Use only valid domain names
3. **Incorrect Workflow Position**: Match position to agent role
4. **Poor Examples**: Provide realistic, detailed examples
5. **Tool Mismatch**: Include tools actually used by agent
6. **GitHub Integration Errors**: Align triggers and outputs with workflow

## Validation Tools

Use the provided validation script to check agent specifications:

```bash
# Validate single agent
python3 scripts/validate-claude-agent.py agents/python/solution-architect-python.md

# Validate all agents
python3 scripts/validate-claude-agent.py

# Strict validation (warnings as errors)
python3 scripts/validate-claude-agent.py --strict
```

The validator checks:
- YAML syntax and structure
- Required field presence
- Value validation against enums
- Naming convention compliance
- Cross-agent consistency
- Example format requirements
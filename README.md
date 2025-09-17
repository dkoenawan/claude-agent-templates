# Claude Agent Templates

**Version 1.0.0** - A comprehensive collection of specialized Claude Code agent templates implementing spec-driven development with domain-specific agents, GitHub issue automation, and hexagonal architecture patterns.

## What This Is

Claude Agent Templates provides a structured 9-step GitHub issue-driven development workflow powered by specialized AI agents across multiple technology domains. Each agent handles a specific phase of software development using hexagonal/clean architecture principles, ensuring complete traceability, collaboration, and reliable delivery.

**Key Features:**
- **Domain-Specific Agents** - Specialized agents for Python, .NET, Node.js, Java, and core workflows
- **GitHub Issue Automation** - Complete CI/CD pipeline with automatic agent orchestration
- **Hexagonal Architecture** - Clean architecture patterns implemented across all domain agents
- **Spec-Driven Development** - Comprehensive specifications with Arc42 documentation standards
- **Test-Driven Development** - Complete test suites with >80% coverage requirements
- **Workflow State Management** - 9-step workflow tracking with progress monitoring

## Why Use This

- **Multi-Domain Support** - Specialized agents for Python, .NET, Node.js, Java with domain expertise
- **Enterprise Architecture** - Hexagonal/clean architecture patterns for maintainable codebases
- **Automated Workflows** - GitHub Actions automation for issue classification and agent orchestration
- **Quality Assurance** - Comprehensive validation, testing, and quality gates
- **Scalable Development** - Structured approach that scales from small teams to enterprise
- **Best Practices** - Implements modern software engineering patterns and methodologies

## Quick Start

### 1. Clone this repository
```bash
git clone https://github.com/dkoenawan/claude-agent-templates.git
cd claude-agent-templates
```

#### Install Task (if not already installed)
Task is a cross-platform task runner that provides unified automation. Install it from [taskfile.dev/installation](https://taskfile.dev/installation/) or use the included binary:

```bash
# The repository includes a Task binary in ./bin/task
# It will be used automatically if Task is not in your PATH
```

### 2. Install agents globally

#### Using Task (Recommended)
```bash
# Install all agents using the unified Taskfile
task install

# Or use other Task operations
task list       # Show available and installed agents
task validate   # Verify installation integrity
task clean      # Remove installed agents (with confirmation)
task update     # Update existing installations
task help       # Show detailed help
```

#### Legacy Scripts (Deprecated)
```bash
# Still supported but deprecated
./scripts/install-agents.sh    # Linux/macOS
./scripts/install-agents.bat   # Windows

# Manual installation (not recommended)
mkdir -p ~/.claude/agents
cp agents/core/*.md ~/.claude/agents/
cp agents/python/*.md ~/.claude/agents/
```

### 3. Use agents in Claude Code
```bash
# Claude will automatically suggest relevant agents
# Or invoke directly with the Task tool
```

## Repository Structure

```
claude-agent-templates/
├── agents/                      # Agent templates organized by domain
│   ├── core/                   # Language-agnostic workflow agents
│   │   ├── requirements-analyst.md
│   │   ├── solution-architect.md
│   │   └── documentation.md
│   ├── python/                 # Python-specific development agents
│   │   ├── solution-architect-python.md
│   │   ├── software-engineer-python.md
│   │   └── test-engineer-python.md
│   ├── dotnet/                 # .NET development agents
│   │   ├── solution-architect-dotnet.md
│   │   ├── software-engineer-dotnet.md
│   │   └── test-engineer-dotnet.md
│   ├── nodejs/                 # Node.js development agents
│   │   ├── solution-architect-nodejs.md
│   │   ├── software-engineer-nodejs.md
│   │   └── test-engineer-nodejs.md
│   └── java/                   # Java development agents
│       ├── solution-architect-java.md
│       ├── software-engineer-java.md
│       └── test-engineer-java.md
├── .github/workflows/          # GitHub Actions automation
│   ├── issue-agent-orchestration.yml
│   ├── execute-phase.yml
│   └── validate-agents.yml
├── scripts/                    # Validation and automation scripts
│   ├── validate-claude-agent.py
│   ├── classify-domain.py
│   ├── track-workflow.py
│   └── validate-agent-spec.sh
├── tests/                      # Comprehensive test suites
│   ├── contract/               # Contract tests for agent formats
│   ├── integration/            # Integration tests for workflows
│   ├── unit/                   # Unit tests for validation logic
│   └── performance/            # Performance benchmarks
├── specs/                      # Technical specifications
│   └── 001-refactor-the-agent/ # Agent refactoring specifications
└── docs/                       # Documentation and guides
```

## GitHub Issue-Driven Development Workflow

This repository implements a structured 9-step development workflow where all agents interact through GitHub issues:

1. **User** raises new bug or feature request via GitHub issue
2. **Requirements Analyst** reviews requirements, asks clarifying questions via issue comments
3. **User** provides answers and clarifications in issue
4. **Solution Architect** creates comprehensive implementation plan following best practices
5. **User** reviews and accepts the architectural plan
6. **Test Engineer Python** creates comprehensive unit test strategy with pytest
7. **Software Engineer Python** implements solution with integrated testing (branch management, commits, PR creation)
8. **User** accepts implementation via GitHub issue or PR review
9. **Documentation Agent** performs final documentation updates and repository cleanup

## Available Agents

### Core Workflow Agents (Language-Agnostic)
- **requirements-analyst** - **Step 2** - Analyzes GitHub issues, extracts business requirements with Arc42 standards
- **solution-architect** - **Step 4** - Creates comprehensive implementation plans with architectural decision records
- **documentation** - **Step 9** - Performs final documentation updates with Arc42 compliance

### Domain-Specific Development Agents

#### Python Ecosystem
- **solution-architect-python** - **Step 4** - FastAPI, Django, Flask architecture planning with hexagonal patterns
- **software-engineer-python** - **Step 6** - Implementation with pytest, black, ruff, and modern Python practices
- **test-engineer-python** - **Step 5** - Comprehensive testing strategies with >80% coverage requirements

#### .NET Ecosystem
- **solution-architect-dotnet** - **Step 4** - ASP.NET Core, Entity Framework, clean architecture planning
- **software-engineer-dotnet** - **Step 6** - Implementation with C#, Entity Framework Core, and xUnit testing
- **test-engineer-dotnet** - **Step 5** - .NET testing strategies with xUnit, Moq, and TestContainers

#### Node.js Ecosystem
- **solution-architect-nodejs** - **Step 4** - Express.js, TypeScript, hexagonal architecture planning
- **software-engineer-nodejs** - **Step 6** - Implementation with modern JavaScript/TypeScript and Jest testing
- **test-engineer-nodejs** - **Step 5** - Node.js testing strategies with Jest, Supertest, and async patterns

#### Java Ecosystem
- **solution-architect-java** - **Step 4** - Spring Boot, clean architecture planning with enterprise patterns
- **software-engineer-java** - **Step 6** - Implementation with Spring Framework, JPA, and modern Java features
- **test-engineer-java** - **Step 5** - Java testing strategies with JUnit 5, Mockito, and TestContainers

**Enterprise Features:**
- ✅ Domain-specific expertise with framework specialization
- ✅ Hexagonal/clean architecture implementation patterns
- ✅ Comprehensive validation and quality gates
- ✅ GitHub Actions automation for issue orchestration
- ✅ Workflow state tracking and progress monitoring

*See [agents/README.md](agents/README.md) for complete workflow documentation*

## Automation and Validation

### GitHub Actions Workflows
- **Issue Agent Orchestration** - Automatically classifies issues and assigns appropriate agents
- **Execution Phase** - Manages workflow execution across planning, implementation, testing, and documentation
- **Agent Validation** - Validates agent specifications and ensures consistency

### Validation Tools
```bash
# Validate all agent specifications
python3 scripts/validate-claude-agent.py

# Classify issue domain and recommend agent
python3 scripts/classify-domain.py --title "Create FastAPI authentication" --body "JWT-based auth system"

# Track workflow progress
python3 scripts/track-workflow.py --issue 123 --progress

# Generate workflow status report
python3 scripts/track-workflow.py --report
```

## Documentation

- [Agent Specifications](docs/agent-specifications.md) - Complete agent specification format
- [Domain Specialization](docs/domain-specialization.md) - Guide to domain-specific agents
- [GitHub Actions Setup](docs/github-actions-setup.md) - Setting up automation workflows
- [Troubleshooting Guide](docs/troubleshooting.md) - Common issues and solutions
- [Migration Guide](docs/migration-guide.md) - Upgrading from previous versions
- [Development Workflow](CLAUDE.md) - Trunk-based development process

## Contributing

We follow a trunk-based development approach with short-lived branches. See [CLAUDE.md](CLAUDE.md) for detailed guidelines.

1. Branch from `main`
2. Create focused changes (one agent per PR)
3. Test your agents in real projects
4. Submit pull request
5. Quick review and merge

## License

MIT License - Feel free to use, modify, and distribute these templates. See [LICENSE.md](LICENSE.md) for details.

---

Perfect for developers who want to maximize Claude Code's potential with battle-tested agent configurations.
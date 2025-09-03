# Claude Agent Templates

**Version 0.1.2** - A comprehensive collection of specialized Claude Code agent templates implementing GitHub issue-driven development workflows with enhanced reliability and cross-agent validation.

## What This Is

Claude Agent Templates provides a structured 9-step GitHub issue-driven development workflow powered by specialized AI agents. Each agent handles a specific phase of software development - from requirements analysis to final documentation - ensuring complete traceability, collaboration, and reliable delivery.

**Key Features:**
- **GitHub Issue Integration** - All agents interact through GitHub issues for complete workflow traceability
- **Comprehensive Agent Validation** - Enhanced reliability with mandatory codebase verification and cross-agent validation
- **End-to-End Accountability** - No incomplete deliverables, agents ensure full implementation
- **Error Recovery Protocols** - Self-correction capabilities and structured error handling

## Why Use This

- **Structured Workflow** - 9-step GitHub issue-driven process ensures nothing falls through cracks
- **Reliable Execution** - Enhanced agents with mandatory verification and validation protocols
- **Complete Traceability** - All agent interactions documented through GitHub issues and comments
- **Production-Ready** - Battle-tested agents with comprehensive error handling and self-correction
- **Team Collaboration** - Standardized workflow that scales across teams and projects
- **Continuous Improvement** - Agents evolve based on real-world performance analysis

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
│   └── python/                 # Python-specific development agents
│       ├── test-engineer-python.md
│       └── software-engineer-python.md
├── examples/
│   ├── project-specific/       # Project-type specific examples
│   └── workflows/             # Workflow combinations
├── scripts/                   # Installation and automation scripts
├── docs/                     # Detailed documentation
└── templates/               # Base template for new agents
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
- **requirements-analyst** - **Step 2** - Analyzes GitHub issues, extracts business requirements, asks clarifying questions
- **solution-architect** - **Step 4** - Creates comprehensive implementation plans with hexagonal architecture
- **documentation** - **Step 9** - Performs final documentation updates and repository cleanup

### Python Development Agents  
- **test-engineer-python** - **Step 6** - Creates comprehensive unit test strategies with pytest, focusing on 80% coverage
- **software-engineer-python** - **Step 7** - Implements solutions using hexagonal architecture with branch management and PR creation

**Enhanced Reliability Features:**
- ✅ Mandatory codebase verification before planning
- ✅ Cross-agent validation and state consistency
- ✅ Implementation accountability with end-to-end delivery
- ✅ Error recovery protocols and self-correction capabilities

*See [agents/README.md](agents/README.md) for complete workflow documentation*

## Documentation

- [Agent Writing Guide](docs/agent-guide.md) - How to write effective agents
- [Customization Guide](docs/customization.md) - How to customize templates
- [Contributing Guidelines](docs/contributing.md) - How to contribute
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
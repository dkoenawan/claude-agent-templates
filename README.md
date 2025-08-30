# Claude Agent Templates

**Version 0.1.1** - A curated collection of reusable Claude Code agent templates for software development workflows.

## What This Is

Specialized Claude Code agent templates that you can deploy to any project to get consistent, expert AI assistance for common development tasks like code review, debugging, testing, and documentation.

## Why Use This

- **Consistency** - Same high-quality approach across all your projects
- **Efficiency** - No need to re-explain context or preferences to Claude
- **Specialization** - Purpose-built agents for specific development tasks
- **Continuous Improvement** - Templates evolve based on real-world usage
- **Team Collaboration** - Share proven workflows across your team

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

## Available Agents

### Core Workflow Agents (Language-Agnostic)
- **requirements-analyst** - Translates business requirements into technical specifications
- **solution-architect** - Breaks down complex features into implementable work units
- **documentation** - Performs final documentation updates and repository cleanup

### Python Development Agents
- **test-engineer-python** - Creates comprehensive unit test strategies with pytest
- **software-engineer-python** - Implements solutions using hexagonal architecture principles

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
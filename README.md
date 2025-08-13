# Claude Agent Templates

A curated collection of reusable Claude Code agent templates for software development workflows.

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

### 2. Deploy agents to your project
```bash
# Copy specific agents to your project
cp agents/code-reviewer.md /path/to/your-project/.claude/agents/

# Or install globally for all projects
cp agents/*.md ~/.claude/agents/
```

### 3. Use agents in Claude Code
```bash
# Claude will automatically suggest relevant agents
# Or invoke directly with the Task tool
```

## Repository Structure

```
claude-agent-templates/
├── agents/             # Core agent templates
├── examples/
│   ├── project-specific/   # Project-type specific examples
│   └── workflows/          # Workflow combinations
├── scripts/            # Automation scripts
├── docs/              # Detailed documentation
└── templates/         # Base template for new agents
```

## Available Agents

*Coming soon - agents will be added incrementally*

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

MIT License - Feel free to use, modify, and distribute these templates.

---

Perfect for developers who want to maximize Claude Code's potential with battle-tested agent configurations.
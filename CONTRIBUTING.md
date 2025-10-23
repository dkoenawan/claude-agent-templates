# Contributing to Claude Agent Templates

Thank you for your interest in contributing to Claude Agent Templates! This document provides guidelines for contributing to the project, whether you're fixing bugs, adding features, or improving documentation.

## Table of Contents

- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Development Workflow](#development-workflow)
- [Agent Development](#agent-development)
- [Testing](#testing)
- [Documentation](#documentation)
- [Pull Request Process](#pull-request-process)
- [Code Style](#code-style)
- [Release Process](#release-process)

---

## Getting Started

### Prerequisites

- **Go 1.21+** - For building the `spec-kit-agents` CLI
- **Git** - For version control
- **Claude Code** - For testing agents
- **Python 3.8+** - For validation scripts (optional)
- **Make or Task** - For automation (optional)

### Fork and Clone

```bash
# Fork the repository on GitHub, then:
git clone https://github.com/YOUR_USERNAME/claude-agent-templates.git
cd claude-agent-templates

# Add upstream remote
git remote add upstream https://github.com/dkoenawan/claude-agent-templates.git
```

---

## Development Setup

### 1. Build the CLI Tool

```bash
# Build the spec-kit-agents binary
go build -o bin/spec-kit-agents ./cmd/spec-kit-agents/

# Verify build
./bin/spec-kit-agents version
```

### 2. Install in Development Mode

```bash
# Install locally for testing
./bin/spec-kit-agents install --prefix .

# This creates .claude-agent-templates/ in the current directory
```

### 3. Run Tests

```bash
# Run all unit tests
go test ./... -v

# Run with coverage
go test ./... -cover

# Run validation script
./validate.sh
```

---

## Development Workflow

We follow **trunk-based development** with short-lived feature branches:

### 1. Create Feature Branch

```bash
# Update main
git checkout main
git pull upstream main

# Create feature branch
git checkout -b feature/your-feature-name
```

### 2. Make Changes

- Keep changes focused and atomic
- Write tests first (TDD approach)
- Update documentation as you go
- Follow existing code patterns

### 3. Test Thoroughly

```bash
# Run tests
go test ./... -v

# Validate agents (if modified)
python3 scripts/validate-claude-agent.py

# Test installation
./bin/spec-kit-agents install --dry-run
```

### 4. Commit

```bash
# Stage changes
git add .

# Commit with conventional commit message
git commit -m "feat: add new capability"
```

**Conventional Commit Format:**
- `feat:` - New feature
- `fix:` - Bug fix
- `docs:` - Documentation only
- `test:` - Adding or updating tests
- `refactor:` - Code refactoring
- `chore:` - Maintenance tasks

### 5. Push and Create PR

```bash
# Push to your fork
git push origin feature/your-feature-name

# Create pull request on GitHub
gh pr create --title "feat: Add new capability" --body "Description of changes"
```

---

## Agent Development

### Creating a New Agent

Agents are written in Markdown with YAML frontmatter. Follow the specification format:

```markdown
---
name: agent-name-domain
description: Brief description of what the agent does
domain: python|dotnet|nodejs|java|core
role: analyst|architect|engineer|test-engineer|documentation
spec_version: "1.0"
tools: Bash, Edit, Write, Read, Glob, Grep
model: inherit
color: blue
inputs:
  - Input description
outputs:
  - Output description
validation:
  - Validation criteria
workflow_position: 1-9
examples:
  - context: "Example scenario"
    input: "Example input"
    output: "Expected output"
---

# Agent Name

## Purpose

Clear description of the agent's purpose and responsibilities.

## Capabilities

- Capability 1
- Capability 2

## Instructions

Detailed instructions for how the agent should behave...
```

### Agent Best Practices

1. **Clear Purpose** - Agent should have a single, well-defined responsibility
2. **Comprehensive Examples** - Include at least 2-3 examples with different contexts
3. **Validation Criteria** - Define how to verify the agent's output
4. **Domain Expertise** - Include framework-specific patterns and best practices
5. **Error Handling** - Document how agent handles edge cases

### Testing Agents

```bash
# Validate agent format
python3 scripts/validate-claude-agent.py agents/your-domain/your-agent.md

# Test in Claude Code
# 1. Copy agent to .claude/agents/
# 2. Open Claude Code
# 3. Test with @your-agent
```

---

## Testing

### Unit Tests

Located in `*_test.go` files alongside code:

```go
// internal/version/compare_test.go
func TestCompareVersions(t *testing.T) {
    tests := []struct {
        name string
        v1   string
        v2   string
        want int
    }{
        {"v1 greater", "1.1.0", "1.0.0", 1},
        {"v1 less", "1.0.0", "1.1.0", -1},
        {"equal", "1.0.0", "1.0.0", 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, _ := CompareVersions(tt.v1, tt.v2)
            if got != tt.want {
                t.Errorf("got %d, want %d", got, tt.want)
            }
        })
    }
}
```

### Integration Tests

Integration tests validate end-to-end workflows. Example structure:

```bash
tests/integration/
├── test_install.sh       # Test installation workflow
├── test_upgrade.sh       # Test upgrade workflow
└── test_rollback.sh      # Test rollback workflow
```

### Test Coverage Goals

- **Unit tests**: >80% coverage
- **Integration tests**: All user stories covered
- **Agent tests**: All agents validated

### Running Tests

```bash
# All tests
go test ./... -v

# Specific package
go test ./internal/version -v

# With coverage
go test ./... -cover

# Coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

## Documentation

### Documentation Standards

1. **Code Comments** - Document all exported functions
2. **README Updates** - Update README for user-facing changes
3. **CLAUDE.md** - Update agent usage guidelines
4. **Inline Help** - Keep CLI help text up to date
5. **Examples** - Include working examples

### Documentation Files

- `README.md` - User-facing project overview
- `CLAUDE.md` - Development guide and agent context
- `CONTRIBUTING.md` - This file
- `specs/*/` - Feature specifications
- `docs/` - Detailed documentation

### Writing Good Documentation

```markdown
# Feature Name

## Overview
Brief 1-2 sentence description

## Usage
```bash
# Example command
spec-kit-agents command --flag value
```

## Options
- `--flag` - Description of flag

## Examples
Practical, working examples that users can copy-paste
```

---

## Pull Request Process

### Before Submitting PR

- [ ] All tests pass (`go test ./... -v`)
- [ ] Code follows project style
- [ ] Documentation updated
- [ ] Commit messages follow conventional format
- [ ] PR description explains what/why

### PR Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests pass
- [ ] Manual testing completed

## Checklist
- [ ] Code follows style guidelines
- [ ] Documentation updated
- [ ] Tests pass
- [ ] No breaking changes (or documented)
```

### Review Process

1. **Automated Checks** - CI runs tests and validation
2. **Code Review** - Maintainer reviews code
3. **Testing** - Changes tested in real scenarios
4. **Approval** - At least one maintainer approval required
5. **Merge** - Squash and merge to main

---

## Code Style

### Go Code Style

Follow standard Go conventions:

```go
// Good: Clear function names, proper error handling
func LoadManifest(path string) (*Manifest, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("failed to read manifest: %w", err)
    }

    var manifest Manifest
    if err := json.Unmarshal(data, &manifest); err != nil {
        return nil, fmt.Errorf("failed to parse manifest: %w", err)
    }

    return &manifest, nil
}

// Use structured error messages
return fmt.Errorf("operation failed: %w", err)

// Document exported functions
// LoadManifest loads and parses a version manifest from the specified path.
// Returns an error if the file cannot be read or parsed.
func LoadManifest(path string) (*Manifest, error) {
    // ...
}
```

### Agent Style

```markdown
# Use clear, action-oriented language
✅ "Analyze requirements and create specification"
❌ "You should analyze the requirements"

# Be specific about tools and commands
✅ "Run `pytest tests/` with coverage flags"
❌ "Run tests"

# Include validation criteria
✅ "Verify all tests pass and coverage >80%"
❌ "Make sure it works"
```

### Commit Message Style

```bash
# Good commit messages
feat: add version compatibility checking
fix: resolve path handling on Windows
docs: update installation instructions
test: add unit tests for manifest loading

# Bad commit messages
update stuff
fix bug
changes
wip
```

---

## Release Process

### Version Numbering

We use **semantic versioning**:
- `MAJOR.MINOR.PATCH`
- `MAJOR` - Breaking changes
- `MINOR` - New features (backward compatible)
- `PATCH` - Bug fixes

### Creating a Release

1. **Update Version**
   ```bash
   # Update version in cmd/spec-kit-agents/main.go
   const Version = "1.1.0"
   ```

2. **Update CHANGELOG**
   ```bash
   # Add release notes to CHANGELOG.md
   ```

3. **Create Tag**
   ```bash
   git tag -a v1.1.0 -m "Release v1.1.0"
   git push upstream v1.1.0
   ```

4. **GitHub Actions** - Automatically builds binaries and creates GitHub release

5. **Test Release**
   ```bash
   # Test one-liner installer
   curl -fsSL https://raw.githubusercontent.com/dkoenawan/claude-agent-templates/main/scripts/install.sh | bash
   ```

---

## Project Structure

Understanding the project structure helps navigate and contribute:

```
claude-agent-templates/
├── agents/                         # Agent markdown files
│   ├── core/                      # Language-agnostic agents
│   ├── python/                    # Python-specific agents
│   ├── dotnet/                    # .NET agents
│   ├── nodejs/                    # Node.js agents
│   └── java/                      # Java agents
├── cmd/spec-kit-agents/           # CLI entry point
│   └── main.go
├── internal/                      # Internal packages
│   ├── config/                    # Configuration (paths, logging)
│   ├── install/                   # Installation logic
│   └── version/                   # Version management
├── pkg/models/                    # Public data models
│   ├── manifest.go               # Version manifest
│   └── lock.go                   # Version lock
├── scripts/                       # Automation scripts
│   ├── install.sh                # One-liner installer
│   └── validate-claude-agent.py  # Agent validation
├── specs/                         # Feature specifications
│   └── ###-feature-name/
│       ├── spec.md               # What/why
│       ├── plan.md               # How (technical)
│       └── tasks.md              # Implementation tasks
├── tests/                         # Test suites
│   ├── contract/                 # Agent format tests
│   ├── integration/              # Workflow tests
│   └── unit/                     # Unit tests
└── .specify/                      # Vendored spec-kit
    ├── templates/
    ├── scripts/
    └── version-manifest.json
```

---

## Common Contribution Scenarios

### Fixing a Bug

1. Create issue describing bug
2. Create branch: `fix/issue-123-description`
3. Write failing test
4. Fix bug
5. Verify test passes
6. Submit PR referencing issue

### Adding a Feature

1. Discuss in issue first (for large features)
2. Create spec in `specs/###-feature-name/`
3. Create branch: `feature/###-feature-name`
4. Follow spec-driven development:
   - Write spec.md (what/why)
   - Write plan.md (how)
   - Write tasks.md (implementation steps)
   - Implement with tests
5. Submit PR with spec reference

### Improving Documentation

1. Create branch: `docs/improve-xyz`
2. Make documentation changes
3. Preview locally (for markdown)
4. Submit PR with clear description

### Adding a New Agent

1. Create branch: `feature/agent-name`
2. Create agent file in appropriate domain
3. Follow agent specification format
4. Add comprehensive examples
5. Validate with `validate-claude-agent.py`
6. Test in Claude Code
7. Submit PR with example usage

---

## Getting Help

### Resources

- **CLAUDE.md** - Development workflow and agent context
- **README.md** - Project overview and quick start
- **specs/** - Feature specifications and technical docs
- **GitHub Discussions** - Ask questions, share ideas
- **GitHub Issues** - Report bugs, request features

### Communication Channels

- **GitHub Issues** - Bug reports, feature requests
- **GitHub Discussions** - Questions, ideas, feedback
- **Pull Requests** - Code reviews, technical discussions

### Code of Conduct

- Be respectful and inclusive
- Provide constructive feedback
- Focus on the code, not the person
- Help others learn and improve
- Assume positive intent

---

## Spec-Driven Development Process

This project follows GitHub's spec-driven development methodology:

### 1. Create Specification

```markdown
# specs/004-new-feature/spec.md

## User Scenarios
- User wants to accomplish X
- Given Y, when Z, then outcome

## Requirements
- FR-001: System must do X
- FR-002: System must support Y

## Success Criteria
- SC-001: Measurable outcome
```

### 2. Create Plan

```markdown
# specs/004-new-feature/plan.md

## Technical Approach
- Use technology X for reason Y
- Implement pattern Z

## Implementation Steps
1. Create data model
2. Implement core logic
3. Add CLI commands
```

### 3. Create Tasks

```markdown
# specs/004-new-feature/tasks.md

- [ ] T001: Create data model in pkg/models/
- [ ] T002: Implement core logic in internal/
- [ ] T003: Add CLI command
```

### 4. Implement

Follow tasks, write tests first (TDD), implement features.

### 5. Review & Merge

PR review, testing, merge to main.

---

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

## Thank You!

Thank you for contributing to Claude Agent Templates! Your contributions help make AI-assisted development more accessible and reliable for everyone.

**Questions?** Open a GitHub Discussion or Issue.

**Ready to contribute?** Check out [Good First Issues](https://github.com/dkoenawan/claude-agent-templates/labels/good%20first%20issue).

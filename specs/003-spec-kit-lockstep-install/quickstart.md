# Quickstart: Spec-Kit Lockstep Installation

**Time to Complete**: 5 minutes
**Goal**: Install claude-agent-templates with pinned spec-kit version in a single command
**Audience**: Developers using Claude Code

---

## One-Command Installation

The simplest way to install claude-agent-templates with lockstep spec-kit version management:

```bash
# Option 1: One-liner (after first release)
curl -fsSL https://raw.githubusercontent.com/dkoenawan/claude-agent-templates/main/scripts/install.sh | bash
spec-kit-agents install

# Option 2: Build from source (current)
git clone https://github.com/dkoenawan/claude-agent-templates.git
cd claude-agent-templates
go build -o bin/spec-kit-agents ./cmd/spec-kit-agents/
sudo cp bin/spec-kit-agents /usr/local/bin/  # OR: cp bin/spec-kit-agents ~/.local/bin/
spec-kit-agents install
```

That's it! Everything is automatically configured.

---

## What Just Happened?

When you run `spec-kit-agents install`, the tool:

1. ✅ **Detects installation mode** - Checks if you have existing `.specify/` directory
2. ✅ **Chooses installation prefix** - Uses `spec-kit-agents/` to avoid conflicts
3. ✅ **Copies spec-kit files** - Installs pinned spec-kit version (v0.0.72)
4. ✅ **Sets up Claude integration** - Creates `.claude/agents/` and `.claude/commands/`
5. ✅ **Creates version lock** - Records installation for version tracking
6. ✅ **Verifies installation** - Ensures everything is properly configured

**Result**: Ready to use with Claude Code immediately.

---

## Installation Modes

The installer automatically detects and chooses the best mode:

### Mode 1: Fresh Installation
**Scenario**: No existing `.specify/` directory
**Action**: Installs to `spec-kit-agents/` (clean, isolated)

### Mode 2: Coexist with Existing Spec-Kit
**Scenario**: You already have `.specify/` directory
**Action**: Installs to `spec-kit-agents/` (both versions coexist)

### Mode 3: Global Installation
**Scenario**: You want agents available across all projects
**Action**: Run with `--global` flag

```bash
spec-kit-agents install --global
```

This installs agents to `~/.claude/agents/` for system-wide access.

---

## Installation Options

### Basic Installation
```bash
# Auto-detect mode (recommended)
spec-kit-agents install

# See what would happen (dry run)
spec-kit-agents install --dry-run

# Verbose output
spec-kit-agents install --verbose
```

### Custom Installation
```bash
# Install globally
spec-kit-agents install --global

# Custom installation prefix
spec-kit-agents install --prefix /path/to/dir

# Force reinstall (overwrite existing)
spec-kit-agents install --force

# Quiet mode (errors only)
spec-kit-agents install --quiet
```

---

## Verify Installation

### Check Installation Status
```bash
spec-kit-agents status
```

**Expected Output**:
```
Installation Details:
  Prefix: spec-kit-agents
  Mode: standalone

Installed Components:
  claude-agent-templates: v1.0.0
  spec-kit: v0.0.72

Installation History:
  [2025-10-23] Initial installation
    - claude-agent-templates v1.0.0
    - spec-kit v0.0.72
```

### Check Version Compatibility
```bash
spec-kit-agents check
```

**Expected Output**:
```
✅ All components compatible
  claude-agent-templates: v1.0.0
  spec-kit: v0.0.72 (pinned)

Compatibility: OK
  Min version: 0.0.70
  Max version: 0.1.0
  Breaking versions: none
```

### Check CLI Version
```bash
spec-kit-agents version
```

---

## Using with Claude Code

After installation, all agents and commands are immediately available:

### Available Slash Commands
```bash
# Create a feature specification
/speckit.specify "Add user authentication feature"

# Generate implementation plan
/speckit.plan

# Break down into tasks
/speckit.tasks

# Validate specification quality
/speckit.analyze

# Interactive clarification
/speckit.clarify
```

### Available Agents
```bash
# In Claude Code, reference agents with @
@requirements-analyst help me understand feature requirements
@solution-architect-python design the authentication system
@software-engineer-python implement the auth endpoints
@test-engineer-python create test strategy
@documentation update the API docs
```

### Project Structure After Installation
```
your-project/
├── spec-kit-agents/        # Installation directory
│   ├── .specify/                   # Pinned spec-kit (v0.0.72)
│   ├── agents/                     # Agent specifications
│   └── .version-lock.json          # Version tracking
├── .claude/                         # Claude Code integration
│   ├── agents/                     # Symlinked agents
│   └── commands/                   # Symlinked slash commands
└── specs/                          # Your feature specs (created by commands)
    └── 001-your-feature/
        ├── spec.md
        ├── plan.md
        └── tasks.md
```

---

## Upgrading

When a new claude-agent-templates version is released:

```bash
# Check for updates
spec-kit-agents check

# Update to latest version
spec-kit-agents update

# Update to specific version
spec-kit-agents update --version v1.2.0

# Keep backup during update (default)
spec-kit-agents update --backup
```

The update process:
1. Creates backup of current installation
2. Downloads new version
3. Updates spec-kit to new pinned version
4. Verifies compatibility
5. Updates version lock

If anything fails, automatic rollback restores previous installation.

---

## Version Management

### Check Current Versions
```bash
spec-kit-agents status
```

### Verify Compatibility
```bash
spec-kit-agents check

# Get detailed compatibility report
spec-kit-agents check --detailed

# Output as JSON (for scripting)
spec-kit-agents check --json
```

### View Installation History
```bash
spec-kit-agents status
```

Shows chronological history of all installations and updates.

---

## Troubleshooting

### Installation Failed
```bash
# Check what went wrong
spec-kit-agents install --verbose

# Force clean reinstall
spec-kit-agents install --force

# Check logs
cat ~/spec-kit-agents/.install-log.txt
```

### Commands Not Found
```bash
# Verify installation
spec-kit-agents status

# Check Claude directory
ls -la .claude/commands/
ls -la .claude/agents/

# Reinstall if needed
spec-kit-agents install --force
```

### Version Conflict
```bash
# Check compatibility
spec-kit-agents check

# See what's wrong
spec-kit-agents check --detailed

# Fix automatically
spec-kit-agents check --fix
```

### Rollback After Failed Update
```bash
# Automatic rollback happens on failure
# Manual rollback if needed:
spec-kit-agents rollback

# Rollback to specific installation
spec-kit-agents rollback --installation-id <uuid>
```

---

## Uninstallation

To remove claude-agent-templates:

```bash
# Remove installation directory
rm -rf spec-kit-agents

# Remove Claude integration (optional)
rm -rf .claude/agents
rm -rf .claude/commands

# Remove global installation (if used)
rm -rf ~/.claude/agents/
sudo rm /usr/local/bin/spec-kit-agents
```

---

## Advanced Usage

### Multiple Projects with Different Versions

Each project maintains its own version lock:

```bash
# Project 1
cd ~/project-1
spec-kit-agents install
# Uses spec-kit v0.0.72

# Project 2 (different version)
cd ~/project-2
spec-kit-agents install
# Also uses spec-kit v0.0.72 (pinned in manifest)
```

Both projects get the same tested combination.

### CI/CD Integration

```yaml
# .github/workflows/spec-validation.yml
name: Validate Specs

on: [push, pull_request]

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Install spec-kit-agents
        run: |
          curl -fsSL https://raw.githubusercontent.com/dkoenawan/claude-agent-templates/main/scripts/install.sh | bash
          spec-kit-agents install

      - name: Validate specs
        run: |
          spec-kit-agents check
          # Run your spec validation
```

### Custom Spec-Kit Version (Advanced)

**⚠️ Not Recommended**: This breaks lockstep version management.

If you absolutely need a different spec-kit version:

1. Fork claude-agent-templates
2. Update `.specify/version-manifest.json`
3. Test compatibility thoroughly
4. Use your fork

---

## Next Steps

### 1. Create Your First Spec
```bash
/speckit.specify "Add user authentication with email and password"
```

This creates `specs/001-user-auth/spec.md` with structured requirements.

### 2. Generate Implementation Plan
```bash
/speckit.plan
```

This creates `specs/001-user-auth/plan.md` with technical design.

### 3. Break Down into Tasks
```bash
/speckit.tasks
```

This creates `specs/001-user-auth/tasks.md` with actionable tasks.

### 4. Use Agents for Implementation
```bash
@software-engineer-python implement the authentication endpoints according to spec
@test-engineer-python create comprehensive test suite
```

---

## Getting Help

### CLI Help
```bash
# General help
spec-kit-agents --help

# Command-specific help
spec-kit-agents install --help
spec-kit-agents check --help
spec-kit-agents update --help
```

### Documentation
- **User Guide**: See main `README.md`
- **Development**: See `CONTRIBUTING.md`
- **Agent Usage**: See `CLAUDE.md`

### Support
- **Issues**: https://github.com/dkoenawan/claude-agent-templates/issues
- **Discussions**: https://github.com/dkoenawan/claude-agent-templates/discussions

---

## Summary

You've installed claude-agent-templates with lockstep spec-kit version management:

✅ **Single command installation** - No manual configuration
✅ **Version pinning** - Spec-kit v0.0.72 guaranteed
✅ **Automatic detection** - Chooses best installation mode
✅ **Claude integration** - Agents and commands ready to use
✅ **Upgrade safety** - Automatic backup and rollback
✅ **No conflicts** - Coexists with your existing spec-kit

**Time invested**: 5 minutes
**Value delivered**: Production-ready AI-assisted development workflow

Happy coding! 🚀

---

## Comparison: Before vs After

### Before (Manual Installation)
```bash
git clone https://github.com/dkoenawan/claude-agent-templates.git
cd claude-agent-templates
mkdir -p ~/.claude/agents
cp -r agents/**/*.md ~/.claude/agents/
cd /path/to/project
cp -r /path/to/claude-agent-templates/.specify .specify
mkdir -p .claude/commands
cp .specify/templates/commands/*.md .claude/commands/
# ... 10+ more manual steps
# ... hope you got the version right
# ... hope you didn't break anything
```

### After (Lockstep Installation)
```bash
spec-kit-agents install
```

One command. Zero errors. Guaranteed compatibility.

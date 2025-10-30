# Quickstart Guide: Global Agent Installation

**Feature**: Global Agent Installation
**Estimated Time**: 5-10 minutes
**Audience**: Developers using claude-agent-templates

## Overview

This guide walks you through installing claude-agent-templates globally, allowing you to access agents and commands from any repository without cloning source code or repeating installation steps.

**What You'll Learn**:
- How to install agents globally using the `--global` flag
- How to verify your installation
- How to use agents across multiple repositories
- How to update your global installation
- How to troubleshoot common issues

**Prerequisites**:
- Claude Code installed and configured
- `spec-kit-agents` binary downloaded (see Installation Options below)
- Basic command-line familiarity

---

## Quick Start (TL;DR)

```bash
# Download and install the binary (one-liner)
curl -fsSL https://raw.githubusercontent.com/dkoenawan/claude-agent-templates/main/scripts/install.sh | bash

# Install agents globally
spec-kit-agents install --global

# Verify installation
spec-kit-agents status

# Start using in any repository
cd /path/to/your/project
# Agents are now available in Claude Code!
```

---

## Step 1: Get the Binary

### Option A: One-Liner Installer (Recommended)

Download and install in one command:

```bash
curl -fsSL https://raw.githubusercontent.com/dkoenawan/claude-agent-templates/main/scripts/install.sh | bash
```

This script:
- Detects your platform (Linux/macOS/Windows)
- Downloads the appropriate binary from GitHub Releases
- Installs to `~/.local/bin/spec-kit-agents`
- Verifies the installation

**Windows (PowerShell)**:
```powershell
iwr -useb https://raw.githubusercontent.com/dkoenawan/claude-agent-templates/main/scripts/install.ps1 | iex
```

### Option B: Manual Download

1. Go to [GitHub Releases](https://github.com/dkoenawan/claude-agent-templates/releases/latest)
2. Download the binary for your platform:
   - Linux: `spec-kit-agents-{version}-linux-amd64`
   - macOS (Intel): `spec-kit-agents-{version}-darwin-amd64`
   - macOS (Apple Silicon): `spec-kit-agents-{version}-darwin-arm64`
   - Windows: `spec-kit-agents-{version}-windows-amd64.exe`

3. Make it executable and move to PATH:

```bash
# Linux/macOS
chmod +x spec-kit-agents-*
sudo mv spec-kit-agents-* /usr/local/bin/spec-kit-agents

# Or user-only install
mkdir -p ~/.local/bin
mv spec-kit-agents-* ~/.local/bin/spec-kit-agents
```

4. Verify it's in your PATH:

```bash
spec-kit-agents version
```

### Option C: Build from Source

```bash
git clone https://github.com/dkoenawan/claude-agent-templates.git
cd claude-agent-templates
go build -o bin/spec-kit-agents ./cmd/spec-kit-agents/
sudo cp bin/spec-kit-agents /usr/local/bin/
```

---

## Step 2: Install Agents Globally

Run the global installation command:

```bash
spec-kit-agents install --global
```

**What happens**:
1. Binary extracts embedded agent files and templates
2. Agents are copied to `~/.claude/agents/` with `cat-` prefix
3. Commands are copied to `~/.claude/commands/` with `speckit.` prefix
4. Version lock is created at `~/.claude/.version-lock.json`
5. Installation is verified

**Expected output**:
```
Starting spec-kit-lockstep installation...
Source files verified
Installation mode: fresh installation
Installing spec-kit-agents v2.1.0 with spec-kit v0.0.72
Copying spec-kit files...
Setting up Claude Code integration...
Integrated with Claude Code: 50 agents, 8 commands
Version lock created
Installation verified

Installation complete!
  Location: /home/user/.claude
  spec-kit-agents: v2.1.0
  spec-kit: v0.0.72
  Files installed: 60
  Agents available: 50 (prefix: cat-)
  Commands available: 8 (prefix: speckit.)

Claude Code is now configured!
  Agents: ~/.claude/agents/cat-*.md
  Commands: Use /speckit.specify, /speckit.plan, /speckit.tasks
```

**Installation time**: Typically 5-30 seconds

---

## Step 3: Verify Installation

Check that agents are installed correctly:

```bash
spec-kit-agents status
```

**Expected output**:
```
Installation Details:
  Location: /home/user/.claude
  Type: global
  Mode: standalone

Installed Components:
  claude-agent-templates: v2.1.0
  spec-kit: v0.0.72

Installation History:
  [2025-10-30T14:00:00Z] Initial installation
    - claude-agent-templates v2.1.0
    - spec-kit v0.0.72

Status: ✓ Installation healthy
```

**Verify agents are accessible**:

```bash
# List installed agents
ls ~/.claude/agents/cat-*

# View an agent
cat ~/.claude/agents/cat-requirements-analyst.md
```

**Verify commands are available**:

```bash
# List installed commands
ls ~/.claude/commands/speckit.*
```

---

## Step 4: Use Agents in Any Repository

Navigate to any project and start using agents:

```bash
cd ~/projects/my-app
# You can now use slash commands in Claude Code:
# /speckit.specify - Create a feature specification
# /speckit.plan - Generate implementation plan
# /speckit.tasks - Break down into tasks
```

**In Claude Code**:
1. Open any repository in Claude Code
2. Type `/` to see available commands
3. You should see `speckit.specify`, `speckit.plan`, `speckit.tasks`, etc.
4. Agents like `cat-requirements-analyst`, `cat-solution-architect-python` are available

**Example workflow**:
```
User: /speckit.specify Add user authentication with email/password

Claude: [Creates specification using cat-requirements-analyst agent]

User: /speckit.plan

Claude: [Generates implementation plan using cat-solution-architect-python agent]

User: /speckit.tasks

Claude: [Breaks down into actionable tasks]
```

---

## Step 5: Update Your Installation

Check for updates periodically:

```bash
# Check if updates are available
spec-kit-agents check
```

**If update available**:
```bash
# Update to latest version
spec-kit-agents update
```

**What happens during update**:
1. Current installation is backed up automatically
2. New version is downloaded from GitHub Releases
3. Agents and commands are updated
4. Version lock is updated with history
5. Old backups are cleaned up (keeps last 3)

**If something goes wrong**:
- Automatic rollback restores previous version
- You can manually rollback with: `spec-kit-agents rollback`

**Update time**: Typically 1-2 minutes

---

## Advanced Usage

### Installation Options

```bash
# Dry run (see what would be done)
spec-kit-agents install --global --dry-run

# Force reinstall (overwrite existing)
spec-kit-agents install --global --force

# Quiet mode (minimal output)
spec-kit-agents install --global --quiet

# Verbose mode (detailed logging)
spec-kit-agents install --global --verbose
```

### Checking Versions

```bash
# Show CLI tool version
spec-kit-agents version

# Show installation status with all details
spec-kit-agents status --detailed

# Check version compatibility
spec-kit-agents check
```

### Offline Installation

If you have limited internet or work in air-gapped environments:

1. Download the binary (includes embedded files)
2. Run `spec-kit-agents install --global` without internet
3. Installation works offline using embedded files

**Note**: Updates require internet or a pre-downloaded package.

---

## Troubleshooting

### Issue: Binary not found

**Error**: `spec-kit-agents: command not found`

**Solution**:
```bash
# Check if binary is in PATH
which spec-kit-agents

# If not found, add to PATH
export PATH="$PATH:$HOME/.local/bin"

# Make permanent (add to ~/.bashrc or ~/.zshrc)
echo 'export PATH="$PATH:$HOME/.local/bin"' >> ~/.bashrc
```

### Issue: Permission denied

**Error**: `permission denied: ~/.claude/`

**Solution**:
```bash
# Ensure ~/.claude/ is writable
chmod 755 ~/.claude
```

### Issue: Installation fails with "source files not found"

**Error**: `source verification failed: required directory not found: .specify/`

**Cause**: Binary doesn't have embedded files (built incorrectly)

**Solution**:
- Download the official release binary (has embedded files)
- If building from source, ensure you build with: `go build -tags embed`

### Issue: Agents not appearing in Claude Code

**Problem**: Commands work but agents don't show up

**Solution**:
1. Verify agents are installed:
   ```bash
   ls ~/.claude/agents/cat-*
   ```

2. Restart Claude Code

3. Check Claude Code configuration:
   - Ensure Claude Code is looking in `~/.claude/agents/`

### Issue: Update fails

**Error**: Update downloads but installation fails

**Solution**:
- Automatic rollback should restore previous version
- Check status: `spec-kit-agents status`
- If needed, manual rollback: `spec-kit-agents rollback`

---

## Comparison: Global vs Repository-Local

| Feature | Global Installation | Repository-Local |
|---------|---------------------|------------------|
| **Installation** | Once for all repos | Per repository |
| **Access** | Any repository | Only that repository |
| **Updates** | Update once, affects all | Update per repository |
| **Source** | Embedded in binary | From cloned repository |
| **Use Case** | General development | Custom modifications |
| **Setup Time** | < 1 minute | ~2 minutes per repo |

**When to use Global**:
- You work across multiple repositories
- You want consistent agents everywhere
- You don't need custom agent modifications
- You want easy updates

**When to use Repository-Local**:
- You need custom agent modifications for a specific project
- You want project-specific versions
- You're contributing to claude-agent-templates

**Can I use both?**:
Yes! You can have a global installation and repository-local installations. Repository-local takes precedence when you're in that repository.

---

## Next Steps

Now that you have global agents installed:

1. **Try the workflow**:
   ```bash
   cd your-project
   # In Claude Code:
   # /speckit.specify - Start with a feature description
   # /speckit.plan - Generate technical implementation plan
   # /speckit.tasks - Get actionable task breakdown
   ```

2. **Explore available agents**:
   ```bash
   ls ~/.claude/agents/
   # Core agents: cat-requirements-analyst, cat-documentation
   # Python: cat-solution-architect-python, cat-software-engineer-python
   # .NET: cat-solution-architect-dotnet, cat-software-engineer-dotnet
   # Node.js: cat-solution-architect-nodejs, cat-software-engineer-nodejs
   # Java: cat-solution-architect-java, cat-software-engineer-java
   ```

3. **Read the full documentation**:
   - [GitHub Repository](https://github.com/dkoenawan/claude-agent-templates)
   - [Agent Specifications](../../docs/agent-specifications.md)
   - [Domain Specialization](../../docs/domain-specialization.md)

4. **Keep updated**:
   ```bash
   # Check for updates monthly
   spec-kit-agents check

   # Update when available
   spec-kit-agents update
   ```

---

## FAQ

**Q: Do I still need to clone the repository?**

A: No! With global installation, the binary contains all files. You only need to clone if you want to contribute or customize agents.

**Q: How much disk space does it use?**

A: ~1MB for agents in `~/.claude/`, plus ~5MB for the binary. Total < 10MB.

**Q: Can I uninstall it?**

A: Yes! Run `spec-kit-agents uninstall` to remove the global installation. Or manually delete `~/.claude/agents/cat-*` and `~/.claude/commands/speckit.*`.

**Q: Will this break my existing setup?**

A: No. Global installation is additive - it doesn't modify your repositories or existing installations.

**Q: What if I have custom agents?**

A: Custom agents in `~/.claude/agents/` work alongside installed agents. Use different prefixes to avoid conflicts (e.g., `my-custom-agent.md` vs `cat-requirements-analyst.md`).

**Q: How do updates work?**

A: Run `spec-kit-agents update`. It downloads the latest version, backs up your current installation, updates files, and can rollback automatically if anything fails.

**Q: Can I install specific versions?**

A: Yes! Use `spec-kit-agents update --version v2.0.1` to install a specific version.

**Q: Does this work on Windows?**

A: Yes! The binary works on Linux, macOS, and Windows. Paths are adjusted automatically (e.g., `C:\Users\{user}\.claude\` on Windows).

---

## Getting Help

- **Issues**: [GitHub Issues](https://github.com/dkoenawan/claude-agent-templates/issues)
- **Discussions**: [GitHub Discussions](https://github.com/dkoenawan/claude-agent-templates/discussions)
- **Documentation**: [Full Documentation](../../README.md)

---

**Installation Complete!** 🎉

You now have claude-agent-templates installed globally and accessible from any repository. Start using `/speckit.*` commands in Claude Code to accelerate your development workflow.

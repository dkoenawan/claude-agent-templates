# Installation Scripts

This directory contains installation scripts to set up Claude Agent Templates globally on your system.

## Scripts

### `install-agents.sh` (Linux/macOS)
Bash script that installs agents to `~/.claude/agents` directory.

**Usage:**
```bash
chmod +x install-agents.sh
./install-agents.sh
```

### `install-agents.bat` (Windows)
Batch script that installs agents to `%USERPROFILE%\.claude\agents` directory.

**Usage:**
- Double-click the file in File Explorer, or
- Run from Command Prompt: `install-agents.bat`
- Run from PowerShell: `.\install-agents.bat`

## What the Scripts Do

1. Create global `.claude/agents` directory in your home folder
2. Copy agent templates from this repository to the global directory
3. Make agents available system-wide for Claude Code

## Available Agents

- **business-requirements-analyst**: Translates business requirements to technical specs
- **solution-architect**: Breaks down complex features into implementable work units

## Using the Agents

After installation, use agents in Claude Code:
1. Run `claude` to start Claude Code
2. Use `/agents` command to see available agents
3. Reference agents in your requests
# Installation Scripts - Migration Notice

**⚠️ DEPRECATION NOTICE: This directory contains legacy installation scripts that have been migrated to use the new Taskfile-based system.**

## Recommended Approach (New)

Use the unified Taskfile-based system for cross-platform compatibility:

```bash
# From the repository root
task install    # Install all agents
task list       # Show available and installed agents
task validate   # Verify installation integrity
task clean      # Remove installed agents
task update     # Update existing installations
task help       # Show detailed help
```

**Benefits of the new system:**
- Single interface works on Linux, macOS, and Windows
- Enhanced functionality beyond basic installation
- Consistent behavior across all platforms
- Better error handling and progress feedback

## Legacy Scripts (Deprecated)

These scripts are maintained for backward compatibility but will be removed in a future version:

### `install-agents.sh` (Linux/macOS)
Bash wrapper script that calls the Taskfile system.

**Usage:**
```bash
./install-agents.sh
```

### `install-agents.bat` (Windows)
Batch wrapper script that calls the Taskfile system.

**Usage:**
- Double-click the file in File Explorer, or
- Run from Command Prompt: `install-agents.bat`
- Run from PowerShell: `.\install-agents.bat`

## Migration Guide

### For New Users
Simply use `task install` instead of the legacy scripts.

### For Existing Users
Your existing scripts will continue to work, but you'll see deprecation warnings. To migrate:

1. **Install Task** (if not already installed):
   - Visit [taskfile.dev/installation](https://taskfile.dev/installation/)
   - Or use the included binary: `./bin/task`

2. **Use new commands**:
   - Replace `./scripts/install-agents.sh` with `task install`
   - Use `task help` to see all available operations

3. **Update scripts/CI/CD**:
   - Update any automation to use `task install` instead of legacy scripts

## What the System Does

1. **Install**: Creates global `.claude/agents` directory and copies agent templates
2. **List**: Shows available agents in repository and installation status
3. **Validate**: Verifies all agents are properly installed
4. **Clean**: Removes installed agents with confirmation prompt
5. **Update**: Refreshes existing installations with latest versions
6. **Help**: Shows comprehensive usage information

## Available Agents

### Core Workflow Agents (Language-Agnostic)
- **requirements-analyst**: Translates business requirements to technical specs
- **solution-architect**: Breaks down complex features into implementable work units  
- **documentation**: Performs final documentation updates and repository cleanup

### Python Development Agents
- **test-engineer-python**: Creates comprehensive unit test strategies with pytest
- **software-engineer-python**: Implements solutions using hexagonal architecture principles

## Using the Agents

After installation, use agents in Claude Code:
1. Run `claude` to start Claude Code
2. Use `/agents` command to see available agents
3. Reference agents in your requests

## Task Installation

### System Installation
Install Task globally from [taskfile.dev/installation](https://taskfile.dev/installation/)

### Project Binary
The repository includes a Task binary at `./bin/task` which will be used automatically if Task is not found in your system PATH.
# Claude Code Custom Commands

This directory contains custom slash commands for Claude Code that enhance the development workflow with git operations and other common tasks.

## Available Commands

### Git Commands (`/git/`)

#### `/git/clean-up-local`
Performs a complete local branch cleanup:
- Switches to the main branch
- Removes all local branches except main
- Fetches the latest changes from remote

**Usage:**
```
/git/clean-up-local
```

#### `/git/commit`
Creates conventional commits by analyzing git diff:
- Reads the current git diff to understand changes
- Analyzes the type and scope of changes
- Creates a conventional commit message following standard format
- Commits all staged changes

**Usage:**
```
/git/commit [optional message]
```

If no message is provided, the command will analyze the diff and generate an appropriate conventional commit message.

## Command Structure

Commands follow the Claude Code slash command standards:
- Located in `.claude/commands/` directory structure
- Organized by category (e.g., `git/`, `docker/`, etc.)
- Written in Markdown format with YAML frontmatter
- Support argument placeholders and tool restrictions

## Installation

Commands are automatically installed when running:
```bash
task install
```

This copies all command files from the `commands/` directory to your `~/.claude/commands/` directory, preserving the directory structure.

## Creating Custom Commands

To create new commands:
1. Create a new `.md` file in the appropriate subdirectory
2. Follow the Claude Code command format with frontmatter
3. Run `task install` to deploy the command
4. Use the command in Claude Code with the appropriate path

## Example Command Format

```markdown
---
allowed-tools: Bash(git:*)
argument-hint: [expected arguments]
description: Brief command description
---

# Command Title

Command implementation and instructions here.
```

For more information on creating custom commands, see the [Claude Code documentation](https://docs.anthropic.com/en/docs/claude-code/slash-commands).
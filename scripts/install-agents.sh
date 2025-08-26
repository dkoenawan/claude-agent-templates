#!/bin/bash

echo "Installing Claude Agent Templates..."


# Create global .claude directory if it doesn't exist
if [ ! -d "$HOME/.claude" ]; then
    echo "Creating ~/.claude directory..."
    mkdir -p "$HOME/.claude"
fi

# Create global agents directory if it doesn't exist  
if [ ! -d "$HOME/.claude/agents" ]; then
    echo "Creating ~/.claude/agents directory..."
    mkdir -p "$HOME/.claude/agents"
fi

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
AGENTS_SOURCE="$SCRIPT_DIR/../agents"

# Copy agent files
echo "Copying agent files to ~/.claude/agents..."

if [ -f "$AGENTS_SOURCE/core/requirements-analyst.md" ]; then
    cp "$AGENTS_SOURCE/core/requirements-analyst.md" "$HOME/.claude/agents/"
    echo "✓ Installed requirements-analyst agent"
else
    echo "✗ Warning: core/requirements-analyst.md not found"
fi

if [ -f "$AGENTS_SOURCE/core/solution-architect.md" ]; then
    cp "$AGENTS_SOURCE/core/solution-architect.md" "$HOME/.claude/agents/"
    echo "✓ Installed solution-architect agent"
else
    echo "✗ Warning: core/solution-architect.md not found"
fi

if [ -f "$AGENTS_SOURCE/python/test-engineer-python.md" ]; then
    cp "$AGENTS_SOURCE/python/test-engineer-python.md" "$HOME/.claude/agents/"
    echo "✓ Installed test-engineer-python agent"
else
    echo "✗ Warning: python/test-engineer-python.md not found"
fi

if [ -f "$AGENTS_SOURCE/python/software-engineer-python.md" ]; then
    cp "$AGENTS_SOURCE/python/software-engineer-python.md" "$HOME/.claude/agents/"
    echo "✓ Installed software-engineer-python agent"
else
    echo "✗ Warning: python/software-engineer-python.md not found"
fi

if [ -f "$AGENTS_SOURCE/core/documentation.md" ]; then
    cp "$AGENTS_SOURCE/core/documentation.md" "$HOME/.claude/agents/"
    echo "✓ Installed documentation agent"
else
    echo "✗ Warning: core/documentation.md not found"

fi

echo ""
echo "Installation complete!"
echo ""
echo "Available agents:"
echo "- requirements-analyst: Translates business requirements to technical specs"
echo "- solution-architect: Breaks down complex features into implementable work units"
echo "- test-engineer-python: Creates comprehensive unit test strategies with pytest"
echo "- software-engineer-python: Implements solutions with hexagonal architecture"
echo "- documentation: Performs final documentation updates and cleanup"

echo ""
echo "To use these agents in Claude Code:"
echo "1. Run 'claude' to start Claude Code"
echo "2. Use '/agents' command to see available agents"
echo "3. Use the agents by referencing them in your requests"
echo ""
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

if [ -f "$AGENTS_SOURCE/business-requirements-analyst.md" ]; then
    cp "$AGENTS_SOURCE/business-requirements-analyst.md" "$HOME/.claude/agents/"
    echo "✓ Installed business-requirements-analyst agent"
else
    echo "✗ Warning: business-requirements-analyst.md not found"
fi

if [ -f "$AGENTS_SOURCE/solution-architect.md" ]; then

    cp "$AGENTS_SOURCE/solution-architect.md" "$HOME/.claude/agents/"
    echo "✓ Installed solution-architect agent"
else
    echo "✗ Warning: solution-architect.md not found"
fi

echo ""
echo "Installation complete!"
echo ""
echo "Available agents:"
echo "- business-requirements-analyst: Translates business requirements to technical specs"
echo "- solution-architect: Breaks down complex features into implementable work units"
echo ""
echo "To use these agents in Claude Code:"
echo "1. Run 'claude' to start Claude Code"
echo "2. Use '/agents' command to see available agents"
echo "3. Use the agents by referencing them in your requests"
echo ""
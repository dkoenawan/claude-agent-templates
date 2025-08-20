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

# Copy all agent files
echo "Copying all agent files to ~/.claude/agents..."

if [ -d "$AGENTS_SOURCE" ]; then
    # Copy all .md files from agents directory, excluding README.md
    find "$AGENTS_SOURCE" -name "*.md" -not -name "README.md" -exec cp {} "$HOME/.claude/agents/" \;
    
    # List installed agents
    AGENT_COUNT=$(find "$HOME/.claude/agents" -name "*.md" | wc -l)
    echo "✓ Installed $AGENT_COUNT agents:"
    find "$HOME/.claude/agents" -name "*.md" -exec basename {} .md \; | sed 's/^/  - /'
else
    echo "✗ Error: agents directory not found at $AGENTS_SOURCE"
    exit 1
fi

echo ""
echo "Installation complete!"
echo ""
echo "Available agents:"
echo "- business-requirements-analyst: Translates business requirements to technical specs"
echo "- solution-architect: Breaks down complex features into implementable work units"
echo "- software-engineer-python: Implements solutions using hexagonal architecture principles"
echo "- documentation: Performs post-implementation documentation and cleanup"
echo ""
echo "To use these agents in Claude Code:"
echo "1. Run 'claude' to start Claude Code"
echo "2. Use '/agents' command to see available agents"
echo "3. Use the agents by referencing them in your requests"
echo ""
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

# Link to agents directory for automatic updates
echo "Setting up agents directory link..."

if [ -d "$AGENTS_SOURCE" ]; then
    # Remove existing agents directory if it exists
    if [ -d "$HOME/.claude/agents" ] || [ -L "$HOME/.claude/agents" ]; then
        rm -rf "$HOME/.claude/agents"
    fi
    
    # Create symlink to this repository's agents directory
    ln -s "$AGENTS_SOURCE" "$HOME/.claude/agents"
    
    # Count and list available agents
    AGENT_COUNT=$(find "$AGENTS_SOURCE" -name "*.md" -not -name "README.md" | wc -l)
    echo "✓ Linked to $AGENT_COUNT agents:"
    find "$AGENTS_SOURCE" -name "*.md" -not -name "README.md" -exec basename {} .md \; | sed 's/^/  - /'
else
    echo "✗ Error: agents directory not found at $AGENTS_SOURCE"
    exit 1
fi

echo ""
echo "Installation complete!"
echo ""
echo "Agents are now linked and will automatically stay up-to-date with this repository."
echo ""
echo "To use these agents in Claude Code:"
echo "1. Run 'claude' to start Claude Code"
echo "2. Use '/agents' command to see available agents"
echo "3. Use the agents by referencing them in your requests"
echo ""
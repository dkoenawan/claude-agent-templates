#!/bin/bash

echo "Installing Claude Agent Templates..."

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
AGENTS_SOURCE="$SCRIPT_DIR/../agents"

# Create global .claude directory structure
mkdir -p "$HOME/.claude/agents"

# Find and copy all agent files
echo "Discovering and copying agent files to ~/.claude/agents..."

installed_count=0
missing_count=0

# Find all .md files in agents subdirectories
while IFS= read -r -d '' agent_file; do
    # Get the filename without path
    filename=$(basename "$agent_file")
    
    # Skip README files
    if [[ "$filename" == "README.md" ]]; then
        continue
    fi
    
    # Copy to global agents directory
    if cp "$agent_file" "$HOME/.claude/agents/"; then
        echo "✓ Installed $filename"
        ((installed_count++))
    else
        echo "✗ Failed to install $filename"
        ((missing_count++))
    fi
done < <(find "$AGENTS_SOURCE" -name "*.md" -type f -print0)

if [ $installed_count -eq 0 ]; then
    echo "✗ No agent files found in $AGENTS_SOURCE"
    exit 1
fi

echo ""
echo "Installation complete! Installed $installed_count agents."
echo ""
echo "To use these agents in Claude Code:"
echo "1. Run 'claude' to start Claude Code"
echo "2. Use '/agents' command to see available agents"
echo "3. Use the agents by referencing them in your requests"
echo ""
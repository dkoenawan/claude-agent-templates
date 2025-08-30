#!/bin/bash

# Backward compatibility wrapper for install-agents.sh
# This script has been migrated to use Taskfile for cross-platform consistency.
# 
# DEPRECATION NOTICE: This script will be removed in a future version.
# Please migrate to using: task install
#
# For more information about the new Taskfile-based automation, run: task help

echo "⚠️  DEPRECATION NOTICE"
echo "   This script is deprecated and will be removed in a future version."
echo "   Please migrate to the new Taskfile-based system:"
echo "   - Install Task: https://taskfile.dev/installation/"
echo "   - Run: task install"
echo "   - For help: task help"
echo ""

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$SCRIPT_DIR/.."

# Check if Task is available and use it if possible
if command -v task >/dev/null 2>&1; then
    echo "Using Taskfile-based installation..."
    cd "$PROJECT_ROOT"
    exec task install
fi

# Check if Task binary is available in project
if [ -f "$PROJECT_ROOT/bin/task" ]; then
    echo "Using project Task binary..."
    cd "$PROJECT_ROOT"
    exec "$PROJECT_ROOT/bin/task" install
fi

# Fallback to original implementation
echo "Task not found, using legacy implementation..."
echo "Consider installing Task for better cross-platform support."
echo ""

echo "Installing Claude Agent Templates..."

AGENTS_SOURCE="$SCRIPT_DIR/../agents"

# Create global .claude directory structure
mkdir -p "$HOME/.claude/agents"

# Find and copy all agent files
echo "Discovering and copying agent files to ~/.claude/agents..."

installed_count=0
missing_count=0

# Find all .md files in agents subdirectories
find "$AGENTS_SOURCE" -name "*.md" -type f | while read -r agent_file; do
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
done

# Count installed files for feedback
installed_count=$(find "$HOME/.claude/agents" -name "*.md" -type f | wc -l)

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
#!/bin/bash

set -e

AGENT_FILE="$1"
ERRORS=0

if [ -z "$AGENT_FILE" ]; then
    echo "Usage: $0 <agent-specification-file>"
    exit 1
fi

if [ ! -f "$AGENT_FILE" ]; then
    echo "Error: File $AGENT_FILE not found"
    exit 1
fi

echo "Validating agent specification: $AGENT_FILE"

# Check for required sections
REQUIRED_SECTIONS=(
    "name:"
    "description:"
    "domain:"
    "capabilities:"
    "tools:"
    "workflow:"
    "constraints:"
    "examples:"
)

for section in "${REQUIRED_SECTIONS[@]}"; do
    if ! grep -q "^$section" "$AGENT_FILE"; then
        echo "❌ Missing required section: $section"
        ((ERRORS++))
    else
        echo "✓ Found section: $section"
    fi
done

# Check for proper YAML/Markdown structure
if ! grep -q "^---$" "$AGENT_FILE"; then
    echo "⚠️ Warning: No YAML frontmatter detected"
fi

# Check for domain-specific requirements
DOMAIN=$(grep "^domain:" "$AGENT_FILE" | cut -d: -f2 | xargs)
if [ ! -z "$DOMAIN" ]; then
    echo "Domain detected: $DOMAIN"

    case "$DOMAIN" in
        python)
            if ! grep -q "pytest\|unittest\|python" "$AGENT_FILE"; then
                echo "⚠️ Warning: Python domain agent should mention Python tools"
            fi
            ;;
        dotnet)
            if ! grep -q "\.NET\|C#\|dotnet\|xunit\|nunit" "$AGENT_FILE"; then
                echo "⚠️ Warning: .NET domain agent should mention .NET tools"
            fi
            ;;
        nodejs)
            if ! grep -q "node\|npm\|jest\|mocha\|javascript\|typescript" "$AGENT_FILE"; then
                echo "⚠️ Warning: Node.js domain agent should mention Node.js tools"
            fi
            ;;
        java)
            if ! grep -q "java\|junit\|maven\|gradle" "$AGENT_FILE"; then
                echo "⚠️ Warning: Java domain agent should mention Java tools"
            fi
            ;;
    esac
fi

# Check for workflow states
if grep -q "^workflow:" "$AGENT_FILE"; then
    for state in "input" "processing" "output" "validation"; do
        if ! grep -q "$state" "$AGENT_FILE"; then
            echo "⚠️ Warning: Workflow should include $state state"
        fi
    done
fi

if [ $ERRORS -eq 0 ]; then
    echo "✅ Validation successful!"
    exit 0
else
    echo "❌ Validation failed with $ERRORS errors"
    exit 1
fi
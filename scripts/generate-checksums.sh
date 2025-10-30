#!/bin/bash
# Generate SHA256 checksums for all embedded files
# This script creates a checksums file used during build validation

set -e

OUTPUT_FILE="internal/embed/checksums.json"
TEMP_FILE="${OUTPUT_FILE}.tmp"

echo "Generating checksums for embedded files..."

# Start JSON array
echo "{" > "$TEMP_FILE"
echo '  "checksums": {' >> "$TEMP_FILE"

# Find all agent files
FIRST=true
for file in agents/**/*.md; do
    if [ -f "$file" ]; then
        # Compute SHA256
        if command -v sha256sum >/dev/null 2>&1; then
            CHECKSUM=$(sha256sum "$file" | cut -d' ' -f1)
        elif command -v shasum >/dev/null 2>&1; then
            CHECKSUM=$(shasum -a 256 "$file" | cut -d' ' -f1)
        else
            echo "ERROR: No SHA256 utility found (sha256sum or shasum required)"
            exit 1
        fi

        # Add comma separator except for first entry
        if [ "$FIRST" = true ]; then
            FIRST=false
        else
            echo "," >> "$TEMP_FILE"
        fi

        # Add checksum entry
        echo -n "    \"$file\": \"$CHECKSUM\"" >> "$TEMP_FILE"
    fi
done

# Find all .specify files
for file in .specify/**/*; do
    if [ -f "$file" ]; then
        # Compute SHA256
        if command -v sha256sum >/dev/null 2>&1; then
            CHECKSUM=$(sha256sum "$file" | cut -d' ' -f1)
        elif command -v shasum >/dev/null 2>&1; then
            CHECKSUM=$(shasum -a 256 "$file" | cut -d' ' -f1)
        else
            echo "ERROR: No SHA256 utility found"
            exit 1
        fi

        # Add comma separator
        echo "," >> "$TEMP_FILE"

        # Add checksum entry
        echo -n "    \"$file\": \"$CHECKSUM\"" >> "$TEMP_FILE"
    fi
done

# Close JSON
echo "" >> "$TEMP_FILE"
echo "  }," >> "$TEMP_FILE"
echo "  \"generated_at\": \"$(date -u +%Y-%m-%dT%H:%M:%SZ)\"" >> "$TEMP_FILE"
echo "}" >> "$TEMP_FILE"

# Move temp file to output
mv "$TEMP_FILE" "$OUTPUT_FILE"

FILE_COUNT=$(grep -c '".*":' "$OUTPUT_FILE" || echo "0")
echo "Generated checksums for $FILE_COUNT files → $OUTPUT_FILE"

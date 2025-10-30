#!/bin/bash
# Build spec-kit-agents with embedded files
# This script builds the binary with all source files embedded using Go's embed package

set -e

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Building spec-kit-agents with embedded files...${NC}"

# Get version from git or default
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

echo -e "${YELLOW}Version: $VERSION${NC}"
echo -e "${YELLOW}Build Time: $BUILD_TIME${NC}"
echo -e "${YELLOW}Commit: $COMMIT${NC}"

# Generate checksums before building
if [ -f "scripts/generate-checksums.sh" ]; then
    echo -e "${GREEN}Generating checksums...${NC}"
    bash scripts/generate-checksums.sh
fi

# Build with ldflags to inject version info
LDFLAGS="-X 'main.Version=$VERSION' -X 'main.BuildTime=$BUILD_TIME' -X 'main.Commit=$COMMIT'"

# Create output directory
mkdir -p bin

# Build for current platform
echo -e "${GREEN}Building for current platform...${NC}"
go build -ldflags="$LDFLAGS" -o bin/spec-kit-agents ./cmd/spec-kit-agents/

# Check binary size
SIZE=$(du -h bin/spec-kit-agents | cut -f1)
echo -e "${GREEN}Binary size: $SIZE${NC}"

# Verify embed worked (binary should be > 1MB with embedded files)
SIZE_BYTES=$(stat -f%z bin/spec-kit-agents 2>/dev/null || stat -c%s bin/spec-kit-agents 2>/dev/null)
if [ "$SIZE_BYTES" -lt 1048576 ]; then
    echo -e "${RED}WARNING: Binary size is unusually small ($SIZE). Embedded files may be missing.${NC}"
fi

echo -e "${GREEN}Build complete: bin/spec-kit-agents${NC}"

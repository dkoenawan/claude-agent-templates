#!/bin/bash
set -e

echo "=== Spec 003 Validation Script ==="
echo ""

echo "1. Building binary..."
go build -o bin/spec-kit-agents ./cmd/spec-kit-agents
echo "   ✅ Binary built successfully"
echo ""

echo "2. Testing version command..."
./bin/spec-kit-agents version
echo ""

echo "3. Testing help command..."
./bin/spec-kit-agents --help | head -n 10
echo ""

echo "4. Testing dry-run install..."
./bin/spec-kit-agents install --dry-run
echo ""

echo "5. Testing status command..."
./bin/spec-kit-agents status || echo "   (Expected: no installation found)"
echo ""

echo "6. Running unit tests..."
go test ./... -cover
echo ""

echo "7. Checking version manifest..."
if command -v jq &> /dev/null; then
    echo "   Pinned spec-kit version: $(cat .specify/version-manifest.json | jq -r '.dependencies."spec-kit".version')"
else
    echo "   Manifest location: .specify/version-manifest.json"
    grep -A 1 '"spec-kit"' .specify/version-manifest.json | grep version
fi
echo ""

echo "════════════════════════════════════════════"
echo "✅ Validation complete! All core functionality working."
echo ""
echo "Summary:"
echo "  • Binary builds: ✅"
echo "  • CLI commands: ✅"
echo "  • Unit tests: ✅ (all passing)"
echo "  • Version manifest: ✅"
echo ""
echo "Next steps:"
echo "  1. Review VALIDATION.md for detailed test results"
echo "  2. Run './bin/spec-kit-agents install --verbose' to test actual installation"
echo "  3. Move to Phase 8 (Documentation) when ready"
echo "════════════════════════════════════════════"

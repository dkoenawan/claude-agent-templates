#!/bin/bash
# Version injection helper for build process
# Extracts version information for use in builds

set -e

# Get version from git tags
get_version() {
    if git describe --tags --exact-match 2>/dev/null; then
        # Exact tag match
        git describe --tags --exact-match
    elif git describe --tags --always 2>/dev/null; then
        # Tag + commits ahead
        git describe --tags --always --dirty
    else
        # No git or no tags
        echo "v0.0.0-dev"
    fi
}

# Get build timestamp in RFC3339 format
get_build_time() {
    date -u +%Y-%m-%dT%H:%M:%SZ
}

# Get commit hash
get_commit() {
    git rev-parse --short HEAD 2>/dev/null || echo "unknown"
}

# Output format based on argument
case "${1:-version}" in
    version)
        get_version
        ;;
    buildtime)
        get_build_time
        ;;
    commit)
        get_commit
        ;;
    all)
        echo "VERSION=$(get_version)"
        echo "BUILD_TIME=$(get_build_time)"
        echo "COMMIT=$(get_commit)"
        ;;
    ldflags)
        VERSION=$(get_version)
        BUILD_TIME=$(get_build_time)
        COMMIT=$(get_commit)
        echo "-X 'main.Version=$VERSION' -X 'main.BuildTime=$BUILD_TIME' -X 'main.Commit=$COMMIT'"
        ;;
    *)
        echo "Usage: $0 {version|buildtime|commit|all|ldflags}"
        exit 1
        ;;
esac

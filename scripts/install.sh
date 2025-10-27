#!/usr/bin/env bash
# One-liner installer for spec-kit-agents
# Usage: curl -fsSL https://raw.githubusercontent.com/dkoenawan/claude-agent-templates/main/scripts/install.sh | bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
REPO="dkoenawan/claude-agent-templates"
BINARY_NAME="spec-kit-agents"
INSTALL_DIR="$HOME/.local/bin"
GITHUB_API="https://api.github.com/repos/$REPO/releases/latest"

# Detect OS and architecture
detect_platform() {
    local os
    local arch

    # Detect OS
    case "$(uname -s)" in
        Linux*)     os="linux" ;;
        Darwin*)    os="darwin" ;;
        MINGW*|MSYS*|CYGWIN*) os="windows" ;;
        *)
            echo -e "${RED}Error: Unsupported operating system$(NC)"
            exit 1
            ;;
    esac

    # Detect architecture
    case "$(uname -m)" in
        x86_64|amd64)  arch="amd64" ;;
        arm64|aarch64) arch="arm64" ;;
        *)
            echo -e "${RED}Error: Unsupported architecture: $(uname -m)${NC}"
            exit 1
            ;;
    esac

    echo "${os}-${arch}"
}

# Download and install binary
install_binary() {
    local platform="$1"
    local version="$2"

    echo -e "${GREEN}Installing spec-kit-agents...${NC}"
    echo "  Platform: $platform"
    echo "  Version: $version"

    # Construct download URL
    local binary_name="${BINARY_NAME}-${version}-${platform}"
    if [[ "$platform" == windows-* ]]; then
        binary_name="${binary_name}.exe"
    fi

    local download_url="https://github.com/$REPO/releases/download/$version/$binary_name"

    echo "  Download URL: $download_url"

    # Create install directory
    mkdir -p "$INSTALL_DIR"

    # Download binary
    local temp_file=$(mktemp)
    echo -e "${YELLOW}Downloading...${NC}"

    if command -v curl &> /dev/null; then
        curl -fsSL "$download_url" -o "$temp_file"
    elif command -v wget &> /dev/null; then
        wget -q "$download_url" -O "$temp_file"
    else
        echo -e "${RED}Error: Neither curl nor wget is available${NC}"
        exit 1
    fi

    # Make binary executable
    chmod +x "$temp_file"

    # Move to install directory
    local install_path="$INSTALL_DIR/$BINARY_NAME"
    mv "$temp_file" "$install_path"

    echo -e "${GREEN}✓ Binary installed to $install_path${NC}"

    # Check if install directory is in PATH
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        echo ""
        echo -e "${YELLOW}Warning: $INSTALL_DIR is not in your PATH${NC}"
        echo "Add the following line to your shell configuration file:"
        echo ""
        echo "  export PATH=\"\$PATH:$INSTALL_DIR\""
        echo ""
    fi

    # Verify installation
    if command -v "$BINARY_NAME" &> /dev/null; then
        echo -e "${GREEN}✓ Installation verified${NC}"
        "$BINARY_NAME" version
    else
        echo -e "${YELLOW}Installation complete. Run the following to use:${NC}"
        echo "  $install_path"
    fi
}

# Main installation flow
main() {
    echo ""
    echo "========================================="
    echo "  spec-kit-agents Installer"
    echo "========================================="
    echo ""

    # Detect platform
    platform=$(detect_platform)
    echo "Detected platform: $platform"

    # Get latest release version
    echo "Fetching latest release..."

    if ! command -v curl &> /dev/null && ! command -v wget &> /dev/null; then
        echo -e "${RED}Error: Neither curl nor wget is available${NC}"
        exit 1
    fi

    # Fetch latest release version from GitHub API
    if command -v curl &> /dev/null; then
        version=$(curl -fsSL "$GITHUB_API" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    elif command -v wget &> /dev/null; then
        version=$(wget -qO- "$GITHUB_API" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    else
        echo -e "${RED}Error: Neither curl nor wget is available${NC}"
        exit 1
    fi

    if [ -z "$version" ]; then
        echo -e "${RED}Error: Could not fetch latest release version${NC}"
        echo "Please check your internet connection or install manually:"
        echo "  https://github.com/$REPO/releases"
        exit 1
    fi

    echo "Latest version: $version"
    echo ""

    # Install the binary
    install_binary "$platform" "$version"
}

main "$@"

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// GetHomeDir returns the user's home directory in a cross-platform way
func GetHomeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return home, nil
}

// GetDefaultInstallDir returns the default installation directory
// Returns ~/.claude-agent-templates on Unix-like systems
// Returns %USERPROFILE%\.claude-agent-templates on Windows
func GetDefaultInstallDir() (string, error) {
	home, err := GetHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".claude-agent-templates"), nil
}

// GetClaudeDir returns the Claude Code configuration directory
// Returns ~/.claude on Unix-like systems
// Returns %USERPROFILE%\.claude on Windows
func GetClaudeDir() (string, error) {
	home, err := GetHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".claude"), nil
}

// GetVersionLockPath returns the path to the version lock file
// If prefix is empty, uses default installation directory
func GetVersionLockPath(prefix string) (string, error) {
	if prefix == "" {
		defaultDir, err := GetDefaultInstallDir()
		if err != nil {
			return "", err
		}
		prefix = defaultDir
	}
	return filepath.Join(prefix, ".version-lock.json"), nil
}

// GetVersionManifestPath returns the path to the version manifest file
// Relative to the repository root or installation directory
func GetVersionManifestPath(prefix string) string {
	if prefix == "" {
		prefix = "."
	}
	return filepath.Join(prefix, ".specify", "version-manifest.json")
}

// GetInstallLogPath returns the path to the installation log file
func GetInstallLogPath(prefix string) (string, error) {
	if prefix == "" {
		defaultDir, err := GetDefaultInstallDir()
		if err != nil {
			return "", err
		}
		prefix = defaultDir
	}
	return filepath.Join(prefix, ".install-log.txt"), nil
}

// EnsureDir creates a directory and all parent directories if they don't exist
func EnsureDir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	return nil
}

// PathExists checks if a path exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsDirectory checks if a path exists and is a directory
func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// NormalizePath normalizes a file path for the current platform
func NormalizePath(path string) string {
	return filepath.Clean(path)
}

// JoinPath joins path elements and normalizes for the current platform
func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}

// GetPlatform returns a string describing the current platform
func GetPlatform() string {
	return fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
}

// IsWindows returns true if running on Windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsUnix returns true if running on Unix-like systems (Linux, macOS, BSD)
func IsUnix() bool {
	return !IsWindows()
}

// ExpandPath expands ~ to the user's home directory
func ExpandPath(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	home, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	if len(path) == 1 {
		return home, nil
	}

	return filepath.Join(home, path[1:]), nil
}

// ToAbsolutePath converts a path to an absolute path
func ToAbsolutePath(path string) (string, error) {
	// First expand ~ if present
	expanded, err := ExpandPath(path)
	if err != nil {
		return "", err
	}

	// Then convert to absolute path
	absPath, err := filepath.Abs(expanded)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for %s: %w", expanded, err)
	}

	return absPath, nil
}

// DetectSpecifyDir checks if .specify/ directory exists in current directory
func DetectSpecifyDir() bool {
	return IsDirectory(".specify")
}

// DetermineInstallPrefix determines the installation prefix based on existing setup
// Returns ".claude-agent-templates" if .specify/ exists, otherwise "."
func DetermineInstallPrefix() string {
	if DetectSpecifyDir() {
		return ".claude-agent-templates"
	}
	return "."
}

// GetClaudeCommandsDir returns the path to Claude commands directory
func GetClaudeCommandsDir() (string, error) {
	claudeDir, err := GetClaudeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(claudeDir, "commands"), nil
}

// GetClaudeAgentsDir returns the path to Claude agents directory
func GetClaudeAgentsDir() (string, error) {
	claudeDir, err := GetClaudeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(claudeDir, "agents"), nil
}

// GetClaudeSkillsDir returns the path to Claude skills directory
func GetClaudeSkillsDir() (string, error) {
	claudeDir, err := GetClaudeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(claudeDir, "skills"), nil
}

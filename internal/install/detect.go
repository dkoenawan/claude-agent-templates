package install

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dkoenawan/claude-agent-templates/internal/config"
)

// DetectExistingInstallation checks for existing .specify/ directory
// Returns true if .specify/ exists in current directory
func DetectExistingInstallation() bool {
	return config.DetectSpecifyDir()
}

// DetermineInstallPrefix determines where to install based on existing setup
// Returns:
//   - ".claude-agent-templates" if .specify/ exists (coexist mode)
//   - "." if no .specify/ exists (standalone mode)
func DetermineInstallPrefix() string {
	return config.DetermineInstallPrefix()
}

// DetectClaudeDir checks if .claude/ directory exists
func DetectClaudeDir() (bool, string, error) {
	claudeDir, err := config.GetClaudeDir()
	if err != nil {
		return false, "", err
	}

	exists := config.PathExists(claudeDir)
	return exists, claudeDir, nil
}

// DetectExistingVersionLock checks if a version lock file exists at the given prefix
func DetectExistingVersionLock(prefix string) (bool, string, error) {
	lockPath, err := config.GetVersionLockPath(prefix)
	if err != nil {
		return false, "", err
	}

	exists := config.PathExists(lockPath)
	return exists, lockPath, nil
}

// DetectInstallationMode determines the installation mode based on environment
type InstallationMode struct {
	Mode        string // "fresh", "upgrade", "coexist"
	HasSpecKit  bool
	HasClaude   bool
	HasLock     bool
	Prefix      string
	Description string
}

// DetectMode analyzes the environment and determines the installation mode
func DetectMode(customPrefix string) (*InstallationMode, error) {
	mode := &InstallationMode{}

	// Check for existing .specify/
	mode.HasSpecKit = DetectExistingInstallation()

	// Check for .claude/
	hasClaudeDir, _, err := DetectClaudeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to detect .claude/ directory: %w", err)
	}
	mode.HasClaude = hasClaudeDir

	// Determine prefix
	if customPrefix != "" {
		mode.Prefix = customPrefix
	} else {
		mode.Prefix = DetermineInstallPrefix()
	}

	// Check for existing version lock at determined prefix
	hasLock, _, err := DetectExistingVersionLock(mode.Prefix)
	if err != nil {
		return nil, fmt.Errorf("failed to detect version lock: %w", err)
	}
	mode.HasLock = hasLock

	// Determine mode and description
	if mode.HasLock {
		mode.Mode = "upgrade"
		mode.Description = "Upgrade existing installation"
	} else if mode.HasSpecKit {
		mode.Mode = "coexist"
		mode.Description = fmt.Sprintf("Install alongside existing spec-kit (prefix: %s)", mode.Prefix)
	} else {
		mode.Mode = "fresh"
		mode.Description = "Fresh installation"
	}

	return mode, nil
}

// ValidateInstallationDirectory checks if the installation directory is valid
func ValidateInstallationDirectory(prefix string) error {
	absPath, err := config.ToAbsolutePath(prefix)
	if err != nil {
		return fmt.Errorf("invalid installation path: %w", err)
	}

	// Check if directory exists
	if config.PathExists(absPath) {
		// Ensure it's a directory
		if !config.IsDirectory(absPath) {
			return fmt.Errorf("installation path exists but is not a directory: %s", absPath)
		}

		// Check if directory is writable
		testFile := filepath.Join(absPath, ".write-test")
		if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
			return fmt.Errorf("installation directory is not writable: %s", absPath)
		}
		os.Remove(testFile)
	} else {
		// Try to create directory
		if err := config.EnsureDir(absPath); err != nil {
			return fmt.Errorf("failed to create installation directory: %w", err)
		}
	}

	return nil
}

// DetectGitRepository checks if current directory is a git repository
func DetectGitRepository() bool {
	gitDir := filepath.Join(".", ".git")
	return config.IsDirectory(gitDir)
}

// DetectRepositoryRoot attempts to find the repository root
// Returns the root path or current directory if not in a repo
func DetectRepositoryRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	// Walk up the directory tree looking for .git
	dir := cwd
	for {
		gitDir := filepath.Join(dir, ".git")
		if config.IsDirectory(gitDir) {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root, no .git found
			return cwd, nil
		}
		dir = parent
	}
}

// GetInstallationPaths returns all relevant paths for installation
type InstallationPaths struct {
	Prefix            string
	VersionLock       string
	VersionManifest   string
	InstallLog        string
	ClaudeDir         string
	ClaudeCommands    string
	ClaudeAgents      string
	ClaudeSkills      string
	SpecifyDir        string
	AgentsSourceDir   string
	TemplatesDir      string
}

// GetPaths calculates all installation paths based on the prefix
func GetPaths(prefix string) (*InstallationPaths, error) {
	// Ensure prefix is absolute
	absPrefix, err := config.ToAbsolutePath(prefix)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve installation prefix: %w", err)
	}

	paths := &InstallationPaths{
		Prefix: absPrefix,
	}

	// Version lock path
	paths.VersionLock, err = config.GetVersionLockPath(prefix)
	if err != nil {
		return nil, err
	}

	// Version manifest path
	paths.VersionManifest = config.GetVersionManifestPath(prefix)

	// Install log path
	paths.InstallLog, err = config.GetInstallLogPath(prefix)
	if err != nil {
		return nil, err
	}

	// Claude directories
	paths.ClaudeDir, err = config.GetClaudeDir()
	if err != nil {
		return nil, err
	}

	paths.ClaudeCommands, err = config.GetClaudeCommandsDir()
	if err != nil {
		return nil, err
	}

	paths.ClaudeAgents, err = config.GetClaudeAgentsDir()
	if err != nil {
		return nil, err
	}

	paths.ClaudeSkills, err = config.GetClaudeSkillsDir()
	if err != nil {
		return nil, err
	}

	// Spec-kit and source directories
	paths.SpecifyDir = filepath.Join(prefix, ".specify")
	paths.AgentsSourceDir = filepath.Join(prefix, "agents")
	paths.TemplatesDir = filepath.Join(prefix, ".specify", "templates")

	return paths, nil
}

// VerifySourceFiles checks that all required source files exist before installation
func VerifySourceFiles() error {
	requiredDirs := []string{
		".specify",
		"agents",
	}

	requiredFiles := []string{
		".specify/version-manifest.json",
	}

	for _, dir := range requiredDirs {
		if !config.IsDirectory(dir) {
			return fmt.Errorf("required directory not found: %s (are you running from repository root?)", dir)
		}
	}

	for _, file := range requiredFiles {
		if !config.PathExists(file) {
			return fmt.Errorf("required file not found: %s", file)
		}
	}

	return nil
}

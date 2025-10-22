package install

import (
	"fmt"

	"github.com/dkoenawan/claude-agent-templates/internal/config"
	"github.com/dkoenawan/claude-agent-templates/internal/version"
)

// Options contains installation configuration options
type Options struct {
	Prefix  string
	Global  bool
	Force   bool
	Quiet   bool
	DryRun  bool
}

// InstallationResult contains the results of an installation
type InstallationResult struct {
	Success           bool
	Mode              string
	Prefix            string
	TemplatesVersion  string
	SpecKitVersion    string
	FilesInstalled    int
	ClaudeIntegration *ClaudeIntegrationResult
	Errors            []error
	Warnings          []string
}

// Run executes the installation process
func Run(opts Options, logger *config.Logger) (*InstallationResult, error) {
	result := &InstallationResult{
		Errors:   []error{},
		Warnings: []string{},
	}

	logger.Info("installer", "Starting spec-kit lockstep installation...")

	// Step 1: Verify source files
	logger.Debug("installer", "Verifying source files...")
	if err := VerifySourceFiles(); err != nil {
		return nil, fmt.Errorf("source verification failed: %w", err)
	}
	logger.Success("installer", "Source files verified")

	// Step 2: Detect installation mode
	logger.Debug("installer", "Detecting installation mode...")
	mode, err := DetectMode(opts.Prefix)
	if err != nil {
		return nil, fmt.Errorf("failed to detect installation mode: %w", err)
	}
	result.Mode = mode.Mode
	result.Prefix = mode.Prefix
	logger.Info("installer", "Installation mode: %s", mode.Description)

	// Step 3: Check for existing installation (upgrade scenario)
	if mode.HasLock && !opts.Force {
		logger.Warn("installer", "Existing installation detected")
		logger.Info("installer", "Use --force to overwrite existing installation")
		return nil, fmt.Errorf("installation already exists at %s", mode.Prefix)
	}

	// Step 4: Validate installation directory
	logger.Debug("installer", "Validating installation directory...")
	if err := ValidateInstallationDirectory(mode.Prefix); err != nil {
		return nil, fmt.Errorf("installation directory validation failed: %w", err)
	}

	// Step 5: Get all installation paths
	paths, err := GetPaths(mode.Prefix)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate installation paths: %w", err)
	}

	// Step 6: Load version manifest
	logger.Debug("installer", "Loading version manifest...")
	manifest, err := version.LoadManifestFromPath(".specify/version-manifest.json")
	if err != nil {
		return nil, fmt.Errorf("failed to load version manifest: %w", err)
	}

	specKitVersion, err := version.GetSpecKitVersion(manifest)
	if err != nil {
		return nil, fmt.Errorf("failed to get spec-kit version: %w", err)
	}
	result.SpecKitVersion = specKitVersion
	result.TemplatesVersion = "1.0.0" // TODO: Get from git tag or version file

	logger.Info("installer", "Installing claude-agent-templates v%s with spec-kit v%s",
		result.TemplatesVersion, result.SpecKitVersion)

	if opts.DryRun {
		logger.Info("installer", "Dry run mode - no files will be modified")
		result.Success = true
		return result, nil
	}

	// Step 7: Copy .specify/ directory (spec-kit files)
	logger.Info("installer", "Copying spec-kit files to %s...", paths.SpecifyDir)
	if err := CopySpecKitFiles(".specify", paths.SpecifyDir); err != nil {
		return nil, fmt.Errorf("failed to copy spec-kit files: %w", err)
	}

	specKitFileCount, err := CountFilesRecursive(paths.SpecifyDir)
	if err != nil {
		logger.Warn("installer", "Failed to count spec-kit files: %v", err)
	} else {
		logger.Success("installer", "Copied %d spec-kit files", specKitFileCount)
		result.FilesInstalled += specKitFileCount
	}

	// Step 8: Set up .claude/ directory structure
	logger.Info("installer", "Setting up Claude Code integration...")
	if err := SetupClaudeDirectory(); err != nil {
		return nil, fmt.Errorf("failed to setup Claude directory: %w", err)
	}

	// Step 9: Integrate with Claude Code (copy agents and commands)
	claudeResult, err := IntegrateWithClaude(paths)
	if err != nil {
		return nil, fmt.Errorf("Claude Code integration failed: %w", err)
	}
	result.ClaudeIntegration = claudeResult
	logger.Success("installer", "%s", claudeResult.GetSummary())

	// Step 10: Create version lock
	logger.Info("installer", "Creating version lock...")
	versionLock := version.CreateVersionLock(
		result.TemplatesVersion,
		result.SpecKitVersion,
		paths.Prefix,
	)

	if err := version.SaveVersionLock(versionLock, paths.VersionLock); err != nil {
		return nil, fmt.Errorf("failed to save version lock: %w", err)
	}
	logger.Success("installer", "Version lock created at %s", paths.VersionLock)

	// Step 11: Verify installation
	logger.Info("installer", "Verifying installation...")
	if err := VerifyInstallation(paths); err != nil {
		return nil, fmt.Errorf("installation verification failed: %w", err)
	}
	logger.Success("installer", "Installation verified")

	// Installation complete
	result.Success = true
	logger.Success("installer", "Installation complete!")
	logger.Info("installer", "")
	logger.Info("installer", "Installation Summary:")
	logger.Info("installer", "  Location: %s", paths.Prefix)
	logger.Info("installer", "  claude-agent-templates: v%s", result.TemplatesVersion)
	logger.Info("installer", "  spec-kit: v%s", result.SpecKitVersion)
	logger.Info("installer", "  Files installed: %d", result.FilesInstalled)
	logger.Info("installer", "  Agents available: %d (prefix: cat-)", claudeResult.AgentsCopied)
	logger.Info("installer", "  Commands available: %d (prefix: speckit.)", claudeResult.CommandsCopied)
	logger.Info("installer", "")
	logger.Info("installer", "Claude Code is now configured!")
	logger.Info("installer", "  Agents: ~/.claude/agents/cat-*.md")
	logger.Info("installer", "  Commands: Use /speckit.specify, /speckit.plan, /speckit.tasks")

	return result, nil
}

// VerifyInstallation checks that all required files and directories are in place
func VerifyInstallation(paths *InstallationPaths) error {
	// Check that .specify/ was copied
	if !config.IsDirectory(paths.SpecifyDir) {
		return fmt.Errorf(".specify/ directory not found at %s", paths.SpecifyDir)
	}

	// Check that version manifest exists
	if !config.PathExists(paths.VersionManifest) {
		return fmt.Errorf("version manifest not found at %s", paths.VersionManifest)
	}

	// Check that version lock was created
	if !config.PathExists(paths.VersionLock) {
		return fmt.Errorf("version lock not found at %s", paths.VersionLock)
	}

	// Validate version lock
	lock, err := version.LoadVersionLockFromPath(paths.VersionLock)
	if err != nil {
		return fmt.Errorf("version lock is invalid: %w", err)
	}

	// Verify lock has required components
	if _, err := lock.GetComponent("claude-agent-templates"); err != nil {
		return fmt.Errorf("version lock missing claude-agent-templates component: %w", err)
	}

	if _, err := lock.GetComponent("spec-kit"); err != nil {
		return fmt.Errorf("version lock missing spec-kit component: %w", err)
	}

	// Verify Claude Code integration
	if err := VerifyClaudeIntegration(paths); err != nil {
		return fmt.Errorf("Claude Code integration verification failed: %w", err)
	}

	return nil
}

// Uninstall removes an installation
func Uninstall(prefix string, logger *config.Logger) error {
	logger.Warn("uninstall", "Uninstalling claude-agent-templates from %s", prefix)

	// TODO: Implement uninstallation
	// - Remove .specify/ directory
	// - Remove version lock
	// - Remove .claude/agents/cat-* files
	// - Remove .claude/commands/speckit.* files
	// - Add history entry to version lock before removal

	return fmt.Errorf("uninstall not yet implemented")
}

// Status displays the current installation status
type InstallationStatus struct {
	Installed         bool
	Prefix            string
	TemplatesVersion  string
	SpecKitVersion    string
	InstalledAt       string
	LastVerified      string
	InstallationID    string
	HistoryEntryCount int
}

// GetStatus retrieves the current installation status
func GetStatus(prefix string) (*InstallationStatus, error) {
	status := &InstallationStatus{
		Prefix: prefix,
	}

	// Get paths
	paths, err := GetPaths(prefix)
	if err != nil {
		return nil, fmt.Errorf("failed to get installation paths: %w", err)
	}

	// Check if version lock exists
	if !config.PathExists(paths.VersionLock) {
		status.Installed = false
		return status, nil
	}

	// Load version lock
	lock, err := version.LoadVersionLockFromPath(paths.VersionLock)
	if err != nil {
		return nil, fmt.Errorf("failed to load version lock: %w", err)
	}

	status.Installed = true
	status.InstallationID = lock.InstallationID
	status.InstalledAt = lock.InstalledAt
	status.LastVerified = lock.LastVerified
	status.HistoryEntryCount = len(lock.History)

	// Get component versions
	if comp, err := lock.GetComponent("claude-agent-templates"); err == nil {
		status.TemplatesVersion = comp.Version
	}

	if comp, err := lock.GetComponent("spec-kit"); err == nil {
		status.SpecKitVersion = comp.Version
	}

	return status, nil
}

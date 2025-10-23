package install

import (
	"fmt"

	"github.com/dkoenawan/claude-agent-templates/internal/config"
	"github.com/dkoenawan/claude-agent-templates/internal/version"
)

// UpdateOptions contains update configuration
type UpdateOptions struct {
	TargetVersion string // Specific version to update to (empty = latest from manifest)
	Backup        bool   // Create backup before update (default: true)
	Force         bool   // Force update even if versions match
	SkipVerify    bool   // Skip version verification
}

// UpdateResult contains the results of an update operation
type UpdateResult struct {
	Success          bool
	UpdatedFrom      string
	UpdatedTo        string
	BackupCreated    bool
	BackupID         string
	ComponentsUpdated int
	Warnings         []string
}

// Update updates an existing installation to a new version
func Update(prefix string, opts UpdateOptions, logger *config.Logger) (*UpdateResult, error) {
	result := &UpdateResult{
		Warnings: []string{},
	}

	logger.Info("update", "Starting update process...")

	// Get installation paths
	paths, err := GetPaths(prefix)
	if err != nil {
		return nil, fmt.Errorf("failed to get installation paths: %w", err)
	}

	// Check if installation exists
	if !config.PathExists(paths.VersionLock) {
		return nil, fmt.Errorf("no installation found at %s (run 'install' first)", prefix)
	}

	// Load current version lock
	currentLock, err := version.LoadVersionLockFromPath(paths.VersionLock)
	if err != nil {
		return nil, fmt.Errorf("failed to load current version lock: %w", err)
	}

	// Get current versions
	currentTemplatesComp, err := currentLock.GetComponent("spec-kit-agents")
	if err != nil {
		return nil, fmt.Errorf("failed to get current spec-kit-agents version: %w", err)
	}
	result.UpdatedFrom = currentTemplatesComp.Version

	currentSpecKitComp, err := currentLock.GetComponent("spec-kit")
	if err != nil {
		return nil, fmt.Errorf("failed to get current spec-kit version: %w", err)
	}

	logger.Info("update", "Current versions:")
	logger.Info("update", "  spec-kit-agents: v%s", result.UpdatedFrom)
	logger.Info("update", "  spec-kit: v%s", currentSpecKitComp.Version)

	// Load version manifest to get target version
	manifest, err := version.LoadManifestFromPath(".specify/version-manifest.json")
	if err != nil {
		return nil, fmt.Errorf("failed to load version manifest: %w", err)
	}

	targetSpecKitVersion, err := version.GetSpecKitVersion(manifest)
	if err != nil {
		return nil, fmt.Errorf("failed to get target spec-kit version: %w", err)
	}

	targetTemplatesVersion := "2.0.0" // TODO: Get from git tag or version file
	result.UpdatedTo = targetTemplatesVersion

	logger.Info("update", "Target versions:")
	logger.Info("update", "  spec-kit-agents: v%s", targetTemplatesVersion)
	logger.Info("update", "  spec-kit: v%s", targetSpecKitVersion)

	// Check if update is needed
	if result.UpdatedFrom == result.UpdatedTo && currentSpecKitComp.Version == targetSpecKitVersion {
		if !opts.Force {
			logger.Info("update", "Already at target version, no update needed")
			logger.Info("update", "Use --force to reinstall anyway")
			result.Success = true
			return result, nil
		}
		logger.Warn("update", "Forcing update even though versions match")
		result.Warnings = append(result.Warnings, "forced update with matching versions")
	}

	// Check version compatibility
	if !opts.SkipVerify {
		logger.Info("update", "Checking version compatibility...")
		compatibility, err := version.GetSpecKitCompatibility(manifest)
		if err != nil {
			return nil, fmt.Errorf("failed to get compatibility constraints: %w", err)
		}

		compatResult, err := version.CheckCompatibility(
			targetSpecKitVersion,
			targetSpecKitVersion,
			compatibility.MinVersion,
			compatibility.MaxVersion,
			compatibility.BreakingVersions,
		)
		if err != nil {
			return nil, fmt.Errorf("compatibility check failed: %w", err)
		}

		if !compatResult.IsCompatible() {
			return nil, fmt.Errorf("target version incompatible: %s", compatResult.GetIssuesText())
		}

		logger.Success("update", "Version compatibility verified")
	}

	// Create backup if requested
	var backup *BackupInfo
	if opts.Backup {
		logger.Info("update", "Creating backup before update...")
		backup, err = CreateBackup(prefix, logger)
		if err != nil {
			return nil, fmt.Errorf("failed to create backup: %w", err)
		}
		result.BackupCreated = true
		result.BackupID = backup.BackupID

		// Ensure backup is cleaned up on successful update
		defer func() {
			if result.Success && backup != nil {
				// Keep backup for safety, but could optionally clean up
				logger.Info("update", "Backup preserved at: %s", backup.BackupPath)
			}
		}()
	}

	// Perform update installation
	logger.Info("update", "Updating installation...")

	installOpts := Options{
		Prefix: prefix,
		Force:  true, // Always force for update
		Quiet:  false,
		DryRun: false,
	}

	installResult, err := Run(installOpts, logger)
	if err != nil {
		// Automatic rollback on failure
		if backup != nil {
			return nil, AutoRollbackOnError(backup, err, logger)
		}
		return nil, fmt.Errorf("update failed: %w", err)
	}

	if !installResult.Success {
		if backup != nil {
			return nil, AutoRollbackOnError(backup, fmt.Errorf("update did not complete successfully"), logger)
		}
		return nil, fmt.Errorf("update did not complete successfully")
	}

	// Update version lock with upgrade history
	updatedLock, err := version.LoadVersionLockFromPath(paths.VersionLock)
	if err != nil {
		logger.Warn("update", "Failed to load updated version lock: %v", err)
	} else {
		result.ComponentsUpdated = len(updatedLock.Components)
	}

	result.Success = true
	logger.Success("update", "Update completed successfully")
	logger.Info("update", "Updated from v%s to v%s", result.UpdatedFrom, result.UpdatedTo)

	if result.BackupCreated {
		logger.Info("update", "Backup available at: %s", backup.BackupPath)
		logger.Info("update", "To rollback: spec-kit-agents rollback --backup-id=%s", backup.BackupID)
	}

	return result, nil
}

// CheckForUpdates checks if updates are available
func CheckForUpdates(prefix string, logger *config.Logger) (bool, string, error) {
	// Get installation paths
	paths, err := GetPaths(prefix)
	if err != nil {
		return false, "", fmt.Errorf("failed to get installation paths: %w", err)
	}

	// Check if installation exists
	if !config.PathExists(paths.VersionLock) {
		return false, "no installation found", nil
	}

	// Load current version lock
	currentLock, err := version.LoadVersionLockFromPath(paths.VersionLock)
	if err != nil {
		return false, "", fmt.Errorf("failed to load current version lock: %w", err)
	}

	// Get current spec-kit version
	currentSpecKitComp, err := currentLock.GetComponent("spec-kit")
	if err != nil {
		return false, "", fmt.Errorf("failed to get current spec-kit version: %w", err)
	}

	// Load version manifest to get target version
	manifest, err := version.LoadManifestFromPath(".specify/version-manifest.json")
	if err != nil {
		return false, "", fmt.Errorf("failed to load version manifest: %w", err)
	}

	targetSpecKitVersion, err := version.GetSpecKitVersion(manifest)
	if err != nil {
		return false, "", fmt.Errorf("failed to get target spec-kit version: %w", err)
	}

	// Compare versions
	if currentSpecKitComp.Version == targetSpecKitVersion {
		return false, fmt.Sprintf("already at latest version (v%s)", targetSpecKitVersion), nil
	}

	cmp, err := version.CompareVersions(targetSpecKitVersion, currentSpecKitComp.Version)
	if err != nil {
		return false, "", fmt.Errorf("failed to compare versions: %w", err)
	}

	if cmp > 0 {
		return true, fmt.Sprintf("update available: v%s → v%s", currentSpecKitComp.Version, targetSpecKitVersion), nil
	}

	return false, fmt.Sprintf("current version (v%s) is newer than manifest (v%s)", currentSpecKitComp.Version, targetSpecKitVersion), nil
}

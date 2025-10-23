package install

import (
	"fmt"

	"github.com/dkoenawan/claude-agent-templates/internal/config"
	"github.com/dkoenawan/claude-agent-templates/internal/version"
)

// RollbackOptions contains rollback configuration
type RollbackOptions struct {
	BackupID string // Specific backup to restore (empty = latest)
	Force    bool   // Force rollback even if current install seems OK
}

// RollbackResult contains the results of a rollback operation
type RollbackResult struct {
	Success           bool
	RestoredFromID    string
	PreviousVersion   string
	RestoredVersion   string
	ComponentsRestored int
}

// Rollback restores a previous installation from backup
func Rollback(prefix string, opts RollbackOptions, logger *config.Logger) (*RollbackResult, error) {
	result := &RollbackResult{}

	logger.Info("rollback", "Starting rollback process...")

	// Get installation paths
	paths, err := GetPaths(prefix)
	if err != nil {
		return nil, fmt.Errorf("failed to get installation paths: %w", err)
	}

	// Record current version before rollback
	if config.PathExists(paths.VersionLock) {
		lock, err := version.LoadVersionLockFromPath(paths.VersionLock)
		if err != nil {
			logger.Warn("rollback", "Failed to load current version lock: %v", err)
		} else {
			if comp, err := lock.GetComponent("claude-agent-templates"); err == nil {
				result.PreviousVersion = comp.Version
			}
		}
	}

	// Find backup to restore
	var backup *BackupInfo
	if opts.BackupID != "" {
		// Restore specific backup
		backups, err := ListBackups(prefix)
		if err != nil {
			return nil, fmt.Errorf("failed to list backups: %w", err)
		}

		for _, b := range backups {
			if b.BackupID == opts.BackupID {
				backup = b
				break
			}
		}

		if backup == nil {
			return nil, fmt.Errorf("backup not found: %s", opts.BackupID)
		}
	} else {
		// Restore latest backup
		backup, err = GetLatestBackup(prefix)
		if err != nil {
			return nil, fmt.Errorf("failed to find latest backup: %w", err)
		}
	}

	logger.Info("rollback", "Restoring from backup: %s", backup.BackupID)
	logger.Info("rollback", "Created: %s", backup.CreatedAt.Format("2006-01-02 15:04:05 UTC"))

	// Restore backup
	if err := RestoreBackup(backup, logger); err != nil {
		return nil, fmt.Errorf("failed to restore backup: %w", err)
	}

	result.RestoredFromID = backup.BackupID

	// Load restored version lock
	if config.PathExists(paths.VersionLock) {
		lock, err := version.LoadVersionLockFromPath(paths.VersionLock)
		if err != nil {
			logger.Warn("rollback", "Failed to load restored version lock: %v", err)
		} else {
			if comp, err := lock.GetComponent("claude-agent-templates"); err == nil {
				result.RestoredVersion = comp.Version
			}

			// Update version lock with rollback event
			lock.AddHistoryEntry("rollback", "all", result.RestoredVersion, "success", nil)
			if err := version.SaveVersionLock(lock, paths.VersionLock); err != nil {
				logger.Warn("rollback", "Failed to update version lock: %v", err)
			}

			result.ComponentsRestored = len(lock.Components)
		}
	}

	result.Success = true
	logger.Success("rollback", "Rollback completed successfully")
	logger.Info("rollback", "Restored version: %s", result.RestoredVersion)

	return result, nil
}

// CanRollback checks if rollback is possible
func CanRollback(prefix string) (bool, string, error) {
	// Check if any backups exist
	backups, err := ListBackups(prefix)
	if err != nil {
		return false, "", fmt.Errorf("failed to list backups: %w", err)
	}

	if len(backups) == 0 {
		return false, "no backups available", nil
	}

	// Get latest backup
	latest, err := GetLatestBackup(prefix)
	if err != nil {
		return false, "", err
	}

	message := fmt.Sprintf("can rollback to backup %s (created %s)",
		latest.BackupID, latest.CreatedAt.Format("2006-01-02 15:04:05 UTC"))

	return true, message, nil
}

// AutoRollbackOnError is a helper to automatically rollback on installation failure
func AutoRollbackOnError(backup *BackupInfo, err error, logger *config.Logger) error {
	if err == nil || backup == nil {
		return err
	}

	logger.Error("rollback", "Installation failed: %v", err)
	logger.Warn("rollback", "Attempting automatic rollback...")

	if restoreErr := RestoreBackup(backup, logger); restoreErr != nil {
		logger.Error("rollback", "Automatic rollback failed: %v", restoreErr)
		return fmt.Errorf("installation failed and rollback also failed: %w (rollback error: %v)", err, restoreErr)
	}

	logger.Success("rollback", "Automatic rollback completed")
	return fmt.Errorf("installation failed but was rolled back successfully: %w", err)
}

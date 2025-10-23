package install

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dkoenawan/claude-agent-templates/internal/config"
)

// BackupInfo contains information about a backup
type BackupInfo struct {
	BackupPath    string
	OriginalPath  string
	CreatedAt     time.Time
	BackupID      string
	ComponentName string
}

// CreateBackup creates a backup of an existing installation
func CreateBackup(installPath string, logger *config.Logger) (*BackupInfo, error) {
	// Check if installation exists
	if !config.PathExists(installPath) {
		return nil, fmt.Errorf("installation path does not exist: %s", installPath)
	}

	// Generate backup ID and path
	timestamp := time.Now().UTC().Format("20060102-150405")
	backupID := fmt.Sprintf("backup-%s", timestamp)
	backupPath := installPath + "." + backupID

	logger.Info("backup", "Creating backup of %s...", installPath)
	logger.Debug("backup", "Backup destination: %s", backupPath)

	// Copy installation to backup location
	if err := CopyDirectory(installPath, backupPath); err != nil {
		return nil, fmt.Errorf("failed to create backup: %w", err)
	}

	// Get backup size
	size, err := GetDirectorySize(backupPath)
	if err != nil {
		logger.Warn("backup", "Failed to calculate backup size: %v", err)
	} else {
		logger.Info("backup", "Backup size: %s", FormatSize(size))
	}

	info := &BackupInfo{
		BackupPath:    backupPath,
		OriginalPath:  installPath,
		CreatedAt:     time.Now().UTC(),
		BackupID:      backupID,
		ComponentName: "claude-agent-templates",
	}

	logger.Success("backup", "Backup created: %s", backupPath)

	return info, nil
}

// RestoreBackup restores an installation from a backup
func RestoreBackup(backup *BackupInfo, logger *config.Logger) error {
	logger.Info("backup", "Restoring from backup: %s", backup.BackupID)

	// Verify backup exists
	if !config.PathExists(backup.BackupPath) {
		return fmt.Errorf("backup does not exist: %s", backup.BackupPath)
	}

	// Remove current installation if it exists
	if config.PathExists(backup.OriginalPath) {
		logger.Debug("backup", "Removing current installation...")
		if err := os.RemoveAll(backup.OriginalPath); err != nil {
			return fmt.Errorf("failed to remove current installation: %w", err)
		}
	}

	// Restore backup
	logger.Debug("backup", "Copying backup to original location...")
	if err := CopyDirectory(backup.BackupPath, backup.OriginalPath); err != nil {
		return fmt.Errorf("failed to restore backup: %w", err)
	}

	logger.Success("backup", "Backup restored successfully")

	return nil
}

// CleanupBackup removes a backup
func CleanupBackup(backup *BackupInfo, logger *config.Logger) error {
	if backup == nil || backup.BackupPath == "" {
		return nil
	}

	logger.Debug("backup", "Cleaning up backup: %s", backup.BackupID)

	if config.PathExists(backup.BackupPath) {
		if err := os.RemoveAll(backup.BackupPath); err != nil {
			return fmt.Errorf("failed to cleanup backup: %w", err)
		}
		logger.Info("backup", "Backup cleaned up: %s", backup.BackupID)
	}

	return nil
}

// ListBackups finds all backups for a given installation path
func ListBackups(installPath string) ([]*BackupInfo, error) {
	// Get parent directory
	parentDir := filepath.Dir(installPath)
	baseName := filepath.Base(installPath)

	// Read directory
	entries, err := os.ReadDir(parentDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	backups := []*BackupInfo{}

	// Find backup directories
	backupPrefix := baseName + ".backup-"
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		if len(name) <= len(backupPrefix) {
			continue
		}

		if name[:len(backupPrefix)] == backupPrefix {
			backupID := name[len(baseName)+1:] // Remove base name and dot

			// Get directory info
			backupPath := filepath.Join(parentDir, name)
			info, err := os.Stat(backupPath)
			if err != nil {
				continue
			}

			backups = append(backups, &BackupInfo{
				BackupPath:    backupPath,
				OriginalPath:  installPath,
				CreatedAt:     info.ModTime(),
				BackupID:      backupID,
				ComponentName: "claude-agent-templates",
			})
		}
	}

	return backups, nil
}

// GetLatestBackup returns the most recent backup
func GetLatestBackup(installPath string) (*BackupInfo, error) {
	backups, err := ListBackups(installPath)
	if err != nil {
		return nil, err
	}

	if len(backups) == 0 {
		return nil, fmt.Errorf("no backups found")
	}

	// Find most recent backup
	latest := backups[0]
	for _, backup := range backups[1:] {
		if backup.CreatedAt.After(latest.CreatedAt) {
			latest = backup
		}
	}

	return latest, nil
}

// CleanupOldBackups removes backups older than a certain duration
func CleanupOldBackups(installPath string, keepDuration time.Duration, logger *config.Logger) error {
	backups, err := ListBackups(installPath)
	if err != nil {
		return err
	}

	cutoffTime := time.Now().UTC().Add(-keepDuration)
	removedCount := 0

	for _, backup := range backups {
		if backup.CreatedAt.Before(cutoffTime) {
			logger.Debug("backup", "Removing old backup: %s (created %s)", backup.BackupID, backup.CreatedAt.Format(time.RFC3339))
			if err := CleanupBackup(backup, logger); err != nil {
				logger.Warn("backup", "Failed to cleanup backup %s: %v", backup.BackupID, err)
			} else {
				removedCount++
			}
		}
	}

	if removedCount > 0 {
		logger.Info("backup", "Cleaned up %d old backup(s)", removedCount)
	}

	return nil
}

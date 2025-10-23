package version

import (
	"fmt"
	"os"

	"github.com/dkoenawan/claude-agent-templates/pkg/models"
)

// CreateVersionLock creates a new version lock with installation details
func CreateVersionLock(templatesVersion, specKitVersion, installPath string) *models.VersionLock {
	lock := models.NewVersionLock()

	// Set spec-kit-agents component
	lock.SetComponent("spec-kit-agents", models.Component{
		Version:       templatesVersion,
		InstalledFrom: "git",
		InstallPath:   installPath,
	})

	// Set spec-kit component
	lock.SetComponent("spec-kit", models.Component{
		Version:       specKitVersion,
		InstalledFrom: "vendored",
		InstallPath:   installPath + "/.specify",
	})

	// Add initial history entry
	lock.AddHistoryEntry("install", "all", templatesVersion, "success", nil)

	return lock
}

// LoadVersionLockFromPath loads a version lock from the specified path
func LoadVersionLockFromPath(path string) (*models.VersionLock, error) {
	lock, err := models.LoadVersionLock(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load version lock from %s: %w", path, err)
	}
	return lock, nil
}

// SaveVersionLock saves a version lock to the specified path
func SaveVersionLock(lock *models.VersionLock, path string) error {
	if err := lock.Save(path); err != nil {
		return fmt.Errorf("failed to save version lock to %s: %w", path, err)
	}
	return nil
}

// GetInstalledSpecKitVersion retrieves the installed spec-kit version from lock
func GetInstalledSpecKitVersion(lock *models.VersionLock) (string, error) {
	comp, err := lock.GetComponent("spec-kit")
	if err != nil {
		return "", err
	}
	return comp.Version, nil
}

// UpdateVersionLock updates an existing version lock with new installation info
func UpdateVersionLock(lock *models.VersionLock, templatesVersion, specKitVersion string) {
	// Update components
	if comp, err := lock.GetComponent("spec-kit-agents"); err == nil {
		comp.Version = templatesVersion
		lock.SetComponent("spec-kit-agents", *comp)
	}

	if comp, err := lock.GetComponent("spec-kit"); err == nil {
		comp.Version = specKitVersion
		lock.SetComponent("spec-kit", *comp)
	}

	// Update verification time
	lock.UpdateVerificationTime()

	// Add history entry
	lock.AddHistoryEntry("upgrade", "all", templatesVersion, "success", nil)
}

// CheckVersionLockExists checks if a version lock file exists
func CheckVersionLockExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

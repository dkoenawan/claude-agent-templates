package version

import (
	"fmt"

	"github.com/dkoenawan/claude-agent-templates/internal/config"
	"github.com/dkoenawan/claude-agent-templates/pkg/models"
)

// LoadManifestFromPath loads the version manifest from a specified path
func LoadManifestFromPath(path string) (*models.Manifest, error) {
	manifest, err := models.LoadManifest(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load manifest from %s: %w", path, err)
	}
	return manifest, nil
}

// LoadManifestFromPrefix loads the version manifest from installation prefix
func LoadManifestFromPrefix(prefix string) (*models.Manifest, error) {
	manifestPath := config.GetVersionManifestPath(prefix)
	return LoadManifestFromPath(manifestPath)
}

// GetSpecKitVersion retrieves the pinned spec-kit version from manifest
func GetSpecKitVersion(manifest *models.Manifest) (string, error) {
	dep, err := manifest.GetSpecKitDependency()
	if err != nil {
		return "", err
	}
	return dep.Version, nil
}

// GetSpecKitCompatibility retrieves the compatibility constraints for spec-kit
func GetSpecKitCompatibility(manifest *models.Manifest) (*models.Compatibility, error) {
	dep, err := manifest.GetSpecKitDependency()
	if err != nil {
		return nil, err
	}
	return &dep.Compatibility, nil
}

// ValidateManifestIntegrity validates the manifest against expected values
func ValidateManifestIntegrity(manifest *models.Manifest) error {
	// Ensure name is correct
	if manifest.Name != "claude-agent-templates" {
		return fmt.Errorf("invalid manifest name: %s (expected claude-agent-templates)", manifest.Name)
	}

	// Ensure spec-kit dependency exists
	_, err := manifest.GetSpecKitDependency()
	if err != nil {
		return fmt.Errorf("manifest validation failed: %w", err)
	}

	return nil
}

package models

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"
)

// Manifest represents the version manifest for claude-agent-templates
type Manifest struct {
	Version      string                 `json:"version"`
	Name         string                 `json:"name"`
	Dependencies map[string]Dependency  `json:"dependencies"`
	UpdatePolicy string                 `json:"update_policy,omitempty"`
	LastUpdated  string                 `json:"last_updated,omitempty"`
}

// Dependency represents a dependency in the version manifest
type Dependency struct {
	Version       string        `json:"version"`
	Source        string        `json:"source"`
	InstallPath   string        `json:"install_path"`
	Integrity     string        `json:"integrity,omitempty"`
	Compatibility Compatibility `json:"compatibility,omitempty"`
}

// Compatibility defines version compatibility constraints
type Compatibility struct {
	MinVersion       string   `json:"min_version,omitempty"`
	MaxVersion       string   `json:"max_version,omitempty"`
	BreakingVersions []string `json:"breaking_versions,omitempty"`
}

// LoadManifest loads a version manifest from a JSON file
func LoadManifest(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest: %w", err)
	}

	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("failed to parse manifest JSON: %w", err)
	}

	// Validate manifest
	if err := manifest.Validate(); err != nil {
		return nil, fmt.Errorf("manifest validation failed: %w", err)
	}

	return &manifest, nil
}

// SaveManifest saves a version manifest to a JSON file
func (m *Manifest) SaveManifest(path string) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal manifest: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write manifest: %w", err)
	}

	return nil
}

// Validate checks if the manifest is valid according to the schema
func (m *Manifest) Validate() error {
	// Validate version format (e.g., "1.0")
	versionPattern := regexp.MustCompile(`^[0-9]+\.[0-9]+$`)
	if !versionPattern.MatchString(m.Version) {
		return fmt.Errorf("invalid manifest version format: %s (expected X.Y)", m.Version)
	}

	// Validate name
	if m.Name == "" {
		return fmt.Errorf("manifest name is required")
	}

	// Validate dependencies
	if len(m.Dependencies) == 0 {
		return fmt.Errorf("manifest must have at least one dependency")
	}

	// Validate each dependency
	for name, dep := range m.Dependencies {
		if err := dep.Validate(name); err != nil {
			return fmt.Errorf("dependency %s: %w", name, err)
		}
	}

	// Validate update policy
	if m.UpdatePolicy != "" {
		validPolicies := map[string]bool{"manual": true, "patch": true, "minor": true}
		if !validPolicies[m.UpdatePolicy] {
			return fmt.Errorf("invalid update_policy: %s (must be manual, patch, or minor)", m.UpdatePolicy)
		}
	}

	// Validate last_updated date format if present
	if m.LastUpdated != "" {
		if _, err := time.Parse("2006-01-02", m.LastUpdated); err != nil {
			return fmt.Errorf("invalid last_updated date format: %s (expected YYYY-MM-DD)", m.LastUpdated)
		}
	}

	return nil
}

// Validate checks if a dependency is valid
func (d *Dependency) Validate(name string) error {
	// Validate version (semver format: X.Y.Z)
	semverPattern := regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`)
	if !semverPattern.MatchString(d.Version) {
		return fmt.Errorf("invalid version format: %s (expected X.Y.Z)", d.Version)
	}

	// Validate source
	validSources := map[string]bool{"vendored": true, "git": true, "npm": true}
	if !validSources[d.Source] {
		return fmt.Errorf("invalid source: %s (must be vendored, git, or npm)", d.Source)
	}

	// Validate install path
	if d.InstallPath == "" {
		return fmt.Errorf("install_path is required")
	}

	// Validate integrity hash format if present (sha256-[64 hex chars])
	if d.Integrity != "" {
		integrityPattern := regexp.MustCompile(`^sha256-[a-f0-9]{64}$`)
		if !integrityPattern.MatchString(d.Integrity) {
			return fmt.Errorf("invalid integrity hash format: %s (expected sha256-[64 hex chars])", d.Integrity)
		}
	}

	// Validate compatibility constraints
	if d.Compatibility.MinVersion != "" {
		if !semverPattern.MatchString(d.Compatibility.MinVersion) {
			return fmt.Errorf("invalid min_version format: %s", d.Compatibility.MinVersion)
		}
	}
	if d.Compatibility.MaxVersion != "" {
		if !semverPattern.MatchString(d.Compatibility.MaxVersion) {
			return fmt.Errorf("invalid max_version format: %s", d.Compatibility.MaxVersion)
		}
	}
	for _, v := range d.Compatibility.BreakingVersions {
		if !semverPattern.MatchString(v) {
			return fmt.Errorf("invalid breaking_version format: %s", v)
		}
	}

	return nil
}

// GetSpecKitDependency is a convenience method to get the spec-kit dependency
func (m *Manifest) GetSpecKitDependency() (*Dependency, error) {
	dep, exists := m.Dependencies["spec-kit"]
	if !exists {
		return nil, fmt.Errorf("spec-kit dependency not found in manifest")
	}
	return &dep, nil
}

package models

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/google/uuid"
)

// VersionLock represents the installed component versions and history
type VersionLock struct {
	Version        string              `json:"version"`
	InstallationID string              `json:"installation_id"`
	InstalledAt    string              `json:"installed_at"`
	LastVerified   string              `json:"last_verified,omitempty"`
	Components     map[string]Component `json:"components"`
	History        []HistoryEntry      `json:"history,omitempty"`
}

// Component represents an installed component
type Component struct {
	Version       string `json:"version"`
	InstalledFrom string `json:"installed_from"`
	Commit        string `json:"commit,omitempty"`
	InstallPath   string `json:"install_path"`
}

// HistoryEntry represents a single installation/upgrade event
type HistoryEntry struct {
	Timestamp string `json:"timestamp"`
	Action    string `json:"action"`
	Component string `json:"component"`
	Version   string `json:"version,omitempty"`
	Status    string `json:"status"`
	Error     string `json:"error,omitempty"`
}

// NewVersionLock creates a new version lock with a unique installation ID
func NewVersionLock() *VersionLock {
	now := time.Now().UTC().Format(time.RFC3339)
	return &VersionLock{
		Version:        "1.0",
		InstallationID: uuid.New().String(),
		InstalledAt:    now,
		LastVerified:   now,
		Components:     make(map[string]Component),
		History:        []HistoryEntry{},
	}
}

// LoadVersionLock loads a version lock from a JSON file
func LoadVersionLock(path string) (*VersionLock, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read version lock: %w", err)
	}

	var lock VersionLock
	if err := json.Unmarshal(data, &lock); err != nil {
		return nil, fmt.Errorf("failed to parse version lock JSON: %w", err)
	}

	// Validate lock
	if err := lock.Validate(); err != nil {
		return nil, fmt.Errorf("version lock validation failed: %w", err)
	}

	return &lock, nil
}

// Save saves the version lock to a JSON file
func (vl *VersionLock) Save(path string) error {
	data, err := json.MarshalIndent(vl, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal version lock: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write version lock: %w", err)
	}

	return nil
}

// Validate checks if the version lock is valid
func (vl *VersionLock) Validate() error {
	// Validate version format
	versionPattern := regexp.MustCompile(`^[0-9]+\.[0-9]+$`)
	if !versionPattern.MatchString(vl.Version) {
		return fmt.Errorf("invalid lock version format: %s (expected X.Y)", vl.Version)
	}

	// Validate installation ID (UUID format)
	if _, err := uuid.Parse(vl.InstallationID); err != nil {
		return fmt.Errorf("invalid installation_id format: %s (expected UUID)", vl.InstallationID)
	}

	// Validate installed_at timestamp
	if _, err := time.Parse(time.RFC3339, vl.InstalledAt); err != nil {
		return fmt.Errorf("invalid installed_at timestamp: %s (expected RFC3339)", vl.InstalledAt)
	}

	// Validate last_verified timestamp if present
	if vl.LastVerified != "" {
		if _, err := time.Parse(time.RFC3339, vl.LastVerified); err != nil {
			return fmt.Errorf("invalid last_verified timestamp: %s (expected RFC3339)", vl.LastVerified)
		}
	}

	// Validate components
	if len(vl.Components) == 0 {
		return fmt.Errorf("version lock must have at least one component")
	}

	for name, comp := range vl.Components {
		if err := comp.Validate(name); err != nil {
			return fmt.Errorf("component %s: %w", name, err)
		}
	}

	// Validate history entries
	for i, entry := range vl.History {
		if err := entry.Validate(); err != nil {
			return fmt.Errorf("history entry %d: %w", i, err)
		}
	}

	return nil
}

// Validate checks if a component is valid
func (c *Component) Validate(name string) error {
	// Validate version (semver format)
	semverPattern := regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`)
	if !semverPattern.MatchString(c.Version) {
		return fmt.Errorf("invalid version format: %s (expected X.Y.Z)", c.Version)
	}

	// Validate installed_from
	validSources := map[string]bool{"git": true, "archive": true, "manual": true, "vendored": true, "npm": true}
	if !validSources[c.InstalledFrom] {
		return fmt.Errorf("invalid installed_from: %s", c.InstalledFrom)
	}

	// Validate commit hash format if present (7-40 hex chars)
	if c.Commit != "" {
		commitPattern := regexp.MustCompile(`^[a-f0-9]{7,40}$`)
		if !commitPattern.MatchString(c.Commit) {
			return fmt.Errorf("invalid commit hash format: %s", c.Commit)
		}
	}

	// Validate install path
	if c.InstallPath == "" {
		return fmt.Errorf("install_path is required")
	}

	return nil
}

// Validate checks if a history entry is valid
func (he *HistoryEntry) Validate() error {
	// Validate timestamp
	if _, err := time.Parse(time.RFC3339, he.Timestamp); err != nil {
		return fmt.Errorf("invalid timestamp: %s (expected RFC3339)", he.Timestamp)
	}

	// Validate action
	validActions := map[string]bool{"install": true, "upgrade": true, "verify": true, "rollback": true}
	if !validActions[he.Action] {
		return fmt.Errorf("invalid action: %s (must be install, upgrade, verify, or rollback)", he.Action)
	}

	// Validate component
	validComponents := map[string]bool{"claude-agent-templates": true, "spec-kit": true, "all": true}
	if !validComponents[he.Component] {
		return fmt.Errorf("invalid component: %s", he.Component)
	}

	// Validate version if present
	if he.Version != "" {
		semverPattern := regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`)
		if !semverPattern.MatchString(he.Version) {
			return fmt.Errorf("invalid version format: %s (expected X.Y.Z)", he.Version)
		}
	}

	// Validate status
	validStatuses := map[string]bool{"success": true, "failure": true, "partial": true}
	if !validStatuses[he.Status] {
		return fmt.Errorf("invalid status: %s (must be success, failure, or partial)", he.Status)
	}

	return nil
}

// AddHistoryEntry adds a new entry to the installation history
func (vl *VersionLock) AddHistoryEntry(action, component, version, status string, err error) {
	entry := HistoryEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Action:    action,
		Component: component,
		Version:   version,
		Status:    status,
	}

	if err != nil {
		entry.Error = err.Error()
	}

	vl.History = append(vl.History, entry)
}

// UpdateVerificationTime updates the last_verified timestamp
func (vl *VersionLock) UpdateVerificationTime() {
	vl.LastVerified = time.Now().UTC().Format(time.RFC3339)
}

// GetComponent is a convenience method to get a specific component
func (vl *VersionLock) GetComponent(name string) (*Component, error) {
	comp, exists := vl.Components[name]
	if !exists {
		return nil, fmt.Errorf("component %s not found in version lock", name)
	}
	return &comp, nil
}

// SetComponent sets or updates a component in the version lock
func (vl *VersionLock) SetComponent(name string, comp Component) {
	if vl.Components == nil {
		vl.Components = make(map[string]Component)
	}
	vl.Components[name] = comp
}

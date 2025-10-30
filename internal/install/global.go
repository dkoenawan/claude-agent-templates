package install

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dkoenawan/claude-agent-templates/internal/config"
	"github.com/dkoenawan/claude-agent-templates/pkg/models"
	"github.com/google/uuid"
)

// GlobalInstallation represents a global installation in ~/.claude/
type GlobalInstallation struct {
	InstallationID   string
	ClaudeDir        string
	AgentsDir        string
	CommandsDir      string
	SkillsDir        string
	VersionLockPath  string
	SourceType       string
	Provider         SourceFileProvider
}

// NewGlobalInstallation creates a new GlobalInstallation instance
func NewGlobalInstallation(provider SourceFileProvider) (*GlobalInstallation, error) {
	claudeDir, err := config.GetClaudeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get Claude directory: %w", err)
	}

	return &GlobalInstallation{
		InstallationID:  uuid.New().String(),
		ClaudeDir:       claudeDir,
		AgentsDir:       filepath.Join(claudeDir, "agents"),
		CommandsDir:     filepath.Join(claudeDir, "commands"),
		SkillsDir:       filepath.Join(claudeDir, "skills"),
		VersionLockPath: filepath.Join(claudeDir, ".version-lock.json"),
		SourceType:      provider.GetSourceType(),
		Provider:        provider,
	}, nil
}

// CreateDirectories creates the ~/.claude/ directory structure
func (gi *GlobalInstallation) CreateDirectories() error {
	dirs := []string{
		gi.ClaudeDir,
		gi.AgentsDir,
		gi.CommandsDir,
		gi.SkillsDir,
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return nil
}

// Exists checks if a global installation already exists
func (gi *GlobalInstallation) Exists() bool {
	_, err := os.Stat(gi.VersionLockPath)
	return err == nil
}

// LoadVersionLock loads the existing version lock
func (gi *GlobalInstallation) LoadVersionLock() (*models.VersionLock, error) {
	return models.LoadVersionLock(gi.VersionLockPath)
}

// SaveVersionLock saves the version lock
func (gi *GlobalInstallation) SaveVersionLock(lock *models.VersionLock) error {
	return lock.Save(gi.VersionLockPath)
}

// GetInstallationInfo returns information about the global installation
func (gi *GlobalInstallation) GetInstallationInfo() map[string]string {
	return map[string]string{
		"installation_type": "global",
		"claude_dir":        gi.ClaudeDir,
		"agents_dir":        gi.AgentsDir,
		"commands_dir":      gi.CommandsDir,
		"skills_dir":        gi.SkillsDir,
		"version_lock":      gi.VersionLockPath,
		"source_type":       gi.SourceType,
	}
}

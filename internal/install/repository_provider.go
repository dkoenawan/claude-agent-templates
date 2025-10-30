package install

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// RepositorySourceProvider implements SourceFileProvider using local repository files
type RepositorySourceProvider struct {
	repositoryRoot string
}

// NewRepositorySourceProvider creates a new RepositorySourceProvider
// repositoryRoot should be the absolute path to the repository root
func NewRepositorySourceProvider(repositoryRoot string) *RepositorySourceProvider {
	return &RepositorySourceProvider{
		repositoryRoot: repositoryRoot,
	}
}

// GetAgents returns a filesystem containing agent files from the repository
func (p *RepositorySourceProvider) GetAgents() (fs.FS, error) {
	agentsPath := filepath.Join(p.repositoryRoot, "agents")
	if _, err := os.Stat(agentsPath); err != nil {
		return nil, fmt.Errorf("repository agents directory not found: %w", err)
	}
	return os.DirFS(agentsPath), nil
}

// GetSpecify returns a filesystem containing .specify/ directory from the repository
func (p *RepositorySourceProvider) GetSpecify() (fs.FS, error) {
	specifyPath := filepath.Join(p.repositoryRoot, ".specify")
	if _, err := os.Stat(specifyPath); err != nil {
		return nil, fmt.Errorf("repository .specify directory not found: %w", err)
	}
	return os.DirFS(specifyPath), nil
}

// GetCommands returns a filesystem containing command files from the repository
func (p *RepositorySourceProvider) GetCommands() (fs.FS, error) {
	commandsPath := filepath.Join(p.repositoryRoot, ".specify", "templates", "commands")
	if _, err := os.Stat(commandsPath); err != nil {
		return nil, fmt.Errorf("repository commands directory not found: %w", err)
	}
	return os.DirFS(commandsPath), nil
}

// GetSourceType returns "repository"
func (p *RepositorySourceProvider) GetSourceType() string {
	return "repository"
}

// Validate checks if repository source files are accessible
func (p *RepositorySourceProvider) Validate() error {
	// Check if repository root exists
	if _, err := os.Stat(p.repositoryRoot); err != nil {
		return fmt.Errorf("repository root not accessible: %w", err)
	}

	// Validate each required directory
	if _, err := p.GetAgents(); err != nil {
		return fmt.Errorf("repository agents validation failed: %w", err)
	}

	if _, err := p.GetSpecify(); err != nil {
		return fmt.Errorf("repository .specify validation failed: %w", err)
	}

	if _, err := p.GetCommands(); err != nil {
		return fmt.Errorf("repository commands validation failed: %w", err)
	}

	return nil
}

// GetRepositoryRoot returns the repository root path
func (p *RepositorySourceProvider) GetRepositoryRoot() string {
	return p.repositoryRoot
}

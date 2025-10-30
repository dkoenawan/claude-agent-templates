package embed

import (
	"fmt"
	"io/fs"
)

// EmbeddedSourceProvider implements SourceFileProvider using embedded files
type EmbeddedSourceProvider struct {
	embedded *EmbeddedFiles
}

// NewSourceProvider creates a new EmbeddedSourceProvider
func NewSourceProvider() *EmbeddedSourceProvider {
	return &EmbeddedSourceProvider{
		embedded: New(),
	}
}

// GetAgents returns the embedded agents filesystem
func (p *EmbeddedSourceProvider) GetAgents() (fs.FS, error) {
	return p.embedded.GetAgents()
}

// GetSpecify returns the embedded .specify filesystem
func (p *EmbeddedSourceProvider) GetSpecify() (fs.FS, error) {
	return p.embedded.GetSpecify()
}

// GetCommands returns the embedded commands filesystem
func (p *EmbeddedSourceProvider) GetCommands() (fs.FS, error) {
	return p.embedded.GetCommands()
}

// GetSourceType returns "embedded"
func (p *EmbeddedSourceProvider) GetSourceType() string {
	return "embedded"
}

// Validate checks if embedded files are accessible
func (p *EmbeddedSourceProvider) Validate() error {
	// Try to access each filesystem to ensure they exist
	if _, err := p.GetAgents(); err != nil {
		return fmt.Errorf("embedded agents validation failed: %w", err)
	}

	if _, err := p.GetSpecify(); err != nil {
		return fmt.Errorf("embedded .specify validation failed: %w", err)
	}

	if _, err := p.GetCommands(); err != nil {
		return fmt.Errorf("embedded commands validation failed: %w", err)
	}

	return nil
}

// GetEmbeddedFiles returns the underlying EmbeddedFiles instance for metadata access
func (p *EmbeddedSourceProvider) GetEmbeddedFiles() *EmbeddedFiles {
	return p.embedded
}

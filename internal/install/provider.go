package install

import "io/fs"

// SourceFileProvider is an interface for accessing source files from different origins
// Implementations include EmbeddedSourceProvider (from binary), RepositorySourceProvider (from local repo),
// and DownloadedSourceProvider (from downloaded package)
type SourceFileProvider interface {
	// GetAgents returns a filesystem containing agent markdown files
	GetAgents() (fs.FS, error)

	// GetSpecify returns a filesystem containing .specify/ directory contents
	GetSpecify() (fs.FS, error)

	// GetCommands returns a filesystem containing slash command files
	GetCommands() (fs.FS, error)

	// GetSourceType returns the type of source: "embedded", "repository", or "downloaded"
	GetSourceType() string

	// Validate checks if the source files are complete and valid
	Validate() error
}

package embed

import (
	"crypto/sha256"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"time"
)

// Embed all agent and spec-kit files at build time
//
//go:embed ../../agents ../../.specify
var embeddedFiles embed.FS

// Build-time injected variables (via -ldflags)
var (
	Version   = "unknown"
	BuildTime = "unknown"
	Commit    = "unknown"
)

// EmbeddedFiles provides access to files embedded in the binary at build time
type EmbeddedFiles struct {
	files     embed.FS
	checksums map[string]string
}

// New creates a new EmbeddedFiles instance
func New() *EmbeddedFiles {
	return &EmbeddedFiles{
		files:     embeddedFiles,
		checksums: make(map[string]string),
	}
}

// GetAgents returns a filesystem containing agent definitions
// Returns a sub-filesystem rooted at the agents/ directory
func (e *EmbeddedFiles) GetAgents() (fs.FS, error) {
	agentsFS, err := fs.Sub(e.files, "../../agents")
	if err != nil {
		return nil, fmt.Errorf("embedded agents directory not found: %w", err)
	}
	return agentsFS, nil
}

// GetSpecify returns a filesystem containing spec-kit files
// Returns a sub-filesystem rooted at the .specify/ directory
func (e *EmbeddedFiles) GetSpecify() (fs.FS, error) {
	specifyFS, err := fs.Sub(e.files, "../../.specify")
	if err != nil {
		return nil, fmt.Errorf("embedded .specify directory not found: %w", err)
	}
	return specifyFS, nil
}

// GetCommands returns a filesystem containing slash command files
// Returns command files from .specify/templates/commands/
func (e *EmbeddedFiles) GetCommands() (fs.FS, error) {
	commandsFS, err := fs.Sub(e.files, "../../.specify/templates/commands")
	if err != nil {
		return nil, fmt.Errorf("embedded commands directory not found: %w", err)
	}
	return commandsFS, nil
}

// GetVersion returns the version of embedded files injected at build time
func (e *EmbeddedFiles) GetVersion() string {
	return Version
}

// GetBuildTime returns the build timestamp
// Returns the time when the binary was built
func (e *EmbeddedFiles) GetBuildTime() time.Time {
	t, err := time.Parse(time.RFC3339, BuildTime)
	if err != nil {
		// Return zero time if build time wasn't injected or is invalid
		return time.Time{}
	}
	return t
}

// GetChecksum returns the SHA256 checksum for a specific file
// path should be relative (e.g., "agents/core/requirements-analyst.md")
func (e *EmbeddedFiles) GetChecksum(path string) (string, error) {
	// Normalize path
	path = filepath.ToSlash(filepath.Clean(path))

	// Check cache first
	if checksum, exists := e.checksums[path]; exists {
		return checksum, nil
	}

	// Compute checksum
	fullPath := filepath.Join("../..", path)
	file, err := e.files.Open(fullPath)
	if err != nil {
		return "", fmt.Errorf("file not found in embedded files: %s", path)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("failed to compute checksum: %w", err)
	}

	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	e.checksums[path] = checksum
	return checksum, nil
}

// ListFiles returns a list of all embedded file paths
// If prefix is provided, only returns paths starting with that prefix
func (e *EmbeddedFiles) ListFiles(prefix string) ([]string, error) {
	var files []string

	// Normalize prefix
	if prefix != "" {
		prefix = filepath.ToSlash(filepath.Clean(prefix))
	}

	err := fs.WalkDir(e.files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Normalize path
		path = filepath.ToSlash(path)

		// Remove leading ../../ from path
		path = strings.TrimPrefix(path, "../../")

		// Filter by prefix if provided
		if prefix != "" && !strings.HasPrefix(path, prefix) {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to list embedded files: %w", err)
	}

	return files, nil
}

// GetMetadata returns metadata about the embedded files
func (e *EmbeddedFiles) GetMetadata() (*Metadata, error) {
	files, err := e.ListFiles("")
	if err != nil {
		return nil, err
	}

	var totalSize int64
	for _, path := range files {
		info, err := fs.Stat(e.files, filepath.Join("../..", path))
		if err != nil {
			continue
		}
		totalSize += info.Size()
	}

	return &Metadata{
		Version:    e.GetVersion(),
		BuildTime:  e.GetBuildTime(),
		FileCount:  len(files),
		TotalSize:  totalSize,
		Commit:     Commit,
	}, nil
}

// Metadata contains information about embedded files
type Metadata struct {
	Version    string
	BuildTime  time.Time
	FileCount  int
	TotalSize  int64
	Commit     string
}

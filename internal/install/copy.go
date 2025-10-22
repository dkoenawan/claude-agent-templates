package install

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/dkoenawan/claude-agent-templates/internal/config"
)

// CopyFile copies a single file from src to dst
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", src, err)
	}
	defer sourceFile.Close()

	// Ensure destination directory exists
	destDir := filepath.Dir(dst)
	if err := config.EnsureDir(destDir); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", dst, err)
	}
	defer destFile.Close()

	// Copy contents
	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return fmt.Errorf("failed to copy file contents: %w", err)
	}

	// Copy permissions
	sourceInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat source file: %w", err)
	}

	if err := os.Chmod(dst, sourceInfo.Mode()); err != nil {
		return fmt.Errorf("failed to set file permissions: %w", err)
	}

	return nil
}

// CopyDirectory recursively copies a directory from src to dst
func CopyDirectory(src, dst string) error {
	// Get source directory info
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat source directory %s: %w", src, err)
	}

	// Create destination directory
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return fmt.Errorf("failed to create destination directory %s: %w", dst, err)
	}

	// Read directory entries
	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", src, err)
	}

	// Copy each entry
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// Recursively copy subdirectory
			if err := CopyDirectory(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			// Copy file
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// CopyWithPrefix copies files from srcDir to dstDir, adding a prefix to each filename
func CopyWithPrefix(srcDir, dstDir, prefix string) error {
	// Ensure destination directory exists
	if err := config.EnsureDir(dstDir); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Read source directory
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("failed to read source directory %s: %w", srcDir, err)
	}

	// Copy each file with prefix
	for _, entry := range entries {
		if entry.IsDir() {
			continue // Skip directories, only copy files
		}

		srcPath := filepath.Join(srcDir, entry.Name())
		dstFileName := prefix + entry.Name()
		dstPath := filepath.Join(dstDir, dstFileName)

		if err := CopyFile(srcPath, dstPath); err != nil {
			return fmt.Errorf("failed to copy %s to %s: %w", srcPath, dstPath, err)
		}
	}

	return nil
}

// CopyAgentsWithPrefix copies agent files from source to .claude/agents/ with "cat-" prefix
func CopyAgentsWithPrefix(agentsSourceDir, claudeAgentsDir string) error {
	// Walk through all agent files (including subdirectories)
	return filepath.Walk(agentsSourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Only copy .md files
		if filepath.Ext(path) != ".md" {
			return nil
		}

		// Get relative path from agents source dir
		relPath, err := filepath.Rel(agentsSourceDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		// Create destination filename with "cat-" prefix
		baseName := filepath.Base(relPath)
		dstFileName := "cat-" + baseName
		dstPath := filepath.Join(claudeAgentsDir, dstFileName)

		// Copy file
		if err := CopyFile(path, dstPath); err != nil {
			return fmt.Errorf("failed to copy agent %s: %w", relPath, err)
		}

		return nil
	})
}

// CopyCommandsWithPrefix copies spec-kit commands to .claude/commands/ with "speckit." prefix
func CopyCommandsWithPrefix(templatesDir, claudeCommandsDir string) error {
	commandsDir := filepath.Join(templatesDir, "commands")
	if !config.IsDirectory(commandsDir) {
		// No commands directory, skip
		return nil
	}

	// Read command files
	entries, err := os.ReadDir(commandsDir)
	if err != nil {
		return fmt.Errorf("failed to read commands directory: %w", err)
	}

	// Copy each command file with "speckit." prefix
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		// Only copy .md files
		if filepath.Ext(entry.Name()) != ".md" {
			continue
		}

		srcPath := filepath.Join(commandsDir, entry.Name())

		// Create destination filename with "speckit." prefix
		baseName := strings.TrimSuffix(entry.Name(), ".md")
		dstFileName := "speckit." + baseName + ".md"
		dstPath := filepath.Join(claudeCommandsDir, dstFileName)

		if err := CopyFile(srcPath, dstPath); err != nil {
			return fmt.Errorf("failed to copy command %s: %w", entry.Name(), err)
		}
	}

	return nil
}

// CopySpecKitFiles copies the .specify/ directory to the installation prefix
func CopySpecKitFiles(srcSpecifyDir, dstSpecifyDir string) error {
	// Ensure source exists
	if !config.IsDirectory(srcSpecifyDir) {
		return fmt.Errorf("source .specify/ directory not found: %s", srcSpecifyDir)
	}

	// Copy entire .specify/ directory
	if err := CopyDirectory(srcSpecifyDir, dstSpecifyDir); err != nil {
		return fmt.Errorf("failed to copy .specify/ directory: %w", err)
	}

	return nil
}

// CountFiles counts files in a directory (non-recursive)
func CountFiles(dir string) (int, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			count++
		}
	}

	return count, nil
}

// CountFilesRecursive counts all files in a directory recursively
func CountFilesRecursive(dir string) (int, error) {
	count := 0
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			count++
		}
		return nil
	})
	return count, err
}

// GetDirectorySize calculates the total size of a directory in bytes
func GetDirectorySize(dir string) (int64, error) {
	var size int64
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

// FormatSize formats a byte size into human-readable format
func FormatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

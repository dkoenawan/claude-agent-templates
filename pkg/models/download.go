package models

import (
	"fmt"
	"time"
)

// ArchiveFormat represents supported archive formats
type ArchiveFormat string

const (
	ArchiveFormatTarGz ArchiveFormat = "tar.gz"
	ArchiveFormatZip   ArchiveFormat = "zip"
)

// DownloadPackage represents a downloaded update package
type DownloadPackage struct {
	Version       string        `json:"version"`
	DownloadURL   string        `json:"download_url"`
	LocalPath     string        `json:"local_path"`
	ExtractedPath string        `json:"extracted_path"`
	ArchiveFormat ArchiveFormat `json:"archive_format"`
	Checksum      string        `json:"checksum"`
	Size          int64         `json:"size"`
	DownloadedAt  string        `json:"downloaded_at"`
	FileCount     int           `json:"file_count"`
}

// NewDownloadPackage creates a new DownloadPackage instance
func NewDownloadPackage(version, downloadURL string, format ArchiveFormat) *DownloadPackage {
	return &DownloadPackage{
		Version:       version,
		DownloadURL:   downloadURL,
		ArchiveFormat: format,
		DownloadedAt:  time.Now().UTC().Format(time.RFC3339),
	}
}

// Validate checks if the download package is valid
func (dp *DownloadPackage) Validate() error {
	if dp.Version == "" {
		return fmt.Errorf("version is required")
	}

	if dp.DownloadURL == "" {
		return fmt.Errorf("download URL is required")
	}

	if dp.LocalPath == "" {
		return fmt.Errorf("local path is required")
	}

	if dp.ArchiveFormat != ArchiveFormatTarGz && dp.ArchiveFormat != ArchiveFormatZip {
		return fmt.Errorf("invalid archive format: %s", dp.ArchiveFormat)
	}

	if dp.Checksum == "" {
		return fmt.Errorf("checksum is required")
	}

	if len(dp.Checksum) != 64 {
		return fmt.Errorf("invalid checksum length: %d (expected 64 hex characters)", len(dp.Checksum))
	}

	if dp.Size <= 0 {
		return fmt.Errorf("invalid size: %d", dp.Size)
	}

	return nil
}

// GetSizeHuman returns the size in human-readable format
func (dp *DownloadPackage) GetSizeHuman() string {
	const unit = 1024
	if dp.Size < unit {
		return fmt.Sprintf("%d B", dp.Size)
	}

	div, exp := int64(unit), 0
	for n := dp.Size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(dp.Size)/float64(div), "KMGTPE"[exp])
}

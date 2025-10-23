package models

import (
	"os"
	"path/filepath"
	"testing"
)

func TestManifest_Validate(t *testing.T) {
	tests := []struct {
		name     string
		manifest *Manifest
		wantErr  bool
		errMsg   string
	}{
		{
			name: "valid manifest",
			manifest: &Manifest{
				Version: "1.0",
				Name:    "claude-agent-templates",
				Dependencies: map[string]Dependency{
					"spec-kit": {
						Version:     "0.0.72",
						Source:      "vendored",
						InstallPath: ".specify",
					},
				},
				UpdatePolicy: "manual",
				LastUpdated:  "2025-10-22",
			},
			wantErr: false,
		},
		{
			name: "invalid manifest version",
			manifest: &Manifest{
				Version: "1.0.0",
				Name:    "claude-agent-templates",
				Dependencies: map[string]Dependency{
					"spec-kit": {
						Version:     "0.0.72",
						Source:      "vendored",
						InstallPath: ".specify",
					},
				},
			},
			wantErr: true,
			errMsg:  "invalid manifest version format",
		},
		{
			name: "missing name",
			manifest: &Manifest{
				Version: "1.0",
				Dependencies: map[string]Dependency{
					"spec-kit": {
						Version:     "0.0.72",
						Source:      "vendored",
						InstallPath: ".specify",
					},
				},
			},
			wantErr: true,
			errMsg:  "manifest name is required",
		},
		{
			name: "no dependencies",
			manifest: &Manifest{
				Version:      "1.0",
				Name:         "claude-agent-templates",
				Dependencies: map[string]Dependency{},
			},
			wantErr: true,
			errMsg:  "manifest must have at least one dependency",
		},
		{
			name: "invalid update policy",
			manifest: &Manifest{
				Version: "1.0",
				Name:    "claude-agent-templates",
				Dependencies: map[string]Dependency{
					"spec-kit": {
						Version:     "0.0.72",
						Source:      "vendored",
						InstallPath: ".specify",
					},
				},
				UpdatePolicy: "invalid",
			},
			wantErr: true,
			errMsg:  "invalid update_policy",
		},
		{
			name: "invalid date format",
			manifest: &Manifest{
				Version: "1.0",
				Name:    "claude-agent-templates",
				Dependencies: map[string]Dependency{
					"spec-kit": {
						Version:     "0.0.72",
						Source:      "vendored",
						InstallPath: ".specify",
					},
				},
				LastUpdated: "10/22/2025",
			},
			wantErr: true,
			errMsg:  "invalid last_updated date format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.manifest.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Manifest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDependency_Validate(t *testing.T) {
	tests := []struct {
		name       string
		dependency *Dependency
		depName    string
		wantErr    bool
		errMsg     string
	}{
		{
			name: "valid dependency",
			dependency: &Dependency{
				Version:     "0.0.72",
				Source:      "vendored",
				InstallPath: ".specify",
			},
			depName: "spec-kit",
			wantErr: false,
		},
		{
			name: "invalid version format",
			dependency: &Dependency{
				Version:     "0.0",
				Source:      "vendored",
				InstallPath: ".specify",
			},
			depName: "spec-kit",
			wantErr: true,
			errMsg:  "invalid version format",
		},
		{
			name: "invalid source",
			dependency: &Dependency{
				Version:     "0.0.72",
				Source:      "invalid",
				InstallPath: ".specify",
			},
			depName: "spec-kit",
			wantErr: true,
			errMsg:  "invalid source",
		},
		{
			name: "missing install path",
			dependency: &Dependency{
				Version: "0.0.72",
				Source:  "vendored",
			},
			depName: "spec-kit",
			wantErr: true,
			errMsg:  "install_path is required",
		},
		{
			name: "invalid integrity hash",
			dependency: &Dependency{
				Version:     "0.0.72",
				Source:      "vendored",
				InstallPath: ".specify",
				Integrity:   "invalid-hash",
			},
			depName: "spec-kit",
			wantErr: true,
			errMsg:  "invalid integrity hash format",
		},
		{
			name: "valid integrity hash",
			dependency: &Dependency{
				Version:     "0.0.72",
				Source:      "vendored",
				InstallPath: ".specify",
				Integrity:   "sha256-0000000000000000000000000000000000000000000000000000000000000000",
			},
			depName: "spec-kit",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dependency.Validate(tt.depName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dependency.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoadManifest(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()

	// Create a valid manifest file
	validManifestPath := filepath.Join(tmpDir, "valid-manifest.json")
	validManifestContent := `{
  "version": "1.0",
  "name": "claude-agent-templates",
  "dependencies": {
    "spec-kit": {
      "version": "0.0.72",
      "source": "vendored",
      "install_path": ".specify"
    }
  }
}`
	if err := os.WriteFile(validManifestPath, []byte(validManifestContent), 0644); err != nil {
		t.Fatalf("Failed to create test manifest file: %v", err)
	}

	// Create an invalid manifest file (bad JSON)
	invalidManifestPath := filepath.Join(tmpDir, "invalid-manifest.json")
	invalidManifestContent := `{ "version": "1.0", "name": "test",`
	if err := os.WriteFile(invalidManifestPath, []byte(invalidManifestContent), 0644); err != nil {
		t.Fatalf("Failed to create test manifest file: %v", err)
	}

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "valid manifest",
			path:    validManifestPath,
			wantErr: false,
		},
		{
			name:    "invalid JSON",
			path:    invalidManifestPath,
			wantErr: true,
		},
		{
			name:    "nonexistent file",
			path:    filepath.Join(tmpDir, "nonexistent.json"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manifest, err := LoadManifest(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadManifest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && manifest == nil {
				t.Error("LoadManifest() returned nil manifest")
			}
		})
	}
}

func TestManifest_SaveManifest(t *testing.T) {
	tmpDir := t.TempDir()
	manifestPath := filepath.Join(tmpDir, "test-manifest.json")

	manifest := &Manifest{
		Version: "1.0",
		Name:    "claude-agent-templates",
		Dependencies: map[string]Dependency{
			"spec-kit": {
				Version:     "0.0.72",
				Source:      "vendored",
				InstallPath: ".specify",
			},
		},
		UpdatePolicy: "manual",
		LastUpdated:  "2025-10-22",
	}

	err := manifest.SaveManifest(manifestPath)
	if err != nil {
		t.Errorf("SaveManifest() error = %v", err)
		return
	}

	// Verify file was created and can be loaded
	loadedManifest, err := LoadManifest(manifestPath)
	if err != nil {
		t.Errorf("Failed to load saved manifest: %v", err)
		return
	}

	if loadedManifest.Version != manifest.Version {
		t.Errorf("Saved manifest version = %s, want %s", loadedManifest.Version, manifest.Version)
	}
	if loadedManifest.Name != manifest.Name {
		t.Errorf("Saved manifest name = %s, want %s", loadedManifest.Name, manifest.Name)
	}
}

func TestManifest_GetSpecKitDependency(t *testing.T) {
	tests := []struct {
		name         string
		manifest     *Manifest
		wantVersion  string
		wantErr      bool
	}{
		{
			name: "spec-kit dependency exists",
			manifest: &Manifest{
				Dependencies: map[string]Dependency{
					"spec-kit": {
						Version:     "0.0.72",
						Source:      "vendored",
						InstallPath: ".specify",
					},
				},
			},
			wantVersion: "0.0.72",
			wantErr:     false,
		},
		{
			name: "spec-kit dependency missing",
			manifest: &Manifest{
				Dependencies: map[string]Dependency{
					"other": {
						Version:     "1.0.0",
						Source:      "git",
						InstallPath: "/tmp",
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dep, err := tt.manifest.GetSpecKitDependency()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpecKitDependency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && dep.Version != tt.wantVersion {
				t.Errorf("GetSpecKitDependency() version = %s, want %s", dep.Version, tt.wantVersion)
			}
		})
	}
}

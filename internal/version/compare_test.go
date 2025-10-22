package version

import (
	"testing"
)

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		name    string
		v1      string
		v2      string
		want    int
		wantErr bool
	}{
		{
			name: "v1 greater than v2",
			v1:   "0.0.72",
			v2:   "0.0.70",
			want: 1,
		},
		{
			name: "v1 less than v2",
			v1:   "0.0.68",
			v2:   "0.0.72",
			want: -1,
		},
		{
			name: "v1 equals v2",
			v1:   "0.0.72",
			v2:   "0.0.72",
			want: 0,
		},
		{
			name: "major version difference",
			v1:   "1.0.0",
			v2:   "0.0.72",
			want: 1,
		},
		{
			name: "minor version difference",
			v1:   "0.1.0",
			v2:   "0.0.72",
			want: 1,
		},
		{
			name:    "invalid v1",
			v1:      "invalid",
			v2:      "0.0.72",
			wantErr: true,
		},
		{
			name:    "invalid v2",
			v1:      "0.0.72",
			v2:      "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareVersions(tt.v1, tt.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareVersions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Normalize comparison result to -1, 0, or 1
				if got > 0 {
					got = 1
				} else if got < 0 {
					got = -1
				}
				if got != tt.want {
					t.Errorf("CompareVersions(%q, %q) = %d, want %d", tt.v1, tt.v2, got, tt.want)
				}
			}
		})
	}
}

func TestInRange(t *testing.T) {
	tests := []struct {
		name       string
		version    string
		minVersion string
		maxVersion string
		want       bool
		wantErr    bool
	}{
		{
			name:       "version in range",
			version:    "0.0.72",
			minVersion: "0.0.70",
			maxVersion: "0.1.0",
			want:       true,
		},
		{
			name:       "version below range",
			version:    "0.0.68",
			minVersion: "0.0.70",
			maxVersion: "0.1.0",
			want:       false,
		},
		{
			name:       "version above range",
			version:    "0.2.0",
			minVersion: "0.0.70",
			maxVersion: "0.1.0",
			want:       false,
		},
		{
			name:       "version equals min",
			version:    "0.0.70",
			minVersion: "0.0.70",
			maxVersion: "0.1.0",
			want:       true,
		},
		{
			name:       "version equals max",
			version:    "0.1.0",
			minVersion: "0.0.70",
			maxVersion: "0.1.0",
			want:       true,
		},
		{
			name:       "no constraints",
			version:    "0.0.72",
			minVersion: "",
			maxVersion: "",
			want:       true,
		},
		{
			name:       "only min constraint",
			version:    "0.0.72",
			minVersion: "0.0.70",
			maxVersion: "",
			want:       true,
		},
		{
			name:       "only max constraint",
			version:    "0.0.72",
			minVersion: "",
			maxVersion: "0.1.0",
			want:       true,
		},
		{
			name:       "invalid version",
			version:    "invalid",
			minVersion: "0.0.70",
			maxVersion: "0.1.0",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InRange(tt.version, tt.minVersion, tt.maxVersion)
			if (err != nil) != tt.wantErr {
				t.Errorf("InRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("InRange(%q, %q, %q) = %v, want %v", tt.version, tt.minVersion, tt.maxVersion, got, tt.want)
			}
		})
	}
}

func TestIsBreakingVersion(t *testing.T) {
	tests := []struct {
		name             string
		version          string
		breakingVersions []string
		want             bool
		wantErr          bool
	}{
		{
			name:             "is breaking version",
			version:          "0.1.0",
			breakingVersions: []string{"0.1.0", "0.2.0"},
			want:             true,
		},
		{
			name:             "not breaking version",
			version:          "0.0.72",
			breakingVersions: []string{"0.1.0", "0.2.0"},
			want:             false,
		},
		{
			name:             "empty breaking versions list",
			version:          "0.0.72",
			breakingVersions: []string{},
			want:             false,
		},
		{
			name:             "nil breaking versions list",
			version:          "0.0.72",
			breakingVersions: nil,
			want:             false,
		},
		{
			name:             "invalid version",
			version:          "invalid",
			breakingVersions: []string{"0.1.0"},
			wantErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsBreakingVersion(tt.version, tt.breakingVersions)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsBreakingVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("IsBreakingVersion(%q, %v) = %v, want %v", tt.version, tt.breakingVersions, got, tt.want)
			}
		})
	}
}

func TestCheckCompatibility(t *testing.T) {
	tests := []struct {
		name             string
		installedVersion string
		requiredVersion  string
		minVersion       string
		maxVersion       string
		breakingVersions []string
		wantCompatible   bool
		wantMismatch     bool
		wantIssues       int
	}{
		{
			name:             "exact match - compatible",
			installedVersion: "0.0.72",
			requiredVersion:  "0.0.72",
			minVersion:       "0.0.70",
			maxVersion:       "0.1.0",
			breakingVersions: []string{},
			wantCompatible:   true,
			wantMismatch:     false,
			wantIssues:       0,
		},
		{
			name:             "version mismatch but in range",
			installedVersion: "0.0.71",
			requiredVersion:  "0.0.72",
			minVersion:       "0.0.70",
			maxVersion:       "0.1.0",
			breakingVersions: []string{},
			wantCompatible:   true,
			wantMismatch:     true,
			wantIssues:       1,
		},
		{
			name:             "version below minimum",
			installedVersion: "0.0.68",
			requiredVersion:  "0.0.72",
			minVersion:       "0.0.70",
			maxVersion:       "0.1.0",
			breakingVersions: []string{},
			wantCompatible:   false,
			wantMismatch:     true,
			wantIssues:       2,
		},
		{
			name:             "version above maximum",
			installedVersion: "0.2.0",
			requiredVersion:  "0.0.72",
			minVersion:       "0.0.70",
			maxVersion:       "0.1.0",
			breakingVersions: []string{},
			wantCompatible:   false,
			wantMismatch:     true,
			wantIssues:       2,
		},
		{
			name:             "breaking version",
			installedVersion: "0.1.0",
			requiredVersion:  "0.0.72",
			minVersion:       "0.0.70",
			maxVersion:       "0.2.0",
			breakingVersions: []string{"0.1.0"},
			wantCompatible:   false,
			wantMismatch:     true,
			wantIssues:       2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CheckCompatibility(
				tt.installedVersion,
				tt.requiredVersion,
				tt.minVersion,
				tt.maxVersion,
				tt.breakingVersions,
			)
			if err != nil {
				t.Errorf("CheckCompatibility() error = %v", err)
				return
			}

			if result.Compatible != tt.wantCompatible {
				t.Errorf("CheckCompatibility() Compatible = %v, want %v", result.Compatible, tt.wantCompatible)
			}

			if result.VersionMismatch != tt.wantMismatch {
				t.Errorf("CheckCompatibility() VersionMismatch = %v, want %v", result.VersionMismatch, tt.wantMismatch)
			}

			if len(result.Issues) != tt.wantIssues {
				t.Errorf("CheckCompatibility() Issues count = %d, want %d\nIssues: %v", len(result.Issues), tt.wantIssues, result.Issues)
			}
		})
	}
}

func TestCompatibilityResult_IsCompatible(t *testing.T) {
	tests := []struct {
		name   string
		result *CompatibilityResult
		want   bool
	}{
		{
			name: "compatible result",
			result: &CompatibilityResult{
				Compatible: true,
			},
			want: true,
		},
		{
			name: "incompatible result",
			result: &CompatibilityResult{
				Compatible: false,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.result.IsCompatible(); got != tt.want {
				t.Errorf("IsCompatible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompatibilityResult_HasIssues(t *testing.T) {
	tests := []struct {
		name   string
		result *CompatibilityResult
		want   bool
	}{
		{
			name: "has issues",
			result: &CompatibilityResult{
				Issues: []string{"issue 1", "issue 2"},
			},
			want: true,
		},
		{
			name: "no issues",
			result: &CompatibilityResult{
				Issues: []string{},
			},
			want: false,
		},
		{
			name: "nil issues",
			result: &CompatibilityResult{
				Issues: nil,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.result.HasIssues(); got != tt.want {
				t.Errorf("HasIssues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompatibilityResult_GetIssuesText(t *testing.T) {
	tests := []struct {
		name   string
		result *CompatibilityResult
		want   string
	}{
		{
			name: "with issues",
			result: &CompatibilityResult{
				Issues: []string{"issue 1", "issue 2"},
			},
			want: "Compatibility issues:\n  1. issue 1\n  2. issue 2\n",
		},
		{
			name: "no issues",
			result: &CompatibilityResult{
				Issues: []string{},
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.result.GetIssuesText(); got != tt.want {
				t.Errorf("GetIssuesText() = %q, want %q", got, tt.want)
			}
		})
	}
}

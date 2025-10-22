package version

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

// CompareVersions compares two semantic version strings
// Returns:
//   - positive value if v1 > v2
//   - 0 if v1 == v2
//   - negative value if v1 < v2
func CompareVersions(v1, v2 string) (int, error) {
	ver1, err := semver.NewVersion(v1)
	if err != nil {
		return 0, fmt.Errorf("invalid version v1 '%s': %w", v1, err)
	}

	ver2, err := semver.NewVersion(v2)
	if err != nil {
		return 0, fmt.Errorf("invalid version v2 '%s': %w", v2, err)
	}

	return ver1.Compare(ver2), nil
}

// InRange checks if a version is within the specified range (inclusive)
// Returns true if minVersion <= version <= maxVersion
func InRange(version, minVersion, maxVersion string) (bool, error) {
	if minVersion == "" && maxVersion == "" {
		return true, nil // No constraints means any version is valid
	}

	v, err := semver.NewVersion(version)
	if err != nil {
		return false, fmt.Errorf("invalid version '%s': %w", version, err)
	}

	// Check minimum version
	if minVersion != "" {
		minVer, err := semver.NewVersion(minVersion)
		if err != nil {
			return false, fmt.Errorf("invalid min version '%s': %w", minVersion, err)
		}
		if v.LessThan(minVer) {
			return false, nil
		}
	}

	// Check maximum version
	if maxVersion != "" {
		maxVer, err := semver.NewVersion(maxVersion)
		if err != nil {
			return false, fmt.Errorf("invalid max version '%s': %w", maxVersion, err)
		}
		if v.GreaterThan(maxVer) {
			return false, nil
		}
	}

	return true, nil
}

// IsBreakingVersion checks if a version is in the list of breaking versions
func IsBreakingVersion(version string, breakingVersions []string) (bool, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return false, fmt.Errorf("invalid version '%s': %w", version, err)
	}

	for _, breakingVer := range breakingVersions {
		bv, err := semver.NewVersion(breakingVer)
		if err != nil {
			return false, fmt.Errorf("invalid breaking version '%s': %w", breakingVer, err)
		}
		if v.Equal(bv) {
			return true, nil
		}
	}

	return false, nil
}

// CheckCompatibility checks if an installed version is compatible with the required version
// based on the compatibility constraints (min, max, breaking versions)
func CheckCompatibility(installedVersion, requiredVersion, minVersion, maxVersion string, breakingVersions []string) (*CompatibilityResult, error) {
	result := &CompatibilityResult{
		InstalledVersion: installedVersion,
		RequiredVersion:  requiredVersion,
		Compatible:       true,
		Issues:           []string{},
	}

	// Check if installed version is in valid range
	inRange, err := InRange(installedVersion, minVersion, maxVersion)
	if err != nil {
		return nil, err
	}

	if !inRange {
		result.Compatible = false
		if minVersion != "" && maxVersion != "" {
			result.Issues = append(result.Issues, fmt.Sprintf("installed version %s is outside allowed range %s - %s", installedVersion, minVersion, maxVersion))
		} else if minVersion != "" {
			result.Issues = append(result.Issues, fmt.Sprintf("installed version %s is below minimum version %s", installedVersion, minVersion))
		} else if maxVersion != "" {
			result.Issues = append(result.Issues, fmt.Sprintf("installed version %s is above maximum version %s", installedVersion, maxVersion))
		}
	}

	// Check if installed version is a known breaking version
	isBreaking, err := IsBreakingVersion(installedVersion, breakingVersions)
	if err != nil {
		return nil, err
	}

	if isBreaking {
		result.Compatible = false
		result.Issues = append(result.Issues, fmt.Sprintf("installed version %s is a known breaking version", installedVersion))
	}

	// Compare with required version
	cmp, err := CompareVersions(installedVersion, requiredVersion)
	if err != nil {
		return nil, err
	}

	if cmp != 0 {
		result.VersionMismatch = true
		if cmp < 0 {
			result.Issues = append(result.Issues, fmt.Sprintf("installed version %s is older than required version %s", installedVersion, requiredVersion))
		} else {
			result.Issues = append(result.Issues, fmt.Sprintf("installed version %s is newer than required version %s", installedVersion, requiredVersion))
		}
	}

	return result, nil
}

// CompatibilityResult represents the result of a compatibility check
type CompatibilityResult struct {
	InstalledVersion string   `json:"installed_version"`
	RequiredVersion  string   `json:"required_version"`
	Compatible       bool     `json:"compatible"`
	VersionMismatch  bool     `json:"version_mismatch"`
	Issues           []string `json:"issues,omitempty"`
}

// IsCompatible is a convenience method to check compatibility
func (cr *CompatibilityResult) IsCompatible() bool {
	return cr.Compatible
}

// HasIssues returns true if there are any compatibility issues
func (cr *CompatibilityResult) HasIssues() bool {
	return len(cr.Issues) > 0
}

// GetIssuesText returns a formatted string of all issues
func (cr *CompatibilityResult) GetIssuesText() string {
	if !cr.HasIssues() {
		return ""
	}

	text := "Compatibility issues:\n"
	for i, issue := range cr.Issues {
		text += fmt.Sprintf("  %d. %s\n", i+1, issue)
	}
	return text
}

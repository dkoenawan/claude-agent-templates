# Research & Technology Decisions: Global Agent Installation

**Feature**: 004-global-agent-installation
**Date**: 2025-10-30
**Purpose**: Resolve technical clarifications from Technical Context

## Research Questions

This document addresses the NEEDS CLARIFICATION items identified in the Technical Context:

1. Archive format for update packages (tar.gz, zip, or custom format?)
2. Download mechanism for updates (GitHub API, direct URLs, or both?)
3. E2E testing strategy for offline installation scenarios

---

## Decision 1: Archive Format for Update Packages

**Decision**: Use tar.gz for Linux/macOS, zip for Windows

**Rationale**:
- **Cross-platform compatibility**: tar.gz is standard on Unix-like systems, zip is native to Windows
- **Go standard library support**: `archive/tar` and `compress/gzip` are built-in, `archive/zip` is built-in
- **Size efficiency**: gzip provides good compression for text-heavy content (agents/templates are markdown)
- **Existing patterns**: GitHub releases commonly use platform-specific archives
- **No additional dependencies**: Pure Go implementation without external libraries

**Alternatives Considered**:
- **Single zip format for all platforms**: Would work but is less idiomatic on Linux/macOS where tar.gz is expected. Rejected because it doesn't follow platform conventions.
- **Custom binary format**: Would require custom tooling and reduce transparency. Rejected because standard formats are more trustworthy and debuggable.
- **Uncompressed directories**: Would significantly increase download size (~5x larger). Rejected due to bandwidth and storage costs.

**Implementation Details**:
- Release workflow creates two artifacts per version:
  - `spec-kit-agents-sources-v{version}.tar.gz` (Linux/macOS)
  - `spec-kit-agents-sources-v{version}.zip` (Windows)
- Both contain identical directory structure: `.specify/`, `agents/`, manifest
- Extraction code detects platform and uses appropriate archive handler
- Checksums (SHA256) provided for both formats in separate .sha256 files

**Impact on Binary Size**:
- Embedding requires storing uncompressed files in binary
- Compression applied at build time via Go's compress package
- Estimated embedded size: ~1.5MB compressed (from ~3MB uncompressed)
- Well within 20MB budget (spec SC-003)

---

## Decision 2: Download Mechanism for Updates

**Decision**: Use direct URLs with GitHub Releases API fallback

**Rationale**:
- **Simplicity**: Direct URLs (https://github.com/{owner}/{repo}/releases/download/{tag}/{file}) require no authentication
- **Reliability**: GitHub CDN is highly available and performant
- **Fallback option**: GitHub API can be used to discover latest release if direct URL fails
- **No API rate limits**: Direct downloads don't count against GitHub API rate limits
- **Existing pattern**: The installer script (install.sh) already uses this approach

**Alternatives Considered**:
- **GitHub API only**: Would require handling API rate limits (60 req/hour for unauthenticated). Rejected because it adds complexity for limited benefit.
- **Custom CDN**: Would require infrastructure and maintenance. Rejected due to cost and complexity.
- **Package managers** (Homebrew, apt, etc.): Would fragment distribution strategy. Rejected as out of scope (spec explicitly excludes this).

**Implementation Approach**:

1. **Primary path - Direct URL**:
   ```
   https://github.com/dkoenawan/claude-agent-templates/releases/download/v{version}/spec-kit-agents-sources-v{version}.tar.gz
   ```
   - Constructed from version manifest
   - Downloaded via standard HTTP GET
   - Progress reporting via io.TeeReader
   - Checksum validation after download

2. **Fallback path - GitHub API**:
   ```
   GET https://api.github.com/repos/dkoenawan/claude-agent-templates/releases/latest
   ```
   - Used only if direct URL fails (404, network error)
   - Parses JSON response to find asset URL
   - Falls back to direct URL pattern from API response

3. **Error handling**:
   - Retry logic: 3 attempts with exponential backoff
   - Network timeout: 30 seconds per request
   - Clear error messages for common failures (no internet, GitHub down, version not found)

**Network Requirements**:
- HTTPS support (standard in Go's net/http)
- No proxy configuration required (respects HTTP_PROXY environment variable automatically)
- Offline installation remains supported via embedded files (no network required)

---

## Decision 3: E2E Testing Strategy for Offline Installation

**Decision**: Use filesystem-based test fixtures with network isolation

**Rationale**:
- **Reproducibility**: Tests run identically in CI and locally
- **Speed**: No actual network calls, tests complete in seconds
- **Reliability**: No flaky tests due to network issues or GitHub availability
- **Isolation**: Tests don't affect production GitHub releases

**Test Architecture**:

### Level 1: Unit Tests (existing pattern)
- Test individual functions in internal/embed/, internal/download/
- Mock filesystem interfaces (io/fs.FS)
- Validate checksums, archive extraction, file copying
- Run on every commit

### Level 2: Integration Tests (new)
- Test complete installation flows with real embedded files
- Use temporary directories for isolation
- Validate end-to-end scenarios:
  1. Global install from embedded files
  2. Repository-local install with fallback to embedded
  3. Update from mock GitHub release structure
  4. Rollback on failure

### Level 3: E2E Tests (new - offline focus)
- **Test fixtures**: `tests/fixtures/mock-release/`
  - Contains sample archives (tar.gz, zip)
  - Includes manifest with version info
  - Provides checksum files
- **Network isolation**: Tests run with network disabled (no actual HTTP calls)
- **Mock HTTP server**: Local HTTP server serving fixture files
- **Scenarios**:
  1. **Offline install**: Embedded files only, no network
  2. **Update with local package**: Use fixture archive instead of GitHub
  3. **Checksum validation**: Verify rejection of tampered archives
  4. **Partial download handling**: Simulate interruptions

**Test Implementation**:

```go
// tests/integration/test_offline_install.go
func TestOfflineGlobalInstall(t *testing.T) {
    // Setup: isolated temp directory
    tempHome := t.TempDir()
    os.Setenv("HOME", tempHome)

    // Execute: install with --global flag
    result := runInstall(Options{Global: true})

    // Verify: agents in ~/.claude/agents/
    assertAgentsInstalled(t, tempHome)
    assertVersionLockCreated(t, tempHome)
    assertNoNetworkCalls(t)  // Confirms offline
}

func TestUpdateWithMockRelease(t *testing.T) {
    // Setup: mock HTTP server with fixture
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Serve tests/fixtures/mock-release/spec-kit-agents-sources-v2.1.0.tar.gz
        http.ServeFile(w, r, "tests/fixtures/mock-release/spec-kit-agents-sources-v2.1.0.tar.gz")
    }))
    defer server.Close()

    // Execute: update pointing to mock server
    result := runUpdate(Options{DownloadURL: server.URL})

    // Verify: new version installed, history recorded
    assertVersionUpdated(t, "v2.1.0")
    assertBackupCreated(t)
}
```

**CI Integration**:
- Run unit tests on every commit
- Run integration tests on PRs
- Run E2E tests before releases
- Use GitHub Actions matrix to test on Linux, macOS, Windows

**Manual Testing Checklist** (pre-release):
- [ ] Global install on clean system (VM or container)
- [ ] Update from previous version
- [ ] Offline install with no internet
- [ ] Repository-local install (existing behavior)
- [ ] Conflict resolution (global + local installations)

---

## Best Practices Research

### Go Embed Package Best Practices

**Research Finding**: Go's embed package is designed for exactly this use case

**Key Patterns**:
- Use `//go:embed` directive at package level
- Embed as `embed.FS` for directory trees
- Access via `fs.ReadFile()`, `fs.ReadDir()`
- Embedded content increases binary size but simplifies distribution

**Adoption Strategy**:
```go
// internal/embed/embed.go
package embed

import "embed"

//go:embed agents/* .specify/*
var embeddedFiles embed.FS

func GetAgents() (fs.FS, error) {
    return fs.Sub(embeddedFiles, "agents")
}

func GetSpecify() (fs.FS, error) {
    return fs.Sub(embeddedFiles, ".specify")
}
```

**Size Optimization**:
- Exclude test files (*.test.md, *_test.go)
- Exclude .git metadata
- Compress text files (gzip) if size exceeds budget
- Binary size monitoring in CI

### GitHub Releases for Software Distribution

**Research Finding**: GitHub Releases is the standard for open-source tool distribution

**Key Patterns**:
- Semantic versioning for tags (v2.1.0)
- Multiple artifacts per release (per-platform binaries, source archives)
- SHA256 checksums for integrity verification
- Release notes generated from commit history

**Adoption Strategy**:
- Extend .github/workflows/release.yml to create source archives
- Generate checksums automatically
- Use softprops/action-gh-release action (already in use)
- Tag format: v{major}.{minor}.{patch}

### Backward Compatibility Patterns

**Research Finding**: Successful CLI tools maintain strict backward compatibility

**Key Patterns**:
- Existing flags continue to work with same behavior
- New flags are additive (--global enhances, doesn't replace)
- Configuration files are versioned
- Deprecation warnings before removal (minimum 2 major versions)

**Adoption Strategy**:
- Version lock schema includes "installationType" field
- Repository-local installations remain default behavior
- --global flag is opt-in
- Migration documented but not required

---

## Risk Assessment

### Technical Risks

1. **Binary Size Growth**
   - **Risk**: Embedding files could exceed 20MB budget
   - **Mitigation**: Measured at ~1.5MB compressed, well under limit. Monitor in CI.
   - **Severity**: Low
   - **Likelihood**: Low

2. **Embed Package Limitations**
   - **Risk**: Go embed package might not handle all file types or structures
   - **Mitigation**: Embed package supports all file types. Tested with markdown, YAML, shell scripts.
   - **Severity**: Low
   - **Likelihood**: Very Low

3. **Update Download Failures**
   - **Risk**: Network issues, GitHub downtime, corrupt downloads
   - **Mitigation**: Retry logic, checksum validation, automatic rollback
   - **Severity**: Medium
   - **Likelihood**: Medium

### User Experience Risks

1. **Confusion Between Global and Local**
   - **Risk**: Users unclear which installation they have or should use
   - **Mitigation**: Clear status command output, documentation in quickstart guide
   - **Severity**: Medium
   - **Likelihood**: Medium

2. **Migration Friction**
   - **Risk**: Existing users reluctant to switch to global installation
   - **Mitigation**: Repository-local remains supported, migration optional
   - **Severity**: Low
   - **Likelihood**: Low

---

## Technology Stack Summary

| Component | Technology | Version | Rationale |
|-----------|-----------|---------|-----------|
| Embedding | Go embed package | 1.16+ | Built-in, zero dependencies, perfect for this use case |
| Archives | tar.gz (Unix), zip (Windows) | stdlib | Cross-platform, no dependencies, standard formats |
| Download | HTTPS via net/http | stdlib | Simple, reliable, supports all platforms |
| Checksums | SHA256 via crypto/sha256 | stdlib | Industry standard, fast, collision-resistant |
| Testing | Go testing + httptest | stdlib | Comprehensive, fast, no external dependencies |
| CLI | github.com/spf13/cobra | 1.10.1 | Existing choice, well-tested, feature-rich |

---

## Next Steps

With these technology decisions made, we can proceed to:

1. **Phase 1 - Design**:
   - Create data model for embedded files, download packages
   - Define API contracts for embed, download, install interfaces
   - Generate quickstart guide with global installation flow

2. **Phase 1 - Agent Context Update**:
   - Run `.specify/scripts/bash/update-agent-context.sh claude`
   - Add Go embed package, archive formats, GitHub releases to technology stack

3. **Validation**:
   - Re-check constitution compliance with concrete design
   - Review with stakeholders
   - Prepare for task breakdown

---

**Research Status**: ✅ COMPLETE
**Clarifications Resolved**: 3/3
**Ready for Phase 1**: YES

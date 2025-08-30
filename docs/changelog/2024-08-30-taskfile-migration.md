# Taskfile Migration Plan

## Executive Summary

This document outlines the architectural plan for migrating the Claude Agent Templates repository from platform-specific shell scripts to a unified Taskfile-based automation system. This migration will provide a consistent cross-platform experience while enhancing functionality beyond the current basic installation scripts.

## Current State Assessment

### Existing Scripts Analysis
- **install-agents.sh** (Linux/macOS): 67 lines of bash script
- **install-agents.bat** (Windows): 64 lines of batch script
- **Functionality**: Both scripts provide identical agent installation capabilities
- **Maintenance Overhead**: Duplicate logic across two platform-specific implementations
- **Limited Scope**: Only handles basic installation, no listing, cleaning, or validation

### Key Pain Points
1. **Code Duplication**: Same logic implemented twice in different scripting languages
2. **Platform Fragmentation**: Users need different commands based on their operating system
3. **Limited Functionality**: No advanced operations like listing installed agents or cleanup
4. **Documentation Split**: Platform-specific installation instructions

## Migration Architecture

### Design Principles
- **Cross-Platform Compatibility**: Single interface works on Linux, macOS, and Windows
- **Backward Compatibility**: Existing script interfaces remain functional during transition
- **Enhanced Functionality**: Expand beyond basic installation to include management operations
- **Hexagonal Architecture**: Clear separation of domain logic, application services, and infrastructure

### Target Architecture

```yaml
# Taskfile.yml Structure
version: '3'

vars:
  AGENTS_DIR: ./agents
  CLAUDE_DIR: ~/.claude/agents

tasks:
  install:     # Core installation functionality
  list:        # List available/installed agents
  clean:       # Remove installed agents
  validate:    # Verify agent installation
  update:      # Update existing agents
  help:        # Display usage information
```

## Implementation Plan

### Phase 1: Foundation (Work Unit 1-2)
**Deliverable**: Basic Taskfile.yml with core installation logic

1. **Create Taskfile.yml Foundation**
   - Set up v3 schema with cross-platform variables
   - Configure path handling for different operating systems
   - Establish error handling patterns

2. **Migrate Core Installation Logic**
   - Convert existing script functionality to Taskfile tasks
   - Implement agent discovery and copying logic
   - Add progress reporting and error handling

**Acceptance Criteria**:
- `task install` replicates existing script behavior
- Works identically on Linux, macOS, and Windows
- Maintains same performance characteristics

### Phase 2: Enhanced Operations (Work Unit 3)
**Deliverable**: Advanced management tasks

3. **Implement Advanced Operations**
   - `task list`: Display available and installed agents
   - `task clean`: Remove installed agents with confirmation
   - `task validate`: Verify agent files and permissions
   - `task update`: Refresh existing agent installations

**Acceptance Criteria**:
- All new tasks work across platforms
- Proper error handling and user feedback
- Help documentation for each task

### Phase 3: Compatibility & Testing (Work Unit 4-5)
**Deliverable**: Backward compatibility and comprehensive testing

4. **Create Backward Compatibility Layer**
   - Wrapper scripts that call Taskfile tasks
   - Maintain existing `scripts/install-agents.*` interfaces
   - Deprecation warnings with migration guidance

5. **Implement Testing Framework**
   - Automated testing across Linux, macOS, Windows
   - Integration tests for all task operations
   - Performance benchmarks against original scripts

**Acceptance Criteria**:
- Existing scripts continue to work unchanged
- All platforms pass automated tests
- Performance within 10% of original scripts

### Phase 4: Documentation & Rollout (Work Unit 6)
**Deliverable**: Updated documentation and migration guides

6. **Update Documentation**
   - Unified installation instructions using `task install`
   - Migration guide for existing users
   - Troubleshooting section for common issues
   - Update README.md with new unified approach

**Acceptance Criteria**:
- All documentation reflects new Taskfile approach
- Clear migration path for existing users
- Comprehensive troubleshooting coverage

## Technical Implementation Details

### Cross-Platform Considerations
- **Path Handling**: Use Taskfile variables for OS-specific paths
- **Shell Compatibility**: Leverage Taskfile's built-in cross-platform support
- **File Operations**: Use portable commands (cp, mkdir, etc.)
- **Error Handling**: Consistent exit codes across platforms

### Performance Requirements
- Installation time within 10% of current scripts
- Minimal dependencies (only Task binary required)
- Efficient file operations with progress feedback

### Security Considerations
- Maintain existing file permission handling
- Validate agent files before installation
- Clear error messages for permission issues

## Risk Mitigation

### Technical Risks
1. **Cross-Platform Compatibility**
   - **Risk**: Different behavior on different operating systems
   - **Mitigation**: Comprehensive testing matrix, use of Taskfile built-ins

2. **Performance Regression**
   - **Risk**: Slower execution than existing scripts
   - **Mitigation**: Performance benchmarking, optimization of critical paths

3. **User Adoption**
   - **Risk**: Users continue using old scripts
   - **Mitigation**: Backward compatibility, clear migration benefits, gradual deprecation

### Migration Risks
1. **Breaking Changes**
   - **Risk**: Existing workflows disrupted
   - **Mitigation**: Maintain wrapper scripts, phased deprecation approach

2. **Documentation Gaps**
   - **Risk**: Users confused during transition
   - **Mitigation**: Comprehensive migration guide, clear examples

## Success Metrics

### Functional Metrics
- All existing script functionality preserved
- 100% test coverage across all platforms
- Zero breaking changes during transition period

### User Experience Metrics
- Single command interface (`task install`) works on all platforms
- Enhanced functionality (list, clean, validate) available
- Reduced platform-specific documentation

### Technical Metrics
- Performance within 10% of original scripts
- Reduced codebase complexity (single source vs. dual implementation)
- Improved maintainability through unified approach

## Timeline and Dependencies

### Prerequisites
- Task binary installation across development environments
- Testing infrastructure for multiple platforms

### Implementation Order
1. **Foundation** (1-2 days): Core Taskfile.yml and basic installation
2. **Enhancement** (1-2 days): Advanced operations and features
3. **Compatibility** (1 day): Wrapper scripts and backward compatibility
4. **Testing** (2-3 days): Comprehensive cross-platform testing
5. **Documentation** (1 day): Updated guides and migration instructions

### Total Estimated Effort: 6-9 days

## Conclusion

This migration plan provides a systematic approach to modernizing the repository's automation infrastructure while maintaining backward compatibility and enhancing user experience. The phased implementation approach minimizes risk while delivering clear value through unified cross-platform operations and enhanced functionality.

The resulting system will be more maintainable, feature-rich, and provide a consistent experience regardless of the user's operating system, aligning with modern development practices and user expectations.
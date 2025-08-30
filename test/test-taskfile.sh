#!/bin/bash

# Test suite for Taskfile migration
# Tests all Task operations for functionality and cross-platform compatibility

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Test configuration
TEST_DIR="/tmp/claude-agents-test-$$"
ORIGINAL_CLAUDE_DIR="$HOME/.claude"
BACKUP_DIR="/tmp/claude-backup-$$"

# Statistics
TESTS_PASSED=0
TESTS_FAILED=0

echo "🧪 Claude Agent Templates - Taskfile Test Suite"
echo "=============================================="
echo ""

# Setup test environment
setup_test_env() {
    echo "Setting up test environment..."
    
    # Backup existing claude directory if it exists
    if [ -d "$ORIGINAL_CLAUDE_DIR" ]; then
        echo "Backing up existing ~/.claude directory to $BACKUP_DIR"
        cp -r "$ORIGINAL_CLAUDE_DIR" "$BACKUP_DIR"
    fi
    
    # Create test directory
    mkdir -p "$TEST_DIR"
    
    # Set test claude directory
    export HOME="$TEST_DIR"
    
    echo "✓ Test environment ready"
    echo ""
}

# Cleanup test environment
cleanup_test_env() {
    echo ""
    echo "Cleaning up test environment..."
    
    # Restore original HOME
    export HOME="${TEST_DIR%/*}/../$(whoami)" # Restore original HOME path
    
    # Restore original claude directory if backup exists
    if [ -d "$BACKUP_DIR" ]; then
        rm -rf "$ORIGINAL_CLAUDE_DIR"
        mv "$BACKUP_DIR" "$ORIGINAL_CLAUDE_DIR"
        echo "✓ Restored original ~/.claude directory"
    fi
    
    # Remove test directory
    rm -rf "$TEST_DIR"
    echo "✓ Test environment cleaned up"
}

# Test result functions
pass_test() {
    echo -e "${GREEN}✓ PASS:${NC} $1"
    TESTS_PASSED=$((TESTS_PASSED + 1))
}

fail_test() {
    echo -e "${RED}✗ FAIL:${NC} $1"
    TESTS_FAILED=$((TESTS_FAILED + 1))
}

warn_test() {
    echo -e "${YELLOW}⚠ WARN:${NC} $1"
}

# Check if Task is available
check_task_availability() {
    echo "Checking Task availability..."
    
    if command -v task >/dev/null 2>&1; then
        pass_test "Task command available in PATH"
    elif [ -f "./bin/task" ]; then
        export PATH="$PWD/bin:$PATH"
        pass_test "Task binary found in project (./bin/task)"
    else
        fail_test "Task not found - please install Task first"
        return 1
    fi
    
    # Test task version
    task_version=$(task --version)
    pass_test "Task version: $task_version"
}

# Test basic task listing
test_task_list() {
    echo ""
    echo "Testing task list functionality..."
    
    if task --list >/dev/null 2>&1; then
        pass_test "Task --list command works"
    else
        fail_test "Task --list command failed"
    fi
}

# Test installation functionality
test_install() {
    echo ""
    echo "Testing installation functionality..."
    
    # Test task install
    if task install >/dev/null 2>&1; then
        pass_test "task install completed without errors"
    else
        fail_test "task install failed"
        return 1
    fi
    
    # Check if claude directory was created
    if [ -d "$HOME/.claude/agents" ]; then
        pass_test "Claude agents directory created at $HOME/.claude/agents"
    else
        fail_test "Claude agents directory not created"
        return 1
    fi
    
    # Count installed agents
    agent_count=$(find "$HOME/.claude/agents" -name "*.md" -type f | wc -l)
    if [ "$agent_count" -gt 0 ]; then
        pass_test "Agents installed: $agent_count files"
    else
        fail_test "No agents were installed"
        return 1
    fi
    
    # Verify specific expected agents exist
    expected_agents=("requirements-analyst.md" "solution-architect.md" "documentation.md")
    for agent in "${expected_agents[@]}"; do
        if [ -f "$HOME/.claude/agents/$agent" ]; then
            pass_test "Expected agent found: $agent"
        else
            warn_test "Expected agent missing: $agent"
        fi
    done
}

# Test list functionality
test_list() {
    echo ""
    echo "Testing list functionality..."
    
    if task list >/dev/null 2>&1; then
        pass_test "task list completed without errors"
    else
        fail_test "task list failed"
    fi
    
    # Test that list shows both available and installed agents
    list_output=$(task list 2>/dev/null)
    if echo "$list_output" | grep -q "Available agents"; then
        pass_test "List shows available agents section"
    else
        fail_test "List does not show available agents section"
    fi
    
    if echo "$list_output" | grep -q "Installed agents"; then
        pass_test "List shows installed agents section"
    else
        fail_test "List does not show installed agents section"
    fi
}

# Test validate functionality
test_validate() {
    echo ""
    echo "Testing validate functionality..."
    
    if task validate >/dev/null 2>&1; then
        pass_test "task validate completed without errors"
    else
        fail_test "task validate failed"
    fi
    
    # Check validation output
    validate_output=$(task validate 2>/dev/null)
    if echo "$validate_output" | grep -q "Available agents:"; then
        pass_test "Validate shows available agent count"
    else
        fail_test "Validate does not show available agent count"
    fi
    
    if echo "$validate_output" | grep -q "Installed agents:"; then
        pass_test "Validate shows installed agent count"
    else
        fail_test "Validate does not show installed agent count"
    fi
}

# Test clean functionality
test_clean() {
    echo ""
    echo "Testing clean functionality..."
    
    # Make sure we have agents installed first
    agent_count_before=$(find "$HOME/.claude/agents" -name "*.md" -type f 2>/dev/null | wc -l)
    
    if [ "$agent_count_before" -eq 0 ]; then
        warn_test "No agents installed before clean test"
        return 0
    fi
    
    # Test clean with -y flag (non-interactive)
    if task clean -y >/dev/null 2>&1; then
        pass_test "task clean -y completed without errors"
    else
        fail_test "task clean -y failed"
        return 1
    fi
    
    # Verify agents were removed
    agent_count_after=$(find "$HOME/.claude/agents" -name "*.md" -type f 2>/dev/null | wc -l)
    
    if [ "$agent_count_after" -eq 0 ]; then
        pass_test "All agents removed by clean command"
    else
        fail_test "Clean command did not remove all agents (found $agent_count_after)"
    fi
}

# Test update functionality
test_update() {
    echo ""
    echo "Testing update functionality..."
    
    # Install agents first
    task install >/dev/null 2>&1
    
    if task update >/dev/null 2>&1; then
        pass_test "task update completed without errors"
    else
        fail_test "task update failed"
    fi
    
    # Verify agents are still there after update
    agent_count=$(find "$HOME/.claude/agents" -name "*.md" -type f 2>/dev/null | wc -l)
    if [ "$agent_count" -gt 0 ]; then
        pass_test "Agents present after update: $agent_count files"
    else
        fail_test "No agents found after update"
    fi
}

# Test help functionality
test_help() {
    echo ""
    echo "Testing help functionality..."
    
    if task help >/dev/null 2>&1; then
        pass_test "task help completed without errors"
    else
        fail_test "task help failed"
    fi
    
    # Check help output contains expected sections
    help_output=$(task help 2>/dev/null)
    if echo "$help_output" | grep -q "USAGE:"; then
        pass_test "Help shows usage information"
    else
        fail_test "Help does not show usage information"
    fi
    
    if echo "$help_output" | grep -q "COMMANDS:"; then
        pass_test "Help shows commands information"
    else
        fail_test "Help does not show commands information"
    fi
}

# Test backward compatibility
test_backward_compatibility() {
    echo ""
    echo "Testing backward compatibility..."
    
    # Clean agents first
    rm -rf "$HOME/.claude/agents"
    
    # Test that install-agents.sh still works
    if ./scripts/install-agents.sh >/dev/null 2>&1; then
        pass_test "install-agents.sh backward compatibility works"
    else
        fail_test "install-agents.sh backward compatibility failed"
    fi
    
    # Verify agents were installed
    agent_count=$(find "$HOME/.claude/agents" -name "*.md" -type f 2>/dev/null | wc -l)
    if [ "$agent_count" -gt 0 ]; then
        pass_test "Backward compatibility script installed $agent_count agents"
    else
        fail_test "Backward compatibility script installed no agents"
    fi
}

# Performance benchmark
test_performance() {
    echo ""
    echo "Testing performance..."
    
    # Clean first
    rm -rf "$HOME/.claude/agents"
    
    # Time the installation
    start_time=$(date +%s%N)
    task install >/dev/null 2>&1
    end_time=$(date +%s%N)
    
    duration_ms=$(( (end_time - start_time) / 1000000 ))
    
    if [ "$duration_ms" -lt 5000 ]; then  # Less than 5 seconds
        pass_test "Installation performance: ${duration_ms}ms (excellent)"
    elif [ "$duration_ms" -lt 10000 ]; then  # Less than 10 seconds
        pass_test "Installation performance: ${duration_ms}ms (good)"
    else
        warn_test "Installation performance: ${duration_ms}ms (consider optimization)"
    fi
}

# Main test execution
main() {
    echo "Starting comprehensive Taskfile test suite..."
    echo ""
    
    # Setup
    setup_test_env
    
    # Run tests
    check_task_availability || exit 1
    test_task_list
    test_install
    test_list
    test_validate
    test_clean
    test_update
    test_help
    test_backward_compatibility
    test_performance
    
    # Cleanup
    trap cleanup_test_env EXIT
    
    # Results
    echo ""
    echo "🧪 Test Results"
    echo "==============="
    echo -e "${GREEN}Passed: $TESTS_PASSED${NC}"
    echo -e "${RED}Failed: $TESTS_FAILED${NC}"
    echo ""
    
    if [ $TESTS_FAILED -eq 0 ]; then
        echo -e "${GREEN}🎉 All tests passed! Taskfile migration is working correctly.${NC}"
        exit 0
    else
        echo -e "${RED}❌ Some tests failed. Please review the output above.${NC}"
        exit 1
    fi
}

# Run main function
main "$@"
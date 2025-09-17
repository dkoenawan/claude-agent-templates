# Troubleshooting Guide

This guide provides solutions for common issues when working with Claude Agent Templates, including validation errors, workflow problems, and configuration issues.

## Table of Contents

- [Agent Validation Issues](#agent-validation-issues)
- [GitHub Actions Problems](#github-actions-problems)
- [Domain Classification Issues](#domain-classification-issues)
- [Workflow State Problems](#workflow-state-problems)
- [Installation Issues](#installation-issues)
- [Performance Issues](#performance-issues)
- [Common Error Messages](#common-error-messages)

## Agent Validation Issues

### Missing Required Fields

**Problem**: Agent validation fails with "Missing required field" errors.

**Symptoms**:
```
❌ agents/python/my-agent.md: Missing required field: examples
❌ agents/python/my-agent.md: Missing required field: workflow_position
```

**Solution**:
1. Check the agent specification format in [Agent Specifications Guide](agent-specifications.md)
2. Ensure all required fields are present:
   ```yaml
   ---
   name: agent-name
   description: Agent description
   domain: python
   role: engineer
   spec_version: "1.0"
   tools: Bash, Edit, Write
   model: inherit
   color: blue
   inputs: [...]
   outputs: [...]
   validation: [...]
   dependencies: [...]
   workflow_position: 6
   github_integration: {...}
   examples: [...]
   ---
   ```

### Invalid YAML Syntax

**Problem**: YAML frontmatter contains syntax errors.

**Symptoms**:
```
❌ agents/python/my-agent.md: Invalid YAML syntax: mapping values are not allowed here
```

**Solution**:
1. Validate YAML syntax online or with a YAML validator
2. Common issues:
   - Missing quotes around strings with special characters
   - Incorrect indentation (use 2 spaces)
   - Colon in unquoted strings
   ```yaml
   # Wrong
   description: Agent for API: REST development

   # Correct
   description: "Agent for API: REST development"
   ```

### Tools Format Issues

**Problem**: Tools field validation fails.

**Symptoms**:
```
❌ agents/python/my-agent.md: Tools must be a list or string
```

**Solution**:
Use either format:
```yaml
# String format (recommended)
tools: Bash, Edit, MultiEdit, Write, Read

# Array format
tools:
  - Bash
  - Edit
  - MultiEdit
  - Write
  - Read
```

### Invalid Domain or Role Values

**Problem**: Domain or role contains invalid values.

**Symptoms**:
```
❌ agents/python/my-agent.md: Invalid domain: javascript
❌ agents/python/my-agent.md: Invalid role: developer
```

**Solution**:
Use only valid values:
```yaml
# Valid domains
domain: python | dotnet | nodejs | java | core

# Valid roles
role: analyst | architect | engineer | test-engineer | documentation
```

## GitHub Actions Problems

### Workflow Not Triggering

**Problem**: GitHub Actions workflows don't trigger on issue events.

**Symptoms**:
- No workflow runs appear in Actions tab
- Issues don't get automatic labels
- Agent assignment comments don't appear

**Solution**:
1. **Check GitHub Actions are enabled**:
   - Go to repository Settings → Actions → General
   - Ensure "Allow all actions and reusable workflows" is selected

2. **Verify trigger conditions**:
   ```yaml
   on:
     issues:
       types: [opened, edited, labeled]  # Must match your use case
   ```

3. **Check issue labels**:
   - Issue must have `agent-ready` label to trigger orchestration
   - Add label manually: `gh issue edit 123 --add-label "agent-ready"`

4. **Review workflow file syntax**:
   ```bash
   # Validate workflow syntax
   cat .github/workflows/issue-agent-orchestration.yml | yaml-validate
   ```

### Permissions Errors

**Problem**: Workflow fails with permissions errors.

**Symptoms**:
```
Error: Resource not accessible by integration
Error: Forbidden
```

**Solution**:
1. **Check repository permissions**:
   ```yaml
   permissions:
     issues: write
     contents: write
     pull-requests: write
     actions: write
   ```

2. **Verify GITHUB_TOKEN scope**:
   - Go to Settings → Actions → General → Workflow permissions
   - Select "Read and write permissions"

3. **For private repositories**, may need Personal Access Token:
   ```yaml
   - uses: actions/github-script@v7
     with:
       github-token: ${{ secrets.PERSONAL_ACCESS_TOKEN }}
   ```

### Classification Script Failures

**Problem**: Domain classification script fails or returns incorrect results.

**Symptoms**:
```
Error: No module named 'yaml'
domain=unknown
phase=unknown
```

**Solution**:
1. **Install Python dependencies**:
   ```yaml
   - name: Install dependencies
     run: |
       pip install pyyaml markdown requests
   ```

2. **Check script permissions**:
   ```bash
   chmod +x scripts/classify-domain.py
   ```

3. **Test classification locally**:
   ```bash
   python3 scripts/classify-domain.py \
     --title "Create FastAPI authentication" \
     --body "Need JWT-based auth with FastAPI"
   ```

## Domain Classification Issues

### Incorrect Domain Detection

**Problem**: Issues are classified into wrong technology domains.

**Symptoms**:
- Python issues get `domain:java` label
- Generic issues classified as specific technology

**Solution**:
1. **Improve issue description**:
   ```markdown
   # Better issue description
   Title: Implement FastAPI user authentication
   Body: Create JWT-based authentication system using FastAPI, SQLAlchemy, and pytest for testing.
   ```

2. **Update classification keywords**:
   Edit `scripts/classify-domain.py`:
   ```python
   self.domain_keywords = {
       'python': [
           'python', 'fastapi', 'django', 'flask', 'pytest',
           'sqlalchemy', 'pydantic'  # Add more keywords
       ]
   }
   ```

3. **Use explicit labels**:
   Add domain label manually: `gh issue edit 123 --add-label "domain:python"`

### Low Confidence Scores

**Problem**: Classification confidence is too low, leading to incorrect assignments.

**Symptoms**:
```
domain=python
confidence=0.3
```

**Solution**:
1. **Add more domain-specific keywords** to issue description
2. **Mention specific frameworks and tools**
3. **Include file extensions** (`.py`, `.cs`, `.js`)
4. **Update classification weights** in the script

## Workflow State Problems

### State Transition Failures

**Problem**: Workflow state doesn't advance properly.

**Symptoms**:
```bash
$ python3 scripts/track-workflow.py --issue 123 --progress
Issue #123 Progress:
  Step: 2/9
  Progress: 22.2%
  Current State: requirements-analysis
  # Stuck at same state
```

**Solution**:
1. **Check valid transitions**:
   ```python
   # Valid state transitions
   'requirements-analysis' → 'requirements-ready'
   'requirements-ready' → 'plan-approved'
   'plan-approved' → 'tests-planned'
   ```

2. **Force state update if needed**:
   ```bash
   python3 scripts/track-workflow.py \
     --issue 123 \
     --state requirements-ready \
     --agent requirements-analyst
   ```

3. **Reset workflow state**:
   ```bash
   # Remove workflow state file to reset
   rm .workflow-state.json
   ```

### Blocked Issues Detection

**Problem**: Issues appear blocked but should be progressing.

**Symptoms**:
```bash
$ python3 scripts/track-workflow.py --blocked
Blocked Issues:
  Issue #123: requirements-analysis (25.5h stalled)
```

**Solution**:
1. **Manual intervention required** - Check issue for blockers
2. **Update state manually** if work was completed outside workflow
3. **Adjust blocking threshold** in script if too sensitive

## Installation Issues

### Agent Installation Failures

**Problem**: Agents fail to install to Claude directory.

**Symptoms**:
```
Error: Permission denied
Error: Directory not found
```

**Solution**:
1. **Check directory permissions**:
   ```bash
   mkdir -p ~/.claude/agents
   chmod 755 ~/.claude/agents
   ```

2. **Use Task installation** (recommended):
   ```bash
   task install
   ```

3. **Manual installation**:
   ```bash
   cp agents/**/*.md ~/.claude/agents/
   ```

### Missing Dependencies

**Problem**: Validation scripts fail due to missing Python packages.

**Symptoms**:
```
ModuleNotFoundError: No module named 'yaml'
```

**Solution**:
```bash
# Install required packages
pip3 install pyyaml markdown requests

# Or use requirements file if available
pip3 install -r requirements.txt
```

### Script Permission Issues

**Problem**: Scripts are not executable.

**Symptoms**:
```
Permission denied: ./scripts/validate-agent-spec.sh
```

**Solution**:
```bash
# Make scripts executable
chmod +x scripts/*.sh
chmod +x scripts/*.py

# Or run directly with interpreter
python3 scripts/validate-claude-agent.py
bash scripts/validate-agent-spec.sh
```

## Performance Issues

### Slow Validation

**Problem**: Agent validation takes too long.

**Symptoms**:
- Validation script runs for several minutes
- CI/CD pipeline times out

**Solution**:
1. **Validate specific agents only**:
   ```bash
   python3 scripts/validate-claude-agent.py agents/python/
   ```

2. **Use parallel validation** in CI:
   ```yaml
   strategy:
     matrix:
       domain: [python, dotnet, nodejs, java, core]
   ```

3. **Cache dependencies**:
   ```yaml
   - uses: actions/cache@v3
     with:
       path: ~/.cache/pip
       key: validation-deps-${{ runner.os }}
   ```

### Workflow Timeout Issues

**Problem**: GitHub Actions workflows timeout.

**Symptoms**:
```
Error: The job running on runner GitHub Actions X has exceeded the maximum execution time of 360 minutes.
```

**Solution**:
1. **Reduce workflow complexity**
2. **Add timeout settings**:
   ```yaml
   jobs:
     validate:
       timeout-minutes: 30
   ```

3. **Use workflow concurrency**:
   ```yaml
   concurrency:
     group: ${{ github.workflow }}-${{ github.ref }}
     cancel-in-progress: true
   ```

## Common Error Messages

### "Agent specification not found"

**Problem**: Workflow can't locate agent specification file.

**Solution**:
1. **Check file exists**:
   ```bash
   ls agents/python/solution-architect-python.md
   ```

2. **Verify naming convention**:
   - File name should match agent name
   - Use correct domain directory

3. **Check file permissions**:
   ```bash
   chmod 644 agents/**/*.md
   ```

### "Invalid state transition"

**Problem**: Attempting invalid workflow state transition.

**Solution**:
1. **Check current state**:
   ```bash
   python3 scripts/track-workflow.py --issue 123 --get-state
   ```

2. **View valid next states**:
   ```bash
   python3 scripts/track-workflow.py --issue 123 --progress
   ```

3. **Follow proper sequence**:
   ```
   issue-created → requirements-analysis → requirements-ready →
   plan-approved → tests-planned → implementation-ready →
   implementation-complete → user-accepted → documentation-complete
   ```

### "No module named 'github'"

**Problem**: GitHub API library not installed.

**Solution**:
```bash
pip3 install PyGithub requests
```

### "YAML frontmatter missing"

**Problem**: Agent file doesn't start with YAML frontmatter.

**Solution**:
Ensure file starts with:
```yaml
---
name: agent-name
# ... other required fields
---

# Agent content below
```

## Getting Additional Help

### Debug Mode

Enable debug logging for detailed troubleshooting:

```bash
# Set debug environment variables
export ACTIONS_STEP_DEBUG=true
export ACTIONS_RUNNER_DEBUG=true

# Run validation with verbose output
python3 scripts/validate-claude-agent.py --format json

# Track workflow with detailed output
python3 scripts/track-workflow.py --issue 123 --format json
```

### Log Analysis

Check GitHub Actions logs for detailed error information:

```bash
# List recent workflow runs
gh run list --workflow="validate-agents.yml"

# View specific run logs
gh run view <run-id> --log
```

### Community Support

1. **Check existing GitHub issues** for similar problems
2. **Create new issue** with:
   - Error message
   - Steps to reproduce
   - Environment details
   - Relevant configuration files

3. **Include debug information**:
   - Workflow run logs
   - Agent validation output
   - Classification script results
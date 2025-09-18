# Migration Guide

This guide helps you migrate from previous versions of Claude Agent Templates to version 1.0.0, which introduces significant changes including domain-specific agents, new specification formats, and GitHub automation.

## Table of Contents

- [Overview of Changes](#overview-of-changes)
- [Breaking Changes](#breaking-changes)
- [Migration Steps](#migration-steps)
- [Specification Format Updates](#specification-format-updates)
- [New Features](#new-features)
- [Backward Compatibility](#backward-compatibility)
- [Troubleshooting Migration](#troubleshooting-migration)

## Overview of Changes

Version 1.0.0 represents a major architectural upgrade with the following key changes:

### Major Additions
- **Domain-Specific Agents**: 12 new agents across Python, .NET, Node.js, and Java domains
- **GitHub Actions Automation**: Complete CI/CD pipeline for issue-driven development
- **Enhanced Specifications**: Structured YAML frontmatter with comprehensive metadata
- **Validation Framework**: Comprehensive validation tools and quality gates
- **Workflow Management**: 9-step workflow tracking with state management

### Repository Structure Changes
```
# Before (v0.1.x)
agents/
├── core/
└── python/

# After (v1.0.0)
agents/
├── core/
├── python/
├── dotnet/
├── nodejs/
└── java/
```

## Breaking Changes

### 1. Agent Specification Format

**Before (v0.1.x)**:
```yaml
---
name: software-engineer-python
description: Expert Python software engineer...
tools: Bash, Edit, MultiEdit, Write, Read
model: inherit
color: blue
---
```

**After (v1.0.0)**:
```yaml
---
name: software-engineer-python
description: Expert Python software engineer...
domain: python
role: engineer
spec_version: "1.0"
tools: Bash, Edit, MultiEdit, Write, Read, Glob, Grep, LS, WebFetch, WebSearch, NotebookEdit, TodoWrite, BashOutput, KillBash
model: inherit
color: blue
inputs:
  - GitHub issues with plan-approved label
  - Architectural plans from solution-architect-python
outputs:
  - Implemented Python solution following hexagonal architecture
  - GitHub pull requests with implementation
validation:
  - Code quality and PEP 8 compliance
  - Test coverage requirements (>80%)
dependencies:
  - Python 3.11+ runtime environment
  - pytest testing framework
workflow_position: 6
github_integration:
  triggers: ["plan-approved", "tests-planned"]
  outputs: ["implementation-complete", "ready-for-review"]
  permissions: ["contents:write", "pull_requests:write"]
examples:
  - context: User has a plan-approved issue that needs implementation
    input: "The architect has approved the plan for issue #123"
    output: "Implement the approved architectural plan"
---
```

### 2. Repository Structure

**New directories and files**:
- `agents/dotnet/` - .NET specific agents
- `agents/nodejs/` - Node.js specific agents
- `agents/java/` - Java specific agents
- `.github/workflows/` - GitHub Actions automation
- `scripts/` - Validation and automation scripts
- `tests/` - Comprehensive test suites
- `specs/` - Technical specifications

### 3. Installation Process

**Before**:
```bash
./scripts/install-agents.sh
```

**After**:
```bash
task install  # Recommended
# OR
./scripts/install-agents.sh  # Still supported
```

## Migration Steps

### Step 1: Backup Current Installation

```bash
# Backup existing agents
mkdir -p backup/agents
cp -r ~/.claude/agents/* backup/agents/

# Backup any custom configurations
cp -r agents/ backup/repository-agents/
```

### Step 2: Update Repository

```bash
# Pull latest changes
git fetch origin
git checkout main
git pull origin main

# Or clone fresh copy
git clone https://github.com/dkoenawan/claude-agent-templates.git
cd claude-agent-templates
```

### Step 3: Validate Current Agents

```bash
# Check which agents need updating
python3 scripts/validate-claude-agent.py

# Review validation results
python3 scripts/validate-claude-agent.py --format json > validation-results.json
```

### Step 4: Update Agent Specifications

For each existing agent that fails validation:

1. **Add required fields**:
   ```yaml
   domain: python  # or appropriate domain
   role: engineer  # or appropriate role
   spec_version: "1.0"
   inputs: [...]
   outputs: [...]
   validation: [...]
   dependencies: [...]
   workflow_position: 6  # appropriate position
   github_integration: {...}
   examples: [...]
   ```

2. **Update tools list** (if needed):
   ```yaml
   tools: Bash, Edit, MultiEdit, Write, Read, Glob, Grep, LS, WebFetch, WebSearch, NotebookEdit, TodoWrite, BashOutput, KillBash
   ```

3. **Add examples** (required):
   ```yaml
   examples:
     - context: "Description of when to use this agent"
       input: "Example input or trigger"
       output: "Expected output or result"
   ```

### Step 5: Install Updated Agents

```bash
# Install using Task (recommended)
task install

# Verify installation
task list
task validate
```

### Step 6: Test Migration

```bash
# Test agent validation
python3 scripts/validate-claude-agent.py

# Test domain classification
python3 scripts/classify-domain.py \
  --title "Test FastAPI development" \
  --body "Create API with Python FastAPI"

# Test workflow tracking
python3 scripts/track-workflow.py --issue 1 --state issue-created
python3 scripts/track-workflow.py --issue 1 --progress
```

## Specification Format Updates

### Required Field Additions

Add these required fields to existing agent specifications:

```yaml
# Domain classification
domain: python|dotnet|nodejs|java|core

# Role classification
role: analyst|architect|engineer|test-engineer|documentation

# Specification version
spec_version: "1.0"

# Input specifications
inputs:
  - List of expected inputs
  - Be specific about GitHub issue labels
  - Include prerequisite artifacts

# Output specifications
outputs:
  - List of expected outputs
  - Include artifacts created
  - Specify GitHub label updates

# Validation criteria
validation:
  - Quality requirements
  - Coverage thresholds
  - Completeness checks

# Dependencies
dependencies:
  - Runtime requirements
  - Tool dependencies
  - Access requirements

# Workflow position (1-9)
workflow_position: 6

# GitHub integration
github_integration:
  triggers: ["triggering-labels"]
  outputs: ["output-labels"]
  permissions: ["required-permissions"]

# Examples (at least one required)
examples:
  - context: "When to use this agent"
    input: "Example input"
    output: "Expected output"
```

### Field Validation Rules

**Domain Values**:
- `core` - Language-agnostic agents
- `python` - Python ecosystem agents
- `dotnet` - .NET ecosystem agents
- `nodejs` - Node.js ecosystem agents
- `java` - Java ecosystem agents

**Role Values**:
- `analyst` - Requirements analysis
- `architect` - Solution architecture
- `engineer` - Implementation
- `test-engineer` - Testing strategy
- `documentation` - Documentation

**Workflow Positions**:
- `2` - Requirements analysis
- `4` - Solution architecture
- `5` - Test planning
- `6` - Implementation
- `9` - Documentation

### Example Migration

**Before**:
```yaml
---
name: test-engineer-python
description: Use this agent to create comprehensive unit test plans...
tools: Bash, Glob, Grep, LS, Read, WebFetch, TodoWrite, WebSearch, BashOutput, KillBash
model: inherit
color: green
---
```

**After**:
```yaml
---
name: test-engineer-python
description: Use this agent to create comprehensive unit test plans for Python projects after architectural planning is complete.
domain: python
role: test-engineer
spec_version: "1.0"
tools: Bash, Glob, Grep, LS, Read, WebFetch, TodoWrite, WebSearch, BashOutput, KillBash
model: inherit
color: green
inputs:
  - GitHub issues with plan-approved label
  - Architectural plans with hexagonal architecture
  - Implementation specifications and acceptance criteria
outputs:
  - Comprehensive test strategy document
  - Test implementation with >80% coverage
  - Pytest fixtures and mock configurations
  - GitHub issue updates with tests-planned label
validation:
  - Test coverage requirements verification
  - Python testing best practices compliance
  - Hexagonal architecture test isolation verification
dependencies:
  - Python 3.11+ runtime environment
  - pytest testing framework
  - pytest-cov for coverage reporting
  - pytest-mock for mocking
workflow_position: 5
github_integration:
  triggers: ["plan-approved"]
  outputs: ["tests-planned"]
  permissions: ["contents:write", "issues:write"]
examples:
  - context: User needs test planning after solution architect creates implementation plan
    input: "The solution architect has completed the plan for the user authentication system"
    output: "Analyze the architectural plan and create a comprehensive unit test strategy with pytest fixtures and coverage recommendations"
---
```

## New Features

### 1. Domain-Specific Agents

Take advantage of new specialized agents:

**Python Domain**:
- `solution-architect-python` - Python architecture planning
- `software-engineer-python` - Enhanced implementation
- `test-engineer-python` - Enhanced testing

**.NET Domain**:
- `solution-architect-dotnet` - ASP.NET Core architecture
- `software-engineer-dotnet` - C# implementation
- `test-engineer-dotnet` - .NET testing strategies

**Node.js Domain**:
- `solution-architect-nodejs` - TypeScript/Express architecture
- `software-engineer-nodejs` - Modern JavaScript implementation
- `test-engineer-nodejs` - Node.js testing patterns

**Java Domain**:
- `solution-architect-java` - Spring Boot architecture
- `software-engineer-java` - Java implementation
- `test-engineer-java` - Java testing frameworks

### 2. GitHub Actions Integration

Set up automated workflows:

1. **Copy workflow files**:
   ```bash
   cp -r .github/workflows/ /path/to/your/project/.github/
   ```

2. **Create required labels**:
   ```bash
   # Use the provided script
   bash docs/github-actions-setup.md  # Follow label creation section
   ```

3. **Test automation**:
   - Create test issue
   - Add `agent-ready` label
   - Observe automated classification and assignment

### 3. Validation Tools

Use new validation capabilities:

```bash
# Comprehensive agent validation
python3 scripts/validate-claude-agent.py

# Domain classification testing
python3 scripts/classify-domain.py --title "Your issue title" --body "Issue description"

# Workflow progress tracking
python3 scripts/track-workflow.py --issue 123 --progress
```

## Backward Compatibility

### Maintained Compatibility

- **Existing agent files** continue to work but with validation warnings
- **Installation scripts** (`.sh` and `.bat`) still functional
- **Basic workflow** remains the same
- **Claude Code integration** unchanged

### Deprecated Features

- **Legacy specification format** (works but generates warnings)
- **Manual installation scripts** (prefer Task-based installation)
- **Hardcoded agent assignments** (use automated classification)

### Gradual Migration

You can migrate gradually:

1. **Start with validation** to identify issues
2. **Update agents one by one** as needed
3. **Add new domain agents** when relevant
4. **Enable GitHub automation** when ready

## Troubleshooting Migration

### Common Migration Issues

**1. Validation Failures**

```bash
# Check specific validation issues
python3 scripts/validate-claude-agent.py agents/python/my-agent.md

# Fix YAML syntax issues
# Ensure all required fields present
# Use correct enumerated values
```

**2. Installation Problems**

```bash
# Clear existing installation
rm -rf ~/.claude/agents/*

# Reinstall with Task
task clean
task install

# Verify installation
task list
```

**3. GitHub Actions Not Working**

```bash
# Check workflow files copied correctly
ls -la .github/workflows/

# Verify required labels exist
gh label list | grep "agent-ready"

# Test classification manually
python3 scripts/classify-domain.py --title "Test" --body "Python development"
```

### Migration Validation

After migration, run these checks:

```bash
# 1. Validate all agents
python3 scripts/validate-claude-agent.py

# 2. Check installation
task validate

# 3. Test classification
python3 scripts/classify-domain.py \
  --title "Create FastAPI authentication" \
  --body "JWT-based auth with FastAPI and SQLAlchemy"

# 4. Test workflow tracking
python3 scripts/track-workflow.py --issue 1 --state issue-created
python3 scripts/track-workflow.py --report

# 5. Verify agent availability in Claude Code
# Should see all agents available when using Task tool
```

### Rollback Procedure

If migration causes issues:

1. **Restore agent backup**:
   ```bash
   rm -rf ~/.claude/agents/*
   cp -r backup/agents/* ~/.claude/agents/
   ```

2. **Revert repository changes**:
   ```bash
   git checkout v0.1.1  # Previous version
   ```

3. **Use legacy installation**:
   ```bash
   ./scripts/install-agents.sh
   ```

### Getting Help

For migration issues:

1. **Check troubleshooting guide** for specific error messages
2. **Review validation output** for detailed error descriptions
3. **Test with minimal examples** to isolate issues
4. **Create GitHub issue** with migration details if needed

Include in support requests:
- Previous version used
- Validation output
- Error messages
- Steps attempted
- Environment details (OS, Python version, etc.)
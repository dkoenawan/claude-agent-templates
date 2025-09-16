# Agent Chaining Patterns & Workflow Documentation

## Table of Contents
1. [Overview](#overview)
2. [Core Chaining Patterns](#core-chaining-patterns)
3. [The 3-Agent Solution Architecture Chain](#the-3-agent-solution-architecture-chain)
4. [Real Workflow Examples](#real-workflow-examples)
5. [Best Practices for Agent Coordination](#best-practices-for-agent-coordination)
6. [Troubleshooting Guide](#troubleshooting-guide)
7. [Advanced Patterns](#advanced-patterns)

## Overview

Agent chaining is a powerful pattern that enables complex software development workflows by coordinating specialized AI agents in sequence or parallel. Each agent focuses on its area of expertise while passing context and results to subsequent agents in the chain.

### Key Benefits
- **Specialization**: Each agent excels at specific tasks
- **Modularity**: Chains can be customized for different workflows
- **Traceability**: GitHub issues track the entire process
- **Quality**: Multiple specialized reviews improve outcomes
- **Efficiency**: Parallel execution where possible

### Core Concepts
- **Sequential Chaining**: Agents execute one after another
- **Parallel Chaining**: Multiple agents work simultaneously
- **Conditional Chaining**: Agents trigger based on conditions
- **Recursive Chaining**: Agents can call themselves or others iteratively

## Core Chaining Patterns

### 1. Sequential Chain Pattern
The most common pattern where agents execute in a defined order, each building on the previous agent's output.

```
User Request → Agent A → Agent B → Agent C → Result
```

**Example**: Requirements → Architecture → Testing → Implementation

### 2. Parallel Chain Pattern
Multiple agents work simultaneously on different aspects of the same problem.

```
          ┌→ Agent A →┐
User →────┼→ Agent B →┼──→ Aggregator → Result
          └→ Agent C →┘
```

**Example**: Parallel analysis of frontend, backend, and database requirements

### 3. Conditional Chain Pattern
Agents are triggered based on specific conditions or decision points.

```
User → Agent A → Decision → [Condition Met] → Agent B
                          → [Else] → Agent C
```

**Example**: If tests fail → test-engineer-python, else → software-engineer-python

### 4. Recursive Chain Pattern
Agents can trigger themselves or other agents multiple times until a condition is met.

```
User → Agent A → [Not Complete] → Agent A (refined)
               → [Complete] → Result
```

**Example**: Iterative refinement of requirements until clarity achieved

## The 3-Agent Solution Architecture Chain

This is our primary workflow for turning business requirements into production-ready code with comprehensive testing.

### Chain Structure
```
requirements-analyst → solution-architect → test-engineer-python → software-engineer-python
```

### Detailed Flow

#### Phase 1: Requirements Analysis
**Agent**: `requirements-analyst`
**Input**: GitHub issue with business requirements
**Process**:
1. Analyzes the issue description
2. Identifies ambiguities or gaps
3. Posts clarifying questions as issue comments
4. Waits for user responses
5. Synthesizes final requirements document

**Output**: Clear technical requirements with acceptance criteria
**Labels Applied**: `requirements-gathering`, `requirements-ready`

#### Phase 2: Solution Architecture
**Agent**: `solution-architect`
**Input**: Requirements-ready GitHub issue
**Process**:
1. Reviews finalized requirements
2. Analyzes existing codebase
3. Designs hexagonal architecture solution
4. Creates atomic work functions
5. Posts comprehensive plan as issue comment

**Output**: Detailed implementation plan with architecture diagrams
**Labels Applied**: `architecture-planning`, `plan-approved`

#### Phase 3: Test Strategy
**Agent**: `test-engineer-python`
**Input**: Plan-approved GitHub issue
**Process**:
1. Reviews architectural plan
2. Designs comprehensive test strategy
3. Creates pytest fixtures and mocks
4. Defines coverage targets (80% minimum)
5. Posts test plan as issue comment

**Output**: Complete test strategy with example tests
**Labels Applied**: `test-planning`, `tests-planned`

#### Phase 4: Implementation
**Agent**: `software-engineer-python`
**Input**: Tests-planned GitHub issue
**Process**:
1. Creates feature/bugfix branch
2. Implements solution following architecture
3. Writes comprehensive tests
4. Ensures 80% test coverage
5. Creates pull request with detailed description

**Output**: Pull request with tested implementation
**Labels Applied**: `implementation`, `pr-created`

### Example GitHub Issue Flow

```yaml
Issue #123: Add user authentication system

1. User creates issue with requirements
2. requirements-analyst adds comment:
   "I need clarification on:
    - Should we support OAuth providers?
    - What session timeout is required?
    - Do we need 2FA support?"

3. User responds with clarifications

4. solution-architect adds comment:
   "Architecture Plan:
    - Hexagonal architecture with ports/adapters
    - JWT tokens with Redis session store
    - OAuth2 integration via adapter pattern
    [detailed implementation plan]"

5. User approves plan

6. test-engineer-python adds comment:
   "Test Strategy:
    - Unit tests for auth service (30 tests)
    - Integration tests for OAuth providers (10 tests)
    - E2E tests for login flows (5 tests)
    - Fixtures for mock users and tokens
    Target coverage: 85%"

7. software-engineer-python:
   - Creates branch: feature/123-user-authentication
   - Implements solution with tests
   - Creates PR #456 linking to issue #123
```

## Real Workflow Examples

### Example 1: Bug Fix Workflow

**Scenario**: Production bug causing payment failures

```bash
# Step 1: User creates issue
gh issue create --title "Payment processing fails for amounts over $1000" \
  --body "Users report payment failures when amount exceeds $1000. Error: 'Invalid amount format'"

# Step 2: Launch requirements analyst
claude "Use requirements-analyst to analyze issue #45 and identify root cause"

# Step 3: After requirements are clear, launch solution architect
claude "Use solution-architect to create fix plan for issue #45"

# Step 4: Create test strategy
claude "Use test-engineer-python to create test plan for the payment bug fix in issue #45"

# Step 5: Implement the fix
claude "Use software-engineer-python to implement the approved fix for issue #45"
```

### Example 2: New Feature Development

**Scenario**: Adding real-time notifications

```bash
# Step 1: Create comprehensive feature request
gh issue create --title "Add real-time notification system" \
  --body "## Requirements
  - Support email, SMS, and push notifications
  - User preference management
  - Delivery tracking and retry logic
  - Rate limiting per user"

# Step 2: Sequential agent chain
claude "Use requirements-analyst to break down the notification system requirements in issue #67"

# After requirements clarification
claude "Use solution-architect to design the notification system architecture for issue #67"

# After architecture approval
claude "Use test-engineer-python to create comprehensive test strategy for issue #67"

# Finally implement
claude "Use software-engineer-python to implement the notification system from issue #67"
```

### Example 3: Parallel Analysis Workflow

**Scenario**: Analyzing multiple system components simultaneously

```bash
# Launch multiple analysts in parallel for different aspects
claude "Run these agents in parallel:
1. Use requirements-analyst on issue #89 focusing on API requirements
2. Use framework-architect to evaluate notification service options
3. Use minimal-work-identifier to define MVP scope"

# Aggregate results and proceed
claude "Based on the parallel analysis results, use solution-architect to create unified plan for issue #89"
```

### Example 4: Iterative Refinement Workflow

**Scenario**: Complex feature requiring multiple refinement cycles

```bash
# Initial analysis
claude "Use requirements-analyst on issue #101 about the recommendation engine"

# User provides partial answers, needs more analysis
claude "Use requirements-analyst again on issue #101 to dive deeper into ML model requirements"

# After multiple rounds, proceed with architecture
claude "Now that requirements are clear, use solution-architect for issue #101"

# Architecture reveals additional considerations
claude "Use adjacent-work-recommender to identify related improvements for the recommendation system"

# Refine architecture with new insights
claude "Update the architecture plan for issue #101 incorporating the adjacent work recommendations"
```

## Best Practices for Agent Coordination

### 1. Clear Handoffs
- Always ensure the previous agent's output is complete before proceeding
- Use GitHub issue comments for traceability
- Apply appropriate labels at each stage

### 2. Context Preservation
- Pass issue numbers explicitly to each agent
- Reference previous agent outputs in prompts
- Maintain comment threads in issues

### 3. Error Handling
```bash
# Always check agent completion status
claude "Use requirements-analyst on issue #123"
# If blocked or incomplete:
claude "Review the requirements-analyst output and address any blocking issues"
```

### 4. Parallel Execution
When tasks are independent, run agents in parallel:
```bash
claude "Run in parallel:
1. test-engineer-python for backend tests on issue #45
2. test-engineer-python for frontend tests on issue #45
3. documentation agent to update API docs"
```

### 5. Validation Checkpoints
- After requirements: Verify completeness
- After architecture: Confirm feasibility
- After tests: Validate coverage targets
- After implementation: Ensure all tests pass

### 6. Label Management
Proper labeling enables workflow tracking:
```yaml
requirements-gathering → requirements-ready
architecture-planning → plan-approved
test-planning → tests-planned
implementation → pr-created
pr-review → pr-approved
merged → documentation-needed
```

## Troubleshooting Guide

### Common Issues and Solutions

#### 1. Agent Produces Incomplete Output
**Symptoms**: Agent stops mid-task or provides partial results
**Solution**:
```bash
# Re-run with more specific instructions
claude "Use solution-architect on issue #45, ensure you complete all sections including database design"
```

#### 2. Context Lost Between Agents
**Symptoms**: Second agent doesn't reference first agent's work
**Solution**:
```bash
# Explicitly reference previous outputs
claude "Use test-engineer-python on issue #45, specifically create tests for the architecture plan in comment #3"
```

#### 3. Circular Dependencies
**Symptoms**: Agents waiting on each other
**Solution**:
```bash
# Break the cycle with explicit sequencing
claude "First, complete requirements analysis for issue #45"
claude "Now proceed with architecture regardless of pending questions"
```

#### 4. Label Conflicts
**Symptoms**: Multiple agents applying conflicting labels
**Solution**:
```bash
# Check current labels
gh issue view 45 --json labels

# Reset labels if needed
gh issue edit 45 --remove-label "requirements-gathering" --add-label "requirements-ready"
```

#### 5. Agent Skips GitHub Integration
**Symptoms**: Agent works locally but doesn't update issue
**Solution**:
```bash
# Ensure gh CLI is authenticated
gh auth status

# Explicitly instruct GitHub interaction
claude "Use solution-architect on issue #45 and post the plan as an issue comment using gh"
```

#### 6. Parallel Execution Conflicts
**Symptoms**: Agents interfere with each other's work
**Solution**:
```bash
# Ensure agents work on different aspects
claude "Run in parallel:
1. test-engineer-python for unit tests only
2. documentation agent for README updates only"
```

### Debug Commands

```bash
# View agent execution history
gh issue view 45 --comments

# Check label progression
gh issue list --label "requirements-ready" --state open

# Verify PR linkage
gh pr list --search "linked:45"

# Review agent outputs
gh issue view 45 --json comments --jq '.comments[] | select(.author.login == "github-actions")'
```

## Advanced Patterns

### 1. Branching Chains
Different paths based on issue type:
```python
if "bug" in issue_labels:
    chain = ["requirements-analyst", "test-engineer-python", "software-engineer-python"]
elif "feature" in issue_labels:
    chain = ["requirements-analyst", "solution-architect", "test-engineer-python", "software-engineer-python"]
elif "refactor" in issue_labels:
    chain = ["solution-architect", "test-engineer-python", "software-engineer-python"]
```

### 2. Quality Gates
Enforce standards between agents:
```bash
# After test planning, verify coverage targets
claude "Review test plan in issue #45 and confirm it meets 80% coverage target"

# Only proceed if confirmed
claude "Use software-engineer-python to implement issue #45"
```

### 3. Feedback Loops
Iterative improvement cycles:
```bash
# Initial implementation
claude "Use software-engineer-python for issue #45"

# Review and refine
claude "Use code-reviewer agent on PR #123"

# If issues found, loop back
claude "Use software-engineer-python to address review comments on PR #123"
```

### 4. Cross-Repository Chains
Coordinate changes across multiple repositories:
```bash
# Analyze impact across repos
claude "Use requirements-analyst to identify cross-repo dependencies for issue #45"

# Create issues in related repos
gh issue create --repo related/service --title "Update API for issue main#45"

# Coordinate implementation
claude "Use software-engineer-python to implement coordinated changes across repos"
```

### 5. Time-Boxed Chains
Enforce time limits for each stage:
```bash
# Set deadline for requirements
claude "Use requirements-analyst on issue #45, complete within 30 minutes"

# Escalate if blocked
claude "If requirements not ready, use minimal-work-identifier to define MVP scope"
```

## Conclusion

Agent chaining transforms complex software development into manageable, specialized workflows. By understanding these patterns and best practices, you can:

- Build more reliable software with comprehensive testing
- Maintain clear traceability through GitHub issues
- Scale development efforts through parallel execution
- Ensure quality through specialized expertise

Remember: The key to successful agent chaining is clear communication, proper sequencing, and maintaining context throughout the workflow.
# GitHub Issue Workflow Labels

## Overview

The Claude Agent Templates repository uses a standardized labeling system to track GitHub issues through the 9-step development workflow. Each agent is responsible for applying specific labels to indicate workflow progression and enable proper agent handoffs.

## Workflow Progression Labels

### Step 2: Requirements Analysis
- **Label**: `requirements-ready`
- **Applied by**: Requirements Analyst
- **Indicates**: Business requirements fully clarified and documented
- **Next Step**: Triggers Solution Architect review

### Step 4: Solution Architecture  
- **Label**: `plan-approved`
- **Applied by**: Solution Architect (after user approval)
- **Indicates**: Implementation plan created and approved by user
- **Next Step**: Triggers Test Engineer planning

### Step 6: Test Planning
- **Label**: `tests-planned`
- **Applied by**: Test Engineer Python
- **Indicates**: Comprehensive unit test strategy completed
- **Next Step**: Triggers Software Engineer implementation

### Step 7: Implementation
- **Label**: `implementation-complete`
- **Applied by**: Software Engineer Python
- **Indicates**: Code implemented, tested, and PR created
- **Next Step**: User review via GitHub issue or PR

### Step 9: Documentation
- **Label**: `documentation-complete`
- **Applied by**: Documentation Agent
- **Indicates**: Final documentation updates completed
- **Next Step**: Issue closure with `resolved` label

## Status Labels

### Issue States
- **Label**: `resolved`
- **Applied by**: Documentation Agent
- **Indicates**: Complete workflow finished, issue can be closed

### Issue Types (Applied at creation)
- `bug`: Bug fix requests
- `feature`: New feature requests  
- `enhancement`: Improvements to existing features
- `documentation`: Documentation-only changes

### Priority Labels
- `priority-high`: Critical issues requiring immediate attention
- `priority-medium`: Standard priority issues
- `priority-low`: Nice-to-have improvements

## Agent Responsibility Matrix

| Agent | Primary Label Applied | Secondary Labels |
|-------|---------------------|------------------|
| Requirements Analyst | `requirements-ready` | Issue type validation |
| Solution Architect | `plan-approved` | Priority assessment |
| Test Engineer Python | `tests-planned` | Coverage indicators |  
| Software Engineer Python | `implementation-complete` | Branch references |
| Documentation Agent | `documentation-complete`, `resolved` | Final cleanup |

## Error Recovery Labels

### Blocked States
- **Label**: `blocked-requirements`
- **Applied by**: Any agent
- **Indicates**: Cannot proceed due to unclear requirements
- **Action**: Return to Requirements Analyst

- **Label**: `blocked-architecture`
- **Applied by**: Any agent  
- **Indicates**: Cannot proceed due to architectural issues
- **Action**: Return to Solution Architect

- **Label**: `blocked-implementation`
- **Applied by**: Software Engineer
- **Indicates**: Implementation blocked by external dependencies
- **Action**: User intervention required

## Usage Guidelines

### For Agents
1. **Always apply your workflow label** when completing your phase
2. **Verify previous agent labels** before proceeding
3. **Add blocking labels** if unable to proceed
4. **Include label references** in GitHub issue comments

### For Users
1. **Monitor workflow progression** through labels
2. **Intervene on blocking labels** to provide clarification
3. **Review issues by label** to understand current status
4. **Use labels for issue filtering** and project management

## Label Application Examples

### Successful Workflow
```
[Initial] bug, priority-medium
[Step 2] bug, priority-medium, requirements-ready  
[Step 4] bug, priority-medium, requirements-ready, plan-approved
[Step 6] bug, priority-medium, requirements-ready, plan-approved, tests-planned
[Step 7] bug, priority-medium, requirements-ready, plan-approved, tests-planned, implementation-complete
[Step 9] bug, priority-medium, requirements-ready, plan-approved, tests-planned, implementation-complete, documentation-complete, resolved
```

### Blocked Workflow
```
[Initial] feature, priority-high
[Step 2] feature, priority-high, blocked-requirements
[Resolution] feature, priority-high, requirements-ready
[Continue workflow...]
```

## Implementation Notes

- All agents MUST update labels when completing their workflow phase
- Labels provide visibility into workflow state for users and other agents
- Blocking labels trigger error recovery processes
- Final `resolved` label indicates successful workflow completion
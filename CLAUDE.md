# Claude Agent Templates - Development Guidelines

## GitHub Issue-Driven Development Workflow

This repository implements a structured 9-step development workflow where all agents interact through GitHub issues to ensure complete traceability and collaboration.

### Workflow Overview
1. **User** raises new bug or feature request via GitHub issue
2. **Requirements Analyst** reviews requirements, asks clarifying questions via issue comments
3. **User** provides answers and clarifications in issue
4. **Solution Architect** creates comprehensive implementation plan following best practices
5. **User** reviews and accepts the architectural plan
6. **Test Engineer Python** creates comprehensive unit test strategy with pytest
7. **Software Engineer Python** implements solution with integrated testing (branch management, commits, PR creation)
8. **User** accepts implementation via GitHub issue or PR review
9. **Documentation Agent** performs final documentation updates and repository cleanup

### Agent Interaction Flow
All agents operate exclusively through GitHub issues using `gh` commands for complete workflow traceability.

## Available Agents

### requirements-analyst
**Workflow Position**: Step 2 - Requirements analysis and clarification
Analyzes GitHub issues to extract business requirements and asks clarifying questions via issue comments. Labels issues as "requirements-ready" when complete.

**Tools:** Bash, Edit, MultiEdit, Write, NotebookEdit

### solution-architect  
**Workflow Position**: Step 4 - Architectural planning
Creates comprehensive implementation plans for requirements-ready issues. Designs hexagonal architecture solutions and posts detailed plans via issue comments. Labels issues as "plan-approved" after user acceptance.

**Tools:** Bash, Glob, Grep, LS, Read, WebFetch, TodoWrite, WebSearch, BashOutput, KillBash

### test-engineer-python
**Workflow Position**: Step 6 - Test strategy and planning
Creates comprehensive unit test plans for plan-approved issues using pytest. Analyzes implementation plans to design test coverage strategies, fixtures, and mocking approaches. Labels issues as "tests-planned" when complete.

**Tools:** Bash, Glob, Grep, LS, Read, WebFetch, TodoWrite, WebSearch, BashOutput, KillBash

### software-engineer-python
**Workflow Position**: Step 7 - Implementation and PR creation
Implements tests-planned issues using hexagonal architecture principles. Manages feature/bugfix branches, creates signed commits, and submits pull requests with comprehensive testing.

**Tools:** All tools for comprehensive software development

### documentation
**Workflow Position**: Step 9 - Final documentation and cleanup
Performs post-implementation documentation updates and repository cleanup after user accepts implementation. Updates README, API docs, closes issues, and cleans up branches.

**Tools:** Read, Glob, Grep, LS, WebFetch, WebSearch, Write, Edit, MultiEdit, Bash, TodoWrite

## Contributing Workflow

This repository follows a **trunk-based development** approach for maintaining clean, collaborative development:

### Branch Strategy
- **Main branch** (`main`) is the single source of truth
- Create **short-lived feature branches** from `main`
- Branch naming: `feature/agent-name` or `fix/issue-description`
- Keep branches focused on a single agent or improvement

### Development Process
1. **Branch from main**: `git checkout -b feature/new-agent`
2. **Make focused changes**: Work on one agent or improvement at a time
3. **Test locally**: Ensure agents work as expected
4. **Commit, push, and create PR**: Always commit changes, push to remote, and create PR to `main`
5. **Review & merge**: Quick review cycle, merge to main
6. **Delete branch**: Clean up after merge

### Best Practices
- **Small, focused PRs** - One agent or improvement per PR
- **Clear commit messages** - Describe what the agent does
- **Test your agents** - Verify they work in real projects
- **Update documentation** - Keep README current

### Agent Development
- Follow existing template structure
- Use descriptive names and clear descriptions
- Include appropriate tool restrictions
- Test with multiple project types before contributing

This approach ensures continuous integration while maintaining high quality and collaborative development.
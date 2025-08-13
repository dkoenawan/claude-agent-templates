# Claude Agent Templates - Development Guidelines

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
4. **Create PR**: Submit pull request to `main`
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
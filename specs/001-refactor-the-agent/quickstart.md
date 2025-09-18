# Quickstart: Agent Refactoring for Spec-Driven Development

## Overview
This quickstart guide demonstrates the refactored agent system that follows specification-driven development with GitHub Issues interface. Agents are now specialized by role and domain/tech stack, with automated validation and GitHub Actions integration.

## Prerequisites
- GitHub repository with issues enabled
- GitHub CLI (`gh`) installed and authenticated
- Task runner installed
- Access to Claude Code CLI or compatible agent system

## Quick Setup

### 1. Install Agent Specifications
```bash
# Clone or update the agent templates repository
git clone https://github.com/your-org/claude-agent-templates.git
cd claude-agent-templates

# Install agents to your environment
task install

# Verify installation
task list
task validate
```

### 2. Install GitHub Workflows (Optional)
```bash
# Install automation workflows to your repository
task install-workflows

# Or manually copy workflow files
cp .github/workflows/agent-*.yml /path/to/your/repo/.github/workflows/
```

## Basic Usage

### 1. Create a GitHub Issue
Create an issue in your repository with appropriate content. The agent system will automatically detect the issue type and trigger the appropriate agent.

**Example Issue for Python Development**:
```markdown
Title: Add user authentication system

Body:
## Description
We need to implement a user authentication system for our Python web application.

## Requirements
- User registration and login
- JWT token-based authentication
- Password reset functionality
- Role-based access control

## Acceptance Criteria
- Users can register with email and password
- Users can log in and receive JWT tokens
- Protected routes require valid tokens
- Admin users have additional permissions
```

### 2. Agent Processing Flow
1. **Issue Detection**: GitHub Actions detect the new issue
2. **Agent Selection**: System selects `requirements-analyst` based on issue content
3. **Requirements Analysis**: Agent processes issue and posts clarifying questions
4. **User Response**: User answers questions in issue comments
5. **Architecture Planning**: `solution-architect-python` creates implementation plan
6. **Test Planning**: `test-engineer-python` designs test strategy
7. **Implementation**: `software-engineer-python` implements the solution
8. **Documentation**: `documentation` agent updates project documentation

### 3. Monitor Progress
```bash
# Check workflow status
gh workflow list

# View issue comments and labels
gh issue view 123

# Check agent performance metrics
task metrics
```

## Agent Specialization Examples

### Domain-Specific Agents

#### Python Stack
- **requirements-analyst** (domain-agnostic)
- **solution-architect-python** (Python-specific architecture patterns)
- **test-engineer-python** (pytest, Python testing best practices)
- **software-engineer-python** (Python implementation)

#### .NET Stack
- **requirements-analyst** (domain-agnostic)
- **solution-architect-dotnet** (C#/.NET architecture patterns)
- **test-engineer-dotnet** (xUnit, .NET testing practices)
- **software-engineer-dotnet** (C#/.NET implementation)

### Agent Selection Logic
```yaml
# Automatic agent selection based on issue content and labels
issue_labels:
  - "python" → solution-architect-python
  - "dotnet" → solution-architect-dotnet
  - "nodejs" → solution-architect-nodejs
  - "requirements" → requirements-analyst (any domain)
  - "documentation" → documentation (any domain)
```

## Validation and Quality Assurance

### Automatic Validation
```bash
# Agent specifications are validated on PR creation
# GitHub Actions run validation workflows
# Agents validate issue content before processing
```

### Manual Validation
```bash
# Validate specific agent
task validate-agent requirements-analyst

# Validate all agents
task validate

# Test agent processing on sample issues
task test-agents
```

## Troubleshooting

### Common Issues

#### Agent Not Triggered
**Symptoms**: Issue created but no agent responds
**Solutions**:
1. Check issue has required labels
2. Verify GitHub Actions are enabled
3. Check agent specification for input requirements
4. Review workflow logs in GitHub Actions

#### Validation Failures
**Symptoms**: Agent reports validation errors
**Solutions**:
1. Review issue format requirements
2. Add missing required sections
3. Ensure proper labeling
4. Check previous workflow steps are completed

#### Performance Issues
**Symptoms**: Slow agent responses
**Solutions**:
1. Check GitHub API rate limits
2. Review agent performance metrics
3. Optimize issue content length
4. Consider load balancing for high-volume usage

### Debug Commands
```bash
# Check agent installation
task doctor

# Verbose validation output
task validate --verbose

# Test specific workflow step
task test-step requirements-analysis --issue 123

# View detailed metrics
task metrics --detailed --timeframe 24h
```

## Integration Testing

### End-to-End Test Scenarios

#### Scenario 1: Python Feature Development
1. Create issue with Python feature request
2. Verify `requirements-analyst` processes issue
3. Respond to clarifying questions
4. Verify `solution-architect-python` creates plan
5. Verify `test-engineer-python` creates test strategy
6. Verify `software-engineer-python` implements solution
7. Verify `documentation` agent updates docs

#### Scenario 2: Cross-Domain Documentation
1. Create documentation improvement issue
2. Verify `documentation` agent processes regardless of tech stack
3. Verify output quality and completeness

#### Scenario 3: Multi-Technology Project
1. Create issue affecting both Python and .NET components
2. Verify appropriate domain-specific agents are triggered
3. Verify coordination between different technology stacks

### Performance Benchmarks
- **Agent Selection**: < 5 seconds
- **Requirements Analysis**: < 2 minutes
- **Architecture Planning**: < 5 minutes
- **Test Planning**: < 3 minutes
- **Basic Implementation**: < 10 minutes
- **Documentation Updates**: < 2 minutes

## Next Steps

### Advanced Configuration
- Custom agent specialization for your domain
- Integration with existing CI/CD pipelines
- Custom validation rules for your workflow
- Performance optimization for high-volume usage

### Extension Points
- Custom domain agents (e.g., mobile, DevOps, security)
- Integration with external tools and services
- Custom GitHub Actions for specialized workflows
- Metrics collection and analytics

### Community
- Contribute agent templates for new domains
- Share best practices and patterns
- Report issues and feature requests
- Participate in agent template development

## Success Metrics

After following this quickstart, you should achieve:
- ✅ **Consistent Workflow**: All issues follow standardized 9-step process
- ✅ **Domain Specialization**: Technology-specific expertise applied appropriately
- ✅ **Automated Validation**: Specification compliance automatically enforced
- ✅ **Performance Tracking**: Metrics collection for continuous improvement
- ✅ **GitHub Integration**: Seamless workflow within existing GitHub processes
- ✅ **Quality Assurance**: Reduced errors through automated validation and testing
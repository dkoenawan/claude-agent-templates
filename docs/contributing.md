# Contributing Guidelines

Thank you for contributing to Claude Agent Templates! This guide will help you create high-quality agent templates that integrate seamlessly with our GitHub issue-driven development workflow and benefit the entire community.

## Getting Started

### Prerequisites
- Experience using Claude Code with agents
- Understanding of our GitHub issue-driven development workflow
- Familiarity with GitHub issues, comments, and labels for agent coordination
- Understanding of the specific domain/technology you're creating an agent for
- Familiarity with our [Agent Writing Guide](agent-guide.md)

### Development Environment
1. Fork the repository
2. Clone your fork locally
3. Install [Task](https://taskfile.dev/installation/) for cross-platform automation
4. Run `task install` to set up agents locally
5. Create a feature branch from `main`

## Contribution Process

### 1. Choose Your Contribution Type

#### New Agent Template
- Identify a gap in existing agents
- Focus on specific, well-defined use cases
- Ensure it's different enough from existing agents

#### Agent Improvement  
- Enhance existing agent functionality
- Fix bugs or unclear instructions
- Improve performance or accuracy

#### Documentation Update
- Clarify existing documentation
- Add examples or use cases
- Fix typos or formatting issues

### 2. Follow Our Development Workflow

#### Branch Strategy
```bash
# Create feature branch
git checkout main
git pull origin main
git checkout -b feature/agent-name

# Or for improvements
git checkout -b fix/improve-code-reviewer
```

#### Branch Naming Conventions
- `feature/agent-name` - New agents
- `fix/issue-description` - Bug fixes or improvements  
- `docs/topic-name` - Documentation updates
- `refactor/component-name` - Code restructuring

### 3. Agent Development Standards

#### Template Structure
Every agent must follow this structure:
```markdown
---
name: descriptive-agent-name
description: "Clear, concise description of when to use this agent"
tools: Tool1, Tool2, Tool3  # Only tools actually needed
---

Your agent's system prompt here.
```

#### Naming Conventions
- Use lowercase with hyphens: `code-reviewer`, `test-writer`
- Be specific: `python-security-auditor` not `security-agent`
- Avoid redundant words: `reviewer` not `code-review-agent`

#### Description Requirements
```markdown
# Good descriptions
description: "Specialized React component code reviewer focusing on hooks, performance, and TypeScript best practices"
description: "Python test writer using pytest, creating comprehensive unit and integration tests"

# Poor descriptions  
description: "Reviews code"
description: "Helps with testing stuff"
```

## Quality Standards

### Agent Requirements

#### ✅ Must Have
- **Clear purpose** - Specific, well-defined role in the 9-step workflow
- **GitHub integration** - Uses `gh` commands for issue interaction and workflow coordination
- **Verification protocols** - Mandatory codebase verification before making assumptions
- **Cross-agent validation** - Checks and validates previous agent work
- **Focused scope** - Not trying to do everything
- **Appropriate tools** - Only tools actually needed for workflow step
- **Tested functionality** - Works reliably in real projects with GitHub issues
- **Error recovery** - Self-correction capabilities when analysis errors occur
- **Professional tone** - Constructive and helpful

#### ❌ Must Avoid
- **Assumption-based analysis** - Must verify file existence and content before planning
- **Incomplete deliverables** - All agents must ensure full completion of their workflow step
- **Workflow isolation** - Agents must coordinate through GitHub issues and comments
- **Overly broad scope** - "Does everything" agents
- **Unclear instructions** - Vague or ambiguous prompts
- **Excessive tool permissions** - More tools than needed
- **Duplicate functionality** - Too similar to existing agents
- **Untested behavior** - Not validated in real scenarios with GitHub workflow
- **Condescending tone** - Negative or judgmental language

### Testing Requirements

#### 1. GitHub Workflow Integration Testing
Test your agent with:
- **GitHub issue scenarios** - Various issue types (bugs, features, questions)
- **Multi-agent coordination** - How it hands off to/receives from other agents
- **Issue commenting** - Proper use of GitHub comments for coordination
- **Label management** - Correct application of workflow labels
- **Error scenarios** - How it handles and recovers from analysis errors

#### 2. Multi-Project Testing
Test your agent with:
- **Different project sizes** - Small scripts to large applications
- **Various architectures** - Monoliths, microservices, libraries
- **Different team contexts** - Solo projects, team codebases
- **Edge cases** - Empty files, legacy code, incomplete features, missing dependencies
- **Codebase verification** - Projects with misleading directory structures

#### 3. Real-World Validation
```markdown
## Testing Checklist
- [ ] Tested GitHub issue workflow in at least 3 different repositories
- [ ] Verified codebase analysis accuracy before making plans
- [ ] Tested cross-agent handoffs and state validation
- [ ] Handles edge cases gracefully with proper error recovery
- [ ] Provides consistent, helpful output through GitHub issues
- [ ] Respects tool limitations and workflow boundaries
- [ ] Works with team workflows and issue templates
- [ ] Performance is acceptable with large codebases
- [ ] Self-correction works when provided user feedback
```

#### 3. Documentation Testing
- Instructions are clear and complete
- Examples work as described
- Links are functional
- Code samples are correct

## Submission Process

### 1. Pre-Submission Checklist
- [ ] Agent follows template structure
- [ ] Tested thoroughly in multiple contexts
- [ ] Documentation is complete and accurate
- [ ] No duplicate functionality with existing agents
- [ ] Follows naming conventions
- [ ] Tool permissions are minimal and appropriate

### 2. Pull Request Requirements

#### PR Title Format
```
Add [agent-name] agent for [specific use case]
Fix [agent-name] agent [specific issue]
Update [component] documentation
```

#### PR Description Template
```markdown
## Description
Brief description of what this PR adds/changes.

## Agent Details
- **Name**: agent-name
- **Purpose**: Specific use case this agent addresses
- **Tools Used**: List of tools and why they're needed
- **Target Users**: Who will benefit from this agent

## Testing
- [ ] Tested on multiple projects
- [ ] Edge cases handled
- [ ] Documentation tested
- [ ] Performance validated

## Examples
Provide examples of the agent in action (screenshots/outputs if helpful).

## Related Issues
Fixes #123
Relates to #456
```

### 3. Review Process

#### What We Look For
1. **Functionality** - Does the agent work as intended?
2. **Quality** - Is the code/documentation well-written?
3. **Uniqueness** - Does it add value beyond existing agents?
4. **Testing** - Is it thoroughly validated?
5. **Documentation** - Is it well-documented?

#### Feedback and Iteration
- Address reviewer feedback promptly
- Be open to suggestions and improvements
- Update documentation as needed
- Retest after making changes

## Community Standards

### Code of Conduct
- Be respectful and constructive in all interactions
- Focus on the work, not the person
- Welcome newcomers and help them learn
- Celebrate diversity of perspectives and approaches

### Communication Guidelines
- **Issues**: Use for bugs, feature requests, and questions
- **Discussions**: Use for broader conversations and ideas
- **Pull Requests**: Keep focused and include context
- **Reviews**: Be constructive and educational

## Advanced Contributions

### 1. GitHub Issue-Driven Workflow Positions
Help organize agents by their position in the 9-step workflow:

**Core Workflow Agents (Language-Agnostic):**
- **Step 2: Requirements Analysis** - Business requirement extraction and clarification
- **Step 4: Solution Architecture** - Technical planning and work unit breakdown
- **Step 9: Documentation** - Final documentation updates and repository cleanup

**Technology-Specific Implementation Agents:**
- **Step 6: Test Engineering** - Test strategy and coverage planning (Python, JavaScript, etc.)
- **Step 7: Software Engineering** - Implementation with branch management and PR creation

**Specialized Domain Agents:**
- **Security** - Vulnerability scanning, compliance, auditing (integrates at various steps)
- **DevOps** - Deployment, infrastructure, monitoring (integrates at implementation step)
- **Performance** - Optimization and profiling (integrates at architecture/implementation steps)

**Agent Coordination Requirements:**
- All agents must use GitHub issues for workflow state tracking
- Each agent must validate previous agent work before proceeding
- Proper workflow labels must be applied (`requirements-ready`, `plan-approved`, `tests-planned`, etc.)
- Cross-agent handoffs must be explicit and verified

### 2. Example Projects
Create comprehensive examples showing agents in action:
```
examples/
├── react-app/
│   ├── .claude/agents/
│   ├── src/
│   └── README.md
└── python-api/
    ├── .claude/agents/
    ├── app/
    └── README.md
```

### 3. Automation and Tools
We use a unified Taskfile-based system for repository automation:

**Available Commands:**
- `task install` - Install all agents to ~/.claude/agents
- `task list` - Show available and installed agents
- `task validate` - Verify agent installation
- `task clean` - Remove installed agents
- `task help` - Show detailed usage information

**GitHub Integration Requirements:**
All agents must support GitHub issue-driven workflows:
- Use `gh issue comment` for workflow communication
- Apply appropriate workflow labels (`requirements-ready`, `plan-approved`, etc.)
- Validate previous agent work before proceeding
- Handle error scenarios with self-correction capabilities

**Legacy Compatibility:**
Existing platform-specific scripts (`scripts/install-agents.sh`, `scripts/install-agents.bat`) are maintained for backward compatibility but new development should use the Taskfile system.

## Recognition

### Contributor Recognition
- Contributors are credited in agent files and documentation
- Outstanding contributions are highlighted in releases
- Regular contributors may be invited as maintainers

### Impact Tracking
We track:
- Agent usage and adoption
- Community feedback and iterations
- Real-world impact and success stories

## Getting Help

### Resources
- [Agent Writing Guide](agent-guide.md) - How to write effective agents
- [Customization Guide](customization.md) - Adapting agents for specific needs
- [Development Workflow](../CLAUDE.md) - Our trunk-based development process

### Support Channels
- **Issues** - Bug reports and feature requests
- **Discussions** - Questions and community support
- **Email** - For sensitive or private matters

### Mentorship
New contributors can request mentorship from experienced community members. We're here to help you succeed!

---

Thank you for helping make Claude Code more powerful and accessible for developers everywhere! 🚀
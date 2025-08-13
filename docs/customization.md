# Customization Guide

How to adapt and customize agent templates for your specific projects, teams, and workflows.

## Understanding Template Structure

Every agent template consists of:
- **YAML frontmatter** - Configuration (name, description, tools)
- **System prompt** - The agent's behavior definition
- **Context-specific instructions** - Project or domain knowledge

## Customization Levels

### 1. Project-Specific Customization

#### File Locations
- **Global agents**: `~/.claude/agents/` (available everywhere)
- **Project agents**: `.claude/agents/` (project-specific, overrides global)

#### Example: Customizing Code Reviewer for React Project
```markdown
---
name: react-code-reviewer
description: "Code reviewer specialized in React, TypeScript, and modern frontend practices"
tools: Read, Grep, Glob, Bash
---

You are a senior React developer with expertise in TypeScript, modern React patterns, and frontend performance.

## Project Context
This is a React application using:
- TypeScript with strict mode
- React 18 with hooks and concurrent features
- Tailwind CSS for styling
- Jest and React Testing Library for testing
- ESLint and Prettier for code quality

## Review Focus
- React hooks usage and dependency arrays
- TypeScript type safety and inference
- Component composition patterns
- Performance optimizations (memo, callback, useMemo)
- Accessibility (a11y) compliance
- Testing strategy and coverage

## Standards
- Use functional components with hooks
- Prefer composition over inheritance
- Implement proper error boundaries
- Follow React naming conventions
- Ensure responsive design patterns
```

### 2. Team/Organization Customization

#### Shared Conventions
Create organization-wide agents that encode your team's standards:

```markdown
---
name: acme-corp-reviewer
description: "Code reviewer following ACME Corp development standards"
---

You are a code reviewer enforcing ACME Corp's development standards.

## Coding Standards
- Follow ACME Corp Style Guide v2.1
- Use approved libraries from internal registry
- Implement standard error handling patterns
- Include required security headers
- Follow ACME logging format specifications

## Architecture Patterns
- Use ACME's microservice template
- Implement circuit breaker patterns
- Follow ACME's API versioning strategy
- Use approved authentication mechanisms

## Review Checklist
1. Security review against ACME threat model
2. Performance impact assessment
3. Monitoring and alerting implementation
4. Documentation completeness
5. Compliance with ACME governance policies
```

### 3. Technology Stack Customization

#### Language-Specific Variants
```markdown
---
name: python-django-reviewer
description: "Python Django application code reviewer"
tools: Read, Edit, Bash, Grep
---

You are a Python expert specializing in Django web applications.

## Technology Stack
- Python 3.11+
- Django 4.2+ with REST Framework
- PostgreSQL database
- Celery for background tasks
- Redis for caching
- Docker for containerization

## Review Areas
- Django best practices and security
- Database query optimization
- Async view implementation
- Middleware and signal usage
- Testing with pytest-django
- API design and serialization
```

## Common Customization Patterns

### 1. Adding Domain Knowledge
```markdown
## Domain Context
This is a financial trading application handling:
- Real-time market data processing
- High-frequency trading algorithms
- Regulatory compliance (MiFID II, Dodd-Frank)
- Risk management systems

## Special Considerations
- Latency requirements (<1ms for critical paths)
- Data accuracy and auditability
- Regulatory reporting requirements
- Security and access controls
```

### 2. Tool Restrictions by Environment
```markdown
# Development environment
tools: Read, Write, Edit, Bash, Grep, Glob

# Production environment  
tools: Read, Grep, Glob  # Read-only for safety
```

### 3. Workflow Integration
```markdown
## CI/CD Integration
After code review:
1. Run `npm run lint` and fix any issues
2. Execute `npm test` and ensure all tests pass
3. Update CHANGELOG.md if needed
4. Verify Docker build succeeds
5. Check bundle size impact with `npm run analyze`
```

## Creating Project Templates

### 1. Project Template Structure
```
project-templates/
├── react-frontend/
│   ├── code-reviewer.md
│   ├── test-writer.md
│   └── deployment-helper.md
├── python-backend/
│   ├── api-reviewer.md
│   ├── security-auditor.md
│   └── performance-optimizer.md
└── setup.sh
```

### 2. Template Deployment Script
```bash
#!/bin/bash
# deploy-template.sh

PROJECT_TYPE=$1
TARGET_DIR=${2:-.claude/agents}

if [ -z "$PROJECT_TYPE" ]; then
    echo "Usage: $0 <project-type> [target-directory]"
    exit 1
fi

mkdir -p "$TARGET_DIR"
cp project-templates/$PROJECT_TYPE/*.md "$TARGET_DIR/"
echo "Deployed $PROJECT_TYPE templates to $TARGET_DIR"
```

## Environment-Specific Configurations

### Development vs Production
```markdown
## Development Mode
- More verbose output and explanations
- Suggest learning resources
- Include experimental features
- Allow broader tool access

## Production Mode  
- Focus on critical issues only
- Conservative recommendations
- Restrict to read-only operations
- Emphasize stability and security
```

### Team Size Adaptations
```markdown
## Small Team (1-5 developers)
- More educational explanations
- Broader scope per agent
- Direct communication style
- Less formal processes

## Large Team (20+ developers)
- Enforce strict standards
- Focused, specialized agents
- Formal review processes
- Integration with team tools
```

## Advanced Customization Techniques

### 1. Dynamic Context Loading
```markdown
Before starting analysis, check for project-specific configuration:
1. Read `.claude/project-config.md` for custom instructions
2. Parse `package.json` or equivalent for dependencies
3. Check for existing documentation standards
4. Adapt behavior based on project characteristics
```

### 2. Multi-Agent Workflows
```markdown
## Agent Orchestration
1. **Analyzer Agent** - Initial code analysis
2. **Specialist Agent** - Domain-specific review  
3. **Validator Agent** - Final quality check
4. **Reporter Agent** - Generate summary report
```

### 3. Conditional Behavior
```markdown
## Conditional Logic
If project uses TypeScript:
- Enforce strict type checking
- Review type definitions
- Check for proper inference

If project is open source:
- Review README and documentation
- Check license compatibility
- Ensure public API clarity
```

## Testing Customizations

### 1. Validation Checklist
- [ ] Agent responds appropriately to project context
- [ ] Custom instructions are followed correctly
- [ ] Tool restrictions are respected
- [ ] Output matches expected format
- [ ] Performance is acceptable for project size

### 2. A/B Testing Approach
```markdown
## Testing Strategy
1. Deploy original template version
2. Create customized version with changes
3. Run both on same codebase
4. Compare results and gather feedback
5. Iterate based on effectiveness
```

## Sharing Customizations

### 1. Contributing Back to Templates
- Abstract project-specific details
- Document the use case clearly
- Provide examples and test cases
- Follow contribution guidelines

### 2. Internal Template Registry
```
organization-agents/
├── backend/
├── frontend/ 
├── devops/
├── security/
└── compliance/
```

## Best Practices

### Do:
- ✅ Start with base templates and iterate
- ✅ Document customization rationale
- ✅ Test thoroughly in real scenarios
- ✅ Version control your customizations
- ✅ Share successful patterns with team

### Don't:
- ❌ Over-customize without testing
- ❌ Hard-code project-specific details in shared templates
- ❌ Create too many similar variants
- ❌ Ignore tool permission implications
- ❌ Forget to update documentation
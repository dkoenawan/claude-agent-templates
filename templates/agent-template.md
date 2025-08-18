# Agent Template

Use this template as a starting point for creating new Claude Code agents.

## Template Structure

```markdown
---
name: your-agent-name
description: "Clear, specific description of when this agent should be used"
tools: Read, Edit, Bash  # Optional - list only tools you actually need
---

Your agent's system prompt defining their role, expertise, and behavior.

## Role Definition
You are a [specific role] with expertise in [domain/technology].

## Core Responsibilities  
Your primary responsibilities include:
- [Specific task 1]
- [Specific task 2]
- [Specific task 3]

## Approach and Methodology
When [working on task]:
1. [Step 1 - analysis/understanding]
2. [Step 2 - planning/strategy]  
3. [Step 3 - implementation/execution]
4. [Step 4 - validation/verification]

## Quality Standards
Focus on:
- **[Quality aspect 1]** - [Specific criteria]
- **[Quality aspect 2]** - [Specific criteria]
- **[Quality aspect 3]** - [Specific criteria]

## Output Format
[Specify how the agent should structure its responses]

[Include any domain-specific context, constraints, or considerations]
```

## Customization Guide

### 1. Choose Your Agent Name
- Use lowercase with hyphens: `code-reviewer`, `test-writer`
- Be specific and descriptive: `python-security-auditor`
- Avoid generic names: `helper`, `assistant`, `agent`

### 2. Write a Clear Description
Your description should answer:
- **What** does this agent do?
- **When** should someone use it?
- **Why** is it specialized for this task?

#### Good Examples:
```yaml
description: "Specialized React code reviewer focusing on hooks, performance, and TypeScript best practices"
description: "Python test writer creating comprehensive pytest-based unit and integration tests"
description: "Security auditor identifying vulnerabilities and security best practice violations"
```

#### Poor Examples:
```yaml
description: "Helps with code"
description: "General purpose assistant"
description: "Does code review and other stuff"
```

### 3. Select Appropriate Tools
Only include tools your agent actually needs:

#### Analysis-Only Agents
```yaml
tools: Read, Grep, Glob
```

#### Code Modification Agents  
```yaml
tools: Read, Edit, Grep, Glob
```

#### Full Development Agents
```yaml
tools: Read, Write, Edit, Bash, Grep, Glob
```

### 4. Craft Your System Prompt

#### Structure Your Prompt:
1. **Role Definition** - Who is the agent?
2. **Expertise Areas** - What are they expert in?
3. **Core Responsibilities** - What do they do?
4. **Methodology** - How do they approach tasks?
5. **Quality Standards** - What do they focus on?
6. **Output Format** - How do they communicate results?

#### Example System Prompt:
```markdown
You are a senior software engineer specializing in Python web development and API design.

Your expertise includes:
- Python best practices and performance optimization
- RESTful API design and implementation
- Database integration and query optimization
- Testing strategies with pytest
- Security considerations for web applications

## Core Responsibilities
- Review Python code for quality, performance, and security
- Suggest architectural improvements for scalability
- Identify potential bugs and edge cases
- Ensure adherence to Python and web development best practices
- Recommend appropriate testing strategies

## Review Methodology
1. **Architecture Analysis** - Understand overall structure and data flow
2. **Code Quality Review** - Check for readability, maintainability, and performance
3. **Security Assessment** - Identify potential vulnerabilities
4. **Testing Evaluation** - Assess test coverage and quality
5. **Documentation Review** - Ensure code is properly documented

## Quality Standards
Focus on:
- **Pythonic Code** - Follows Python idioms and conventions
- **Performance** - Efficient algorithms and data structures
- **Security** - Input validation, authentication, and authorization
- **Testability** - Code that's easy to test and well-covered
- **Maintainability** - Clear, readable, and well-organized code

## Output Format
Provide feedback as:
1. **Summary** - High-level assessment
2. **Critical Issues** - Security vulnerabilities and bugs (with file:line references)
3. **Improvements** - Suggestions for better code quality
4. **Best Practices** - Recommendations for following conventions
5. **Next Steps** - Actionable items for the developer

Always explain the reasoning behind your suggestions and provide specific examples.
```

## Testing Your Agent

### 1. Basic Functionality Test
Create a simple test scenario and verify:
- Agent understands its role
- Responds appropriately to requests
- Uses only specified tools
- Provides helpful output

### 2. Edge Cases
Test with:
- Empty or minimal codebases
- Legacy/poorly structured code
- Large, complex projects
- Different programming languages (if applicable)

### 3. Real-World Validation
Use your agent on actual projects you're familiar with to ensure:
- Responses are accurate and helpful
- Suggestions are actionable
- Behavior is consistent
- Performance is acceptable

## Example Agents

### Code Reviewer Agent
```markdown
---
name: code-reviewer
description: "General-purpose code reviewer for quality, security, and best practices across multiple languages"
tools: Read, Grep, Glob
---

You are a senior software engineer with broad experience in code review and software quality assurance.

[System prompt continues...]
```

### Test Writer Agent
```markdown
---
name: test-writer
description: "Specialized test writing agent creating comprehensive unit, integration, and end-to-end tests"
tools: Read, Write, Edit, Bash
---

You are a test automation specialist focused on creating comprehensive, maintainable test suites.

[System prompt continues...]
```

### Documentation Writer Agent
```markdown
---
name: documentation-writer
description: "Technical documentation specialist for APIs, README files, and code documentation"
tools: Read, Write, Edit, Grep
---

You are a technical writer specializing in developer documentation and API documentation.

[System prompt continues...]
```

## Best Practices Checklist

Before submitting your agent:

### Agent Design
- [ ] Name is specific and descriptive
- [ ] Description clearly explains when to use the agent
- [ ] Tools are minimal and appropriate
- [ ] System prompt is focused and clear
- [ ] Role and responsibilities are well-defined

### Quality Assurance
- [ ] Tested on multiple projects/scenarios  
- [ ] Handles edge cases gracefully
- [ ] Provides consistent, helpful output
- [ ] Performance is acceptable
- [ ] Documentation is complete

### Community Standards
- [ ] Follows template structure
- [ ] Uses professional, constructive tone
- [ ] Provides educational value
- [ ] Doesn't duplicate existing functionality
- [ ] Includes examples where helpful

## Need Help?

- Read our [Agent Writing Guide](../docs/agent-guide.md) for detailed guidance
- Check [Contributing Guidelines](../docs/contributing.md) for submission process
- Look at existing agents in `/agents` for inspiration
- Open an issue for questions or feedback

Happy agent building! 🤖
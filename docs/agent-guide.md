# Agent Writing Guide

A comprehensive guide to creating effective Claude Code agents that deliver consistent, high-quality results.

## Agent Structure

Every agent follows this Markdown format with YAML frontmatter:

```markdown
---
name: agent-name
description: Clear description of when this agent should be invoked
tools: Read, Edit, Bash  # Optional - defaults to all tools
---

Your agent's system prompt defining their role, capabilities, and approach.
```

## Writing Effective System Prompts

### 1. Define Clear Role and Expertise
```markdown
You are a senior software engineer specializing in code security and quality assurance.
```

### 2. Specify Core Responsibilities
```markdown
Your primary responsibilities are:
- Identify security vulnerabilities and potential exploits
- Review code for performance bottlenecks
- Ensure adherence to coding standards and best practices
- Suggest architectural improvements
```

### 3. Define Problem-Solving Approach
```markdown
When reviewing code:
1. First, understand the overall architecture and data flow
2. Analyze security implications of each component
3. Check for common vulnerability patterns (SQL injection, XSS, etc.)
4. Evaluate performance characteristics
5. Provide specific, actionable recommendations
```

### 4. Set Quality Standards
```markdown
Focus on:
- **Critical issues first** - Security vulnerabilities and bugs
- **Specific feedback** - Line numbers and exact changes needed
- **Educational value** - Explain the "why" behind recommendations
- **Actionable suggestions** - Provide concrete solutions
```

## Tool Selection Strategy

### Minimal Tool Sets (Recommended)
Be strategic about tool access to create focused agents:

```yaml
# Code reviewer - focused on analysis
tools: Read, Grep, Glob

# Bug fixer - needs modification capability  
tools: Read, Edit, Bash, Grep

# Test writer - needs creation and execution
tools: Read, Write, Edit, Bash
```

### Why Limit Tools?
- **Focus** - Prevents agent from scope creep
- **Performance** - Faster decision making
- **Predictability** - Clearer behavior patterns
- **Security** - Reduces potential for unintended actions

## Agent Personality and Behavior

### Professional Tone
```markdown
You approach code review with a constructive, educational mindset. 
Provide clear explanations and avoid condescending language.
```

### Proactive Behavior
```markdown
Proactively suggest improvements beyond the immediate request.
If you notice related issues, address them in your analysis.
```

### Context Awareness
```markdown
Always consider the broader codebase context. Check related files
and dependencies before making recommendations.
```

## Common Patterns

### 1. Analysis-First Agents
```markdown
Before making any changes:
1. Read and understand the existing codebase
2. Identify patterns and conventions
3. Analyze the specific request in context
4. Plan your approach
```

### 2. Validation Agents
```markdown
After completing work:
1. Review changes for correctness
2. Run relevant tests
3. Check for unintended side effects
4. Verify adherence to project standards
```

### 3. Educational Agents
```markdown
When providing solutions:
1. Explain the reasoning behind your approach
2. Reference best practices and standards
3. Provide alternative solutions when appropriate
4. Link to relevant documentation
```

## Testing Your Agents

### 1. Multi-Project Testing
Test agents across different:
- Programming languages
- Project sizes
- Architectural patterns
- Team structures

### 2. Edge Case Scenarios
- Empty or minimal codebases
- Legacy code with technical debt
- Partially completed features
- Complex dependency chains

### 3. Real-World Usage
- Use agents in your daily workflow
- Gather feedback from team members
- Monitor for unexpected behaviors
- Iterate based on actual results

## Best Practices

### Do:
- ✅ Write specific, actionable system prompts
- ✅ Test agents thoroughly before sharing
- ✅ Use descriptive names and clear descriptions
- ✅ Limit tool access appropriately
- ✅ Include context about when to use the agent
- ✅ Provide examples of expected behavior

### Don't:
- ❌ Create overly broad, general-purpose agents
- ❌ Grant unnecessary tool permissions
- ❌ Use vague or ambiguous language
- ❌ Forget to test edge cases
- ❌ Create agents that duplicate existing functionality
- ❌ Make agents too verbose or too terse

## Example: Well-Designed Agent

```markdown
---
name: security-auditor
description: "Specialized security code review agent. Use for security-focused code analysis, vulnerability detection, and security best practice enforcement."
tools: Read, Grep, Glob
---

You are a cybersecurity expert specializing in application security and code review.

Your mission is to identify security vulnerabilities, potential attack vectors, and security best practice violations in codebases.

## Analysis Approach
1. **Threat Modeling** - Understand data flow and trust boundaries
2. **Vulnerability Scanning** - Check for OWASP Top 10 and language-specific issues  
3. **Code Pattern Analysis** - Identify dangerous patterns and anti-patterns
4. **Dependency Review** - Assess third-party library security

## Focus Areas
- Input validation and sanitization
- Authentication and authorization mechanisms
- Data encryption and storage security
- Error handling and information disclosure
- Injection vulnerabilities (SQL, XSS, etc.)
- Insecure dependencies and configurations

## Output Format
Provide findings as:
1. **Severity Level** (Critical/High/Medium/Low)
2. **Location** (file:line)
3. **Vulnerability Description**
4. **Exploitation Scenario**
5. **Remediation Steps**
6. **Prevention Strategies**

Always explain the security implications and provide educational context.
```

This agent is focused, specific, and provides clear guidance on its purpose and behavior.
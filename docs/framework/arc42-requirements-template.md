# Arc42 Requirements Analysis Template

This template guides requirements analysts in gathering structured information following arc42 principles.

## Section 1: Introduction & Goals

### Requirements Overview
- **Business Objective**: [What business problem does this solve?]
- **Target Users**: [Who will use this system/feature?]
- **Success Criteria**: [How will success be measured?]

### Quality Goals
**Rank the top 3 quality requirements (1=highest priority):**
- [ ] Performance/Efficiency
- [ ] Security
- [ ] Usability/User Experience
- [ ] Reliability/Availability
- [ ] Maintainability
- [ ] Scalability
- [ ] Testability

### Key Stakeholders
- **Primary Users**: [Direct system users]
- **Business Stakeholders**: [Decision makers, sponsors]
- **Technical Stakeholders**: [Developers, operators]
- **Affected Systems**: [External systems impacted]

## Section 2: Constraints

### Organizational Constraints
- **Team Size/Structure**: [Available resources]
- **Timeline**: [Delivery expectations]
- **Budget**: [Resource limitations]
- **Process Requirements**: [Development methodology, approval processes]

### Technical Constraints
- **Existing Systems**: [Systems that must be integrated with]
- **Technology Stack**: [Required or prohibited technologies]
- **Performance Requirements**: [Specific performance targets]
- **Security Requirements**: [Compliance or security standards]

### External Constraints
- **Regulatory**: [Legal or compliance requirements]
- **Standards**: [Industry standards to follow]
- **Third-party Dependencies**: [External services or libraries]

## Section 3: Context & Scope

### Business Context
- **Primary Use Cases**: [Main user scenarios]
- **Business Processes**: [Workflows this system supports]
- **External Entities**: [Other systems or users that interact]

### Technical Context
- **Integration Points**: [APIs, databases, services to connect with]
- **Data Sources**: [Where data comes from]
- **Communication Protocols**: [How systems communicate]

### System Boundaries
- **In Scope**: [What this system will do]
- **Out of Scope**: [What this system will NOT do]
- **Future Considerations**: [Potential future expansions]

## Critical Thinking Checklist

### Problem Validation
- [ ] Is this addressing the actual root problem?
- [ ] Have we considered alternative solutions?
- [ ] Are we solving a symptom vs the underlying issue?
- [ ] Does this align with the project's core purpose?

### Requirement Quality
- [ ] Are requirements specific and measurable?
- [ ] Have we identified all affected stakeholders?
- [ ] Are there conflicting requirements that need resolution?
- [ ] Have we considered edge cases and error scenarios?

### Business Value
- [ ] What's the business impact if we don't implement this?
- [ ] How does this contribute to organizational goals?
- [ ] Are there simpler ways to achieve the same outcome?
- [ ] What are the costs vs benefits?

## Probing Questions Framework

### Understanding the "Why"
- "What problem are you trying to solve?"
- "How are users currently handling this situation?"
- "What happens if we don't implement this feature?"
- "How will you measure success?"

### Exploring Alternatives
- "Have you considered other approaches?"
- "What would be the simplest solution?"
- "Are there existing tools that could address this?"
- "Could this be solved with configuration vs code?"

### Validating Assumptions
- "How do you know users want this feature?"
- "What evidence supports this approach?"
- "Have you validated this with actual users?"
- "What are the risks of this approach?"

### Challenging Scope
- "Is this the minimum viable solution?"
- "What parts could be delivered later?"
- "Which requirements are must-have vs nice-to-have?"
- "Could we start with a simpler version?"

## Output Format

When requirements analysis is complete, provide:

```markdown
## Arc42 Requirements Summary

### 1. Introduction & Goals
[Structured overview of business objectives and quality goals]

### 2. Constraints
[Documented limitations and requirements]

### 3. Context & Scope
[Clear system boundaries and integration points]

### Critical Analysis
[Summary of challenges raised and alternatives considered]

### Next Steps
[Specific items for solution architect to address]

---
**Status**: requirements-ready
**Agent**: Requirements Analyst | **Arc42 Sections**: 1, 2, 3 | **Timestamp**: [ISO timestamp]
```
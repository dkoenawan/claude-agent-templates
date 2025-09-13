# Arc42 Documentation Template

This template guides documentation agents in creating comprehensive final documentation following arc42 principles.

## Complete Arc42 Documentation Structure

### Section 1: Introduction & Goals
[Compiled from Requirements Analyst]
- Requirements overview
- Quality goals with priorities
- Key stakeholders

### Section 2: Constraints
[Compiled from Requirements Analyst]
- Organizational constraints
- Technical constraints
- External constraints

### Section 3: Context & Scope
[Compiled from Requirements Analyst]
- Business context
- Technical context
- System boundaries

### Section 4: Solution Strategy
[Compiled from Solution Architect]
- Architectural approach
- Quality achievement strategy
- Top-level decomposition

### Section 5: Building Block View
[Compiled from Solution Architect + Implementation Results]
- System overview
- Component architecture
- Implementation details

### Section 6: Runtime View
[Implementation-specific scenarios]
- Key use cases
- Error handling flows
- Integration scenarios

### Section 7: Deployment View
[Based on actual implementation]
- Infrastructure setup
- Deployment pipeline
- Environment configuration

### Section 8: Crosscutting Concepts
[Implementation patterns used]
- Domain patterns
- Technical patterns
- Security concepts

### Section 9: Architectural Decisions
[Compiled and updated from implementation]
- Original architectural decisions
- Implementation adjustments
- Lessons learned

### Section 10: Quality Requirements
[Validation of quality goals]
- Quality metrics achieved
- Testing results
- Performance benchmarks

### Section 11: Risks & Technical Debt
[Post-implementation assessment]
- Identified risks during implementation
- Technical debt created
- Mitigation recommendations

### Section 12: Glossary
[Complete terminology]
- Domain terms
- Technical terms
- Project-specific acronyms

## Documentation Quality Standards

### Completeness Checklist
- [ ] All sections contain relevant information
- [ ] Implementation matches architectural plan
- [ ] Quality goals are validated with evidence
- [ ] All decisions are documented with rationale
- [ ] Risks and technical debt are identified

### Clarity Standards
- [ ] Language appropriate for target audience
- [ ] Technical terms defined in glossary
- [ ] Diagrams support text explanations
- [ ] Examples provided for complex concepts

### Maintenance Guidelines
- [ ] Version control information included
- [ ] Update responsibilities assigned
- [ ] Review schedule established
- [ ] Change tracking process defined

## Integration with Codebase

### Code Documentation Links
- Link architecture documentation to code structure
- Reference implementation files for each component
- Include code examples for key patterns
- Document API interfaces and contracts

### README Updates
- Update main README with architecture overview
- Link to detailed arc42 documentation
- Provide quick-start guides
- Include contribution guidelines

### Developer Onboarding
- Architecture overview for new developers
- Code organization explanation
- Development workflow documentation
- Testing and deployment procedures

## Output Format Templates

### Main Documentation File
```markdown
# [Project Name] - Architecture Documentation

## Arc42 Architecture Documentation

This document follows the [arc42 template](https://arc42.org/) for architecture documentation.

### Table of Contents
1. [Introduction & Goals](#1-introduction--goals)
2. [Constraints](#2-constraints)
3. [Context & Scope](#3-context--scope)
4. [Solution Strategy](#4-solution-strategy)
5. [Building Block View](#5-building-block-view)
6. [Runtime View](#6-runtime-view)
7. [Deployment View](#7-deployment-view)
8. [Crosscutting Concepts](#8-crosscutting-concepts)
9. [Architectural Decisions](#9-architectural-decisions)
10. [Quality Requirements](#10-quality-requirements)
11. [Risks & Technical Debt](#11-risks--technical-debt)
12. [Glossary](#12-glossary)

[Detailed sections follow...]
```

### README Integration
```markdown
## Architecture

This project follows hexagonal architecture principles. For detailed architecture documentation, see:

- **[Architecture Overview](docs/architecture.md)** - Complete arc42 documentation
- **[Development Guide](docs/contributing.md)** - How to contribute to this project
- **[API Documentation](docs/api.md)** - Interface specifications

### Quick Start
[Implementation-specific quick start guide]

### Project Structure
[Code organization following architectural layers]
```

### Final Status Update
```markdown
## Arc42 Documentation Complete

### Documentation Created
- [x] Complete arc42 architecture documentation
- [x] Updated README with architecture overview
- [x] Developer onboarding materials
- [x] API documentation updates

### Quality Validation
- [x] All arc42 sections completed
- [x] Implementation matches architecture
- [x] Quality goals validated
- [x] Technical debt documented

### Maintenance Plan
- **Review Schedule**: [Quarterly/As needed]
- **Update Triggers**: [Major changes, new features]
- **Ownership**: [Team/Individual responsible]

---
**Status**: documentation-complete
**Agent**: Documentation | **Arc42 Sections**: Complete | **Timestamp**: [ISO timestamp]
```
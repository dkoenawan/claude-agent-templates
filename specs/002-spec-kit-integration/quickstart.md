# Quickstart Guide: Spec-Kit Integration

**Feature**: Spec-Kit Integration and Project Pivot
**Audience**: Non-technical users and developers new to the project
**Time**: 15 minutes

---

## What You'll Learn

By the end of this guide, you'll be able to:
- Create a new feature specification using `/speckit.specify`
- Generate an implementation plan using `/speckit.plan`
- Understand the spec-driven development workflow
- Use Claude skills for specialized tasks

---

## Prerequisites

Before starting, ensure you have:

1. **Git** installed (version 2.x or higher)
   ```bash
   git --version
   ```

2. **Claude Code** installed and configured
   - [Installation guide](https://docs.claude.com/en/docs/claude-code)

3. **Repository cloned**
   ```bash
   git clone https://github.com/your-org/claude-agent-templates
   cd claude-agent-templates
   ```

4. **On the main branch** with latest changes
   ```bash
   git checkout main
   git pull origin main
   ```

---

## Step 1: Create Your First Feature Specification (5 min)

### Using Claude Code

1. **Open Claude Code** in your repository directory

2. **Describe your feature** using the `/speckit.specify` command:
   ```
   /speckit.specify Add a dark mode toggle to the user settings page
   ```

3. **Review the generated spec**:
   - Claude will create a new branch (e.g., `003-dark-mode-toggle`)
   - Generate `specs/003-dark-mode-toggle/spec.md` with:
     - User scenarios and acceptance criteria
     - Functional requirements
     - Success criteria
     - Scope boundaries

4. **Answer clarification questions** (if any):
   - Claude may ask 1-3 targeted questions
   - Select from the provided options or provide custom answers
   - Example: "Should dark mode persist across sessions?" → Choose "Yes" or "No"

5. **Verify the spec**:
   ```bash
   cat specs/003-dark-mode-toggle/spec.md
   ```

**What just happened?**
- Created a feature branch following naming conventions
- Generated a complete specification without writing any code
- Validated the spec meets quality standards
- Ready for technical planning

---

## Step 2: Generate Implementation Plan (5 min)

### Using Claude Code

1. **Run the planning command**:
   ```
   /speckit.plan
   ```

2. **Wait for artifacts to generate**:
   Claude will create:
   - `plan.md` - Implementation strategy and technical decisions
   - `research.md` - Technology choices and rationale
   - `data-model.md` - Data structures and entities
   - `contracts/` - API interfaces (if applicable)
   - `quickstart.md` - Getting started guide (like this one!)

3. **Review the technical context**:
   ```bash
   # See the implementation approach
   cat specs/003-dark-mode-toggle/plan.md

   # Check technology decisions
   cat specs/003-dark-mode-toggle/research.md
   ```

4. **Examine the data model**:
   ```bash
   cat specs/003-dark-mode-toggle/data-model.md
   ```

**What just happened?**
- Researched best practices for your feature
- Made technology decisions (frameworks, libraries, patterns)
- Designed data structures
- Created API contracts
- All without writing implementation code

---

## Step 3: Understand the Workflow (2 min)

The spec-driven development workflow has these phases:

```
1. Specify     → Create spec.md describing WHAT and WHY
   ↓
2. Plan        → Generate plan.md describing technical HOW
   ↓
3. Clarify     → Resolve any ambiguities (optional)
   ↓
4. Tasks       → Break down into implementation tasks
   ↓
5. Implement   → Write code following the plan
   ↓
6. Review      → Validate against spec and plan
   ↓
7. Document    → Update docs and finalize
```

**Current Status**: You're at step 2 (Planning complete)

**Next Steps**:
- Run `/speckit.tasks` to break down into concrete tasks
- Run `/speckit.implement` to execute the implementation
- Or continue reading to learn about skills

---

## Step 4: Use Claude Skills (3 min)

Skills are reusable workflows for common tasks. Let's try a few:

### Skill 1: Domain Classifier

Classify what technology domain your feature belongs to:

```
Use the domain-classifier skill to analyze this feature description:
"Build a REST API for user management with authentication"
```

**Expected Output**:
```json
{
  "domain": "python",  // or "nodejs", "dotnet", "java"
  "confidence": 0.85,
  "rationale": "REST API keywords suggest backend development..."
}
```

### Skill 2: Agent Migrator

Migrate an existing agent to spec-kit format:

```
Use the agent-migrator skill to migrate: agents/python/software-engineer-python.md
```

**Expected Output**:
- Adds spec-kit frontmatter fields
- Preserves existing content
- Validates the migrated format

### Skill 3: Contract Generator

Generate API contracts from your spec:

```
Use the contract-generator skill on: specs/003-dark-mode-toggle/spec.md
```

**Expected Output**:
- OpenAPI YAML file in `contracts/`
- Endpoints derived from functional requirements

**What are skills good for?**
- Automating repetitive tasks
- Applying best practices consistently
- Generating boilerplate code
- Classifying and analyzing features

---

## Common Tasks

### View All Available Commands

```bash
# List all spec-kit commands
ls .claude/commands/speckit.*.md

# See command documentation
cat .claude/commands/speckit.specify.md
```

### Check Feature Status

```bash
# See your current branch
git branch --show-current

# List all features
ls specs/

# View workflow state
cat .workflow-state.json
```

### Validate Your Work

```bash
# Validate agent specifications
python3 scripts/validate-claude-agent.py

# Run tests
python3 -m unittest discover tests/ -v
```

### Get Help

```bash
# View project documentation
cat README.md
cat CLAUDE.md

# Check specific feature docs
cat specs/003-dark-mode-toggle/quickstart.md
```

---

## Troubleshooting

### "Command not found: /speckit.specify"

**Solution**: Ensure you're running Claude Code in the repository directory where `.claude/commands/` exists.

### "No specification found on current branch"

**Solution**: You need to run `/speckit.specify` first to create a spec before running `/speckit.plan`.

### "Feature already exists"

**Solution**: Check `specs/` directory for existing features. Use a different feature name or switch to the existing branch:
```bash
git checkout 003-dark-mode-toggle
```

### "Invalid feature description"

**Solution**: Provide more detail. Good descriptions include:
- What the feature does
- Who it's for
- Why it's needed

Example: ❌ "Add feature"
Example: ✅ "Add OAuth2 authentication to allow users to log in with Google and GitHub"

---

## Next Steps

Now that you've completed the quickstart:

### For Non-Technical Users

1. **Create specifications** for features you need
2. **Collaborate with developers** by reviewing generated plans
3. **Validate implementations** against your original spec
4. **Iterate** using `/speckit.clarify` to refine requirements

### For Developers

1. **Generate task breakdowns** using `/speckit.tasks`
2. **Implement features** following the plan
3. **Write tests** based on acceptance criteria in spec
4. **Create PRs** linked to your feature branch
5. **Update documentation** as features evolve

### Advanced Topics

- **Creating custom skills**: See `.claude/skills/README.md`
- **Defining project constitution**: Run `/speckit.constitution`
- **Integrating with GitHub Actions**: See `.github/workflows/`
- **Multi-domain development**: See `agents/` directory

---

## Learning Resources

### Documentation

- [Spec-Kit User Guide](../../docs/framework/spec-kit-guide.md)
- [Claude Skills Documentation](../../docs/framework/skills-guide.md)
- [Agent Development Guide](../../docs/framework/agent-development.md)
- [GitHub Spec-Kit](https://github.com/github/spec-kit)

### Examples

- [Example Feature Spec](../../examples/feature-spec-example.md)
- [Example Implementation Plan](../../examples/plan-example.md)
- [Example Claude Skill](../../examples/skill-example.md)

### Support

- **Issues**: [GitHub Issues](https://github.com/your-org/claude-agent-templates/issues)
- **Discussions**: [GitHub Discussions](https://github.com/your-org/claude-agent-templates/discussions)
- **Documentation**: [docs/ directory](../../docs/)

---

## Summary

You've learned how to:

✅ Create feature specifications with `/speckit.specify`
✅ Generate implementation plans with `/speckit.plan`
✅ Use Claude skills for specialized tasks
✅ Navigate the spec-driven development workflow
✅ Troubleshoot common issues

**Remember**: The spec-kit approach keeps you focused on WHAT and WHY (specification) before diving into HOW (implementation). This leads to better-designed, more maintainable software.

---

**Questions?** Check the [documentation](../../docs/) or [open an issue](https://github.com/your-org/claude-agent-templates/issues).

**Ready to build?** Run `/speckit.tasks` to start implementing your feature!

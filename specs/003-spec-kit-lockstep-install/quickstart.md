# Quickstart: Spec-Kit Lockstep Installation

**Time to Complete**: 5-10 minutes
**Goal**: Install claude-agent-templates with pinned spec-kit version
**Audience**: Developers using Claude Code

---

## Choose Your Installation Method

Select the installation method that fits your workflow:

- **[Option 1: Global Installation](#option-1-global-installation)** - Agents available across all projects
- **[Option 2: Project-Local Installation](#option-2-project-local-installation)** - Isolated to single project
- **[Option 3: Coexist with Existing Spec-Kit](#option-3-coexist-with-existing-spec-kit)** - Already using spec-kit

---

## Option 1: Global Installation

**Best for**: Using claude-agent-templates across multiple projects

### Step 1: Install Agents Globally

```bash
# Clone the repository
git clone https://github.com/dkoenawan/claude-agent-templates.git
cd claude-agent-templates

# Install agents to global Claude directory
mkdir -p ~/.claude/agents
cp -r agents/core/*.md ~/.claude/agents/
cp -r agents/python/*.md ~/.claude/agents/
cp -r agents/dotnet/*.md ~/.claude/agents/
cp -r agents/nodejs/*.md ~/.claude/agents/
cp -r agents/java/*.md ~/.claude/agents/
```

### Step 2: Add Spec-Kit to Each Project

```bash
# Navigate to your project
cd /path/to/your/project

# Copy spec-kit with pinned version
cp -r /path/to/claude-agent-templates/.specify .specify

# Create .claude directory structure
mkdir -p .claude/{commands,agents}

# COPY (not symlink) slash commands with namespace prefix
# This preserves your existing custom commands
cp .specify/templates/commands/specify.md .claude/commands/speckit.specify.md
cp .specify/templates/commands/plan.md .claude/commands/speckit.plan.md
cp .specify/templates/commands/tasks.md .claude/commands/speckit.tasks.md

# COPY agents with "cat-" prefix (claude-agent-templates)
# This prevents conflicts with your custom agents
cd /path/to/claude-agent-templates
for agent in agents/**/*.md; do
    filename=$(basename "$agent")
    cp "$agent" "/path/to/your/project/.claude/agents/cat-$filename"
done
```

### Step 3: Verify Installation

```bash
# Check spec-kit version
cat .specify/version-manifest.json | jq -r '.dependencies["spec-kit"].version'

# Expected output: 0.0.72 (or current pinned version)
```

### Step 4: Use in Claude Code

```bash
# Open Claude Code in your project
# Agents are automatically available
# Run a slash command
/speckit.specify "Add user authentication feature"
```

**Result**: Agents available globally, spec-kit version managed per-project

---

## Option 2: Project-Local Installation

**Best for**: Single project or testing claude-agent-templates

### Step 1: Clone into Project

```bash
# Navigate to your project
cd /path/to/your/project

# Clone as subdirectory
git clone https://github.com/dkoenawan/claude-agent-templates.git .claude-agent-templates
```

### Step 2: Integrate with Claude Code

```bash
# Create .claude directory structure
mkdir -p .claude/{commands,agents}

# COPY (not symlink) agents with "cat-" prefix to avoid conflicts
for agent in .claude-agent-templates/agents/**/*.md; do
    filename=$(basename "$agent")
    cp "$agent" ".claude/agents/cat-$filename"
done

# COPY (not symlink) slash commands with namespace prefix
cp .claude-agent-templates/.specify/templates/commands/specify.md .claude/commands/speckit.specify.md
cp .claude-agent-templates/.specify/templates/commands/plan.md .claude/commands/speckit.plan.md
cp .claude-agent-templates/.specify/templates/commands/tasks.md .claude/commands/speckit.tasks.md
```

### Step 3: Verify Installation

```bash
# Check directory structure
ls .claude/agents/
# Expected output:
#   cat-requirements-analyst.md
#   cat-solution-architect-python.md
#   cat-software-engineer-python.md
#   ... (and any of your own custom agents)

ls .claude/commands/
# Expected output:
#   speckit.specify.md
#   speckit.plan.md
#   speckit.tasks.md
#   ... (and any of your own custom commands)

# Check spec-kit version
cat .claude-agent-templates/.specify/version-manifest.json | jq -r '.dependencies["spec-kit"].version'
```

### Step 4: Use in Claude Code

Open Claude Code in your project - agents and commands are automatically discovered.

**Result**: Everything isolated to this project

---

## Option 3: Coexist with Existing Spec-Kit

**Best for**: Already using spec-kit, want to add claude-agent-templates

### Step 1: Check Existing Installation

```bash
# Check if you already have spec-kit
ls -la .specify/

# Expected output:
# .specify/
# ├── memory/
# ├── scripts/
# └── templates/

# Check your spec-kit version
cat .specify/memory/constitution.md  # Or other identifying file
```

### Step 2: Install to Separate Directory

```bash
# Clone claude-agent-templates to separate directory
git clone https://github.com/dkoenawan/claude-agent-templates.git .claude-agent-templates

# Your project structure:
# .specify/                    # Your existing spec-kit
# .claude-agent-templates/     # Our installation
#   ├── .specify/              # Our vendored spec-kit (isolated)
#   └── agents/
```

### Step 3: Choose Spec-Kit Version

**Important Decision**: Which spec-kit version to use?

#### Option A: Use Our Vendored Spec-Kit (Recommended)

```bash
# Update .claude/commands to point to our spec-kit
mkdir -p .claude/commands
ln -s $(pwd)/.claude-agent-templates/.specify/templates/commands/*.md .claude/commands/

# Keep your .specify/ for other tools
# Use .claude-agent-templates/.specify/ for Claude Code
```

**Pros**: Guaranteed compatibility with our agents
**Cons**: Two spec-kit installations

#### Option B: Use Your Existing Spec-Kit (Advanced)

```bash
# Check version compatibility
YOUR_VERSION=$(cat .specify/version-manifest.json | jq -r '.dependencies["spec-kit"].version' 2>/dev/null || echo "unknown")
OUR_MIN_VERSION=$(cat .claude-agent-templates/.specify/version-manifest.json | jq -r '.dependencies["spec-kit"].compatibility.min_version')
OUR_MAX_VERSION=$(cat .claude-agent-templates/.specify/version-manifest.json | jq -r '.dependencies["spec-kit"].compatibility.max_version')

echo "Your version: $YOUR_VERSION"
echo "Compatible range: $OUR_MIN_VERSION - $OUR_MAX_VERSION"

# If compatible, use your version
# Update .claude/commands to point to your .specify/
ln -s $(pwd)/.specify/templates/commands/*.md .claude/commands/
```

**Pros**: Single spec-kit installation
**Cons**: Risk of version incompatibility

### Step 4: Integrate Agents

```bash
# Add agents to .claude/agents with "cat-" prefix to avoid conflicts
mkdir -p .claude/agents
for agent in .claude-agent-templates/agents/**/*.md; do
    filename=$(basename "$agent")
    cp "$agent" ".claude/agents/cat-$filename"
done
```

### Step 5: Verify Both Coexist

```bash
# Check directory structure
tree -L 2 .
# Expected output:
# .
# ├── .specify/                    # Your existing spec-kit
# ├── .claude-agent-templates/     # Our installation
# │   ├── .specify/                # Our vendored spec-kit
# │   └── agents/
# └── .claude/                      # Claude Code integration
#     ├── agents/ -> ...
#     └── commands/ -> ...

# Verify commands work
/speckit.specify "Test feature"
```

**Result**: Both spec-kit versions coexist, Claude Code uses the one you choose

---

## Verification Checklist

After installation, verify everything works:

- [ ] Claude Code recognizes agents (`/agents` command shows them)
- [ ] Slash commands available (`/speckit.specify`, `/speckit.plan`, `/speckit.tasks`)
- [ ] Spec-kit version matches expected (check `version-manifest.json`)
- [ ] Can create a test spec: `/speckit.specify "test feature"`
- [ ] Spec file created in `specs/001-test-feature/spec.md`

---

## Troubleshooting

### Commands Not Found

**Symptom**: `/speckit.specify` shows "Command not found"

**Solution**:
```bash
# Ensure .claude/commands/ exists and has command files
ls -la .claude/commands/

# If empty, copy commands:
cp .specify/templates/commands/specify.md .claude/commands/speckit.specify.md
cp .specify/templates/commands/plan.md .claude/commands/speckit.plan.md
cp .specify/templates/commands/tasks.md .claude/commands/speckit.tasks.md
# OR (if using .claude-agent-templates)
cp .claude-agent-templates/.specify/templates/commands/specify.md .claude/commands/speckit.specify.md
cp .claude-agent-templates/.specify/templates/commands/plan.md .claude/commands/speckit.plan.md
cp .claude-agent-templates/.specify/templates/commands/tasks.md .claude/commands/speckit.tasks.md
```

### Agents Not Appearing

**Symptom**: Agents don't show in Claude Code

**Solution**:
```bash
# Check ~/.claude/agents/ (global) or .claude/agents/ (local)
ls -la ~/.claude/agents/
ls -la .claude/agents/

# Reinstall agents with "cat-" prefix
cp -r agents/**/*.md ~/.claude/agents/  # Global (no prefix needed)
# OR for local installation with prefix
for agent in .claude-agent-templates/agents/**/*.md; do
    filename=$(basename "$agent")
    cp "$agent" ".claude/agents/cat-$filename"
done
```

### Need to Update After claude-agent-templates Upgrade

**Symptom**: Installed new version of claude-agent-templates but commands/agents are outdated

**Solution**: Re-copy files (since we use copies, not symlinks)
```bash
# Navigate to project
cd /path/to/your/project

# Update slash commands
cp .claude-agent-templates/.specify/templates/commands/specify.md .claude/commands/speckit.specify.md
cp .claude-agent-templates/.specify/templates/commands/plan.md .claude/commands/speckit.plan.md
cp .claude-agent-templates/.specify/templates/commands/tasks.md .claude/commands/speckit.tasks.md

# Update agents (remove old ones first)
rm .claude/agents/cat-*.md
for agent in .claude-agent-templates/agents/**/*.md; do
    filename=$(basename "$agent")
    cp "$agent" ".claude/agents/cat-$filename"
done

echo "✅ Updated to latest claude-agent-templates version"
```

### Version Conflict

**Symptom**: Error about spec-kit version mismatch

**Solution**:
```bash
# Check installed version
cat .claude-agent-templates/.version-lock.json | jq -r '.components["spec-kit"].version'

# Check required version
cat .claude-agent-templates/.specify/version-manifest.json | jq -r '.dependencies["spec-kit"].version'

# If mismatch, reinstall:
rm -rf .claude-agent-templates
git clone https://github.com/dkoenawan/claude-agent-templates.git .claude-agent-templates
```

---

## Next Steps

### 1. Create Your First Spec

```bash
/speckit.specify "Add user authentication with email and password"
```

This creates `specs/001-user-auth/spec.md` with structured requirements.

### 2. Generate Implementation Plan

```bash
/speckit.plan
```

This creates `specs/001-user-auth/plan.md` with technical design.

### 3. Break Down into Tasks

```bash
/speckit.tasks
```

This creates `specs/001-user-auth/tasks.md` with actionable tasks.

### 4. Explore Agents

In Claude Code, type:
```
@requirements-analyst help me understand feature requirements
@solution-architect-python design the authentication system
@software-engineer-python implement the auth endpoints
```

---

## Upgrading

When a new claude-agent-templates version is released:

### Global Installation

```bash
cd /path/to/claude-agent-templates
git pull origin main

# Re-copy agents
cp -r agents/**/*.md ~/.claude/agents/

# Update spec-kit in each project
cd /path/to/your/project
cp -r /path/to/claude-agent-templates/.specify .specify
```

### Project-Local Installation

```bash
cd /path/to/your/project/.claude-agent-templates
git pull origin main

# Symlinks automatically point to new files
# If using copies, re-copy:
cp .specify/templates/commands/*.md ../.claude/commands/
```

---

## Getting Help

- **Documentation**: See `CLAUDE.md` for detailed agent usage
- **Issues**: https://github.com/dkoenawan/claude-agent-templates/issues
- **Discussions**: https://github.com/dkoenawan/claude-agent-templates/discussions

---

## Summary

You've installed claude-agent-templates with lockstep spec-kit version management. Key points:

✅ **Agents** are available in Claude Code
✅ **Slash commands** (`/speckit.*`) work in your project
✅ **Spec-kit version** is pinned and managed
✅ **No version conflicts** with other projects or your existing spec-kit

**Time invested**: 5-10 minutes
**Value delivered**: Structured development workflow with AI-assisted agents

Happy coding! 🚀

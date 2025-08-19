# Claude Agent Templates

This directory contains specialized agent templates for Claude Code.

## Available Agents

### business-requirements-analyst
**Purpose:** Translates high-level business requirements into detailed technical specifications. Bridges the gap between stakeholder needs and development tasks by breaking down complex business processes into implementable features.

**Tools Available:** Bash, Edit, MultiEdit, Write, NotebookEdit

**Use Cases:**
- Converting business requirements to technical specs
- Breaking down complex business processes
- Bridging stakeholder needs and development tasks
- Feature requirement analysis

### solution-architect
**Purpose:** Breaks down complex technical requirements into discrete, implementable work units while considering existing system constraints and technical debt. Applies SOLID, DRY, and KISS principles for clean architecture.

**Tools Available:** Bash, Glob, Grep, LS, Read, WebFetch, TodoWrite, WebSearch, BashOutput, KillBash

**Use Cases:**
- Architectural planning and design
- Breaking down complex features into tasks
- System constraint analysis
- Technical debt assessment
- Clean code architecture guidance

## Installation

Use the installation scripts in the `scripts/` directory to install these agents globally:

- **Linux/macOS:** Run `scripts/install-agents.sh`
- **Windows:** Run `scripts/install-agents.bat`

## Usage in Claude Code

After installation, you can use these agents by referencing them in your Claude Code sessions:

1. Start Claude Code: `claude`
2. View available agents: `/agents`
3. Use agents by referencing them in your requests

The agents will help you with specialized tasks according to their purposes and tool sets.
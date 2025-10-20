# API Contracts: Spec-Kit Integration

**Feature**: 002-spec-kit-integration
**Date**: 2025-10-20

## Overview

This directory contains interface contracts for the spec-kit integration feature. Since this is primarily a CLI/scripting project, the "API" consists of:

1. **Bash Script Interfaces**: Input/output specifications for shell scripts
2. **Python Script Interfaces**: Command-line arguments and JSON output formats
3. **Slash Command Contracts**: Expected inputs and outputs for Claude Code commands
4. **Skill Contracts**: Parameter and return value specifications for Claude skills

---

## Contract Organization

```
contracts/
├── README.md                    # This file
├── bash-scripts.yaml           # Bash script interfaces
├── python-scripts.yaml         # Python validation script interfaces
├── slash-commands.yaml         # Claude Code slash command contracts
└── skills.yaml                 # Claude skill contracts
```

---

## Contract Principles

1. **Text-Based I/O**: All interfaces use text input/output for debuggability
2. **JSON Support**: Machine-readable output available via `--json` flag
3. **Exit Codes**: 0 = success, non-zero = failure with error message to stderr
4. **Idempotent**: Same input always produces same output
5. **Version Compatibility**: Breaking changes require version increment in contract

---

## See Also

- [bash-scripts.yaml](bash-scripts.yaml) - Shell script contracts
- [python-scripts.yaml](python-scripts.yaml) - Python script contracts
- [slash-commands.yaml](slash-commands.yaml) - Slash command contracts
- [skills.yaml](skills.yaml) - Claude skill contracts

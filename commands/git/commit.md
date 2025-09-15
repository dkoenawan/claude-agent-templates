---
allowed-tools: Bash(git diff:*), Bash(git add:*), Bash(git commit:*)
argument-hint: [optional commit message]
description: Read git diff and create conventional commit with proper formatting
---

# Git Conventional Commit

This command helps create properly formatted conventional commits by:
1. Reading the current git diff to understand changes
2. Analyzing the type and scope of changes
3. Creating a conventional commit message following standard format
4. Committing all staged changes

## Usage
```
/git/commit [optional message]
```

If no message is provided, the command will analyze the diff and generate an appropriate conventional commit message.

## Implementation

First, check the current git status and diff:

```bash
git status
git diff --staged
```

If no changes are staged, stage all changes:
```bash
git add .
```

Analyze the changes and create a conventional commit message following the format:
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

Common types:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `test`: Adding missing tests or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools

If arguments are provided, use them as the commit message. Otherwise, generate based on the diff analysis.
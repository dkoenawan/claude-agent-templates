---
allowed-tools: Bash(git checkout:*), Bash(git branch:*), Bash(git fetch:*)
description: Switch to main, remove all local branches except main, and fetch latest
---

# Git Local Cleanup

This command performs a complete local branch cleanup:
1. Switches to the main branch
2. Removes all local branches except main
3. Fetches the latest changes from remote

## Usage
```
/git/clean-up-local
```

## Implementation
Switch to main branch and clean up all local branches:

```bash
git checkout main
git branch | grep -v "main" | xargs -r git branch -D
git fetch
```

This ensures your local repository is clean with only the main branch and up-to-date with the remote repository.
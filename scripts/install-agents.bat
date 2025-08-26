@echo off
setlocal enabledelayedexpansion

echo Installing Claude Agent Templates...

REM Get the script directory
set "SCRIPT_DIR=%~dp0"
set "AGENTS_SOURCE=%SCRIPT_DIR%..\agents"

REM Create global .claude directory structure
if not exist "%USERPROFILE%\.claude\agents" (
    mkdir "%USERPROFILE%\.claude\agents"
)

REM Find and copy all agent files
echo Discovering and copying agent files to %USERPROFILE%\.claude\agents...

set installed_count=0
set missing_count=0

REM Copy all .md files from agents subdirectories (excluding README files)
for /R "%AGENTS_SOURCE%" %%f in (*.md) do (
    set "filename=%%~nxf"
    
    REM Skip README files
    if not "!filename!"=="README.md" (
        copy "%%f" "%USERPROFILE%\.claude\agents\" >nul 2>&1
        if !errorlevel! equ 0 (
            echo ✓ Installed !filename!
            set /a installed_count+=1
        ) else (
            echo ✗ Failed to install !filename!
            set /a missing_count+=1
        )
    )
)

if !installed_count! equ 0 (
    echo ✗ No agent files found in %AGENTS_SOURCE%
    pause
    exit /b 1
)

echo.
echo Installation complete! Installed !installed_count! agents.
echo.
echo To use these agents in Claude Code:
echo 1. Run 'claude' to start Claude Code
echo 2. Use '/agents' command to see available agents
echo 3. Use the agents by referencing them in your requests
echo.
pause
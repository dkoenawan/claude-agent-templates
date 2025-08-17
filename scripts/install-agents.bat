@echo off
setlocal enabledelayedexpansion

echo Installing Claude Agent Templates...

REM Check if we're in a git repository
git rev-parse --is-inside-work-tree >nul 2>&1
if errorlevel 1 (
    echo Error: This script must be run from within a git repository.
    echo Please navigate to your project directory and try again.
    pause
    exit /b 1
)

REM Create .claude directory if it doesn't exist
if not exist ".claude" (
    echo Creating .claude directory...
    mkdir .claude
)

REM Create agents directory if it doesn't exist
if not exist ".claude\agents" (
    echo Creating .claude\agents directory...
    mkdir .claude\agents
)

REM Get the script directory
set "SCRIPT_DIR=%~dp0"
set "AGENTS_SOURCE=%SCRIPT_DIR%..\..claude\agents"

REM Copy agent files
echo Copying agent files...

if exist "%AGENTS_SOURCE%\business-requirements-analyst.md" (
    copy "%AGENTS_SOURCE%\business-requirements-analyst.md" ".claude\agents\" >nul
    echo ✓ Installed business-requirements-analyst agent
) else (
    echo ✗ Warning: business-requirements-analyst.md not found
)

if exist "%AGENTS_SOURCE%\solution-architect.md" (
    copy "%AGENTS_SOURCE%\solution-architect.md" ".claude\agents\" >nul
    echo ✓ Installed solution-architect agent
) else (
    echo ✗ Warning: solution-architect.md not found
)

echo.
echo Installation complete!
echo.
echo Available agents:
echo - business-requirements-analyst: Translates business requirements to technical specs
echo - solution-architect: Breaks down complex features into implementable work units
echo.
echo To use these agents in Claude Code:
echo 1. Run 'claude' to start Claude Code
echo 2. Use '/agents' command to see available agents
echo 3. Use the agents by referencing them in your requests
echo.
pause
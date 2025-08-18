@echo off
setlocal enabledelayedexpansion

echo Installing Claude Agent Templates...


REM Create global .claude directory if it doesn't exist
if not exist "%USERPROFILE%\.claude" (
    echo Creating %USERPROFILE%\.claude directory...
    mkdir "%USERPROFILE%\.claude"
)

REM Create global agents directory if it doesn't exist
if not exist "%USERPROFILE%\.claude\agents" (
    echo Creating %USERPROFILE%\.claude\agents directory...
    mkdir "%USERPROFILE%\.claude\agents"

)

REM Get the script directory
set "SCRIPT_DIR=%~dp0"

set "AGENTS_SOURCE=%SCRIPT_DIR%..\agents"

REM Copy agent files
echo Copying agent files to %USERPROFILE%\.claude\agents...

if exist "%AGENTS_SOURCE%\business-requirements-analyst.md" (
    copy "%AGENTS_SOURCE%\business-requirements-analyst.md" "%USERPROFILE%\.claude\agents\" >nul

    echo ✓ Installed business-requirements-analyst agent
) else (
    echo ✗ Warning: business-requirements-analyst.md not found
)

if exist "%AGENTS_SOURCE%\solution-architect.md" (
    copy "%AGENTS_SOURCE%\solution-architect.md" "%USERPROFILE%\.claude\agents\" >nul

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
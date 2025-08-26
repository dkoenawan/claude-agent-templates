@echo off
setlocal enabledelayedexpansion

echo Installing Claude Agent Templates...


REM Create global .claude directory if it doesn't exist
if not exist "%USERPROFILE%\.claude" (
    echo Creating %USERPROFILE%\.claude directory...
    mkdir "%USERPROFILE%\.claude"
)

REM Remove existing agents directory/link if it exists
if exist "%USERPROFILE%\.claude\agents" (
    rmdir /s /q "%USERPROFILE%\.claude\agents" 2>nul
)

REM Get the script directory
set "SCRIPT_DIR=%~dp0"

set "AGENTS_SOURCE=%SCRIPT_DIR%..\agents"

REM Copy agent files
echo Copying agent files to %USERPROFILE%\.claude\agents...

if exist "%AGENTS_SOURCE%\core\requirements-analyst.md" (
    copy "%AGENTS_SOURCE%\core\requirements-analyst.md" "%USERPROFILE%\.claude\agents\" >nul
    echo ✓ Installed requirements-analyst agent
) else (
    echo ✗ Warning: core\requirements-analyst.md not found
)

if exist "%AGENTS_SOURCE%\core\solution-architect.md" (
    copy "%AGENTS_SOURCE%\core\solution-architect.md" "%USERPROFILE%\.claude\agents\" >nul
    echo ✓ Installed solution-architect agent
) else (
    echo ✗ Warning: core\solution-architect.md not found
)

if exist "%AGENTS_SOURCE%\python\test-engineer-python.md" (
    copy "%AGENTS_SOURCE%\python\test-engineer-python.md" "%USERPROFILE%\.claude\agents\" >nul
    echo ✓ Installed test-engineer-python agent
) else (
    echo ✗ Warning: python\test-engineer-python.md not found
)

if exist "%AGENTS_SOURCE%\python\software-engineer-python.md" (
    copy "%AGENTS_SOURCE%\python\software-engineer-python.md" "%USERPROFILE%\.claude\agents\" >nul
    echo ✓ Installed software-engineer-python agent
) else (
    echo ✗ Warning: python\software-engineer-python.md not found
)

if exist "%AGENTS_SOURCE%\core\documentation.md" (
    copy "%AGENTS_SOURCE%\core\documentation.md" "%USERPROFILE%\.claude\agents\" >nul
    echo ✓ Installed documentation agent
) else (
    echo ✗ Warning: core\documentation.md not found
)

echo.
echo Installation complete!
echo.

echo Available agents:
echo - requirements-analyst: Translates business requirements to technical specs
echo - solution-architect: Breaks down complex features into implementable work units
echo - test-engineer-python: Creates comprehensive unit test strategies with pytest
echo - software-engineer-python: Implements solutions with hexagonal architecture
echo - documentation: Performs final documentation updates and cleanup

echo.
echo To use these agents in Claude Code:
echo 1. Run 'claude' to start Claude Code
echo 2. Use '/agents' command to see available agents
echo 3. Use the agents by referencing them in your requests
echo.
pause
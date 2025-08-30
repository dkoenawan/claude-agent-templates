@echo off
setlocal enabledelayedexpansion

REM Backward compatibility wrapper for install-agents.bat
REM This script has been migrated to use Taskfile for cross-platform consistency.
REM 
REM DEPRECATION NOTICE: This script will be removed in a future version.
REM Please migrate to using: task install
REM
REM For more information about the new Taskfile-based automation, run: task help

echo ⚠️ DEPRECATION NOTICE
echo   This script is deprecated and will be removed in a future version.
echo   Please migrate to the new Taskfile-based system:
echo   - Install Task: https://taskfile.dev/installation/
echo   - Run: task install
echo   - For help: task help
echo.

REM Get the script directory
set "SCRIPT_DIR=%~dp0"
set "PROJECT_ROOT=%SCRIPT_DIR%.."

REM Check if Task is available in PATH and use it if possible
where task >nul 2>&1
if !errorlevel! equ 0 (
    echo Using Taskfile-based installation...
    cd /d "%PROJECT_ROOT%"
    task install
    pause
    exit /b 0
)

REM Check if Task binary is available in project
if exist "%PROJECT_ROOT%\bin\task.exe" (
    echo Using project Task binary...
    cd /d "%PROJECT_ROOT%"
    "%PROJECT_ROOT%\bin\task.exe" install
    pause
    exit /b 0
)

REM Fallback to original implementation
echo Task not found, using legacy implementation...
echo Consider installing Task for better cross-platform support.
echo.

echo Installing Claude Agent Templates...

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
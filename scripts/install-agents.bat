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

REM Link to agents directory for automatic updates
echo Setting up agents directory link...

if exist "%AGENTS_SOURCE%" (
    REM Create symbolic link to agents directory (requires Windows Vista+ and admin rights or Developer Mode)
    mklink /D "%USERPROFILE%\.claude\agents" "%AGENTS_SOURCE%" >nul
    if !ERRORLEVEL! equ 0 (
        REM Count and list available agents
        set /a AGENT_COUNT=0
        for %%f in ("%AGENTS_SOURCE%\*.md") do (
            if not "%%~nxf"=="README.md" (
                set /a AGENT_COUNT+=1
            )
        )
        echo ✓ Linked to !AGENT_COUNT! agents:
        for %%f in ("%AGENTS_SOURCE%\*.md") do (
            if not "%%~nxf"=="README.md" (
                echo   - %%~nf
            )
        )
    ) else (
        echo ✗ Warning: Could not create symbolic link. Trying junction...
        mklink /J "%USERPROFILE%\.claude\agents" "%AGENTS_SOURCE%" >nul
        if !ERRORLEVEL! equ 0 (
            echo ✓ Created junction to agents directory
        ) else (
            echo ✗ Error: Could not create link or junction to agents directory
            echo   Please run as administrator or enable Developer Mode
            exit /b 1
        )
    )
) else (
    echo ✗ Error: agents directory not found at %AGENTS_SOURCE%
    exit /b 1
)

echo.
echo Installation complete!
echo.
echo Agents are now linked and will automatically stay up-to-date with this repository.
echo.
echo To use these agents in Claude Code:
echo 1. Run 'claude' to start Claude Code
echo 2. Use '/agents' command to see available agents
echo 3. Use the agents by referencing them in your requests
echo.
pause
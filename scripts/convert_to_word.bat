@echo off
setlocal enabledelayedexpansion

REM Check if pandoc is installed
where pandoc >nul 2>nul
if %ERRORLEVEL% neq 0 (
    echo Pandoc is not installed. Please install it first.
    echo Visit https://pandoc.org/installing.html for installation instructions.
    exit /b 1
)

REM Get the desktop path
set "DESKTOP_PATH=%USERPROFILE%\Desktop"

REM Convert markdown to Word and save to desktop
pandoc docs\project_documentation.md -o "%DESKTOP_PATH%\CarbonQuest_Documentation.docx" ^
    --reference-doc=docs\template.docx ^
    --toc ^
    --toc-depth=3 ^
    --highlight-style=tango ^
    --pdf-engine=xelatex

echo Documentation has been exported to: %DESKTOP_PATH%\CarbonQuest_Documentation.docx 
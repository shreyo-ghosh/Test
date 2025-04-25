#!/bin/bash

# Check if pandoc is installed
if ! command -v pandoc &> /dev/null; then
    echo "Pandoc is not installed. Please install it first."
    echo "Visit https://pandoc.org/installing.html for installation instructions."
    exit 1
fi

# Get the desktop path
DESKTOP_PATH="$HOME/Desktop"

# Convert markdown to Word and save to desktop
pandoc docs/project_documentation.md -o "$DESKTOP_PATH/CarbonQuest_Documentation.docx" \
    --reference-doc=docs/template.docx \
    --toc \
    --toc-depth=3 \
    --highlight-style=tango \
    --pdf-engine=xelatex

echo "Documentation has been exported to: $DESKTOP_PATH/CarbonQuest_Documentation.docx" 
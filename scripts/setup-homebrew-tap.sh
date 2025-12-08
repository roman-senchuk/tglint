#!/bin/bash
set -e

# Script to set up Homebrew tap repository

TAP_NAME="homebrew-tglint"
GITHUB_USER="roman-senchuk"

echo "Setting up Homebrew tap: $TAP_NAME"

# Check if tap repo exists
if [ -d "$TAP_NAME" ]; then
    echo "Directory $TAP_NAME already exists. Removing..."
    rm -rf "$TAP_NAME"
fi

# Create tap directory structure
mkdir -p "$TAP_NAME/Formula"
cd "$TAP_NAME"

# Initialize git repo
git init
git remote add origin "https://github.com/$GITHUB_USER/$TAP_NAME.git" || true

# Copy formula
cp ../Formula/tglint-binary.rb Formula/tglint.rb

echo ""
echo "✅ Homebrew tap setup complete!"
echo ""
echo "Next steps:"
echo "1. Create repository on GitHub: https://github.com/new"
echo "   Repository name: $TAP_NAME"
echo "   Description: Homebrew tap for tglint"
echo "   Visibility: Public"
echo ""
echo "2. Push to GitHub:"
echo "   cd $TAP_NAME"
echo "   git add ."
echo "   git commit -m 'Initial commit: Add tglint formula'"
echo "   git branch -M main"
echo "   git push -u origin main"
echo ""
echo "3. Users can then install with:"
echo "   brew install $GITHUB_USER/$TAP_NAME/tglint"

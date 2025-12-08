#!/bin/bash
set -e

# Script to prepare Homebrew core PR
# This script helps you prepare the formula for submission to Homebrew core

VERSION="1.0.0"
FORMULA_NAME="tglint"
HOMEBREW_CORE_FORK="${HOMEBREW_CORE_FORK:-$HOME/homebrew-core}"

echo "🚀 Preparing tglint for Homebrew Core submission"
echo ""

# Check if homebrew-core fork exists
if [ ! -d "$HOMEBREW_CORE_FORK" ]; then
    echo "❌ Homebrew core fork not found at: $HOMEBREW_CORE_FORK"
    echo ""
    echo "Please fork https://github.com/Homebrew/homebrew-core and clone it:"
    echo "  git clone https://github.com/YOUR_USERNAME/homebrew-core.git $HOMEBREW_CORE_FORK"
    exit 1
fi

# Get SHA256 of source tarball
echo "📦 Getting SHA256 checksum for source tarball..."
TARBALL_URL="https://github.com/roman-senchuk/tglint/archive/v${VERSION}.tar.gz"
SHA256=$(curl -sL "$TARBALL_URL" | shasum -a 256 | cut -d' ' -f1)

echo "✅ SHA256: $SHA256"
echo ""

# Update formula with SHA256
echo "📝 Updating formula with SHA256..."
FORMULA_FILE="Formula/tglint-core.rb"
if [ ! -f "$FORMULA_FILE" ]; then
    echo "❌ Formula file not found: $FORMULA_FILE"
    exit 1
fi

# Create updated formula
sed "s/sha256 \"\"/sha256 \"$SHA256\"/" "$FORMULA_FILE" > /tmp/tglint-formula.rb

# Copy to homebrew-core
echo "📋 Copying formula to homebrew-core..."
cp /tmp/tglint-formula.rb "$HOMEBREW_CORE_FORK/Formula/$FORMULA_NAME.rb"

# Create branch
cd "$HOMEBREW_CORE_FORK"
if git rev-parse --verify "$FORMULA_NAME" >/dev/null 2>&1; then
    echo "⚠️  Branch '$FORMULA_NAME' already exists. Checking it out..."
    git checkout "$FORMULA_NAME"
else
    echo "🌿 Creating branch '$FORMULA_NAME'..."
    git checkout -b "$FORMULA_NAME"
fi

# Stage changes
git add "Formula/$FORMULA_NAME.rb"

echo ""
echo "✅ Preparation complete!"
echo ""
echo "Next steps:"
echo "1. Review the formula:"
echo "   cd $HOMEBREW_CORE_FORK"
echo "   cat Formula/$FORMULA_NAME.rb"
echo ""
echo "2. Test the formula:"
echo "   brew install --build-from-source Formula/$FORMULA_NAME.rb"
echo "   brew test Formula/$FORMULA_NAME.rb"
echo "   brew audit --strict Formula/$FORMULA_NAME.rb"
echo ""
echo "3. Commit and push:"
echo "   git commit -m \"$FORMULA_NAME: add formula\""
echo "   git push origin $FORMULA_NAME"
echo ""
echo "4. Create PR at:"
echo "   https://github.com/Homebrew/homebrew-core/compare"

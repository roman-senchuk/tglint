#!/bin/bash
set -e

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed or not in PATH"
    echo ""
    echo "To install Go:"
    echo "  brew install go"
    echo ""
    echo "Or download from: https://go.dev/dl/"
    echo ""
    echo "After installation, make sure Go is in your PATH:"
    echo "  export PATH=\$PATH:/usr/local/go/bin"
    echo "  # or for Homebrew:"
    echo "  export PATH=\$PATH:\$(brew --prefix go)/bin"
    exit 1
fi

echo "Building tglint..."
go build -o tglint .

echo "Build successful! Binary created: ./tglint"
echo ""
echo "To install globally:"
echo "  go install ."
echo ""
echo "To test:"
echo "  ./tglint --help"

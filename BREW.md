# Homebrew Installation

## Quick Setup

Use the setup script to create your Homebrew tap:

```bash
./scripts/setup-homebrew-tap.sh
```

Then follow the instructions it prints.

## Option 1: Custom Tap (Recommended)

Create a custom Homebrew tap repository and add the formula there.

### Step 1: Create Tap Repository

1. Create a new repository on GitHub: `homebrew-tglint`
2. Clone it locally:
   ```bash
   git clone https://github.com/roman-senchuk/homebrew-tglint.git
   cd homebrew-tglint
   ```

### Step 2: Add Formula

Copy the formula from `Formula/tglint.rb` to your tap repository:

```bash
mkdir -p Formula
cp /path/to/tglint/Formula/tglint.rb Formula/tglint.rb
```

### Step 3: Update Formula

Update the formula with the correct URL and SHA256:

```ruby
class Tglint < Formula
  desc "Formatter and linter for Terragrunt HCL files"
  homepage "https://github.com/roman-senchuk/tglint"
  url "https://github.com/roman-senchuk/tglint/archive/v1.0.0.tar.gz"
  sha256 "REPLACE_WITH_ACTUAL_SHA256"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin/"tglint", "."
  end

  test do
    system "#{bin}/tglint", "--version"
  end
end
```

To get the SHA256:
```bash
curl -sL https://github.com/roman-senchuk/tglint/archive/v1.0.0.tar.gz | shasum -a 256
```

### Step 4: Install from Tap

Users can install with:
```bash
brew install roman-senchuk/tglint/tglint
```

## Option 2: Binary Installation (Current Release Workflow)

The release workflow creates binaries. You can create a formula that uses the binary:

```ruby
class Tglint < Formula
  desc "Formatter and linter for Terragrunt HCL files"
  homepage "https://github.com/roman-senchuk/tglint"
  version "1.0.0"
  license "MIT"

  if OS.mac?
    if Hardware::CPU.intel?
      url "https://github.com/roman-senchuk/tglint/releases/download/v1.0.0/tglint-darwin-amd64"
      sha256 "REPLACE_WITH_SHA256"
    else
      url "https://github.com/roman-senchuk/tglint/releases/download/v1.0.0/tglint-darwin-arm64"
      sha256 "REPLACE_WITH_SHA256"
    end
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/roman-senchuk/tglint/releases/download/v1.0.0/tglint-linux-amd64"
      sha256 "REPLACE_WITH_SHA256"
    else
      url "https://github.com/roman-senchuk/tglint/releases/download/v1.0.0/tglint-linux-arm64"
      sha256 "REPLACE_WITH_SHA256"
    end
  end

  def install
    bin.install Dir["tglint-*"].first => "tglint"
  end

  test do
    system "#{bin}/tglint", "--version"
  end
end
```

## Option 3: Submit to Homebrew Core

For official Homebrew inclusion, you need to:

1. Meet Homebrew's requirements:
   - Stable, maintained project
   - At least 30 stars on GitHub
   - Active development
   - Useful to others

2. Create a pull request to [homebrew-core](https://github.com/Homebrew/homebrew-core)

3. Follow their [formula guidelines](https://docs.brew.sh/Formula-Cookbook)

## Recommended Approach

Start with **Option 1 (Custom Tap)** - it's the easiest and gives you full control. Once the project gains traction, you can submit to Homebrew core.

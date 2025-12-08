# Submitting tglint to Homebrew Core

This guide will help you submit tglint to the official Homebrew core repository.

## Prerequisites

Before submitting, ensure your project meets Homebrew's requirements:

- ✅ **At least 30 stars** on GitHub
- ✅ **Stable and maintained** project
- ✅ **Useful to others** (not just personal use)
- ✅ **Open source** with a clear license (MIT)
- ✅ **No dependencies** on closed-source software

## Step 1: Prepare Your Formula

The formula file is ready at `Formula/tglint-core.rb`. This is the version for Homebrew core submission.

## Step 2: Fork Homebrew Core

1. Go to https://github.com/Homebrew/homebrew-core
2. Click "Fork" to create your fork
3. Clone your fork locally:

```bash
git clone https://github.com/YOUR_USERNAME/homebrew-core.git
cd homebrew-core
```

## Step 3: Create a Branch

```bash
git checkout -b tglint
```

## Step 4: Add the Formula

```bash
# Copy the formula to the Formula directory
cp /path/to/tglint/Formula/tglint-core.rb Formula/tglint.rb
```

## Step 5: Test the Formula Locally

```bash
# Install from your local formula
brew install --build-from-source Formula/tglint.rb

# Or test without installing
brew test Formula/tglint.rb
brew audit --strict Formula/tglint.rb
```

## Step 6: Get the SHA256

Homebrew requires the SHA256 checksum of the source tarball:

```bash
# Download the tarball and get its SHA256
curl -sL https://github.com/roman-senchuk/tglint/archive/v1.0.0.tar.gz | shasum -a 256
```

Update the `sha256` line in the formula with this value.

## Step 7: Update Formula with SHA256

Edit `Formula/tglint.rb` and replace the empty `sha256 ""` with the actual checksum:

```ruby
sha256 "abc123def456..." # Replace with actual SHA256
```

## Step 8: Commit and Push

```bash
git add Formula/tglint.rb
git commit -m "tglint: add formula"
git push origin tglint
```

## Step 9: Create Pull Request

1. Go to https://github.com/Homebrew/homebrew-core
2. Click "New Pull Request"
3. Select your fork and `tglint` branch
4. Fill out the PR template:
   - **Description**: Brief description of tglint
   - **Formula name**: `tglint`
   - **Current version**: `1.0.0`
   - **License**: `MIT`

## Step 10: PR Template

Use this template for your PR:

```markdown
- [x] Have you followed the [guidelines for contributing](https://docs.brew.sh/Adding-Software-to-Homebrew#adding-a-new-formula)?
- [x] Have you ensured that your commits follow the [commit style guide](https://docs.brew.sh/Formula-Cookbook#commit)?
- [x] Have you checked that there aren't other open [pull requests](https://github.com/Homebrew/homebrew-core/pulls) for the same formula update/change?
- [x] Have you built your formula locally with `brew install --build-from-source <formula>`, where `<formula>` is the name of the formula you're submitting?
- [x] Is your test failing and you did not write a test? If so, please explain why.
- [x] Does your build pass `brew audit --strict <formula>` (after doing `brew install --build-from-source <formula>`)? If this is a new formula, does it pass `brew audit --new <formula>`?

---

**Description:**
tglint is a fast, CI-friendly formatter and linter for Terragrunt HCL files.

**Homepage:** https://github.com/roman-senchuk/tglint
**License:** MIT
**Version:** 1.0.0
```

## What Happens Next

1. **Automated checks**: Homebrew's CI will run tests
2. **Review**: Maintainers will review your PR
3. **Feedback**: You may receive feedback or requested changes
4. **Merge**: Once approved, your formula will be merged
5. **Available**: Users can install with `brew install tglint`

## Tips

- **Be patient**: Review can take time
- **Respond quickly**: Address feedback promptly
- **Follow guidelines**: Read [Homebrew's formula guidelines](https://docs.brew.sh/Formula-Cookbook)
- **Test thoroughly**: Make sure everything works before submitting

## Alternative: Start with Custom Tap

If you don't meet the requirements yet (e.g., < 30 stars), start with a custom tap:

```bash
# Users can install from your tap
brew install roman-senchuk/tglint/tglint
```

Then submit to core once you meet the requirements.

## Resources

- [Homebrew Formula Cookbook](https://docs.brew.sh/Formula-Cookbook)
- [Adding Software to Homebrew](https://docs.brew.sh/Adding-Software-to-Homebrew)
- [Homebrew Core Repository](https://github.com/Homebrew/homebrew-core)

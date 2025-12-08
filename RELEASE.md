# Creating a Release

## Prerequisites

- Git repository with GitHub remote configured
- GitHub Actions enabled for the repository

## Creating a Release

### Step 1: Update Version

Update the version in `cmd/version.go`:

```go
const Version = "1.0.0"
```

### Step 2: Commit and Push Changes

```bash
git add cmd/version.go
git commit -m "Bump version to 1.0.0"
git push origin main
```

### Step 3: Create and Push Tag

```bash
# Create an annotated tag
git tag -a v1.0.0 -m "Release version 1.0.0"

# Push the tag to trigger the release workflow
git push origin v1.0.0
```

### Step 4: Release Workflow

The GitHub Actions workflow will automatically:
1. Build binaries for all platforms (Linux, macOS, Windows)
2. Generate SHA256 checksums
3. Create a GitHub release with all artifacts attached

## Version Format

- Use semantic versioning: `v1.0.0`, `v1.1.0`, `v2.0.0`
- Pre-releases: `v1.0.0-beta`, `v1.0.0-rc1` (will be marked as pre-release)
- The `v` prefix is required in the tag name

## Manual Release

If you need to create a release manually:

1. Go to GitHub repository → Releases → Draft a new release
2. Choose tag: `v1.0.0`
3. Upload binaries from `dist/` directory
4. Add release notes

## Release Artifacts

Each release includes:
- Binaries for Linux (AMD64, ARM64)
- Binaries for macOS (AMD64, ARM64)
- Binary for Windows (AMD64)
- SHA256 checksums for all binaries

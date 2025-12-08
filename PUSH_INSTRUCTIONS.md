# How to Push to Your Repository

## Step 1: Create a repository on GitHub/GitLab/etc.

Create a new repository on your Git hosting service (GitHub, GitLab, Bitbucket, etc.)
- **Don't** initialize it with a README, .gitignore, or license (we already have these)

## Step 2: Add remote and push

### If using GitHub:

```bash
# Replace YOUR_USERNAME and YOUR_REPO_NAME with your actual values
git remote add origin https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
git branch -M main
git push -u origin main
```

### If using SSH:

```bash
git remote add origin git@github.com:YOUR_USERNAME/YOUR_REPO_NAME.git
git branch -M main
git push -u origin main
```

### If using GitLab:

```bash
git remote add origin https://gitlab.com/YOUR_USERNAME/YOUR_REPO_NAME.git
git branch -M main
git push -u origin main
```

## Step 3: Verify

After pushing, verify your repository has all the files:
- Check that all source files are present
- Verify the README.md is visible
- Check that .gitignore is working (binary `tglint` should NOT be in the repo)

## Quick Commands Reference

```bash
# Check current remotes
git remote -v

# Change remote URL if needed
git remote set-url origin NEW_URL

# Push to remote
git push origin main

# Push with tags (if you create version tags later)
git push --tags origin main
```

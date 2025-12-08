# tglint

A fast, CI-friendly formatter and linter for Terragrunt HCL files.

## Features

- **Format** `terragrunt.hcl` files recursively with canonical HCL formatting
- **Lint** with configurable rules
- **CI-friendly** with `--check` mode and proper exit codes
- **Fast** - no dependency on Terragrunt or Terraform CLI
- **Independent** - pure Go implementation

## Installation

### Via Go Install

```bash
go install github.com/roman-senchuk/tglint@latest
```

**Note**: Make sure `$HOME/go/bin` (or `$GOPATH/bin`) is in your PATH. Add this to your `~/.zshrc` or `~/.bashrc`:

```bash
export PATH="$HOME/go/bin:$PATH"
```

Then reload your shell:
```bash
source ~/.zshrc  # or source ~/.bashrc
```

### Build from Source Linux

```bash
git clone https://github.com/roman-senchuk/tglint.git
cd tglint
go build -o tglint
sudo mv tglint /usr/local/bin/  # Optional: install system-wide
```

### Build from Source MacOs

```bash
git clone https://github.com/roman-senchuk/tglint.git
cd tglint
go build -o tglint
mv tglint ~/.local/bin/  # Optional: install system-wide
```

### Verify Installation

```bash
tglint --version
```

## Usage

### Format files

Format all `terragrunt.hcl` files recursively (default behavior):

```bash
tglint fmt
```

Format files in a specific directory (recursively):

```bash
tglint fmt ./infrastructure
```

### Check formatting (CI mode)

Check if files are formatted without modifying them:

```bash
tglint fmt --check
```

Exits with code 2 if any files need formatting.

### Lint files

Run linting rules on all `terragrunt.hcl` files:

```bash
tglint lint
```

Lint files in a specific directory:

```bash
tglint lint ./infrastructure
```

Skip specific rules (comma-separated):

```bash
tglint lint --skip-rules remote_state_required,forbid_absolute_paths
```

Exits with code 1 if any violations are found.

## Lint Rules

The following rules are enforced by default:

1. **remote_state_required** - `remote_state` block must be present
2. **terraform_source_required** - `terraform.source` must be set and non-empty
3. **forbid_hardcoded_aws_account_id** - Hardcoded 12-digit AWS account IDs are not allowed
4. **disallow_empty_inputs** - Empty `inputs = {}` is not allowed
5. **forbid_absolute_paths** - Absolute paths in `terraform.source` are not allowed

You can skip specific rules using the `--skip-rules` flag:

```bash
# Skip remote_state requirement (useful for local development)
tglint lint --skip-rules remote_state_required

# Skip multiple rules
tglint lint --skip-rules remote_state_required,forbid_absolute_paths
```


## Exit Codes

| Case | Code |
|------|------|
| Success | 0 |
| Lint issues | 1 |
| Format mismatch (`--check`) | 2 |
| Fatal error | 3 |

## CI Integration

### GitHub Actions

```yaml
name: tglint

on:
  pull_request:
  push:
    branches: [main]

jobs:
  format:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Install tglint
        run: go install github.com/roman-senchuk/tglint@latest
      
      - name: Check formatting
        run: tglint fmt --check
      
      - name: Run linter
        run: tglint lint
```

### GitLab CI

```yaml
tglint:
  image: golang:1.21
  script:
    - go install github.com/roman-senchuk/tglint@latest
    - tglint fmt --check
    - tglint lint
```

## Examples

### Format Output

```bash
$ tglint fmt
FORMAT live/prod/eks/terragrunt.hcl
FORMAT live/staging/vpc/terragrunt.hcl

Formatted 2 file(s) recursively
```

### Check Formatting (CI)

```bash
$ tglint fmt --check
ERROR: live/prod/eks/terragrunt.hcl is not formatted
```

### Lint Output

```bash
$ tglint lint
live/prod/eks/terragrunt.hcl:5:1: terraform.source is required (terraform_source_required)
live/staging/vpc/terragrunt.hcl:12:3: hardcoded AWS account ID detected (forbid_hardcoded_aws_account_id)
```

## Notes

- Automatically skips `.terraform/` and `.terragrunt-cache/` directories
- Respects `.gitignore` files
- Supports optional `.tglintignore` file

## License

MIT

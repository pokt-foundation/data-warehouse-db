<div align="center">
    <img src=".github/banner.png" alt="Pocket Network logo" width="600"/>
    <!-- TODO Rename header -->
    <h1>Backend Go Repo Template</h1>
    <big>Template repository for creating a new backend Go repo</big>
    <div>
    <br/>
        <a href="https://github.com/pokt-foundation/data-warehouse-db/pulse"><img src="https://img.shields.io/github/last-commit/pokt-foundation/data-warehouse-db.svg"/></a>
        <a href="https://github.com/pokt-foundation/data-warehouse-db/pulls"><img src="https://img.shields.io/github/issues-pr/pokt-foundation/data-warehouse-db.svg"/></a>
        <a href="https://github.com/pokt-foundation/data-warehouse-db/issues"><img src="https://img.shields.io/github/issues-closed/pokt-foundation/data-warehouse-db.svg"/></a>
    </div>
</div>
<br/>

# Development

## Pre-Commit Installation

Before starting development work on this repo, `pre-commit` must be installed.

In order to do so, run the command **`make init-pre-commit`** from the repository root.

Once this is done, the following checks will be performed on every commit to the repo and must pass before the commit is allowed:

### 1. Basic checks

- **check-yaml** - Checks YAML files for errors
- **check-merge-conflict** - Ensures there are no merge conflict markers
- **end-of-file-fixer** - Adds a newline to end of files
- **trailing-whitespace** - Trims trailing whitespace
- **no-commit-to-branch** - Ensures commits are not made directly to `main`

### 2. Go-specific checks

- **go-fmt** - Runs `gofmt`
- **go-imports** - Runs `goimports`
- **golangci-lint** - run `golangci-lint run ./...`
- **go-critic** - run `gocritic check ./...`
- **go-build** - run `go build`
- **go-mod-tidy** - run `go mod tidy -v`

### 3. Detect Secrets

Will detect any potential secrets or sensitive information before allowing a commit.

- Test variables that may resemble secrets (random hex strings, etc.) should be prefixed with `test_`
- The inline comment `pragma: allowlist secret` may be added to a line to force acceptance of a false positive

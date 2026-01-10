# File Organizer

A CLI tool that organizes files in a directory into subdirectories based on configurable criteria: file extension, modification date, or file size.

## Installation

### Option 1: Install globally (recommended)

```bash
go install github.com/Goku-kun/fileorg/cmd/fileorg@latest
```

This installs the binary to `$GOPATH/bin` (usually `~/go/bin`).

Make sure `~/go/bin` is in your PATH. Add this to your `~/.zshrc` or `~/.bashrc`:

```bash
export PATH="$PATH:$HOME/go/bin"
```

Then reload your shell:

```bash
source ~/.zshrc  # or source ~/.bashrc
```

### Option 2: Build locally

```bash
# Clone and navigate to the project root
cd /path/to/fileorg/root

# Build the binary
go build -o fileorg ./cmd/fileorg

# Move to a directory in your PATH
sudo mv fileorg /usr/local/bin/

# Or add current directory to PATH temporarily
export PATH="$PATH:$(pwd)"
```

### Option 3: Run without installing

```bash
go run ./cmd/fileorg [flags] <directory>
```

## Usage

```bash
fileorg [flags] <directory>
```

### Flags

| Flag        | Short | Description                                                     |
| ----------- | ----- | --------------------------------------------------------------- |
| `--by`      |       | Organization strategy: `extension` (default), `date`, or `size` |
| `--dry-run` |       | Preview changes without moving files                            |
| `--verbose` |       | Show detailed output                                            |

### Examples

```bash
# Organize by file extension (default)
fileorg ~/Downloads

# Preview what would happen (dry run)
fileorg --dry-run ~/Downloads

# Organize by modification date (YYYY-MM folders)
fileorg --by date ~/Downloads

# Organize by file size (small/medium/large folders)
fileorg --by size ~/Downloads

# Verbose output
fileorg --verbose ~/Downloads

# Combine flags
fileorg --dry-run --verbose --by date ~/Downloads
```

## Organization Strategies

### Extension (default)

Groups files by their extension:

```
Downloads/
├── pdf/
│   └── report.pdf
├── jpg/
│   └── photo.jpg
└── misc/
    └── README
```

### Date

Groups files by modification date (YYYY-MM):

```
Downloads/
├── 2026-01/
│   └── recent.pdf
└── 2025-06/
    └── old-file.txt
```

### Size

Groups files by size:

| Category | Size Range    |
| -------- | ------------- |
| small    | < 1 MB        |
| medium   | 1 MB - 100 MB |
| large    | > 100 MB      |

```
Downloads/
├── small/
│   └── notes.txt
├── medium/
│   └── document.pdf
└── large/
    └── video.mp4
```

## Features

- Skips hidden files (starting with `.`)
- Handles filename collisions (appends `_1`, `_2`, etc.)
- Dry-run mode for safe previews
- Graceful error handling
- Exit codes: 0 (success), 1 (error)

## Development

```bash
# Run tests
go test ./internal/organizer/...

# Run tests with verbose output
go test -v ./internal/organizer/...

# Build
go build ./cmd/fileorg

# Format code
go fmt ./...

# Vet code
go vet ./...
```

## Project Structure

```
fileorg/
├── cmd/fileorg/main.go           # CLI entry point
├── internal/organizer/
│   ├── file.go                   # FileInfo struct
│   ├── organizer.go              # Core organization logic
│   ├── organizer_test.go         # Tests for safePath
│   ├── strategies.go             # Strategy interface + implementations
│   └── strategies_test.go        # Strategy tests
├── go.mod
└── README.md
```

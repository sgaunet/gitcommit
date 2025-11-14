# gitcommit - Git Commit Date Setter

[![Go Report Card](https://goreportcard.com/badge/github.com/sgaunet/gitcommit)](https://goreportcard.com/report/github.com/sgaunet/gitcommit)
[![GitHub release](https://img.shields.io/github/release/sgaunet/gitcommit.svg)](https://github.com/sgaunet/gitcommit/releases/latest)
![GitHub Downloads](https://img.shields.io/github/downloads/sgaunet/gitcommit/total)
![Coverage Badge](https://raw.githubusercontent.com/wiki/sgaunet/gitcommit/coverage-badge.svg)
[![linter](https://github.com/sgaunet/gitcommit/actions/workflows/coverage.yml/badge.svg)](https://github.com/sgaunet/gitcommit/actions/workflows/coverage.yml)
[![coverage](https://github.com/sgaunet/gitcommit/actions/workflows/coverage.yml/badge.svg)](https://github.com/sgaunet/gitcommit/actions/workflows/coverage.yml)
[![Snapshot Build](https://github.com/sgaunet/gitcommit/actions/workflows/snapshot.yml/badge.svg)](https://github.com/sgaunet/gitcommit/actions/workflows/snapshot.yml)
[![Release Build](https://github.com/sgaunet/gitcommit/actions/workflows/release.yml/badge.svg)](https://github.com/sgaunet/gitcommit/actions/workflows/release.yml)
[![License](https://img.shields.io/github/license/sgaunet/gitcommit.svg)](LICENSE)

A command-line tool that enables developers to create Git commits with specific historical dates while maintaining chronological integrity.

## Features

- ✨ Create commits with custom dates in format `YYYY-MM-DD HH:MM:SS`
- 🔒 Validates dates maintain chronological order (no backdating before last commit)
- 🌍 Automatic local timezone detection
- 🚀 Fast and lightweight (Go stdlib only, no external dependencies)
- 📝 Clear, actionable error messages
- ✅ POSIX-compliant CLI interface (`--help`, `--version`)

## Installation

### From Source

```bash
git clone https://github.com/sgaunet/gitcommit.git
cd gitcommit
make install
```

### Verify Installation

```bash
gitcommit --version
```

## Quick Start

```bash
# Stage your changes
git add .

# Create commit with custom date
gitcommit "2025-02-05 20:19:19" "Historical commit message"
```

## Usage

```bash
gitcommit <date> <message>
```

**Arguments:**
- `<date>`: Date and time in format `YYYY-MM-DD HH:MM:SS`
- `<message>`: Commit message text

**Flags:**
- `--help, -h`: Show usage information
- `--version, -v`: Show version number

## Examples

```bash
# Basic usage
gitcommit "2025-02-05 20:19:19" "Add feature X"

# Backdating offline work
gitcommit "2025-01-15 10:00:00" "Work done offline"

# First commit in new repository
git init
git add README.md
gitcommit "2025-01-01 00:00:00" "Initial commit"
```

## Requirements

- **Go**: 1.21 or later (for building)
- **Git**: 2.20 or later
- **Platforms**: Linux, macOS, Windows

## Date Format Rules

- ✅ Format: `YYYY-MM-DD HH:MM:SS` (24-hour time)
- ✅ Date must be after the last commit in the repository
- ✅ Future dates are allowed
- ✅ Empty repositories accept any date
- ❌ Dates equal to or before the last commit are rejected

## Development

```bash
# Build
make build

# Run tests
make test

# Run linter
make lint

# Format code
make fmt

# View all commands
make help
```

## Testing

```bash
# Run all tests with coverage
make coverage

# Run specific test
go test -v ./internal/datetime/...
```

## Architecture

```
cmd/gitcommit/          # CLI entry point
internal/
  ├── cli/              # CLI logic and orchestration
  ├── datetime/         # Date parsing, validation, formatting
  └── git/              # Git repository operations
tests/
  ├── integration/      # Integration tests
  └── testdata/         # Test fixtures
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Use Cases

- 📅 Organizing repository history chronologically
- 💼 Backdating work done offline
- 🔄 Migrating commits from other version control systems
- 📊 Maintaining accurate project timelines

## Troubleshooting

**Error: "Invalid date format"**
- Use exact format: `YYYY-MM-DD HH:MM:SS`
- Example: `2025-02-05 20:19:19`

**Error: "Chronology violation"**
- Ensure date is after your last commit
- Check: `git log -1 --format="%aI"`

**Error: "Not a Git repository"**
- Run from inside a Git repository
- Or initialize: `git init`

## Performance

- Date parsing & validation: <10ms
- Git operations: <100ms
- Total operation: <200ms (p95)

## Credits

Built with ❤️ using Go standard library only.

---

**Version**: 1.0.0
**Status**: Production Ready

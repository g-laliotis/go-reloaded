# Contributing to go-reloaded

Thank you for your interest in contributing to go-reloaded! This document provides guidelines for contributing to the project.

## Development Setup

### Prerequisites
- Go 1.22 or higher
- Git

### Getting Started
1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/g-laliotis/go-reloaded.git
   cd go-reloaded
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Development Workflow

### Running Tests
```bash
# Run all tests
make test

# Run specific test types
make testcases                   # Comprehensive test cases (24 cases)
go test -v ./internal/...        # Unit tests
make bench                       # Benchmark tests
make clear-cache                 # Clear Go cache

# Docker testing
make docker-build                # Build Docker image locally
make docker-run INPUT=testdata/samples/sample.txt OUTPUT=result.txt

# Using published image (no build needed)
docker run --rm -v "$(pwd):/data" ghcr.io/g-laliotis/go-reloaded:latest /data/testdata/samples/sample.txt /data/result.txt
```

### Code Quality
Before submitting changes:
```bash
make fmt         # Format code
make vet         # Static analysis
make testcases   # Run comprehensive test cases
make test        # Run all tests
```

## Project Structure

```
go-reloaded/
├── .github/
│   ├── ISSUE_TEMPLATE/       # Issue templates
│   ├── workflows/            # CI/CD workflows
│   └── PULL_REQUEST_TEMPLATE.md # PR template
├── cmd/go-reloaded/          # CLI application
├── docs/                     # GitHub Pages site
│   ├── .nojekyll             # Bypass Jekyll processing
│   ├── index.html            # Main site page
│   └── README.md             # Docs readme
├── internal/
│   ├── transformations/      # Core transformation agents
│   └── version/              # Version and build information
├── testdata/                 # Test data and fixtures
│   ├── comprehensive/        # Comprehensive test cases (24 cases)
│   └── samples/              # Sample files for demos
├── *_test.go                 # Various test files
├── AGENTS.md                 # Technical specification
├── CHANGELOG.md              # Version history
├── CODE_OF_CONDUCT.md        # Community guidelines
├── CONTRIBUTING.md           # This file
├── Makefile                  # Build automation
├── README.md                 # Project documentation
└── SECURITY.md               # Security policy
```

## Adding New Features

### Adding a New Agent
1. Create agent file in `internal/transformations/`
2. Implement the `Agent` interface:
   ```go
   type Agent interface {
       Name() string
       Process(input string) string
   }
   ```
3. Add to pipeline in `pipeline.go`
4. Add unit tests
5. Update `AGENTS.md` documentation

### Adding Test Cases
1. Add input/output files to `testdata/comprehensive/`
2. Follow naming convention: `caseXX_name_input.txt` / `caseXX_name_expected_output.txt`
3. Test cases are automatically discovered by `testcases_test.go`

## Code Style

- Follow standard Go conventions
- Use meaningful variable and function names
- Add comments for exported functions
- Keep functions focused and single-purpose
- Maintain the pipeline architecture pattern

## Testing Guidelines

- All new features must include tests
- Maintain 100% test coverage for core functionality
- Add both positive and negative test cases
- Include edge cases and boundary conditions
- Update documentation when behavior changes

## Submitting Changes

1. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
2. Make your changes
3. Run tests and quality checks:
   ```bash
   make testcases && make fmt && make vet
   ```
4. Commit with descriptive messages:
   ```bash
   git commit -m "feat: add new transformation agent"
   ```
5. Push and create a pull request using our [PR template](.github/PULL_REQUEST_TEMPLATE.md)

## Commit Message Format

Use conventional commits:
- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `test:` - Test additions/changes
- `refactor:` - Code refactoring
- `perf:` - Performance improvements

## Questions?

Feel free to open an issue using our templates:
- **Bug reports**: Use the bug report template
- **Feature requests**: Use the feature request template
- **Security issues**: Follow the [Security Policy](SECURITY.md)
- **General questions**: Open a discussion or issue

## Community

This project follows our [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you agree to uphold this code.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
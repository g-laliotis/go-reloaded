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
go test -v -run TestCases        # Comprehensive test cases
go test -v ./internal/...        # Unit tests
make bench                       # Benchmark tests
```

### Code Quality
Before submitting changes:
```bash
make fmt    # Format code
make vet    # Static analysis
make test   # Ensure all tests pass
```

## Project Structure

```
go-reloaded/
├── .github/workflows/        # CI/CD workflows
├── cmd/go-reloaded/          # CLI application
├── internal/
│   ├── transformations/      # Core transformation agents
│   └── version/              # Version and build information
├── testcases/                # Comprehensive test cases (24 cases)
├── testdata/                 # Test fixtures
├── *_test.go                 # Various test files
├── agents.md                 # Technical specification
├── CHANGELOG.md              # Version history
├── CONTRIBUTING.md           # This file
├── Makefile                  # Build automation
└── README.md                 # Project documentation
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
5. Update `agents.md` documentation

### Adding Test Cases
1. Add input/output files to `testcases/`
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
   make test && make fmt && make vet
   ```
4. Commit with descriptive messages:
   ```bash
   git commit -m "feat: add new transformation agent"
   ```
5. Push and create a pull request

## Commit Message Format

Use conventional commits:
- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `test:` - Test additions/changes
- `refactor:` - Code refactoring
- `perf:` - Performance improvements

## Questions?

Feel free to open an issue for:
- Bug reports
- Feature requests
- Questions about the codebase
- Clarification on requirements

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
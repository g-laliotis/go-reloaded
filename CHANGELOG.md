# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2025-11-08

### Added
- Docker support with multi-stage build
- Dockerfile with Go builder and Alpine runtime stages
- .dockerignore for optimized build context
- Docker commands in Makefile (docker-build, docker-run)
- GitHub Actions workflow for automated Docker publishing to GHCR
- Pre-built images available at `ghcr.io/g-laliotis/go-reloaded:latest`
- Docker usage documentation in README.md
- Docker testing instructions in CONTRIBUTING.md
- Docker deployment option in PRESENTATION_GUIDE.md
- Docker quick start in GitHub Pages
- Enhanced README.md with TL;DR Overview heading and target audience
- Added interactive links in Professional Features section
- Improved professional presentation with compact formatting
- Enhancement request template for better issue management
- Version output example in installation documentation
- Non-technical user guide for Docker usage
- Comprehensive --help flag with detailed usage information and examples

### Changed
- Updated project structure documentation to include Docker files
- Enhanced Makefile with containerization commands
- Improved development workflow with Docker testing options
- Prioritized GHCR usage over local builds in documentation
- Bumped version to 1.1.0 reflecting Docker and professional features
- Enhanced README with professional polish and accessibility

## [1.0.0] - 2025-11-02

### Added
- Initial release of go-reloaded text transformation tool
- Pipeline architecture with 4 core agents:
  - HexBinAgent: Convert hex/binary numbers to decimal
  - CaseConvAgent: Handle (up), (low), (cap) transformations with counts
  - PunctuationAgent: Fix spacing and quote formatting
  - ArticleAgent: Convert a/A to an/An before vowels and 'h'
- Comprehensive CLI interface with file I/O
- Complete test suite with 24 test cases covering edge cases
- Performance benchmarking and monitoring
- Professional documentation (README.md, agents.md)
- Build automation with Makefile
- CI/CD integration with GitHub Pages
- Full audit compliance with all requirements

### Features
- Supports hex/binary to decimal conversion: `1E (hex)` → `30`
- Case transformations: `word (up,2)` → `WORD UP`
- Smart punctuation spacing and quote handling
- Article correction: `a apple` → `an apple`
- Preserves line breaks and handles complex nested scenarios
- Robust error handling and edge case management

### Testing
- 24 comprehensive test cases including audit compliance
- Unit tests for individual agents (internal/transformations/agents_test.go)
- Integration tests for full pipeline (main_test.go, edge_cases_test.go)
- Performance benchmarks (benchmark_test.go)
- Comprehensive test runner (testcases_test.go) with consolidated testdata structure
- Edge case validation and boundary testing
- 100% test coverage for core functionality

### Documentation
- Complete README with installation and usage
- Technical specification in agents.md
- Version tracking and build information (internal/version/)
- Contribution guidelines (CONTRIBUTING.md)
- Version history (CHANGELOG.md)
- Inline code documentation
- Example usage and test cases

### Professional Features
- Security policy (SECURITY.md) with vulnerability reporting process
- Code of Conduct (CODE_OF_CONDUCT.md) based on Contributor Covenant 2.1
- GitHub issue templates for structured bug reports and feature requests
- Pull request template with comprehensive quality checklist
- CI/CD pipeline with automated testing on Go 1.22 & 1.23
- GitHub Pages deployment with professional dark theme
- Machine-readable AGENTS.md specification
- Stakeholder presentation guide (PRESENTATION_GUIDE.md)
- Professional project structure following industry standards
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
- Comprehensive test runner (testcases_test.go)
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
# Security Policy

## Supported Versions

We actively support the following versions of go-reloaded:

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security issue, please follow these steps:

### How to Report

1. **Do NOT create a public GitHub issue** for security vulnerabilities
2. Send an email to: **giorgoslaliotis@gmail.com** with:
   - Description of the vulnerability
   - Steps to reproduce the issue
   - Potential impact assessment
   - Any suggested fixes (if available)

### What to Expect

- **Acknowledgment**: We will acknowledge receipt within 48 hours
- **Initial Assessment**: We will provide an initial assessment within 5 business days
- **Updates**: We will keep you informed of our progress
- **Resolution**: We aim to resolve critical issues within 30 days

### Security Best Practices

When using go-reloaded:

- **Input Validation**: Always validate input files before processing
- **File Permissions**: Ensure proper file permissions for input/output files
- **Resource Limits**: Be aware of memory usage with large files
- **Sandboxing**: Consider running in isolated environments for untrusted input

### Scope

This security policy covers:
- The core go-reloaded CLI application
- All transformation agents
- Build and deployment processes

### Out of Scope

- Third-party dependencies (report to respective maintainers)
- Issues in development/testing tools
- Social engineering attacks

## Security Updates

Security updates will be:
- Released as patch versions (e.g., 1.0.1)
- Documented in [CHANGELOG.md](CHANGELOG.md)
- Announced in release notes

## Contact

For security-related questions or concerns:
- Email: **giorgoslaliotis@gmail.com**
- GitHub: [@g-laliotis](https://github.com/g-laliotis)

---

Thank you for helping keep go-reloaded secure!
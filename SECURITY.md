# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| latest  | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability in appledocs, please report it by emailing security@appledocs.local. 

Please include:
- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested remediation (if any)

We aim to respond within 48 hours and provide a fix within 7 days for critical issues.

## Security Measures

### Container Security

We provide multiple Dockerfile variants with different security profiles:

1. **Dockerfile.secure** - Maximum security using scratch base image
   - No shell or package manager
   - Minimal attack surface
   - Non-root execution
   - Read-only root filesystem

2. **Dockerfile.distroless** - High security with better debugging
   - Google's hardened distroless image
   - Non-root by default
   - No shell access
   - Includes necessary certificates

3. **Dockerfile.alpine** - Balanced security for development
   - Small Alpine Linux base
   - Non-root user configuration
   - Security updates applied
   - Shell access for debugging

### Security Scanning

All images are scanned for vulnerabilities using:
- Trivy for vulnerability detection
- Hadolint for Dockerfile best practices
- gosec for Go code security issues
- govulncheck for Go vulnerability checking

### Runtime Security

Recommended runtime configurations:
- Run as non-root user (UID 1001)
- Drop all Linux capabilities
- Use read-only root filesystem
- Set resource limits (CPU, memory, PIDs)
- No new privileges flag enabled

### Supply Chain Security

- All dependencies verified with `go mod verify`
- SBOM (Software Bill of Materials) generated for each release
- Container images signed with cosign
- Build provenance tracked

## Security Best Practices

### For Users

1. **Always use the latest version** - Security patches are included in updates
2. **Run with minimal privileges** - Use provided security configurations
3. **Scan images before deployment** - Run `make security-scan`
4. **Use secure base images** - Prefer scratch or distroless variants
5. **Monitor runtime behavior** - Enable logging and monitoring

### For Contributors

1. **Dependency Management**
   - Review dependencies before adding
   - Check licenses with `make security-check`
   - Run vulnerability scans on PRs

2. **Code Security**
   - Follow secure coding practices
   - Validate all external inputs
   - Use prepared statements for any queries
   - Implement proper error handling

3. **Secret Management**
   - Never commit secrets to the repository
   - Use environment variables or secret management tools
   - Rotate credentials regularly

## Compliance

This project aims to comply with:
- CIS Docker Benchmark
- NIST 800-190 Container Security Guide
- OWASP Container Security Top 10

## Security Tools

Install security tools:
```bash
# Vulnerability scanning
brew install aquasecurity/trivy/trivy

# Go security analysis
go install github.com/securego/gosec/v2/cmd/gosec@latest
go install golang.org/x/vuln/cmd/govulncheck@latest

# Dockerfile linting
brew install hadolint

# License checking
go install github.com/uw-labs/lichen@latest
```

Run security checks:
```bash
# Quick security check
make security-check

# Full security scan
make security-all

# Scan specific image
./security/security-scan.sh appledocs:secure
```

## Security Configuration

See the following files for security configurations:
- `.trivy.yaml` - Trivy scanner configuration
- `.trivyignore.yaml` - Security scan exclusions
- `security/docker-compose.secure.yml` - Secure deployment configuration
- `.github/workflows/security-scan.yml` - CI/CD security pipeline

## Updates and Patches

Security updates are released as soon as possible after discovery. Users are encouraged to:
1. Watch the repository for security advisories
2. Enable automatic security updates where possible
3. Review the CHANGELOG for security-related updates

## Contact

For security concerns, contact: security@appledocs.local

For general issues, use the GitHub issue tracker.
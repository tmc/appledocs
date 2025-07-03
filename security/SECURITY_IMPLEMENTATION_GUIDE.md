# Security Implementation Guide for Appledocs

## Quick Start Security Checklist

### Immediate Actions (Day 1)
- [ ] Build secure container using `Dockerfile.secure`
- [ ] Run security scan with `./security/security-scan.sh`
- [ ] Review and fix any CRITICAL vulnerabilities
- [ ] Implement non-root user execution
- [ ] Enable read-only root filesystem

### Phase 1: Basic Security (Week 1)
- [ ] Implement multi-stage Docker builds
- [ ] Configure resource limits
- [ ] Set up vulnerability scanning in CI
- [ ] Create security documentation
- [ ] Implement basic monitoring

### Phase 2: Advanced Security (Week 2)
- [ ] Integrate runtime security monitoring
- [ ] Implement security policies
- [ ] Set up automated patching
- [ ] Configure secret management
- [ ] Implement SBOM generation

## Building Secure Containers

### 1. Build the Most Secure Image (Scratch-based)
```bash
# Build the scratch-based image for maximum security
docker build -f Dockerfile.secure -t appledocs:secure .

# Verify the image
docker run --rm appledocs:secure -version
```

### 2. Build with Distroless (Recommended for Production)
```bash
# Build the distroless image
docker build -f Dockerfile.distroless -t appledocs:distroless .

# Test the image
docker run --rm appledocs:distroless -help
```

### 3. Build with Alpine (For Development/Debugging)
```bash
# Build the Alpine-based image
docker build -f Dockerfile.alpine -t appledocs:alpine .

# Run with security options
docker run --rm \
  --security-opt=no-new-privileges:true \
  --cap-drop=ALL \
  --read-only \
  appledocs:alpine
```

## Running Security Scans

### Local Security Scanning
```bash
# Make the script executable
chmod +x security/security-scan.sh

# Run comprehensive security scan
./security/security-scan.sh appledocs:secure

# View the report
cat security-reports/summary-latest.md
```

### Trivy Scanning
```bash
# Install Trivy
brew install aquasecurity/trivy/trivy

# Scan for vulnerabilities
trivy image --severity HIGH,CRITICAL appledocs:secure

# Scan Dockerfile for misconfigurations
trivy config Dockerfile.secure

# Scan with custom config
trivy image --config .trivy.yaml appledocs:secure
```

### Dependency Scanning
```bash
# Install tools
go install golang.org/x/vuln/cmd/govulncheck@latest
go install github.com/sonatype-nexus-community/nancy@latest

# Check for vulnerabilities
govulncheck ./...

# Check dependencies with Nancy
go list -json -deps ./... | nancy sleuth
```

## Secure Container Deployment

### 1. Using Docker Compose with Security
```bash
# Deploy with security constraints
docker-compose -f security/docker-compose.secure.yml up -d

# Verify security settings
docker inspect appledocs-secure | jq '.[0].HostConfig.SecurityOpt'
```

### 2. Running with Maximum Security
```bash
docker run -d \
  --name appledocs-secure \
  --security-opt=no-new-privileges:true \
  --security-opt=seccomp=default \
  --cap-drop=ALL \
  --read-only \
  --tmpfs /tmp:noexec,nosuid,size=100M \
  --memory=512m \
  --cpus=1 \
  --pids-limit=100 \
  --user=1001:1001 \
  -v $(pwd)/output:/app/output:rw \
  -v $(pwd)/cache:/app/cache:rw \
  appledocs:secure
```

### 3. Kubernetes Deployment with Security
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: appledocs
spec:
  template:
    spec:
      securityContext:
        runAsNonRoot: true
        runAsUser: 1001
        fsGroup: 1001
        seccompProfile:
          type: RuntimeDefault
      containers:
      - name: appledocs
        image: appledocs:secure
        securityContext:
          allowPrivilegeEscalation: false
          readOnlyRootFilesystem: true
          capabilities:
            drop:
            - ALL
        resources:
          limits:
            memory: "512Mi"
            cpu: "1000m"
          requests:
            memory: "256Mi"
            cpu: "500m"
```

## CI/CD Integration

### GitHub Actions Setup
```bash
# Create workflow directory
mkdir -p .github/workflows

# Copy the security workflow
cp security/security-scan.yml .github/workflows/

# Commit and push
git add .github/workflows/security-scan.yml
git commit -m "Add security scanning workflow"
git push
```

### GitLab CI Integration
```yaml
# .gitlab-ci.yml
security-scan:
  stage: test
  image: aquasec/trivy:latest
  script:
    - trivy image --exit-code 1 --severity HIGH,CRITICAL $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
  allow_failure: false
```

## Monitoring and Alerting

### 1. Set Up Runtime Security with Falco
```bash
# Install Falco
helm repo add falcosecurity https://falcosecurity.github.io/charts
helm install falco falcosecurity/falco

# Apply custom rules
kubectl apply -f security/falco-rules.yaml
```

### 2. Container Activity Monitoring
```bash
# Monitor container activity
docker events --filter container=appledocs-secure

# Check container logs
docker logs -f appledocs-secure

# Monitor resource usage
docker stats appledocs-secure
```

## Secret Management

### 1. Environment Variables (Basic)
```bash
# Never commit secrets to git
echo "API_KEY=your-secret-key" > .env
echo ".env" >> .gitignore

# Run with env file
docker run --env-file .env appledocs:secure
```

### 2. Docker Secrets (Swarm)
```bash
# Create secret
echo "your-secret-key" | docker secret create api_key -

# Use in service
docker service create \
  --name appledocs \
  --secret api_key \
  appledocs:secure
```

### 3. Kubernetes Secrets
```bash
# Create secret
kubectl create secret generic appledocs-secrets \
  --from-literal=api-key=your-secret-key

# Use in deployment
kubectl apply -f security/k8s-deployment-secure.yaml
```

## Compliance and Reporting

### Generate Compliance Report
```bash
# Run CIS Docker Benchmark
docker run --rm --net host --pid host --userns host \
  --cap-add audit_control \
  -v /var/lib:/var/lib:ro \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  --label docker_bench_security \
  docker/docker-bench-security

# Generate SBOM
syft packages appledocs:secure -o spdx-json > sbom.json

# Scan licenses
lichen --config=.lichen.yaml .
```

### Security Metrics Dashboard
```bash
# Create metrics directory
mkdir -p security-metrics

# Collect metrics
echo "# Security Metrics - $(date)" > security-metrics/report.md
echo "## Vulnerability Count" >> security-metrics/report.md
trivy image --format json appledocs:secure | \
  jq '.Results[].Vulnerabilities | length' >> security-metrics/report.md
```

## Troubleshooting Security Issues

### Common Issues and Solutions

1. **Permission Denied Errors**
   ```bash
   # Fix: Ensure correct ownership
   docker run --user $(id -u):$(id -g) appledocs:secure
   ```

2. **Read-only Filesystem Errors**
   ```bash
   # Fix: Mount writable volumes
   docker run -v $(pwd)/cache:/app/cache:rw appledocs:secure
   ```

3. **Missing Certificates**
   ```bash
   # Fix: Ensure ca-certificates are copied in Dockerfile
   COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
   ```

4. **Network Access Issues**
   ```bash
   # Fix: Check security policies and network configuration
   docker run --network=host appledocs:secure
   ```

## Security Best Practices Summary

### Do's
- ✅ Always use non-root users
- ✅ Implement multi-stage builds
- ✅ Scan images before deployment
- ✅ Use minimal base images
- ✅ Keep dependencies updated
- ✅ Implement resource limits
- ✅ Use read-only root filesystem
- ✅ Drop all unnecessary capabilities
- ✅ Sign and verify images
- ✅ Generate and maintain SBOMs

### Don'ts
- ❌ Never run containers as root
- ❌ Don't include build tools in runtime images
- ❌ Avoid using latest tags
- ❌ Don't store secrets in images
- ❌ Never disable security features
- ❌ Don't ignore vulnerability warnings
- ❌ Avoid excessive privileges
- ❌ Don't skip security scans

## Additional Resources

- [NIST Container Security Guide](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-190.pdf)
- [CIS Docker Benchmark](https://www.cisecurity.org/benchmark/docker)
- [OWASP Container Security](https://cheatsheetseries.owasp.org/cheatsheets/Docker_Security_Cheat_Sheet.html)
- [Kubernetes Security Best Practices](https://kubernetes.io/docs/concepts/security/)

## Support and Updates

For security updates and patches:
1. Watch the repository for security advisories
2. Enable Dependabot alerts
3. Subscribe to security mailing lists
4. Regular security review schedule (monthly)

Remember: Security is an ongoing process, not a one-time implementation.
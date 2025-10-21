# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.1] - 2025-10-21

### Fixed
- Removed unsupported `--log-level` flag from deployment template that caused CrashLoopBackOff
- Fixed operator startup issues in Kubernetes environments

### Changed
- Chart version bumped to 0.1.1 with deployment fixes

## [0.1.0] - 2025-10-21

### Added
- Initial release of HelmReboot Operator
- Automatic restart functionality for failed Flux HelmRelease resources
- Support for timeout error detection and remediation
- Comprehensive test coverage (84.2%)
- Docker multi-platform builds (AMD64/ARM64)
- Helm chart for easy deployment
- GitHub Actions workflows for CI/CD
- OCI-based chart publishing to GitHub Container Registry

### Features
- Monitors HelmRelease resources for timeout failures
- Automatically patches HelmRelease to trigger restart
- Configurable reconciliation intervals
- Metrics endpoint for monitoring
- Health probes for Kubernetes deployment
- RBAC configuration for secure operation
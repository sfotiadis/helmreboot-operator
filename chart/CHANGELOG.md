# Helm Chart Changelog

All notable changes to the HelmReboot Operator Helm chart will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.1] - 2025-10-21

### Added
- **Grafana Dashboard Support**: Optional Grafana dashboard ConfigMap creation
  - Configurable via `monitoring.grafanaDashboard.enabled`
  - Customizable datasource and folder placement
  - Comprehensive operator performance and health monitoring
  - Automatic dashboard deployment with chart installation
- **Enhanced ServiceMonitor**: Improved ServiceMonitor with secure/insecure metrics support
  - Automatically configures HTTPS/HTTP based on `operator.secureMetrics` setting
  - Proper TLS configuration for secure metrics
  - Bearer token authentication when needed
- **BREAKING FIX**: Removed unsupported `--log-level` flag from deployment template
  - The controller manager does not recognize the `--log-level` argument
  - This was causing CrashLoopBackOff in deployed pods
  - Flag has been completely removed from `templates/deployment.yaml`

### Changed
- **BREAKING CHANGE**: Updated metrics configuration for better user experience
  - **Default metrics are now HTTP (port 8080) without authentication** for easy access
  - Added `operator.secureMetrics` flag to optionally enable HTTPS + RBAC (port 8443)
  - Service ports now dynamically adjust based on secure metrics setting
  - RBAC for metrics only created when `operator.secureMetrics=true`
  - This makes metrics immediately accessible after installation without extra setup

## [0.1.0] - 2025-10-21

### Added
- Initial Helm chart release for HelmReboot Operator
- Complete Kubernetes deployment templates:
  - Deployment with manager container
  - Service for metrics endpoint
  - ServiceAccount with proper RBAC
  - ClusterRole and ClusterRoleBinding
  - Configurable values for all components

### Chart Features
- Configurable replica count
- Resource limits and requests
- Pod security context
- Image pull secrets support
- Node selector and affinity rules
- Horizontal Pod Autoscaler (HPA) support
- Ingress configuration
- Comprehensive values.yaml with defaults

### Values Configuration
- `operator.reconcileInterval`: Control reconciliation frequency
- `operator.metricsAddr`: Metrics server bind address  
- `operator.probeAddr`: Health probe bind address
- `operator.leaderElect`: Enable/disable leader election
- `image.repository`: Container image repository
- `image.tag`: Override image tag (defaults to appVersion)
- `resources`: Pod resource limits and requests
- `autoscaling`: HPA configuration
- `ingress`: Ingress controller setup
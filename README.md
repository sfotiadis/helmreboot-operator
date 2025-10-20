# HelmReboot Operator

[![Go Report Card](https://goreportcard.com/badge/github.com/sfotiadis/helmreboot-operator)](https://goreportcard.com/report/github.com/sfotiadis/helmreboot-operator)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GitHub release](https://img.shields.io/github/release/sfotiadis/helmreboot-operator.svg)](https://github.com/sfotiadis/helmreboot-operator/releases)

A Kubernetes operator that automatically restarts failed Flux HelmRelease resources when they encounter timeout errors. This operator monitors HelmRelease objects and triggers reconciliation by adding the `fluxcd.io/reconcileAt` annotation when specific failure conditions are detected.

## Overview

The HelmReboot Operator is designed to solve a common issue in GitOps workflows where Flux HelmRelease resources fail due to temporary network issues, registry timeouts, or other transient errors. Instead of manual intervention, this operator automatically detects these failures and triggers a retry by adding reconciliation annotations.

### Key Features

- **Automatic Recovery**: Detects failed HelmRelease resources and triggers automatic retries
- **Smart Detection**: Only restarts releases that failed due to specific timeout errors
- **Monitoring Ready**: Includes Prometheus metrics and comprehensive logging
- **Lightweight**: Minimal resource footprint with efficient reconciliation loops
- **Secure**: Follows Kubernetes RBAC best practices with minimal required permissions
- **Well Tested**: Comprehensive unit and end-to-end test coverage

## How It Works

The operator continuously monitors all HelmRelease resources in the cluster and:

1. **Watches** for HelmRelease objects with failed conditions
2. **Detects** specific error patterns (e.g., "context deadline exceeded")
3. **Triggers** automatic retry by adding the `fluxcd.io/reconcileAt` annotation
4. **Logs** all restart actions for audit and debugging purposes

### Supported Error Patterns

Currently, the operator handles:
- `context deadline exceeded` - Network timeouts during chart operations
- Additional patterns can be easily configured in the controller logic

## Installation

### Prerequisites

- Kubernetes cluster (v1.20+)
- Flux v2 installed and running
- `kubectl` configured to access your cluster

### Quick Start

1. **Install using kubectl:**
   ```bash
   kubectl apply -f https://raw.githubusercontent.com/sfotiadis/helmreboot-operator/main/config/default/kustomization.yaml
   ```

2. **Or build and deploy from source:**
   ```bash
   git clone https://github.com/sfotiadis/helmreboot-operator.git
   cd helmreboot-operator
   make deploy
   ```

3. **Verify installation:**
   ```bash
   kubectl get pods -n helmreboot-operator-system
   ```

### Helm Installation (Coming Soon)

```bash
helm repo add helmreboot-operator https://sfotiadis.github.io/helmreboot-operator
helm install helmreboot-operator helmreboot-operator/helmreboot-operator
```

## Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `METRICS_BIND_ADDRESS` | Address for metrics server | `:8080` |
| `HEALTH_PROBE_BIND_ADDRESS` | Address for health probes | `:8081` |
| `LEADER_ELECT` | Enable leader election | `false` |

### RBAC Permissions

The operator requires the following permissions:
- `get`, `list`, `watch` on HelmRelease resources
- `patch`, `update` on HelmRelease resources (for adding annotations)

## Monitoring

### Prometheus Metrics

The operator exposes metrics on the `/metrics` endpoint:

- `controller_runtime_reconcile_total` - Total number of reconciliations
- `controller_runtime_reconcile_errors_total` - Total number of reconciliation errors
- `controller_runtime_reconcile_time_seconds` - Time spent in reconciliation

### Health Checks

Health endpoints are available:
- `GET /healthz` - Liveness probe
- `GET /readyz` - Readiness probe

## Development

### Prerequisites

- Go 1.21+
- Docker
- kubectl
- Kubebuilder 3.0+

### Local Development

1. **Clone the repository:**
   ```bash
   git clone https://github.com/sfotiadis/helmreboot-operator.git
   cd helmreboot-operator
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run tests:**
   ```bash
   make test
   ```

4. **Run locally against your cluster:**
   ```bash
   make install run
   ```

### Building

```bash
# Build the binary
make build

# Build the Docker image
make docker-build

# Run tests with coverage
make test

# Run linting
make lint
```

### Testing

The project includes comprehensive testing:

```bash
# Unit tests
make test

# End-to-end tests
make test-e2e

# Integration tests with coverage
make test-integration
```

## Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  HelmRelease    │    │  HelmReboot     │    │  Flux           │
│  (Failed)       │───▶│  Operator       │───▶│  Controller     │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                              │
                              ▼
                       ┌─────────────────┐
                       │  Add reconcile  │
                       │  annotation     │
                       └─────────────────┘
```

### Controller Logic

1. **Watch Phase**: Monitor all HelmRelease resources for status changes
2. **Analysis Phase**: Check if the failure matches known recoverable patterns
3. **Action Phase**: Add `fluxcd.io/reconcileAt` annotation to trigger Flux retry
4. **Monitoring Phase**: Log actions and update metrics


## Roadmap

- [ ] Helm Chart for easy installation
- [ ] Support for additional error patterns
- [ ] Configurable retry delays and limits
- [ ] Dashboard for monitoring restart actions
- [ ] Integration with popular monitoring systems
- [ ] Multi-cluster support

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Flux CD](https://fluxcd.io/) for the excellent GitOps toolkit
- [Kubebuilder](https://kubebuilder.io/) for the operator framework
- [Controller Runtime](https://github.com/kubernetes-sigs/controller-runtime) for the underlying controller libraries

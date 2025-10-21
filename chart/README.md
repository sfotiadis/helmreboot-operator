# HelmReboot Operator Helm Chart

![Version: 0.1.1](https://img.shields.io/badge/Version-0.1.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.1.0](https://img.shields.io/badge/AppVersion-0.1.0-informational?style=flat-square)

A Helm chart for deploying the HelmReboot Operator, a Kubernetes operator that automatically restarts failed Flux HelmRelease resources.

## Description

The HelmReboot Operator monitors Flux HelmRelease resources for timeout failures and automatically triggers restarts by patching the resource. This helps maintain the desired state of your GitOps deployments without manual intervention.

## Prerequisites

- Kubernetes 1.19+
- Helm 3.2.0+
- Flux CD v2 installed in the cluster

## Installation

### Install from OCI Registry (Recommended)

```bash
helm install helmreboot-operator oci://ghcr.io/sfotiadis/helmreboot-operator --version 0.1.1
```

### Install from Source

```bash
git clone https://github.com/sfotiadis/helmreboot-operator.git
cd helmreboot-operator
helm install helmreboot-operator ./chart
```

### Custom Installation

```bash
helm install helmreboot-operator oci://ghcr.io/sfotiadis/helmreboot-operator \
  --version 0.1.1 \
  --namespace flux-system \
  --create-namespace \
  --set operator.reconcileInterval=5m \
  --set resources.limits.memory=256Mi
```

## Uninstallation

```bash
helm uninstall helmreboot-operator
```

## Configuration

The following table lists the configurable parameters of the HelmReboot Operator chart and their default values.

### Basic Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `replicaCount` | Number of replicas for the deployment | `1` |
| `image.repository` | Container image repository | `sfotia2s/helmreboot-operator` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `image.tag` | Image tag (overrides appVersion) | `""` |
| `imagePullSecrets` | Image pull secrets | `[]` |
| `nameOverride` | Override chart name | `""` |
| `fullnameOverride` | Override full name | `""` |

### Operator Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `operator.logLevel` | Log level (info, debug, error) | `info` |
| `operator.metricsAddr` | Metrics server bind address | `:8080` |
| `operator.probeAddr` | Health probe bind address | `:8081` |
| `operator.leaderElect` | Enable leader election for HA | `false` |
| `operator.reconcileInterval` | Check interval for HelmReleases | `10m` |
| `operator.secureMetrics` | Enable secure HTTPS metrics with RBAC | `false` |

### Service Account & RBAC

| Parameter | Description | Default |
|-----------|-------------|---------|
| `serviceAccount.create` | Create service account | `true` |
| `serviceAccount.automount` | Auto-mount service account token | `true` |
| `serviceAccount.annotations` | Service account annotations | `{}` |
| `serviceAccount.name` | Service account name | `""` |
| `rbac.create` | Create RBAC resources | `true` |
| `rbac.annotations` | Additional RBAC annotations | `{}` |

### Pod Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `podAnnotations` | Pod annotations | `{}` |
| `podLabels` | Pod labels | `{}` |
| `podSecurityContext.runAsNonRoot` | Run as non-root user | `true` |
| `securityContext.allowPrivilegeEscalation` | Allow privilege escalation | `false` |
| `securityContext.capabilities.drop` | Drop capabilities | `["ALL"]` |
| `securityContext.readOnlyRootFilesystem` | Read-only root filesystem | `true` |
| `securityContext.runAsNonRoot` | Run as non-root user | `true` |

### Service Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `service.type` | Service type | `ClusterIP` |
| `service.ports` | Service ports configuration | See values.yaml |

### Resource Management

| Parameter | Description | Default |
|-----------|-------------|---------|
| `resources.limits.cpu` | CPU limit | `500m` |
| `resources.limits.memory` | Memory limit | `128Mi` |
| `resources.requests.cpu` | CPU request | `10m` |
| `resources.requests.memory` | Memory request | `64Mi` |

### Health Checks

| Parameter | Description | Default |
|-----------|-------------|---------|
| `livenessProbe.httpGet.path` | Liveness probe path | `/healthz` |
| `livenessProbe.httpGet.port` | Liveness probe port | `8081` |
| `livenessProbe.initialDelaySeconds` | Initial delay for liveness probe | `15` |
| `livenessProbe.periodSeconds` | Period for liveness probe | `20` |
| `readinessProbe.httpGet.path` | Readiness probe path | `/readyz` |
| `readinessProbe.httpGet.port` | Readiness probe port | `8081` |
| `readinessProbe.initialDelaySeconds` | Initial delay for readiness probe | `5` |
| `readinessProbe.periodSeconds` | Period for readiness probe | `10` |

### Autoscaling

| Parameter | Description | Default |
|-----------|-------------|---------|
| `autoscaling.enabled` | Enable horizontal pod autoscaler | `false` |
| `autoscaling.minReplicas` | Minimum number of replicas | `1` |
| `autoscaling.maxReplicas` | Maximum number of replicas | `3` |
| `autoscaling.targetCPUUtilizationPercentage` | Target CPU utilization | `80` |

### Monitoring

| Parameter | Description | Default |
|-----------|-------------|---------|
| `monitoring.serviceMonitor.enabled` | Enable Prometheus ServiceMonitor | `false` |
| `monitoring.serviceMonitor.additionalLabels` | Additional labels for ServiceMonitor | `{}` |
| `monitoring.serviceMonitor.interval` | Scrape interval | `30s` |
| `monitoring.serviceMonitor.scrapeTimeout` | Scrape timeout | `10s` |
| `monitoring.grafanaDashboard.enabled` | Enable Grafana Dashboard ConfigMap | `false` |
| `monitoring.grafanaDashboard.namespace` | Namespace for dashboard ConfigMap | `""` |
| `monitoring.grafanaDashboard.additionalLabels` | Additional labels for dashboard | `{}` |
| `monitoring.grafanaDashboard.config.folder` | Grafana folder for dashboard | `"Kubernetes"` |
| `monitoring.grafanaDashboard.config.datasource` | Prometheus datasource name | `"prometheus"` |

### Node Assignment

| Parameter | Description | Default |
|-----------|-------------|---------|
| `nodeSelector` | Node selector for pod assignment | `{}` |
| `tolerations` | Tolerations for pod assignment | `[]` |
| `affinity` | Affinity rules for pod assignment | `{}` |

### Storage

| Parameter | Description | Default |
|-----------|-------------|---------|
| `volumes` | Additional volumes | `[]` |
| `volumeMounts` | Additional volume mounts | `[]` |

## Examples

### Basic Installation with Easy Metrics Access

```bash
helm install helmreboot-operator oci://ghcr.io/sfotiadis/helmreboot-operator
# Metrics are immediately accessible via HTTP (no authentication required)
kubectl port-forward svc/helmreboot-operator 8080:8080 &
curl http://localhost:8080/metrics
```

### Installation with Secure Metrics

```bash
helm install helmreboot-operator oci://ghcr.io/sfotiadis/helmreboot-operator \
  --set operator.secureMetrics=true
# Metrics require authentication (production recommended)
```

### Installation with Custom Reconcile Interval

```bash
helm install helmreboot-operator oci://ghcr.io/sfotiadis/helmreboot-operator \
  --set operator.reconcileInterval=5m
```

### Installation with Resource Limits

```bash
helm install helmreboot-operator oci://ghcr.io/sfotiadis/helmreboot-operator \
  --set resources.limits.cpu=1000m \
  --set resources.limits.memory=256Mi \
  --set resources.requests.cpu=100m \
  --set resources.requests.memory=128Mi
```

### Installation with Monitoring Enabled

```bash
helm install helmreboot-operator oci://ghcr.io/sfotiadis/helmreboot-operator \
  --set monitoring.serviceMonitor.enabled=true \
  --set monitoring.serviceMonitor.additionalLabels.release=prometheus \
  --set monitoring.grafanaDashboard.enabled=true
```

### Installation with Custom Grafana Datasource

```bash
helm install helmreboot-operator oci://ghcr.io/sfotiadis/helmreboot-operator \
  --set monitoring.grafanaDashboard.enabled=true \
  --set monitoring.grafanaDashboard.config.datasource=my-prometheus \
  --set monitoring.grafanaDashboard.namespace=grafana-system
```

### Installation with High Availability

```bash
helm install helmreboot-operator oci://ghcr.io/sfotiadis/helmreboot-operator \
  --set replicaCount=2 \
  --set operator.leaderElect=true \
  --set autoscaling.enabled=true \
  --set autoscaling.maxReplicas=5
```

### Custom Values File

Create a `values.yaml` file:

```yaml
operator:
  reconcileInterval: 5m
  leaderElect: true

resources:
  limits:
    cpu: 1000m
    memory: 256Mi
  requests:
    cpu: 100m
    memory: 128Mi

monitoring:
  serviceMonitor:
    enabled: true
    additionalLabels:
      release: prometheus

autoscaling:
  enabled: true
  maxReplicas: 5
  targetCPUUtilizationPercentage: 70
```

Then install:

```bash
helm install helmreboot-operator oci://ghcr.io/sfotiadis/helmreboot-operator -f values.yaml
```

## Verifying Installation

Check if the operator is running:

```bash
kubectl get pods -l app.kubernetes.io/name=helmreboot-operator
```

View operator logs:

```bash
kubectl logs -l app.kubernetes.io/name=helmreboot-operator
```

Check metrics endpoint:

```bash
kubectl port-forward svc/helmreboot-operator 8080:8080
curl http://localhost:8080/metrics
```

For secure metrics (when `operator.secureMetrics=true`):

```bash
kubectl port-forward svc/helmreboot-operator 8443:8443
TOKEN=$(kubectl create token helmreboot-operator)
curl -k -H "Authorization: Bearer $TOKEN" https://localhost:8443/metrics
```

## Troubleshooting

### Pod Not Starting

If the pod is not starting, check the logs:

```bash
kubectl logs -l app.kubernetes.io/name=helmreboot-operator
```

Common issues:
- RBAC permissions missing
- Invalid configuration values
- Resource constraints

### Operator Not Working

Verify the operator has the necessary permissions:

```bash
kubectl get clusterrole helmreboot-operator
kubectl get clusterrolebinding helmreboot-operator
```

Check if Flux HelmReleases exist:

```bash
kubectl get helmreleases -A
```

### Resource Issues

Monitor resource usage:

```bash
kubectl top pods -l app.kubernetes.io/name=helmreboot-operator
```

## Contributing

Please refer to the main project repository for contribution guidelines:
https://github.com/sfotiadis/helmreboot-operator

## License

This chart is licensed under the same license as the HelmReboot Operator project.
# kube-node-status

This chart installs `kube-node-status` on a Kubernetes cluster using the Helm package manager.

It exposes a Prometheus metric to the cluster, reporting the status (Ready and NotReady) of nodes.

See changelog [here](https://github.com/torbendury/kube-node-status/blob/main/CHANGELOG.md).

## Source Code

You can find the source code of this chart on [GitHub](https://github.com/torbendury/kube-node-status).

## Installation

```bash
helm repo add kube-node-status https://torbendury.github.io/kube-node-status
helm repo update
helm install kube-node-status kube-node-status/kube-node-status --create-namespace --namespace kube-node-status
```

## Configuration

You can configure the following settings:

| Parameter                                       | Description                                                                                                                       | Default                       |
|-------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------|-------------------------------|
| `replicaCount`                                  | Number of replicas                                                                                                                | `2`                           |
| `image.repository`                              | Image repository                                                                                                                  | `torbendury/kube-node-status` |
| `image.tag`                                     | Image tag. Defaults to the Charts appVersion.                                                                                     | `""`                          |
| `image.pullPolicy`                              | Image pull policy                                                                                                                 | `IfNotPresent`                |
| `imagePullSecrets`                              | Image pull secrets                                                                                                                | `[]`                          |
| `nameOverride`                                  | Override the name of the chart                                                                                                    | `""`                          |
| `fullnameOverride`                              | Override the fullname of the chart                                                                                                | `""`                          |
| `serviceAccount.annotations`                    | Annotations to add to the ServiceAccount                                                                                          | `{}`                          |
| `serviceAccount.name`                           | The name of the ServiceAccount. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `""`                          |
| `serviceAccount.automount`                      | AutomountServiceAccountToken controls whether a service account token should be automatically mounted                             | `true`                        |
| `podAnnotations`                                | Annotations to add to the pod                                                                                                     | `{}`                          |
| `podSecurityContext`                            | Security context for the pod                                                                                                      | `{}`                          |
| `podLabels`                                     | Labels to add to the pod                                                                                                          | `{}`                          |
| `securityContext`                               | Security context for the container                                                                                                | `{}`                          |
| `resources`                                     | Pod resource requests and limits                                                                                                  | `{}`                          |
| `service.type`                                  | Kubernetes Service type                                                                                                           | `ClusterIP`                   |
| `autoscaling.enabled`                           | Enable autoscaling for the deployment                                                                                             | `true`                        |
| `autoscaling.minReplicas`                       | Minimum number of replicas                                                                                                        | `2`                           |
| `autoscaling.maxReplicas`                       | Maximum number of replicas                                                                                                        | `5`                           |
| `autoscaling.targetCPUUtilizationPercentage`    | Target CPU utilization percentage                                                                                                 | `80`                          |
| `autoscaling.targetMemoryUtilizationPercentage` | Target memory utilization percentage                                                                                              | `80`                          |
| `metricScraping.googleManagedPrometheus`        | Enable auto instrumented metric scraping for Google Managed Prometheus                                                            | `false`                       |
| `metricScraping.prometheusOperator`             | Enable auto instrumented metric scraping for self managed Prometheus with Prometheus Operator                                     | `false`                       |

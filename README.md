[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/torbendury/kube-node-status.svg)](https://github.com/torbendury/kube-node-status)
[![Go Report Card](https://goreportcard.com/badge/github.com/torbendury/kube-node-status)](https://goreportcard.com/report/github.com/torbendury/kube-node-status)
![GitHub license](https://img.shields.io/github/license/torbendury/kube-node-status.svg)
![GitHub release](https://img.shields.io/github/release/torbendury/kube-node-status.svg)
[![GitHub latest commit](https://badgen.net/github/last-commit/torbendury/kube-node-status)](https://GitHub.com/torbendury/kube-node-status/commit/)

# kube-node-status

`kube-node-status` is a containerized Golang application which continuously reports the status of Kubernetes nodes.

This becomes useful if you use a [managed version of kube-state-metrics](https://cloud.google.com/kubernetes-engine/docs/how-to/kube-state-metrics) which does not support retrieving Kubernetes node metrics and you are not able to use a self-managed version of it.

**This project is under active development. Use with caution.**

## Installing

Install it by using Helm:

```bash
helm repo add kube-node-status https://torbendury.github.io/kube-node-status
helm repo update
helm install kube-node-status kube-node-status/kube-node-status --create-namespace --namespace kube-node-status
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. Basically, you can do whatever you want with this project, but you have to include the license and the license notice. And if you break something while using this piece of software or anything around it, I'm not responsible for that.

## Release

The application is released via Helm Chart releases using GitHub Actions, GitHub Releases and GitHub Pages.

## Metrics exposed

```prometheus
# HELP kube_node_status The status of the nodes in the cluster (0 = not ready, 1 = ready)
# TYPE kube_node_status gauge
kube_node_status{node="minikube"} 1
# HELP kube_node_status_nodes_not_ready The total number of NotReady nodes in the cluster
# TYPE kube_node_status_nodes_not_ready gauge
kube_node_status_nodes_not_ready 0
# HELP kube_node_status_nodes_not_ready_ratio The ratio of NotReady nodes in the cluster
# TYPE kube_node_status_nodes_not_ready_ratio gauge
kube_node_status_nodes_not_ready_ratio 0
# HELP kube_node_status_nodes_ready The total number of Ready nodes in the cluster
# TYPE kube_node_status_nodes_ready gauge
kube_node_status_nodes_ready 1
# HELP kube_node_status_nodes_total The total number of nodes in the cluster
# TYPE kube_node_status_nodes_total gauge
kube_node_status_nodes_total 1
```

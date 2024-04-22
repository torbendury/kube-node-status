[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/torbendury/kube-node-status.svg)](https://github.com/torbendury/kube-node-status)
[![Go Report Card](https://goreportcard.com/badge/github.com/torbendury/kube-node-status)](https://goreportcard.com/report/github.com/torbendury/kube-node-status)
![GitHub license](https://img.shields.io/github/license/torbendury/kube-node-status.svg)
![GitHub release](https://img.shields.io/github/release/torbendury/kube-node-status.svg)
[![GitHub latest commit](https://badgen.net/github/last-commit/torbendury/kube-node-status)](https://GitHub.com/torbendury/kube-node-status/commit/)

# kube-node-status

`kube-node-status` is a containerized Golang application which continuously reports the status of Kubernetes nodes.

This becomes useful if you use a [managed version of kube-state-metrics](https://cloud.google.com/kubernetes-engine/docs/how-to/kube-state-metrics) which does not support retrieving Kubernetes node metrics and you are not able to use a self-managed version of it.

## Installing

Install it by using Helm:

```bash
helm repo add kube-node-status https://torbendury.github.io/kube-node-status
helm repo update
helm install kube-node-status kube-node-status/kube-node-status --create-namespace --namespace kube-node-status
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. Basically, you can do whatever you want with this project, but you have to include the license and the license notice. And if you break something while using this piece of software or anything around it, I'm not responsible for that.


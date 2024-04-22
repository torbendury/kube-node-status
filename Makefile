.RECIPEPREFIX = >
.PHONY: test build savebuild kube kuberun github helm local localproxy

### Variables
RELEASE_IMAGE_NAME := torbendury/kube-node-status

### Run the tests
test:
> go test -v -race ./...
> helm lint helm/kube-node-status
> go mod verify
> go vet ./...
> go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
> go run golang.org/x/vuln/cmd/govulncheck@latest ./...

### Build the release container
build:
> docker build --no-cache -t $(RELEASE_IMAGE_NAME):latest --target release .


savebuild:
> docker image save -o image.tar $(RELEASE_IMAGE_NAME):latest

### Create local test cluster
kube: savebuild
> minikube start
> minikube -p minikube docker-env
> minikube image load image.tar
> sleep 10

kuberun:
> kubectl run --rm -i kube-node-status --image=$(RELEASE_IMAGE_NAME):latest --image-pull-policy=IfNotPresent --port=2112

### Install the Helm Chart
helm:
> helm upgrade --install kube-node-status ./helm/kube-node-status --set image.tag=latest --namespace kube-node-status --create-namespace

### Create a complete local environment
local: build kube

localproxy:
> kubectl port-forward svc/kube-node-status 2112:2112

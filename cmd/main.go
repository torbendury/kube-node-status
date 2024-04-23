package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics(kc *kubernetes.Clientset) {
	go func() {
		for {
			nodes, err := kc.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
			if statusError, isStatus := err.(*errors.StatusError); isStatus {
				logger.Error("error getting nodes", statusError.ErrStatus.Message, err)
			} else if err != nil {
				logger.Error("error getting nodes", err)
				time.Sleep(500 * time.Millisecond)
				continue
			}
			nodesTotal.Set(float64(len(nodes.Items)))
			// ! This might not be the most efficient way to do this, but it's the most idiomatic for the moment
			// ! and it's not a big deal if we are not running in a huge cluster.
			readyNodes := 0
			for _, node := range nodes.Items {
				var status float64
				for _, condition := range node.Status.Conditions {
					if condition.Type == "Ready" && condition.Status == "True" {
						readyNodes++
						status = 1
					} else {
						status = 0
					}
				}
				nodeStatus.WithLabelValues(node.Name).Set(status)
			}
			nodesReady.Set(float64(readyNodes))
			nodesNotReady.Set(float64(len(nodes.Items) - readyNodes))
			time.Sleep(5 * time.Second)
		}
	}()
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var (
	nodesTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kube_node_status_nodes_total",
		Help: "The total number of nodes in the cluster",
	})

	nodesReady = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kube_node_status_nodes_ready",
		Help: "The total number of nodes in the cluster",
	})

	nodesNotReady = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kube_node_status_nodes_not_ready",
		Help: "The total number of nodes in the cluster",
	})

	nodeStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_node_status",
		Help: "The status of the nodes in the cluster (0 = not ready, 1 = ready)",
	}, []string{"node"})

	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	reg = prometheus.NewRegistry()
)

func init() {
	reg.MustRegister(nodesTotal, nodesReady, nodesNotReady, nodeStatus)
}

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	recordMetrics(clientset)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	http.HandleFunc("/healthz", healthz)
	logger.Error("error serving", http.ListenAndServe(":2112", nil))
}

package worker

import (
	"context"
	"fmt"
	"os"
	"time"

	"metric-collector/pkg/api/grpc"
	"metric-collector/pkg/api/kubelet"
	"metric-collector/pkg/decode"
	"metric-collector/pkg/storage"

	"github.com/prometheus/client_golang/prometheus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

type PodCollector struct {
	ClientSet     *kubernetes.Clientset
	KubeletClient *kubelet.KubeletClient
	pod           *grpc.PodManager
}

func NewPodManager(reg prometheus.Registerer, clientset *kubernetes.Clientset, kubeletClient *kubelet.KubeletClient) *PodCollector {
	newpodmetric := grpc.PodManager{}
	return &PodCollector{
		KubeletClient: kubeletClient,
		ClientSet:     clientset,
		pod:           &newpodmetric,
	}
}

func (nc PodCollector) Collect() {
	for {
		nodeName := os.Getenv("NODE_NAME")
		node, err := nc.ClientSet.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
		if err != nil {
			klog.Errorln(err)
		}

		totalCPUQuantity := node.Status.Allocatable["cpu"]
		totalCPU, _ := totalCPUQuantity.AsInt64()
		totalMemoryQuantity := node.Status.Allocatable["memory"]
		totalMemory, _ := totalMemoryQuantity.AsInt64()
		totalStorageQuantity := node.Status.Allocatable["ephemeral-storage"]
		totalStorage, _ := totalStorageQuantity.AsInt64()

		collection, err := PodScrap(nc.KubeletClient, node)
		if err != nil {
			klog.Errorln(err)
		}

		nc.pod.Podmetric = make([]*grpc.PodMetric, 0)

		for _, podMetric := range collection.Metricsbatch.Pods {
			if podMetric.Namespace == "cdi" || podMetric.Namespace == "keti-controller-system" || podMetric.Namespace == "keti-system" || podMetric.Namespace == "kube-flannel" || podMetric.Namespace == "kube-node-lease" || podMetric.Namespace == "kube-public" || podMetric.Namespace == "kube-system" || podMetric.Namespace == "kubevirt" {
				continue
			}

			podmetric_data := grpc.PodMetric{}

			cpuUsage, _ := podMetric.CPUUsageNanoCores.AsInt64()
			cpuPercent := (float64(cpuUsage) * Nano) / float64(totalCPU)
			memoryUsage, _ := podMetric.MemoryUsageBytes.AsInt64()
			memoryPercent := float64(memoryUsage) / float64(totalMemory)
			storageUsage, _ := podMetric.FsUsedBytes.AsInt64()
			storagePercent := float64(storageUsage) / float64(totalStorage)
			prevRx := podMetric.PrevNetworkRxBytes
			prevTx := podMetric.PrevNetworkTxBytes
			networkrxUsage, _ := podMetric.NetworkRxBytes.AsInt64()
			networkrxUsage = networkrxUsage - prevRx
			networktxUsage, _ := podMetric.NetworkTxBytes.AsInt64()
			networktxUsage = networktxUsage - prevTx

			podmetric_data.PodName = podMetric.Name
			podmetric_data.CPUUsage = float32(cpuUsage)
			podmetric_data.TotalCPU = totalCPU
			podmetric_data.CpuPercent = float32(cpuPercent)

			podmetric_data.MemoryUsage = memoryUsage
			podmetric_data.TotalMemory = totalMemory
			podmetric_data.MemoryPercent = float32(memoryPercent)

			podmetric_data.StorageUsage = memoryUsage
			podmetric_data.TotalStorage = totalStorage
			podmetric_data.StoragePercent = float32(storagePercent)

			podmetric_data.NetworkRx = networkrxUsage
			podmetric_data.NetworkTx = networktxUsage

			nc.pod.Podmetric = append(nc.pod.Podmetric, &podmetric_data)
		}

		time.Sleep(time.Second * 5)
	}

}

func PodScrap(kubelet_client *kubelet.KubeletClient, node *v1.Node) (*storage.Collection, error) {
	metrics, err := CollectPod(kubelet_client, node)
	if err != nil {
		klog.Errorf("unable to fully scrape metrics from node %s: %v", node.Name, err)
	}

	var errs []error
	res := &storage.Collection{
		Metricsbatch: metrics,
		ClusterName:  os.Getenv("CLUSTER_NAME"),
	}

	return res, utilerrors.NewAggregate(errs)
}

func CollectPod(kubelet_client *kubelet.KubeletClient, node *v1.Node) (*storage.MetricsBatch, error) {
	summary, err := kubelet_client.GetSummary()
	if err != nil {
		return nil, fmt.Errorf("unable to fetch metrics from Kubelet %s (%s): %v", node.Name, node.Status.Addresses[0].Address, err)
	}

	return decode.DecodePodBatch(summary, prevPods)
}

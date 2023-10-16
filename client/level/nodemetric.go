package level

import (
	"context"
	"time"

	levelv1 "github.com/KETI-Hybrid/keti-controller/apis/level/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type NodeMetricGetter interface {
	NodeMetrics() NodeMetricInterface
}

type NodeMetricInterface interface {
	Create(*levelv1.NodeMetric) (*levelv1.NodeMetric, error)
	Update(*levelv1.NodeMetric) (*levelv1.NodeMetric, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*levelv1.NodeMetric, error)
	List(opts metav1.ListOptions) (*levelv1.NodeMetricList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *levelv1.NodeMetric, err error)
}

type nodemetrics struct {
	client rest.Interface
}

func newNodeMetrics(c *LevelV1Client) *nodemetrics {
	return &nodemetrics{
		client: c.RESTClient(),
	}
}

func (c *nodemetrics) Get(name string, options metav1.GetOptions) (result *levelv1.NodeMetric, err error) {
	result = &levelv1.NodeMetric{}
	err = c.client.Get().
		Resource("nodemetrics").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *nodemetrics) List(opts metav1.ListOptions) (result *levelv1.NodeMetricList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &levelv1.NodeMetricList{}
	err = c.client.Get().
		Resource("nodemetrics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *nodemetrics) Create(collector *levelv1.NodeMetric) (result *levelv1.NodeMetric, err error) {
	result = &levelv1.NodeMetric{}
	err = c.client.Post().
		Resource("nodemetrics").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *nodemetrics) Update(collector *levelv1.NodeMetric) (result *levelv1.NodeMetric, err error) {
	result = &levelv1.NodeMetric{}
	err = c.client.Put().
		Resource("nodemetrics").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *nodemetrics) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nodemetrics").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *nodemetrics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *levelv1.NodeMetric, err error) {
	result = &levelv1.NodeMetric{}
	err = c.client.Patch(pt).
		Resource("nodemetrics").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

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

type PodMetricGetter interface {
	PodMetrics() PodMetricInterface
}

type PodMetricInterface interface {
	Create(*levelv1.PodMetric) (*levelv1.PodMetric, error)
	Update(*levelv1.PodMetric) (*levelv1.PodMetric, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*levelv1.PodMetric, error)
	List(opts metav1.ListOptions) (*levelv1.PodMetricList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *levelv1.PodMetric, err error)
}

type podmetrics struct {
	client rest.Interface
}

func newPodMetrics(c *LevelV1Client) *podmetrics {
	return &podmetrics{
		client: c.RESTClient(),
	}
}

func (c *podmetrics) Get(name string, options metav1.GetOptions) (result *levelv1.PodMetric, err error) {
	result = &levelv1.PodMetric{}
	err = c.client.Get().
		Resource("podmetrics").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *podmetrics) List(opts metav1.ListOptions) (result *levelv1.PodMetricList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &levelv1.PodMetricList{}
	err = c.client.Get().
		Resource("podmetrics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *podmetrics) Create(collector *levelv1.PodMetric) (result *levelv1.PodMetric, err error) {
	result = &levelv1.PodMetric{}
	err = c.client.Post().
		Resource("podmetrics").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *podmetrics) Update(collector *levelv1.PodMetric) (result *levelv1.PodMetric, err error) {
	result = &levelv1.PodMetric{}
	err = c.client.Put().
		Resource("podmetrics").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *podmetrics) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("podmetrics").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *podmetrics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *levelv1.PodMetric, err error) {
	result = &levelv1.PodMetric{}
	err = c.client.Patch(pt).
		Resource("podmetrics").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

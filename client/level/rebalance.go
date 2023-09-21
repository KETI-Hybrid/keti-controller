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

type RebalanceGetter interface {
	Rebalances() RebalanceInterface
}

type RebalanceInterface interface {
	Create(*levelv1.Rebalance) (*levelv1.Rebalance, error)
	Update(*levelv1.Rebalance) (*levelv1.Rebalance, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*levelv1.Rebalance, error)
	List(opts metav1.ListOptions) (*levelv1.RebalanceList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *levelv1.Rebalance, err error)
}

type rebalances struct {
	client rest.Interface
}

func newRebalances(c *LevelV1Client) *rebalances {
	return &rebalances{
		client: c.RESTClient(),
	}
}

func (c *rebalances) Get(name string, options metav1.GetOptions) (result *levelv1.Rebalance, err error) {
	result = &levelv1.Rebalance{}
	err = c.client.Get().
		Resource("rebalances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *rebalances) List(opts metav1.ListOptions) (result *levelv1.RebalanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &levelv1.RebalanceList{}
	err = c.client.Get().
		Resource("rebalances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *rebalances) Create(collector *levelv1.Rebalance) (result *levelv1.Rebalance, err error) {
	result = &levelv1.Rebalance{}
	err = c.client.Post().
		Resource("rebalances").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *rebalances) Update(collector *levelv1.Rebalance) (result *levelv1.Rebalance, err error) {
	result = &levelv1.Rebalance{}
	err = c.client.Put().
		Resource("rebalances").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *rebalances) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("rebalances").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *rebalances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *levelv1.Rebalance, err error) {
	result = &levelv1.Rebalance{}
	err = c.client.Patch(pt).
		Resource("rebalances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

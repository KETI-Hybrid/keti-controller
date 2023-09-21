package resource

import (
	"context"
	"time"

	resourcev1 "github.com/KETI-Hybrid/keti-controller/apis/resource/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type StatefulsetGetter interface {
	Statefulsets() StatefulsetInterface
}

type StatefulsetInterface interface {
	Create(*resourcev1.Statefulset) (*resourcev1.Statefulset, error)
	Update(*resourcev1.Statefulset) (*resourcev1.Statefulset, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*resourcev1.Statefulset, error)
	List(opts metav1.ListOptions) (*resourcev1.StatefulsetList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.Statefulset, err error)
}

type statefulsets struct {
	client rest.Interface
}

func newStatefulsets(c *ResourceV1Client) *statefulsets {
	return &statefulsets{
		client: c.RESTClient(),
	}
}

func (c *statefulsets) Get(name string, options metav1.GetOptions) (result *resourcev1.Statefulset, err error) {
	result = &resourcev1.Statefulset{}
	err = c.client.Get().
		Resource("statefulsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *statefulsets) List(opts metav1.ListOptions) (result *resourcev1.StatefulsetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &resourcev1.StatefulsetList{}
	err = c.client.Get().
		Resource("statefulsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a statefulset and creates it.  Returns the server's representation of the statefulset, and an error, if there is any.
func (c *statefulsets) Create(collector *resourcev1.Statefulset) (result *resourcev1.Statefulset, err error) {
	result = &resourcev1.Statefulset{}
	err = c.client.Post().
		Resource("statefulsets").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a statefulset and updates it. Returns the server's representation of the statefulset, and an error, if there is any.
func (c *statefulsets) Update(collector *resourcev1.Statefulset) (result *resourcev1.Statefulset, err error) {
	result = &resourcev1.Statefulset{}
	err = c.client.Put().
		Resource("statefulsets").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the statefulset and deletes it. Returns an error if one occurs.
func (c *statefulsets) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("statefulsets").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched statefulset.
func (c *statefulsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.Statefulset, err error) {
	result = &resourcev1.Statefulset{}
	err = c.client.Patch(pt).
		Resource("statefulsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

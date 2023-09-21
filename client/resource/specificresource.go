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

type SpecificResourceGetter interface {
	SpecificResources() SpecificResourceInterface
}

type SpecificResourceInterface interface {
	Create(*resourcev1.SpecificResource) (*resourcev1.SpecificResource, error)
	Update(*resourcev1.SpecificResource) (*resourcev1.SpecificResource, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*resourcev1.SpecificResource, error)
	List(opts metav1.ListOptions) (*resourcev1.SpecificResourceList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.SpecificResource, err error)
}

type specificResources struct {
	client rest.Interface
}

func newSpecificResources(c *ResourceV1Client) *specificResources {
	return &specificResources{
		client: c.RESTClient(),
	}
}

func (c *specificResources) Get(name string, options metav1.GetOptions) (result *resourcev1.SpecificResource, err error) {
	result = &resourcev1.SpecificResource{}
	err = c.client.Get().
		Resource("specificResources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *specificResources) List(opts metav1.ListOptions) (result *resourcev1.SpecificResourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &resourcev1.SpecificResourceList{}
	err = c.client.Get().
		Resource("specificResources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a specificResource and creates it.  Returns the server's representation of the specificResource, and an error, if there is any.
func (c *specificResources) Create(collector *resourcev1.SpecificResource) (result *resourcev1.SpecificResource, err error) {
	result = &resourcev1.SpecificResource{}
	err = c.client.Post().
		Resource("specificResources").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a specificResource and updates it. Returns the server's representation of the specificResource, and an error, if there is any.
func (c *specificResources) Update(collector *resourcev1.SpecificResource) (result *resourcev1.SpecificResource, err error) {
	result = &resourcev1.SpecificResource{}
	err = c.client.Put().
		Resource("specificResources").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the specificResource and deletes it. Returns an error if one occurs.
func (c *specificResources) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("specificResources").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched specificResource.
func (c *specificResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.SpecificResource, err error) {
	result = &resourcev1.SpecificResource{}
	err = c.client.Patch(pt).
		Resource("specificResources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

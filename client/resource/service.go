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

type ServiceGetter interface {
	Services() ServiceInterface
}

type ServiceInterface interface {
	Create(*resourcev1.Service) (*resourcev1.Service, error)
	Update(*resourcev1.Service) (*resourcev1.Service, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*resourcev1.Service, error)
	List(opts metav1.ListOptions) (*resourcev1.ServiceList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.Service, err error)
}

type services struct {
	client rest.Interface
}

func newServices(c *ResourceV1Client) *services {
	return &services{
		client: c.RESTClient(),
	}
}

func (c *services) Get(name string, options metav1.GetOptions) (result *resourcev1.Service, err error) {
	result = &resourcev1.Service{}
	err = c.client.Get().
		Resource("services").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *services) List(opts metav1.ListOptions) (result *resourcev1.ServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &resourcev1.ServiceList{}
	err = c.client.Get().
		Resource("services").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a service and creates it.  Returns the server's representation of the service, and an error, if there is any.
func (c *services) Create(collector *resourcev1.Service) (result *resourcev1.Service, err error) {
	result = &resourcev1.Service{}
	err = c.client.Post().
		Resource("services").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a service and updates it. Returns the server's representation of the service, and an error, if there is any.
func (c *services) Update(collector *resourcev1.Service) (result *resourcev1.Service, err error) {
	result = &resourcev1.Service{}
	err = c.client.Put().
		Resource("services").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the service and deletes it. Returns an error if one occurs.
func (c *services) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("services").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched service.
func (c *services) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.Service, err error) {
	result = &resourcev1.Service{}
	err = c.client.Patch(pt).
		Resource("services").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

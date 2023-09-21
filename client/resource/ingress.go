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

type IngressGetter interface {
	Ingresss() IngressInterface
}

type IngressInterface interface {
	Create(*resourcev1.Ingress) (*resourcev1.Ingress, error)
	Update(*resourcev1.Ingress) (*resourcev1.Ingress, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*resourcev1.Ingress, error)
	List(opts metav1.ListOptions) (*resourcev1.IngressList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.Ingress, err error)
}

type ingresss struct {
	client rest.Interface
}

func newIngresss(c *ResourceV1Client) *ingresss {
	return &ingresss{
		client: c.RESTClient(),
	}
}

func (c *ingresss) Get(name string, options metav1.GetOptions) (result *resourcev1.Ingress, err error) {
	result = &resourcev1.Ingress{}
	err = c.client.Get().
		Resource("ingresss").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *ingresss) List(opts metav1.ListOptions) (result *resourcev1.IngressList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &resourcev1.IngressList{}
	err = c.client.Get().
		Resource("ingresss").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *ingresss) Create(collector *resourcev1.Ingress) (result *resourcev1.Ingress, err error) {
	result = &resourcev1.Ingress{}
	err = c.client.Post().
		Resource("ingresss").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *ingresss) Update(collector *resourcev1.Ingress) (result *resourcev1.Ingress, err error) {
	result = &resourcev1.Ingress{}
	err = c.client.Put().
		Resource("ingresss").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *ingresss) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ingresss").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *ingresss) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.Ingress, err error) {
	result = &resourcev1.Ingress{}
	err = c.client.Patch(pt).
		Resource("ingresss").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

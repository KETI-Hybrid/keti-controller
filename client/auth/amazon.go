package auth

import (
	"context"
	"time"

	authv1 "github.com/KETI-Hybrid/keti-controller/apis/auth/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type AmazonGetter interface {
	Amazons() AmazonInterface
}

type AmazonInterface interface {
	Create(*authv1.Amazon) (*authv1.Amazon, error)
	Update(*authv1.Amazon) (*authv1.Amazon, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*authv1.Amazon, error)
	List(opts metav1.ListOptions) (*authv1.AmazonList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.Amazon, err error)
}

type amazons struct {
	client rest.Interface
}

func newAmazons(c *AuthV1Client) *amazons {
	return &amazons{
		client: c.RESTClient(),
	}
}

func (c *amazons) Get(name string, options metav1.GetOptions) (result *authv1.Amazon, err error) {
	result = &authv1.Amazon{}
	err = c.client.Get().
		Resource("amazons").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *amazons) List(opts metav1.ListOptions) (result *authv1.AmazonList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &authv1.AmazonList{}
	err = c.client.Get().
		Resource("amazons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *amazons) Create(collector *authv1.Amazon) (result *authv1.Amazon, err error) {
	result = &authv1.Amazon{}
	err = c.client.Post().
		Resource("amazons").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *amazons) Update(collector *authv1.Amazon) (result *authv1.Amazon, err error) {
	result = &authv1.Amazon{}
	err = c.client.Put().
		Resource("amazons").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *amazons) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("amazons").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *amazons) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.Amazon, err error) {
	result = &authv1.Amazon{}
	err = c.client.Patch(pt).
		Resource("amazons").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

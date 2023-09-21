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

type GoogleGetter interface {
	Googles() GoogleInterface
}

type GoogleInterface interface {
	Create(*authv1.Google) (*authv1.Google, error)
	Update(*authv1.Google) (*authv1.Google, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*authv1.Google, error)
	List(opts metav1.ListOptions) (*authv1.GoogleList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.Google, err error)
}

type googles struct {
	client rest.Interface
}

func newGoogles(c *AuthV1Client) *googles {
	return &googles{
		client: c.RESTClient(),
	}
}

func (c *googles) Get(name string, options metav1.GetOptions) (result *authv1.Google, err error) {
	result = &authv1.Google{}
	err = c.client.Get().
		Resource("googles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *googles) List(opts metav1.ListOptions) (result *authv1.GoogleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &authv1.GoogleList{}
	err = c.client.Get().
		Resource("googles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *googles) Create(collector *authv1.Google) (result *authv1.Google, err error) {
	result = &authv1.Google{}
	err = c.client.Post().
		Resource("googles").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *googles) Update(collector *authv1.Google) (result *authv1.Google, err error) {
	result = &authv1.Google{}
	err = c.client.Put().
		Resource("googles").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *googles) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googles").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *googles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.Google, err error) {
	result = &authv1.Google{}
	err = c.client.Patch(pt).
		Resource("googles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

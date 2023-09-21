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

type NHNGetter interface {
	NHNs() NHNInterface
}

type NHNInterface interface {
	Create(*authv1.NHN) (*authv1.NHN, error)
	Update(*authv1.NHN) (*authv1.NHN, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*authv1.NHN, error)
	List(opts metav1.ListOptions) (*authv1.NHNList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.NHN, err error)
}

type nhns struct {
	client rest.Interface
}

func newNHSs(c *AuthV1Client) *nhns {
	return &nhns{
		client: c.RESTClient(),
	}
}

func (c *nhns) Get(name string, options metav1.GetOptions) (result *authv1.NHN, err error) {
	result = &authv1.NHN{}
	err = c.client.Get().
		Resource("nhns").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *nhns) List(opts metav1.ListOptions) (result *authv1.NHNList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &authv1.NHNList{}
	err = c.client.Get().
		Resource("nhns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *nhns) Create(collector *authv1.NHN) (result *authv1.NHN, err error) {
	result = &authv1.NHN{}
	err = c.client.Post().
		Resource("nhns").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *nhns) Update(collector *authv1.NHN) (result *authv1.NHN, err error) {
	result = &authv1.NHN{}
	err = c.client.Put().
		Resource("nhns").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *nhns) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nhns").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *nhns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.NHN, err error) {
	result = &authv1.NHN{}
	err = c.client.Patch(pt).
		Resource("nhns").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

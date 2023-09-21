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

type AzureGetter interface {
	Azures() AzureInterface
}

type AzureInterface interface {
	Create(*authv1.Azure) (*authv1.Azure, error)
	Update(*authv1.Azure) (*authv1.Azure, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*authv1.Azure, error)
	List(opts metav1.ListOptions) (*authv1.AzureList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.Azure, err error)
}

type azures struct {
	client rest.Interface
}

func newAzures(c *AuthV1Client) *azures {
	return &azures{
		client: c.RESTClient(),
	}
}

func (c *azures) Get(name string, options metav1.GetOptions) (result *authv1.Azure, err error) {
	result = &authv1.Azure{}
	err = c.client.Get().
		Resource("azures").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *azures) List(opts metav1.ListOptions) (result *authv1.AzureList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &authv1.AzureList{}
	err = c.client.Get().
		Resource("azures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *azures) Create(collector *authv1.Azure) (result *authv1.Azure, err error) {
	result = &authv1.Azure{}
	err = c.client.Post().
		Resource("azures").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *azures) Update(collector *authv1.Azure) (result *authv1.Azure, err error) {
	result = &authv1.Azure{}
	err = c.client.Put().
		Resource("azures").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *azures) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azures").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *azures) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.Azure, err error) {
	result = &authv1.Azure{}
	err = c.client.Patch(pt).
		Resource("azures").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

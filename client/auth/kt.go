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

type KTGetter interface {
	KTs() KTInterface
}

type KTInterface interface {
	Create(*authv1.KT) (*authv1.KT, error)
	Update(*authv1.KT) (*authv1.KT, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*authv1.KT, error)
	List(opts metav1.ListOptions) (*authv1.KTList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.KT, err error)
}

type kts struct {
	client rest.Interface
}

func newKTs(c *AuthV1Client) *kts {
	return &kts{
		client: c.RESTClient(),
	}
}

func (c *kts) Get(name string, options metav1.GetOptions) (result *authv1.KT, err error) {
	result = &authv1.KT{}
	err = c.client.Get().
		Resource("kts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *kts) List(opts metav1.ListOptions) (result *authv1.KTList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &authv1.KTList{}
	err = c.client.Get().
		Resource("kts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *kts) Create(collector *authv1.KT) (result *authv1.KT, err error) {
	result = &authv1.KT{}
	err = c.client.Post().
		Resource("kts").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *kts) Update(collector *authv1.KT) (result *authv1.KT, err error) {
	result = &authv1.KT{}
	err = c.client.Put().
		Resource("kts").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *kts) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("kts").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *kts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.KT, err error) {
	result = &authv1.KT{}
	err = c.client.Patch(pt).
		Resource("kts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

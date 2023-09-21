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

type NaverGetter interface {
	Navers() NaverInterface
}

type NaverInterface interface {
	Create(*authv1.Naver) (*authv1.Naver, error)
	Update(*authv1.Naver) (*authv1.Naver, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*authv1.Naver, error)
	List(opts metav1.ListOptions) (*authv1.NaverList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.Naver, err error)
}

type navers struct {
	client rest.Interface
}

func newNavers(c *AuthV1Client) *navers {
	return &navers{
		client: c.RESTClient(),
	}
}

func (c *navers) Get(name string, options metav1.GetOptions) (result *authv1.Naver, err error) {
	result = &authv1.Naver{}
	err = c.client.Get().
		Resource("navers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *navers) List(opts metav1.ListOptions) (result *authv1.NaverList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &authv1.NaverList{}
	err = c.client.Get().
		Resource("navers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *navers) Create(collector *authv1.Naver) (result *authv1.Naver, err error) {
	result = &authv1.Naver{}
	err = c.client.Post().
		Resource("navers").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *navers) Update(collector *authv1.Naver) (result *authv1.Naver, err error) {
	result = &authv1.Naver{}
	err = c.client.Put().
		Resource("navers").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *navers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("navers").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *navers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.Naver, err error) {
	result = &authv1.Naver{}
	err = c.client.Patch(pt).
		Resource("navers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

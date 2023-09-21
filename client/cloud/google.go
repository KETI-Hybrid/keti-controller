package cloud

import (
	"context"
	"time"

	cloudv1 "github.com/KETI-Hybrid/keti-controller/apis/cloud/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type GoogleGetter interface {
	Googles() GoogleInterface
}

type GoogleInterface interface {
	Create(*cloudv1.Google) (*cloudv1.Google, error)
	Update(*cloudv1.Google) (*cloudv1.Google, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*cloudv1.Google, error)
	List(opts metav1.ListOptions) (*cloudv1.GoogleList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cloudv1.Google, err error)
}

type googles struct {
	client rest.Interface
}

func newGoogles(c *CloudV1Client) *googles {
	return &googles{
		client: c.RESTClient(),
	}
}

func (c *googles) Get(name string, options metav1.GetOptions) (result *cloudv1.Google, err error) {
	result = &cloudv1.Google{}
	err = c.client.Get().
		Resource("googles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *googles) List(opts metav1.ListOptions) (result *cloudv1.GoogleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &cloudv1.GoogleList{}
	err = c.client.Get().
		Resource("googles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *googles) Create(collector *cloudv1.Google) (result *cloudv1.Google, err error) {
	result = &cloudv1.Google{}
	err = c.client.Post().
		Resource("googles").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *googles) Update(collector *cloudv1.Google) (result *cloudv1.Google, err error) {
	result = &cloudv1.Google{}
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
func (c *googles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cloudv1.Google, err error) {
	result = &cloudv1.Google{}
	err = c.client.Patch(pt).
		Resource("googles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

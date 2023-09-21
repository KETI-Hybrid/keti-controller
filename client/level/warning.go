package level

import (
	"context"
	"time"

	levelv1 "github.com/KETI-Hybrid/keti-controller/apis/level/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type WarningGetter interface {
	Warnings() WarningInterface
}

type WarningInterface interface {
	Create(*levelv1.Warning) (*levelv1.Warning, error)
	Update(*levelv1.Warning) (*levelv1.Warning, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*levelv1.Warning, error)
	List(opts metav1.ListOptions) (*levelv1.WarningList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *levelv1.Warning, err error)
}

type warnings struct {
	client rest.Interface
}

func newWarnings(c *LevelV1Client) *warnings {
	return &warnings{
		client: c.RESTClient(),
	}
}

func (c *warnings) Get(name string, options metav1.GetOptions) (result *levelv1.Warning, err error) {
	result = &levelv1.Warning{}
	err = c.client.Get().
		Resource("warnings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *warnings) List(opts metav1.ListOptions) (result *levelv1.WarningList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &levelv1.WarningList{}
	err = c.client.Get().
		Resource("warnings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *warnings) Create(collector *levelv1.Warning) (result *levelv1.Warning, err error) {
	result = &levelv1.Warning{}
	err = c.client.Post().
		Resource("warnings").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *warnings) Update(collector *levelv1.Warning) (result *levelv1.Warning, err error) {
	result = &levelv1.Warning{}
	err = c.client.Put().
		Resource("warnings").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *warnings) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("warnings").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *warnings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *levelv1.Warning, err error) {
	result = &levelv1.Warning{}
	err = c.client.Patch(pt).
		Resource("warnings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

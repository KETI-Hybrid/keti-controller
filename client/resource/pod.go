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

type PodGetter interface {
	Pods() PodInterface
}

type PodInterface interface {
	Create(*resourcev1.Pod) (*resourcev1.Pod, error)
	Update(*resourcev1.Pod) (*resourcev1.Pod, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*resourcev1.Pod, error)
	List(opts metav1.ListOptions) (*resourcev1.PodList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.Pod, err error)
}

type pods struct {
	client rest.Interface
}

func newPods(c *ResourceV1Client) *pods {
	return &pods{
		client: c.RESTClient(),
	}
}

func (c *pods) Get(name string, options metav1.GetOptions) (result *resourcev1.Pod, err error) {
	result = &resourcev1.Pod{}
	err = c.client.Get().
		Resource("pods").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *pods) List(opts metav1.ListOptions) (result *resourcev1.PodList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &resourcev1.PodList{}
	err = c.client.Get().
		Resource("pods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *pods) Create(collector *resourcev1.Pod) (result *resourcev1.Pod, err error) {
	result = &resourcev1.Pod{}
	err = c.client.Post().
		Resource("pods").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *pods) Update(collector *resourcev1.Pod) (result *resourcev1.Pod, err error) {
	result = &resourcev1.Pod{}
	err = c.client.Put().
		Resource("pods").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *pods) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("pods").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *pods) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.Pod, err error) {
	result = &resourcev1.Pod{}
	err = c.client.Patch(pt).
		Resource("pods").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

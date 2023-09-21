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

type WatchingGetter interface {
	Watchings() WatchingInterface
}

type WatchingInterface interface {
	Create(*levelv1.Watching) (*levelv1.Watching, error)
	Update(*levelv1.Watching) (*levelv1.Watching, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*levelv1.Watching, error)
	List(opts metav1.ListOptions) (*levelv1.WatchingList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *levelv1.Watching, err error)
}

type watchings struct {
	client rest.Interface
}

func newWatchings(c *LevelV1Client) *watchings {
	return &watchings{
		client: c.RESTClient(),
	}
}

func (c *watchings) Get(name string, options metav1.GetOptions) (result *levelv1.Watching, err error) {
	result = &levelv1.Watching{}
	err = c.client.Get().
		Resource("watchings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *watchings) List(opts metav1.ListOptions) (result *levelv1.WatchingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &levelv1.WatchingList{}
	err = c.client.Get().
		Resource("watchings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *watchings) Create(collector *levelv1.Watching) (result *levelv1.Watching, err error) {
	result = &levelv1.Watching{}
	err = c.client.Post().
		Resource("watchings").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *watchings) Update(collector *levelv1.Watching) (result *levelv1.Watching, err error) {
	result = &levelv1.Watching{}
	err = c.client.Put().
		Resource("watchings").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *watchings) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("watchings").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *watchings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *levelv1.Watching, err error) {
	result = &levelv1.Watching{}
	err = c.client.Patch(pt).
		Resource("watchings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

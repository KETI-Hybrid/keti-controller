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

type DaemonsetGetter interface {
	Daemonsets() DaemonsetInterface
}

type DaemonsetInterface interface {
	Create(*resourcev1.Daemonset) (*resourcev1.Daemonset, error)
	Update(*resourcev1.Daemonset) (*resourcev1.Daemonset, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*resourcev1.Daemonset, error)
	List(opts metav1.ListOptions) (*resourcev1.DaemonsetList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.Daemonset, err error)
}

type daemonsets struct {
	client rest.Interface
}

func newDaemonsets(c *ResourceV1Client) *daemonsets {
	return &daemonsets{
		client: c.RESTClient(),
	}
}

func (c *daemonsets) Get(name string, options metav1.GetOptions) (result *resourcev1.Daemonset, err error) {
	result = &resourcev1.Daemonset{}
	err = c.client.Get().
		Resource("daemonsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *daemonsets) List(opts metav1.ListOptions) (result *resourcev1.DaemonsetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &resourcev1.DaemonsetList{}
	err = c.client.Get().
		Resource("daemonsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *daemonsets) Create(collector *resourcev1.Daemonset) (result *resourcev1.Daemonset, err error) {
	result = &resourcev1.Daemonset{}
	err = c.client.Post().
		Resource("daemonsets").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *daemonsets) Update(collector *resourcev1.Daemonset) (result *resourcev1.Daemonset, err error) {
	result = &resourcev1.Daemonset{}
	err = c.client.Put().
		Resource("daemonsets").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *daemonsets) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("daemonsets").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *daemonsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.Daemonset, err error) {
	result = &resourcev1.Daemonset{}
	err = c.client.Patch(pt).
		Resource("daemonsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

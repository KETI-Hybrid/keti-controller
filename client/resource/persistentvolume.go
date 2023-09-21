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

type PersistentVolumeGetter interface {
	PersistentVolumes() PersistentVolumeInterface
}

type PersistentVolumeInterface interface {
	Create(*resourcev1.PersistentVolume) (*resourcev1.PersistentVolume, error)
	Update(*resourcev1.PersistentVolume) (*resourcev1.PersistentVolume, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*resourcev1.PersistentVolume, error)
	List(opts metav1.ListOptions) (*resourcev1.PersistentVolumeList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.PersistentVolume, err error)
}

type persistentVolumes struct {
	client rest.Interface
}

func newPersistentVolumes(c *ResourceV1Client) *persistentVolumes {
	return &persistentVolumes{
		client: c.RESTClient(),
	}
}

func (c *persistentVolumes) Get(name string, options metav1.GetOptions) (result *resourcev1.PersistentVolume, err error) {
	result = &resourcev1.PersistentVolume{}
	err = c.client.Get().
		Resource("persistentVolumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *persistentVolumes) List(opts metav1.ListOptions) (result *resourcev1.PersistentVolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &resourcev1.PersistentVolumeList{}
	err = c.client.Get().
		Resource("persistentVolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *persistentVolumes) Create(collector *resourcev1.PersistentVolume) (result *resourcev1.PersistentVolume, err error) {
	result = &resourcev1.PersistentVolume{}
	err = c.client.Post().
		Resource("persistentVolumes").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *persistentVolumes) Update(collector *resourcev1.PersistentVolume) (result *resourcev1.PersistentVolume, err error) {
	result = &resourcev1.PersistentVolume{}
	err = c.client.Put().
		Resource("persistentVolumes").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *persistentVolumes) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("persistentVolumes").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *persistentVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.PersistentVolume, err error) {
	result = &resourcev1.PersistentVolume{}
	err = c.client.Patch(pt).
		Resource("persistentVolumes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

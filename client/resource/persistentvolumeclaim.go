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

type PersistentVolumeClaimGetter interface {
	PersistentVolumeClaims() PersistentVolumeClaimInterface
}

type PersistentVolumeClaimInterface interface {
	Create(*resourcev1.PersistentVolumeClaim) (*resourcev1.PersistentVolumeClaim, error)
	Update(*resourcev1.PersistentVolumeClaim) (*resourcev1.PersistentVolumeClaim, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*resourcev1.PersistentVolumeClaim, error)
	List(opts metav1.ListOptions) (*resourcev1.PersistentVolumeClaimList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.PersistentVolumeClaim, err error)
}

type persistentVolumeClaims struct {
	client rest.Interface
}

func newPersistentVolumeClaims(c *ResourceV1Client) *persistentVolumeClaims {
	return &persistentVolumeClaims{
		client: c.RESTClient(),
	}
}

func (c *persistentVolumeClaims) Get(name string, options metav1.GetOptions) (result *resourcev1.PersistentVolumeClaim, err error) {
	result = &resourcev1.PersistentVolumeClaim{}
	err = c.client.Get().
		Resource("persistentVolumeClaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *persistentVolumeClaims) List(opts metav1.ListOptions) (result *resourcev1.PersistentVolumeClaimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &resourcev1.PersistentVolumeClaimList{}
	err = c.client.Get().
		Resource("persistentVolumeClaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *persistentVolumeClaims) Create(collector *resourcev1.PersistentVolumeClaim) (result *resourcev1.PersistentVolumeClaim, err error) {
	result = &resourcev1.PersistentVolumeClaim{}
	err = c.client.Post().
		Resource("persistentVolumeClaims").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *persistentVolumeClaims) Update(collector *resourcev1.PersistentVolumeClaim) (result *resourcev1.PersistentVolumeClaim, err error) {
	result = &resourcev1.PersistentVolumeClaim{}
	err = c.client.Put().
		Resource("persistentVolumeClaims").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *persistentVolumeClaims) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("persistentVolumeClaims").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *persistentVolumeClaims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.PersistentVolumeClaim, err error) {
	result = &resourcev1.PersistentVolumeClaim{}
	err = c.client.Patch(pt).
		Resource("persistentVolumeClaims").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

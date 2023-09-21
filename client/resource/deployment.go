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

type DeploymentGetter interface {
	Deployments() DeploymentInterface
}

type DeploymentInterface interface {
	Create(*resourcev1.Deployment) (*resourcev1.Deployment, error)
	Update(*resourcev1.Deployment) (*resourcev1.Deployment, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*resourcev1.Deployment, error)
	List(opts metav1.ListOptions) (*resourcev1.DeploymentList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.Deployment, err error)
}

type deployments struct {
	client rest.Interface
}

func newDeployments(c *ResourceV1Client) *deployments {
	return &deployments{
		client: c.RESTClient(),
	}
}

func (c *deployments) Get(name string, options metav1.GetOptions) (result *resourcev1.Deployment, err error) {
	result = &resourcev1.Deployment{}
	err = c.client.Get().
		Resource("deployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *deployments) List(opts metav1.ListOptions) (result *resourcev1.DeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &resourcev1.DeploymentList{}
	err = c.client.Get().
		Resource("deployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *deployments) Create(collector *resourcev1.Deployment) (result *resourcev1.Deployment, err error) {
	result = &resourcev1.Deployment{}
	err = c.client.Post().
		Resource("deployments").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *deployments) Update(collector *resourcev1.Deployment) (result *resourcev1.Deployment, err error) {
	result = &resourcev1.Deployment{}
	err = c.client.Put().
		Resource("deployments").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *deployments) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("deployments").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *deployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *resourcev1.Deployment, err error) {
	result = &resourcev1.Deployment{}
	err = c.client.Patch(pt).
		Resource("deployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}

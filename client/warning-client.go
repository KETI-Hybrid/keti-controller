package client

import (
	"context"

	v1 "github.com/KETI-Hybrid/keti-controller/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type WarningInterface interface {
	List(opts metav1.ListOptions) (*v1.WarningList, error)
	Get(name string, options metav1.GetOptions) (*v1.Warning, error)
	Create(*v1.Warning) (*v1.Warning, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

type warningClient struct {
	restClient rest.Interface
}

func (c *warningClient) List(opts metav1.ListOptions) (*v1.WarningList, error) {
	result := v1.WarningList{}
	err := c.restClient.
		Get().
		Resource("warnings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *warningClient) Get(name string, opts metav1.GetOptions) (*v1.Warning, error) {
	result := v1.Warning{}
	err := c.restClient.
		Get().
		Resource("warnings").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *warningClient) Create(warning *v1.Warning) (*v1.Warning, error) {
	result := v1.Warning{}
	err := c.restClient.
		Post().
		Resource("warnings").
		Body(warning).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *warningClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Resource("warnings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}

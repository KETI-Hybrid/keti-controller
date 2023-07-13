package client

import (
	"context"

	v1 "github.com/KETI-Hybrid/keti-controller/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type WatchingInterface interface {
	List(opts metav1.ListOptions) (*v1.WatchingList, error)
	Get(name string, options metav1.GetOptions) (*v1.Watching, error)
	Create(*v1.Watching) (*v1.Watching, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

type watchingClient struct {
	restClient rest.Interface
}

func (c *watchingClient) List(opts metav1.ListOptions) (*v1.WatchingList, error) {
	result := v1.WatchingList{}
	err := c.restClient.
		Get().
		Resource("watchings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *watchingClient) Get(name string, opts metav1.GetOptions) (*v1.Watching, error) {
	result := v1.Watching{}
	err := c.restClient.
		Get().
		Resource("watchings").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *watchingClient) Create(watching *v1.Watching) (*v1.Watching, error) {
	result := v1.Watching{}
	err := c.restClient.
		Post().
		Resource("warnings").
		Body(watching).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *watchingClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Resource("watchings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}

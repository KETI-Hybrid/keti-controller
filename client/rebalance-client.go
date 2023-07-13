package client

import (
	"context"

	v1 "github.com/KETI-Hybrid/keti-controller/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type RebalanceInterface interface {
	List(opts metav1.ListOptions) (*v1.RebalanceList, error)
	Get(name string, options metav1.GetOptions) (*v1.Rebalance, error)
	Create(*v1.Rebalance) (*v1.Rebalance, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

type rebalanceClient struct {
	restClient rest.Interface
}

func (c *rebalanceClient) List(opts metav1.ListOptions) (*v1.RebalanceList, error) {
	result := v1.RebalanceList{}
	err := c.restClient.
		Get().
		Resource("rebalances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *rebalanceClient) Get(name string, opts metav1.GetOptions) (*v1.Rebalance, error) {
	result := v1.Rebalance{}
	err := c.restClient.
		Get().
		Resource("rebalances").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *rebalanceClient) Create(rebalance *v1.Rebalance) (*v1.Rebalance, error) {
	result := v1.Rebalance{}
	err := c.restClient.
		Post().
		Resource("rebalances").
		Body(rebalance).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *rebalanceClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Resource("rebalances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}

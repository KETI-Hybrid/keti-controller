package client

import (
	v1 "github.com/KETI-Hybrid/keti-controller/api/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type KetiV1Interface interface {
	Rebalance() RebalanceInterface
	Watching() WatchingInterface
	Warning() WarningInterface
}

type KetiV1Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*KetiV1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &v1.GroupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &KetiV1Client{restClient: client}, nil
}

func (c *KetiV1Client) Rebalance() RebalanceInterface {
	return &rebalanceClient{
		restClient: c.restClient,
	}
}

func (c *KetiV1Client) Watching() WatchingInterface {
	return &watchingClient{
		restClient: c.restClient,
	}
}

func (c *KetiV1Client) Warning() WarningInterface {
	return &warningClient{
		restClient: c.restClient,
	}
}

package level

import (
	levelv1 "github.com/KETI-Hybrid/keti-controller/apis/level/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type LevelV1Interface interface {
	RESTClient() rest.Interface
	RebalanceGetter
	WarningGetter
	WatchingGetter
}

type LevelV1Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*LevelV1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: levelv1.GroupVersion.Group, Version: levelv1.GroupVersion.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &LevelV1Client{restClient: client}, nil
}

func (c *LevelV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *LevelV1Client) Rebalances() RebalanceInterface {
	return newRebalances(c)
}

func (c *LevelV1Client) Warnings() WarningInterface {
	return newWarnings(c)
}

func (c *LevelV1Client) Watchings() WatchingInterface {
	return newWatchings(c)
}

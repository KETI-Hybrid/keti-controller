package cloud

import (
	cloudv1 "github.com/KETI-Hybrid/keti-controller/apis/cloud/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type CloudV1Interface interface {
	RESTClient() rest.Interface
	AmazonGetter
	AzureGetter
	GoogleGetter
	KTGetter
	NaverGetter
	NHNGetter
}

type CloudV1Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*CloudV1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: cloudv1.GroupVersion.Group, Version: cloudv1.GroupVersion.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &CloudV1Client{restClient: client}, nil
}

func (c *CloudV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *CloudV1Client) Amazons() AmazonInterface {
	return newAmazons(c)
}

func (c *CloudV1Client) Azures() AzureInterface {
	return newAzures(c)
}

func (c *CloudV1Client) Googles() GoogleInterface {
	return newGoogles(c)
}

func (c *CloudV1Client) KTs() KTInterface {
	return newKTs(c)
}

func (c *CloudV1Client) Navers() NaverInterface {
	return newNavers(c)
}

func (c *CloudV1Client) NHNs() NHNInterface {
	return newNHSs(c)
}

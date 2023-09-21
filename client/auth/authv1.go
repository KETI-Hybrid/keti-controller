package auth

import (
	authv1 "github.com/KETI-Hybrid/keti-controller/apis/auth/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type AuthV1Interface interface {
	RESTClient() rest.Interface
	AmazonGetter
	AzureGetter
	GoogleGetter
	KTGetter
	NaverGetter
	NHNGetter
}

type AuthV1Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*AuthV1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: authv1.GroupVersion.Group, Version: authv1.GroupVersion.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &AuthV1Client{restClient: client}, nil
}

func (c *AuthV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *AuthV1Client) Amazons() AmazonInterface {
	return newAmazons(c)
}

func (c *AuthV1Client) Azures() AzureInterface {
	return newAzures(c)
}

func (c *AuthV1Client) Googles() GoogleInterface {
	return newGoogles(c)
}

func (c *AuthV1Client) KTs() KTInterface {
	return newKTs(c)
}

func (c *AuthV1Client) Navers() NaverInterface {
	return newNavers(c)
}

func (c *AuthV1Client) NHNs() NHNInterface {
	return newNHSs(c)
}

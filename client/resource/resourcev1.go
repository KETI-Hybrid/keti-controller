package resource

import (
	resourcev1 "github.com/KETI-Hybrid/keti-controller/apis/resource/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type ResourceV1Interface interface {
	RESTClient() rest.Interface
	DaemonsetGetter
	DeploymentGetter
	IngressGetter
	PersistentVolumeClaimGetter
	PersistentVolumeGetter
	PodGetter
	ServiceGetter
	SpecificResourceGetter
	StatefulsetGetter
}

type ResourceV1Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*ResourceV1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: resourcev1.GroupVersion.Group, Version: resourcev1.GroupVersion.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &ResourceV1Client{restClient: client}, nil
}

func (c *ResourceV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *ResourceV1Client) Daemonsets() DaemonsetInterface {
	return newDaemonsets(c)
}

func (c *ResourceV1Client) Deployments() DeploymentInterface {
	return newDeployments(c)
}

func (c *ResourceV1Client) Ingresses() IngressInterface {
	return newIngresses(c)
}

func (c *ResourceV1Client) PersistentVolumeClaims() PersistentVolumeClaimInterface {
	return newPersistentVolumeClaims(c)
}

func (c *ResourceV1Client) PersistentVolumes() PersistentVolumeInterface {
	return newPersistentVolumes(c)
}

func (c *ResourceV1Client) Pods() PodInterface {
	return newPods(c)
}

func (c *ResourceV1Client) Services() ServiceInterface {
	return newServices(c)
}

func (c *ResourceV1Client) SpecificResources() SpecificResourceInterface {
	return newSpecificResources(c)
}

func (c *ResourceV1Client) Statefulsets() StatefulsetInterface {
	return newStatefulsets(c)
}

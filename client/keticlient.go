package client

import (
	"fmt"

	"github.com/KETI-Hybrid/keti-controller/client/auth"
	"github.com/KETI-Hybrid/keti-controller/client/cloud"
	"github.com/KETI-Hybrid/keti-controller/client/level"
	"github.com/KETI-Hybrid/keti-controller/client/resource"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	AuthV1() auth.AuthV1Interface
	CloudV1() cloud.CloudV1Interface
	LevelV1() level.LevelV1Interface
	ResourceV1() resource.ResourceV1Interface
}

type ClientSet struct {
	authV1     *auth.AuthV1Client
	cloudV1    *cloud.CloudV1Client
	levelV1    *level.LevelV1Client
	resourceV1 *resource.ResourceV1Client
}

func NewForConfig(c *rest.Config) (*ClientSet, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when ratelimiter is not set and qps is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs ClientSet
	var err error
	cs.authV1, err = auth.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cloudV1, err = cloud.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.levelV1, err = level.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.resourceV1, err = resource.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	
	return &cs, nil
}

func (c *ClientSet) AuthV1() auth.AuthV1Interface {
	return c.authV1
}

func (c *ClientSet) CloudV1() cloud.CloudV1Interface {
	return c.cloudV1
}

func (c *ClientSet) LevelV1() level.LevelV1Interface {
	return c.levelV1
}

func (c *ClientSet) ResourceV1() resource.ResourceV1Interface {
	return c.resourceV1
}

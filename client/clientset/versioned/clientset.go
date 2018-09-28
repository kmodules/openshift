// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	appsv1 "github.com/pharmer/openshift/client/clientset/versioned/typed/apps/v1"
	securityv1 "github.com/pharmer/openshift/client/clientset/versioned/typed/security/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AppsV1() appsv1.AppsV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Apps() appsv1.AppsV1Interface
	SecurityV1() securityv1.SecurityV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Security() securityv1.SecurityV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	appsV1     *appsv1.AppsV1Client
	securityV1 *securityv1.SecurityV1Client
}

// AppsV1 retrieves the AppsV1Client
func (c *Clientset) AppsV1() appsv1.AppsV1Interface {
	return c.appsV1
}

// Deprecated: Apps retrieves the default version of AppsClient.
// Please explicitly pick a version.
func (c *Clientset) Apps() appsv1.AppsV1Interface {
	return c.appsV1
}

// SecurityV1 retrieves the SecurityV1Client
func (c *Clientset) SecurityV1() securityv1.SecurityV1Interface {
	return c.securityV1
}

// Deprecated: Security retrieves the default version of SecurityClient.
// Please explicitly pick a version.
func (c *Clientset) Security() securityv1.SecurityV1Interface {
	return c.securityV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.appsV1, err = appsv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.securityV1, err = securityv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.appsV1 = appsv1.NewForConfigOrDie(c)
	cs.securityV1 = securityv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.appsV1 = appsv1.New(c)
	cs.securityV1 = securityv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}

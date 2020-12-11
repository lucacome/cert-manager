/*
Copyright 2020 The cert-manager Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	acmev1 "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/typed/acme/v1"
	acmev1alpha2 "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/typed/acme/v1alpha2"
	acmev1alpha3 "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/typed/acme/v1alpha3"
	acmev1beta1 "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/typed/acme/v1beta1"
	certmanagerv1 "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/typed/certmanager/v1"
	certmanagerv1alpha2 "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/typed/certmanager/v1alpha2"
	certmanagerv1alpha3 "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/typed/certmanager/v1alpha3"
	certmanagerv1beta1 "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/typed/certmanager/v1beta1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AcmeV1alpha2() acmev1alpha2.AcmeV1alpha2Interface
	AcmeV1alpha3() acmev1alpha3.AcmeV1alpha3Interface
	AcmeV1beta1() acmev1beta1.AcmeV1beta1Interface
	AcmeV1() acmev1.AcmeV1Interface
	CertmanagerV1alpha2() certmanagerv1alpha2.CertmanagerV1alpha2Interface
	CertmanagerV1alpha3() certmanagerv1alpha3.CertmanagerV1alpha3Interface
	CertmanagerV1beta1() certmanagerv1beta1.CertmanagerV1beta1Interface
	CertmanagerV1() certmanagerv1.CertmanagerV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	acmeV1alpha2        *acmev1alpha2.AcmeV1alpha2Client
	acmeV1alpha3        *acmev1alpha3.AcmeV1alpha3Client
	acmeV1beta1         *acmev1beta1.AcmeV1beta1Client
	acmeV1              *acmev1.AcmeV1Client
	certmanagerV1alpha2 *certmanagerv1alpha2.CertmanagerV1alpha2Client
	certmanagerV1alpha3 *certmanagerv1alpha3.CertmanagerV1alpha3Client
	certmanagerV1beta1  *certmanagerv1beta1.CertmanagerV1beta1Client
	certmanagerV1       *certmanagerv1.CertmanagerV1Client
}

// AcmeV1alpha2 retrieves the AcmeV1alpha2Client
func (c *Clientset) AcmeV1alpha2() acmev1alpha2.AcmeV1alpha2Interface {
	return c.acmeV1alpha2
}

// AcmeV1alpha3 retrieves the AcmeV1alpha3Client
func (c *Clientset) AcmeV1alpha3() acmev1alpha3.AcmeV1alpha3Interface {
	return c.acmeV1alpha3
}

// AcmeV1beta1 retrieves the AcmeV1beta1Client
func (c *Clientset) AcmeV1beta1() acmev1beta1.AcmeV1beta1Interface {
	return c.acmeV1beta1
}

// AcmeV1 retrieves the AcmeV1Client
func (c *Clientset) AcmeV1() acmev1.AcmeV1Interface {
	return c.acmeV1
}

// CertmanagerV1alpha2 retrieves the CertmanagerV1alpha2Client
func (c *Clientset) CertmanagerV1alpha2() certmanagerv1alpha2.CertmanagerV1alpha2Interface {
	return c.certmanagerV1alpha2
}

// CertmanagerV1alpha3 retrieves the CertmanagerV1alpha3Client
func (c *Clientset) CertmanagerV1alpha3() certmanagerv1alpha3.CertmanagerV1alpha3Interface {
	return c.certmanagerV1alpha3
}

// CertmanagerV1beta1 retrieves the CertmanagerV1beta1Client
func (c *Clientset) CertmanagerV1beta1() certmanagerv1beta1.CertmanagerV1beta1Interface {
	return c.certmanagerV1beta1
}

// CertmanagerV1 retrieves the CertmanagerV1Client
func (c *Clientset) CertmanagerV1() certmanagerv1.CertmanagerV1Interface {
	return c.certmanagerV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.acmeV1alpha2, err = acmev1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.acmeV1alpha3, err = acmev1alpha3.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.acmeV1beta1, err = acmev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.acmeV1, err = acmev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.certmanagerV1alpha2, err = certmanagerv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.certmanagerV1alpha3, err = certmanagerv1alpha3.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.certmanagerV1beta1, err = certmanagerv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.certmanagerV1, err = certmanagerv1.NewForConfig(&configShallowCopy)
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
	cs.acmeV1alpha2 = acmev1alpha2.NewForConfigOrDie(c)
	cs.acmeV1alpha3 = acmev1alpha3.NewForConfigOrDie(c)
	cs.acmeV1beta1 = acmev1beta1.NewForConfigOrDie(c)
	cs.acmeV1 = acmev1.NewForConfigOrDie(c)
	cs.certmanagerV1alpha2 = certmanagerv1alpha2.NewForConfigOrDie(c)
	cs.certmanagerV1alpha3 = certmanagerv1alpha3.NewForConfigOrDie(c)
	cs.certmanagerV1beta1 = certmanagerv1beta1.NewForConfigOrDie(c)
	cs.certmanagerV1 = certmanagerv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.acmeV1alpha2 = acmev1alpha2.New(c)
	cs.acmeV1alpha3 = acmev1alpha3.New(c)
	cs.acmeV1beta1 = acmev1beta1.New(c)
	cs.acmeV1 = acmev1.New(c)
	cs.certmanagerV1alpha2 = certmanagerv1alpha2.New(c)
	cs.certmanagerV1alpha3 = certmanagerv1alpha3.New(c)
	cs.certmanagerV1beta1 = certmanagerv1beta1.New(c)
	cs.certmanagerV1 = certmanagerv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}

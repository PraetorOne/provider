/*
Copyright The Akash Network Authors.

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

package v2beta2

import (
	"context"
	"time"

	v2beta2 "github.com/akash-network/provider/pkg/apis/akash.network/v2beta2"
	scheme "github.com/akash-network/provider/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LeaseParamsServicesGetter has a method to return a LeaseParamsServiceInterface.
// A group's client should implement this interface.
type LeaseParamsServicesGetter interface {
	LeaseParamsServices(namespace string) LeaseParamsServiceInterface
}

// LeaseParamsServiceInterface has methods to work with LeaseParamsService resources.
type LeaseParamsServiceInterface interface {
	Create(ctx context.Context, leaseParamsService *v2beta2.LeaseParamsService, opts v1.CreateOptions) (*v2beta2.LeaseParamsService, error)
	Update(ctx context.Context, leaseParamsService *v2beta2.LeaseParamsService, opts v1.UpdateOptions) (*v2beta2.LeaseParamsService, error)
	UpdateStatus(ctx context.Context, leaseParamsService *v2beta2.LeaseParamsService, opts v1.UpdateOptions) (*v2beta2.LeaseParamsService, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2beta2.LeaseParamsService, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2beta2.LeaseParamsServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta2.LeaseParamsService, err error)
	LeaseParamsServiceExpansion
}

// leaseParamsServices implements LeaseParamsServiceInterface
type leaseParamsServices struct {
	client rest.Interface
	ns     string
}

// newLeaseParamsServices returns a LeaseParamsServices
func newLeaseParamsServices(c *AkashV2beta2Client, namespace string) *leaseParamsServices {
	return &leaseParamsServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the leaseParamsService, and returns the corresponding leaseParamsService object, and an error if there is any.
func (c *leaseParamsServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta2.LeaseParamsService, err error) {
	result = &v2beta2.LeaseParamsService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("leaseparamsservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LeaseParamsServices that match those selectors.
func (c *leaseParamsServices) List(ctx context.Context, opts v1.ListOptions) (result *v2beta2.LeaseParamsServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2beta2.LeaseParamsServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("leaseparamsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested leaseParamsServices.
func (c *leaseParamsServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("leaseparamsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a leaseParamsService and creates it.  Returns the server's representation of the leaseParamsService, and an error, if there is any.
func (c *leaseParamsServices) Create(ctx context.Context, leaseParamsService *v2beta2.LeaseParamsService, opts v1.CreateOptions) (result *v2beta2.LeaseParamsService, err error) {
	result = &v2beta2.LeaseParamsService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("leaseparamsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(leaseParamsService).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a leaseParamsService and updates it. Returns the server's representation of the leaseParamsService, and an error, if there is any.
func (c *leaseParamsServices) Update(ctx context.Context, leaseParamsService *v2beta2.LeaseParamsService, opts v1.UpdateOptions) (result *v2beta2.LeaseParamsService, err error) {
	result = &v2beta2.LeaseParamsService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("leaseparamsservices").
		Name(leaseParamsService.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(leaseParamsService).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *leaseParamsServices) UpdateStatus(ctx context.Context, leaseParamsService *v2beta2.LeaseParamsService, opts v1.UpdateOptions) (result *v2beta2.LeaseParamsService, err error) {
	result = &v2beta2.LeaseParamsService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("leaseparamsservices").
		Name(leaseParamsService.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(leaseParamsService).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the leaseParamsService and deletes it. Returns an error if one occurs.
func (c *leaseParamsServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("leaseparamsservices").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *leaseParamsServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("leaseparamsservices").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched leaseParamsService.
func (c *leaseParamsServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta2.LeaseParamsService, err error) {
	result = &v2beta2.LeaseParamsService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("leaseparamsservices").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

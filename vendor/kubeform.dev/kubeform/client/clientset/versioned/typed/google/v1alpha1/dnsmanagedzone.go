/*
Copyright The Kubeform Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DnsManagedZonesGetter has a method to return a DnsManagedZoneInterface.
// A group's client should implement this interface.
type DnsManagedZonesGetter interface {
	DnsManagedZones(namespace string) DnsManagedZoneInterface
}

// DnsManagedZoneInterface has methods to work with DnsManagedZone resources.
type DnsManagedZoneInterface interface {
	Create(ctx context.Context, dnsManagedZone *v1alpha1.DnsManagedZone, opts v1.CreateOptions) (*v1alpha1.DnsManagedZone, error)
	Update(ctx context.Context, dnsManagedZone *v1alpha1.DnsManagedZone, opts v1.UpdateOptions) (*v1alpha1.DnsManagedZone, error)
	UpdateStatus(ctx context.Context, dnsManagedZone *v1alpha1.DnsManagedZone, opts v1.UpdateOptions) (*v1alpha1.DnsManagedZone, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DnsManagedZone, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DnsManagedZoneList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DnsManagedZone, err error)
	DnsManagedZoneExpansion
}

// dnsManagedZones implements DnsManagedZoneInterface
type dnsManagedZones struct {
	client rest.Interface
	ns     string
}

// newDnsManagedZones returns a DnsManagedZones
func newDnsManagedZones(c *GoogleV1alpha1Client, namespace string) *dnsManagedZones {
	return &dnsManagedZones{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dnsManagedZone, and returns the corresponding dnsManagedZone object, and an error if there is any.
func (c *dnsManagedZones) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DnsManagedZone, err error) {
	result = &v1alpha1.DnsManagedZone{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsmanagedzones").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DnsManagedZones that match those selectors.
func (c *dnsManagedZones) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DnsManagedZoneList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DnsManagedZoneList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsmanagedzones").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dnsManagedZones.
func (c *dnsManagedZones) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dnsmanagedzones").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dnsManagedZone and creates it.  Returns the server's representation of the dnsManagedZone, and an error, if there is any.
func (c *dnsManagedZones) Create(ctx context.Context, dnsManagedZone *v1alpha1.DnsManagedZone, opts v1.CreateOptions) (result *v1alpha1.DnsManagedZone, err error) {
	result = &v1alpha1.DnsManagedZone{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dnsmanagedzones").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dnsManagedZone).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dnsManagedZone and updates it. Returns the server's representation of the dnsManagedZone, and an error, if there is any.
func (c *dnsManagedZones) Update(ctx context.Context, dnsManagedZone *v1alpha1.DnsManagedZone, opts v1.UpdateOptions) (result *v1alpha1.DnsManagedZone, err error) {
	result = &v1alpha1.DnsManagedZone{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnsmanagedzones").
		Name(dnsManagedZone.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dnsManagedZone).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dnsManagedZones) UpdateStatus(ctx context.Context, dnsManagedZone *v1alpha1.DnsManagedZone, opts v1.UpdateOptions) (result *v1alpha1.DnsManagedZone, err error) {
	result = &v1alpha1.DnsManagedZone{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnsmanagedzones").
		Name(dnsManagedZone.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dnsManagedZone).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dnsManagedZone and deletes it. Returns an error if one occurs.
func (c *dnsManagedZones) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsmanagedzones").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dnsManagedZones) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsmanagedzones").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dnsManagedZone.
func (c *dnsManagedZones) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DnsManagedZone, err error) {
	result = &v1alpha1.DnsManagedZone{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dnsmanagedzones").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

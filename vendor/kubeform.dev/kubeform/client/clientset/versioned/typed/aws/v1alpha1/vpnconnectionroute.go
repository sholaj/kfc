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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VpnConnectionRoutesGetter has a method to return a VpnConnectionRouteInterface.
// A group's client should implement this interface.
type VpnConnectionRoutesGetter interface {
	VpnConnectionRoutes(namespace string) VpnConnectionRouteInterface
}

// VpnConnectionRouteInterface has methods to work with VpnConnectionRoute resources.
type VpnConnectionRouteInterface interface {
	Create(ctx context.Context, vpnConnectionRoute *v1alpha1.VpnConnectionRoute, opts v1.CreateOptions) (*v1alpha1.VpnConnectionRoute, error)
	Update(ctx context.Context, vpnConnectionRoute *v1alpha1.VpnConnectionRoute, opts v1.UpdateOptions) (*v1alpha1.VpnConnectionRoute, error)
	UpdateStatus(ctx context.Context, vpnConnectionRoute *v1alpha1.VpnConnectionRoute, opts v1.UpdateOptions) (*v1alpha1.VpnConnectionRoute, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VpnConnectionRoute, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VpnConnectionRouteList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VpnConnectionRoute, err error)
	VpnConnectionRouteExpansion
}

// vpnConnectionRoutes implements VpnConnectionRouteInterface
type vpnConnectionRoutes struct {
	client rest.Interface
	ns     string
}

// newVpnConnectionRoutes returns a VpnConnectionRoutes
func newVpnConnectionRoutes(c *AwsV1alpha1Client, namespace string) *vpnConnectionRoutes {
	return &vpnConnectionRoutes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vpnConnectionRoute, and returns the corresponding vpnConnectionRoute object, and an error if there is any.
func (c *vpnConnectionRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VpnConnectionRoute, err error) {
	result = &v1alpha1.VpnConnectionRoute{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpnconnectionroutes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VpnConnectionRoutes that match those selectors.
func (c *vpnConnectionRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VpnConnectionRouteList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VpnConnectionRouteList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpnconnectionroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vpnConnectionRoutes.
func (c *vpnConnectionRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vpnconnectionroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vpnConnectionRoute and creates it.  Returns the server's representation of the vpnConnectionRoute, and an error, if there is any.
func (c *vpnConnectionRoutes) Create(ctx context.Context, vpnConnectionRoute *v1alpha1.VpnConnectionRoute, opts v1.CreateOptions) (result *v1alpha1.VpnConnectionRoute, err error) {
	result = &v1alpha1.VpnConnectionRoute{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vpnconnectionroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpnConnectionRoute).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vpnConnectionRoute and updates it. Returns the server's representation of the vpnConnectionRoute, and an error, if there is any.
func (c *vpnConnectionRoutes) Update(ctx context.Context, vpnConnectionRoute *v1alpha1.VpnConnectionRoute, opts v1.UpdateOptions) (result *v1alpha1.VpnConnectionRoute, err error) {
	result = &v1alpha1.VpnConnectionRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpnconnectionroutes").
		Name(vpnConnectionRoute.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpnConnectionRoute).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vpnConnectionRoutes) UpdateStatus(ctx context.Context, vpnConnectionRoute *v1alpha1.VpnConnectionRoute, opts v1.UpdateOptions) (result *v1alpha1.VpnConnectionRoute, err error) {
	result = &v1alpha1.VpnConnectionRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpnconnectionroutes").
		Name(vpnConnectionRoute.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpnConnectionRoute).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vpnConnectionRoute and deletes it. Returns an error if one occurs.
func (c *vpnConnectionRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpnconnectionroutes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vpnConnectionRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpnconnectionroutes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vpnConnectionRoute.
func (c *vpnConnectionRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VpnConnectionRoute, err error) {
	result = &v1alpha1.VpnConnectionRoute{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vpnconnectionroutes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

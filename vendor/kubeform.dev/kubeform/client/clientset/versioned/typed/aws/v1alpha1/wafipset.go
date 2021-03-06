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

// WafIpsetsGetter has a method to return a WafIpsetInterface.
// A group's client should implement this interface.
type WafIpsetsGetter interface {
	WafIpsets(namespace string) WafIpsetInterface
}

// WafIpsetInterface has methods to work with WafIpset resources.
type WafIpsetInterface interface {
	Create(ctx context.Context, wafIpset *v1alpha1.WafIpset, opts v1.CreateOptions) (*v1alpha1.WafIpset, error)
	Update(ctx context.Context, wafIpset *v1alpha1.WafIpset, opts v1.UpdateOptions) (*v1alpha1.WafIpset, error)
	UpdateStatus(ctx context.Context, wafIpset *v1alpha1.WafIpset, opts v1.UpdateOptions) (*v1alpha1.WafIpset, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.WafIpset, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WafIpsetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WafIpset, err error)
	WafIpsetExpansion
}

// wafIpsets implements WafIpsetInterface
type wafIpsets struct {
	client rest.Interface
	ns     string
}

// newWafIpsets returns a WafIpsets
func newWafIpsets(c *AwsV1alpha1Client, namespace string) *wafIpsets {
	return &wafIpsets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the wafIpset, and returns the corresponding wafIpset object, and an error if there is any.
func (c *wafIpsets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WafIpset, err error) {
	result = &v1alpha1.WafIpset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafipsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WafIpsets that match those selectors.
func (c *wafIpsets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WafIpsetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WafIpsetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested wafIpsets.
func (c *wafIpsets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("wafipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a wafIpset and creates it.  Returns the server's representation of the wafIpset, and an error, if there is any.
func (c *wafIpsets) Create(ctx context.Context, wafIpset *v1alpha1.WafIpset, opts v1.CreateOptions) (result *v1alpha1.WafIpset, err error) {
	result = &v1alpha1.WafIpset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("wafipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(wafIpset).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a wafIpset and updates it. Returns the server's representation of the wafIpset, and an error, if there is any.
func (c *wafIpsets) Update(ctx context.Context, wafIpset *v1alpha1.WafIpset, opts v1.UpdateOptions) (result *v1alpha1.WafIpset, err error) {
	result = &v1alpha1.WafIpset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafipsets").
		Name(wafIpset.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(wafIpset).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *wafIpsets) UpdateStatus(ctx context.Context, wafIpset *v1alpha1.WafIpset, opts v1.UpdateOptions) (result *v1alpha1.WafIpset, err error) {
	result = &v1alpha1.WafIpset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafipsets").
		Name(wafIpset.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(wafIpset).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the wafIpset and deletes it. Returns an error if one occurs.
func (c *wafIpsets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafipsets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *wafIpsets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafipsets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched wafIpset.
func (c *wafIpsets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WafIpset, err error) {
	result = &v1alpha1.WafIpset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("wafipsets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

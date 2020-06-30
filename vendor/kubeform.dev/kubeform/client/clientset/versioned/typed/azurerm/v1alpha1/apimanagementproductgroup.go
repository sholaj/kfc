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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApiManagementProductGroupsGetter has a method to return a ApiManagementProductGroupInterface.
// A group's client should implement this interface.
type ApiManagementProductGroupsGetter interface {
	ApiManagementProductGroups(namespace string) ApiManagementProductGroupInterface
}

// ApiManagementProductGroupInterface has methods to work with ApiManagementProductGroup resources.
type ApiManagementProductGroupInterface interface {
	Create(ctx context.Context, apiManagementProductGroup *v1alpha1.ApiManagementProductGroup, opts v1.CreateOptions) (*v1alpha1.ApiManagementProductGroup, error)
	Update(ctx context.Context, apiManagementProductGroup *v1alpha1.ApiManagementProductGroup, opts v1.UpdateOptions) (*v1alpha1.ApiManagementProductGroup, error)
	UpdateStatus(ctx context.Context, apiManagementProductGroup *v1alpha1.ApiManagementProductGroup, opts v1.UpdateOptions) (*v1alpha1.ApiManagementProductGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ApiManagementProductGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ApiManagementProductGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiManagementProductGroup, err error)
	ApiManagementProductGroupExpansion
}

// apiManagementProductGroups implements ApiManagementProductGroupInterface
type apiManagementProductGroups struct {
	client rest.Interface
	ns     string
}

// newApiManagementProductGroups returns a ApiManagementProductGroups
func newApiManagementProductGroups(c *AzurermV1alpha1Client, namespace string) *apiManagementProductGroups {
	return &apiManagementProductGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiManagementProductGroup, and returns the corresponding apiManagementProductGroup object, and an error if there is any.
func (c *apiManagementProductGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiManagementProductGroup, err error) {
	result = &v1alpha1.ApiManagementProductGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementproductgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiManagementProductGroups that match those selectors.
func (c *apiManagementProductGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiManagementProductGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiManagementProductGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementproductgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiManagementProductGroups.
func (c *apiManagementProductGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementproductgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a apiManagementProductGroup and creates it.  Returns the server's representation of the apiManagementProductGroup, and an error, if there is any.
func (c *apiManagementProductGroups) Create(ctx context.Context, apiManagementProductGroup *v1alpha1.ApiManagementProductGroup, opts v1.CreateOptions) (result *v1alpha1.ApiManagementProductGroup, err error) {
	result = &v1alpha1.ApiManagementProductGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apimanagementproductgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apiManagementProductGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a apiManagementProductGroup and updates it. Returns the server's representation of the apiManagementProductGroup, and an error, if there is any.
func (c *apiManagementProductGroups) Update(ctx context.Context, apiManagementProductGroup *v1alpha1.ApiManagementProductGroup, opts v1.UpdateOptions) (result *v1alpha1.ApiManagementProductGroup, err error) {
	result = &v1alpha1.ApiManagementProductGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementproductgroups").
		Name(apiManagementProductGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apiManagementProductGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *apiManagementProductGroups) UpdateStatus(ctx context.Context, apiManagementProductGroup *v1alpha1.ApiManagementProductGroup, opts v1.UpdateOptions) (result *v1alpha1.ApiManagementProductGroup, err error) {
	result = &v1alpha1.ApiManagementProductGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementproductgroups").
		Name(apiManagementProductGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apiManagementProductGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the apiManagementProductGroup and deletes it. Returns an error if one occurs.
func (c *apiManagementProductGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementproductgroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiManagementProductGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementproductgroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched apiManagementProductGroup.
func (c *apiManagementProductGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiManagementProductGroup, err error) {
	result = &v1alpha1.ApiManagementProductGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apimanagementproductgroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

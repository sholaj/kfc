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

// NeptuneClusterInstancesGetter has a method to return a NeptuneClusterInstanceInterface.
// A group's client should implement this interface.
type NeptuneClusterInstancesGetter interface {
	NeptuneClusterInstances(namespace string) NeptuneClusterInstanceInterface
}

// NeptuneClusterInstanceInterface has methods to work with NeptuneClusterInstance resources.
type NeptuneClusterInstanceInterface interface {
	Create(ctx context.Context, neptuneClusterInstance *v1alpha1.NeptuneClusterInstance, opts v1.CreateOptions) (*v1alpha1.NeptuneClusterInstance, error)
	Update(ctx context.Context, neptuneClusterInstance *v1alpha1.NeptuneClusterInstance, opts v1.UpdateOptions) (*v1alpha1.NeptuneClusterInstance, error)
	UpdateStatus(ctx context.Context, neptuneClusterInstance *v1alpha1.NeptuneClusterInstance, opts v1.UpdateOptions) (*v1alpha1.NeptuneClusterInstance, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NeptuneClusterInstance, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NeptuneClusterInstanceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NeptuneClusterInstance, err error)
	NeptuneClusterInstanceExpansion
}

// neptuneClusterInstances implements NeptuneClusterInstanceInterface
type neptuneClusterInstances struct {
	client rest.Interface
	ns     string
}

// newNeptuneClusterInstances returns a NeptuneClusterInstances
func newNeptuneClusterInstances(c *AwsV1alpha1Client, namespace string) *neptuneClusterInstances {
	return &neptuneClusterInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the neptuneClusterInstance, and returns the corresponding neptuneClusterInstance object, and an error if there is any.
func (c *neptuneClusterInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NeptuneClusterInstance, err error) {
	result = &v1alpha1.NeptuneClusterInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("neptuneclusterinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NeptuneClusterInstances that match those selectors.
func (c *neptuneClusterInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NeptuneClusterInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NeptuneClusterInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("neptuneclusterinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested neptuneClusterInstances.
func (c *neptuneClusterInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("neptuneclusterinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a neptuneClusterInstance and creates it.  Returns the server's representation of the neptuneClusterInstance, and an error, if there is any.
func (c *neptuneClusterInstances) Create(ctx context.Context, neptuneClusterInstance *v1alpha1.NeptuneClusterInstance, opts v1.CreateOptions) (result *v1alpha1.NeptuneClusterInstance, err error) {
	result = &v1alpha1.NeptuneClusterInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("neptuneclusterinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(neptuneClusterInstance).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a neptuneClusterInstance and updates it. Returns the server's representation of the neptuneClusterInstance, and an error, if there is any.
func (c *neptuneClusterInstances) Update(ctx context.Context, neptuneClusterInstance *v1alpha1.NeptuneClusterInstance, opts v1.UpdateOptions) (result *v1alpha1.NeptuneClusterInstance, err error) {
	result = &v1alpha1.NeptuneClusterInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("neptuneclusterinstances").
		Name(neptuneClusterInstance.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(neptuneClusterInstance).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *neptuneClusterInstances) UpdateStatus(ctx context.Context, neptuneClusterInstance *v1alpha1.NeptuneClusterInstance, opts v1.UpdateOptions) (result *v1alpha1.NeptuneClusterInstance, err error) {
	result = &v1alpha1.NeptuneClusterInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("neptuneclusterinstances").
		Name(neptuneClusterInstance.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(neptuneClusterInstance).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the neptuneClusterInstance and deletes it. Returns an error if one occurs.
func (c *neptuneClusterInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("neptuneclusterinstances").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *neptuneClusterInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("neptuneclusterinstances").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched neptuneClusterInstance.
func (c *neptuneClusterInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NeptuneClusterInstance, err error) {
	result = &v1alpha1.NeptuneClusterInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("neptuneclusterinstances").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

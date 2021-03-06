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

// StoragegatewayCachesGetter has a method to return a StoragegatewayCacheInterface.
// A group's client should implement this interface.
type StoragegatewayCachesGetter interface {
	StoragegatewayCaches(namespace string) StoragegatewayCacheInterface
}

// StoragegatewayCacheInterface has methods to work with StoragegatewayCache resources.
type StoragegatewayCacheInterface interface {
	Create(ctx context.Context, storagegatewayCache *v1alpha1.StoragegatewayCache, opts v1.CreateOptions) (*v1alpha1.StoragegatewayCache, error)
	Update(ctx context.Context, storagegatewayCache *v1alpha1.StoragegatewayCache, opts v1.UpdateOptions) (*v1alpha1.StoragegatewayCache, error)
	UpdateStatus(ctx context.Context, storagegatewayCache *v1alpha1.StoragegatewayCache, opts v1.UpdateOptions) (*v1alpha1.StoragegatewayCache, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.StoragegatewayCache, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.StoragegatewayCacheList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StoragegatewayCache, err error)
	StoragegatewayCacheExpansion
}

// storagegatewayCaches implements StoragegatewayCacheInterface
type storagegatewayCaches struct {
	client rest.Interface
	ns     string
}

// newStoragegatewayCaches returns a StoragegatewayCaches
func newStoragegatewayCaches(c *AwsV1alpha1Client, namespace string) *storagegatewayCaches {
	return &storagegatewayCaches{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the storagegatewayCache, and returns the corresponding storagegatewayCache object, and an error if there is any.
func (c *storagegatewayCaches) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StoragegatewayCache, err error) {
	result = &v1alpha1.StoragegatewayCache{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storagegatewaycaches").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StoragegatewayCaches that match those selectors.
func (c *storagegatewayCaches) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StoragegatewayCacheList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StoragegatewayCacheList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storagegatewaycaches").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storagegatewayCaches.
func (c *storagegatewayCaches) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("storagegatewaycaches").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a storagegatewayCache and creates it.  Returns the server's representation of the storagegatewayCache, and an error, if there is any.
func (c *storagegatewayCaches) Create(ctx context.Context, storagegatewayCache *v1alpha1.StoragegatewayCache, opts v1.CreateOptions) (result *v1alpha1.StoragegatewayCache, err error) {
	result = &v1alpha1.StoragegatewayCache{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("storagegatewaycaches").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storagegatewayCache).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a storagegatewayCache and updates it. Returns the server's representation of the storagegatewayCache, and an error, if there is any.
func (c *storagegatewayCaches) Update(ctx context.Context, storagegatewayCache *v1alpha1.StoragegatewayCache, opts v1.UpdateOptions) (result *v1alpha1.StoragegatewayCache, err error) {
	result = &v1alpha1.StoragegatewayCache{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storagegatewaycaches").
		Name(storagegatewayCache.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storagegatewayCache).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *storagegatewayCaches) UpdateStatus(ctx context.Context, storagegatewayCache *v1alpha1.StoragegatewayCache, opts v1.UpdateOptions) (result *v1alpha1.StoragegatewayCache, err error) {
	result = &v1alpha1.StoragegatewayCache{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storagegatewaycaches").
		Name(storagegatewayCache.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storagegatewayCache).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the storagegatewayCache and deletes it. Returns an error if one occurs.
func (c *storagegatewayCaches) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storagegatewaycaches").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storagegatewayCaches) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storagegatewaycaches").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched storagegatewayCache.
func (c *storagegatewayCaches) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StoragegatewayCache, err error) {
	result = &v1alpha1.StoragegatewayCache{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("storagegatewaycaches").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

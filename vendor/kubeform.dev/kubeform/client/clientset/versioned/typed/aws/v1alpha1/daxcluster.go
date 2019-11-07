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
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DaxClustersGetter has a method to return a DaxClusterInterface.
// A group's client should implement this interface.
type DaxClustersGetter interface {
	DaxClusters(namespace string) DaxClusterInterface
}

// DaxClusterInterface has methods to work with DaxCluster resources.
type DaxClusterInterface interface {
	Create(*v1alpha1.DaxCluster) (*v1alpha1.DaxCluster, error)
	Update(*v1alpha1.DaxCluster) (*v1alpha1.DaxCluster, error)
	UpdateStatus(*v1alpha1.DaxCluster) (*v1alpha1.DaxCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DaxCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.DaxClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DaxCluster, err error)
	DaxClusterExpansion
}

// daxClusters implements DaxClusterInterface
type daxClusters struct {
	client rest.Interface
	ns     string
}

// newDaxClusters returns a DaxClusters
func newDaxClusters(c *AwsV1alpha1Client, namespace string) *daxClusters {
	return &daxClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the daxCluster, and returns the corresponding daxCluster object, and an error if there is any.
func (c *daxClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.DaxCluster, err error) {
	result = &v1alpha1.DaxCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("daxclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DaxClusters that match those selectors.
func (c *daxClusters) List(opts v1.ListOptions) (result *v1alpha1.DaxClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DaxClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("daxclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested daxClusters.
func (c *daxClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("daxclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a daxCluster and creates it.  Returns the server's representation of the daxCluster, and an error, if there is any.
func (c *daxClusters) Create(daxCluster *v1alpha1.DaxCluster) (result *v1alpha1.DaxCluster, err error) {
	result = &v1alpha1.DaxCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("daxclusters").
		Body(daxCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a daxCluster and updates it. Returns the server's representation of the daxCluster, and an error, if there is any.
func (c *daxClusters) Update(daxCluster *v1alpha1.DaxCluster) (result *v1alpha1.DaxCluster, err error) {
	result = &v1alpha1.DaxCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("daxclusters").
		Name(daxCluster.Name).
		Body(daxCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *daxClusters) UpdateStatus(daxCluster *v1alpha1.DaxCluster) (result *v1alpha1.DaxCluster, err error) {
	result = &v1alpha1.DaxCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("daxclusters").
		Name(daxCluster.Name).
		SubResource("status").
		Body(daxCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the daxCluster and deletes it. Returns an error if one occurs.
func (c *daxClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("daxclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *daxClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("daxclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched daxCluster.
func (c *daxClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DaxCluster, err error) {
	result = &v1alpha1.DaxCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("daxclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
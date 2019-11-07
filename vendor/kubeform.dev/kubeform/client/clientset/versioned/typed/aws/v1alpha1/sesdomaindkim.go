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

// SesDomainDkimsGetter has a method to return a SesDomainDkimInterface.
// A group's client should implement this interface.
type SesDomainDkimsGetter interface {
	SesDomainDkims(namespace string) SesDomainDkimInterface
}

// SesDomainDkimInterface has methods to work with SesDomainDkim resources.
type SesDomainDkimInterface interface {
	Create(*v1alpha1.SesDomainDkim) (*v1alpha1.SesDomainDkim, error)
	Update(*v1alpha1.SesDomainDkim) (*v1alpha1.SesDomainDkim, error)
	UpdateStatus(*v1alpha1.SesDomainDkim) (*v1alpha1.SesDomainDkim, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SesDomainDkim, error)
	List(opts v1.ListOptions) (*v1alpha1.SesDomainDkimList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesDomainDkim, err error)
	SesDomainDkimExpansion
}

// sesDomainDkims implements SesDomainDkimInterface
type sesDomainDkims struct {
	client rest.Interface
	ns     string
}

// newSesDomainDkims returns a SesDomainDkims
func newSesDomainDkims(c *AwsV1alpha1Client, namespace string) *sesDomainDkims {
	return &sesDomainDkims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sesDomainDkim, and returns the corresponding sesDomainDkim object, and an error if there is any.
func (c *sesDomainDkims) Get(name string, options v1.GetOptions) (result *v1alpha1.SesDomainDkim, err error) {
	result = &v1alpha1.SesDomainDkim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sesdomaindkims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SesDomainDkims that match those selectors.
func (c *sesDomainDkims) List(opts v1.ListOptions) (result *v1alpha1.SesDomainDkimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SesDomainDkimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sesdomaindkims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sesDomainDkims.
func (c *sesDomainDkims) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sesdomaindkims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sesDomainDkim and creates it.  Returns the server's representation of the sesDomainDkim, and an error, if there is any.
func (c *sesDomainDkims) Create(sesDomainDkim *v1alpha1.SesDomainDkim) (result *v1alpha1.SesDomainDkim, err error) {
	result = &v1alpha1.SesDomainDkim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sesdomaindkims").
		Body(sesDomainDkim).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sesDomainDkim and updates it. Returns the server's representation of the sesDomainDkim, and an error, if there is any.
func (c *sesDomainDkims) Update(sesDomainDkim *v1alpha1.SesDomainDkim) (result *v1alpha1.SesDomainDkim, err error) {
	result = &v1alpha1.SesDomainDkim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sesdomaindkims").
		Name(sesDomainDkim.Name).
		Body(sesDomainDkim).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sesDomainDkims) UpdateStatus(sesDomainDkim *v1alpha1.SesDomainDkim) (result *v1alpha1.SesDomainDkim, err error) {
	result = &v1alpha1.SesDomainDkim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sesdomaindkims").
		Name(sesDomainDkim.Name).
		SubResource("status").
		Body(sesDomainDkim).
		Do().
		Into(result)
	return
}

// Delete takes name of the sesDomainDkim and deletes it. Returns an error if one occurs.
func (c *sesDomainDkims) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sesdomaindkims").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sesDomainDkims) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sesdomaindkims").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sesDomainDkim.
func (c *sesDomainDkims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesDomainDkim, err error) {
	result = &v1alpha1.SesDomainDkim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sesdomaindkims").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
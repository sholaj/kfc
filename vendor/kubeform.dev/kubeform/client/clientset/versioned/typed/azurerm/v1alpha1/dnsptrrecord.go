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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DnsPtrRecordsGetter has a method to return a DnsPtrRecordInterface.
// A group's client should implement this interface.
type DnsPtrRecordsGetter interface {
	DnsPtrRecords(namespace string) DnsPtrRecordInterface
}

// DnsPtrRecordInterface has methods to work with DnsPtrRecord resources.
type DnsPtrRecordInterface interface {
	Create(*v1alpha1.DnsPtrRecord) (*v1alpha1.DnsPtrRecord, error)
	Update(*v1alpha1.DnsPtrRecord) (*v1alpha1.DnsPtrRecord, error)
	UpdateStatus(*v1alpha1.DnsPtrRecord) (*v1alpha1.DnsPtrRecord, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DnsPtrRecord, error)
	List(opts v1.ListOptions) (*v1alpha1.DnsPtrRecordList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DnsPtrRecord, err error)
	DnsPtrRecordExpansion
}

// dnsPtrRecords implements DnsPtrRecordInterface
type dnsPtrRecords struct {
	client rest.Interface
	ns     string
}

// newDnsPtrRecords returns a DnsPtrRecords
func newDnsPtrRecords(c *AzurermV1alpha1Client, namespace string) *dnsPtrRecords {
	return &dnsPtrRecords{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dnsPtrRecord, and returns the corresponding dnsPtrRecord object, and an error if there is any.
func (c *dnsPtrRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.DnsPtrRecord, err error) {
	result = &v1alpha1.DnsPtrRecord{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsptrrecords").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DnsPtrRecords that match those selectors.
func (c *dnsPtrRecords) List(opts v1.ListOptions) (result *v1alpha1.DnsPtrRecordList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DnsPtrRecordList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsptrrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dnsPtrRecords.
func (c *dnsPtrRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dnsptrrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dnsPtrRecord and creates it.  Returns the server's representation of the dnsPtrRecord, and an error, if there is any.
func (c *dnsPtrRecords) Create(dnsPtrRecord *v1alpha1.DnsPtrRecord) (result *v1alpha1.DnsPtrRecord, err error) {
	result = &v1alpha1.DnsPtrRecord{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dnsptrrecords").
		Body(dnsPtrRecord).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dnsPtrRecord and updates it. Returns the server's representation of the dnsPtrRecord, and an error, if there is any.
func (c *dnsPtrRecords) Update(dnsPtrRecord *v1alpha1.DnsPtrRecord) (result *v1alpha1.DnsPtrRecord, err error) {
	result = &v1alpha1.DnsPtrRecord{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnsptrrecords").
		Name(dnsPtrRecord.Name).
		Body(dnsPtrRecord).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dnsPtrRecords) UpdateStatus(dnsPtrRecord *v1alpha1.DnsPtrRecord) (result *v1alpha1.DnsPtrRecord, err error) {
	result = &v1alpha1.DnsPtrRecord{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnsptrrecords").
		Name(dnsPtrRecord.Name).
		SubResource("status").
		Body(dnsPtrRecord).
		Do().
		Into(result)
	return
}

// Delete takes name of the dnsPtrRecord and deletes it. Returns an error if one occurs.
func (c *dnsPtrRecords) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsptrrecords").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dnsPtrRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsptrrecords").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dnsPtrRecord.
func (c *dnsPtrRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DnsPtrRecord, err error) {
	result = &v1alpha1.DnsPtrRecord{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dnsptrrecords").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
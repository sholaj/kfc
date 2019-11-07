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

// NotificationHubsGetter has a method to return a NotificationHubInterface.
// A group's client should implement this interface.
type NotificationHubsGetter interface {
	NotificationHubs(namespace string) NotificationHubInterface
}

// NotificationHubInterface has methods to work with NotificationHub resources.
type NotificationHubInterface interface {
	Create(*v1alpha1.NotificationHub) (*v1alpha1.NotificationHub, error)
	Update(*v1alpha1.NotificationHub) (*v1alpha1.NotificationHub, error)
	UpdateStatus(*v1alpha1.NotificationHub) (*v1alpha1.NotificationHub, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NotificationHub, error)
	List(opts v1.ListOptions) (*v1alpha1.NotificationHubList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NotificationHub, err error)
	NotificationHubExpansion
}

// notificationHubs implements NotificationHubInterface
type notificationHubs struct {
	client rest.Interface
	ns     string
}

// newNotificationHubs returns a NotificationHubs
func newNotificationHubs(c *AzurermV1alpha1Client, namespace string) *notificationHubs {
	return &notificationHubs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the notificationHub, and returns the corresponding notificationHub object, and an error if there is any.
func (c *notificationHubs) Get(name string, options v1.GetOptions) (result *v1alpha1.NotificationHub, err error) {
	result = &v1alpha1.NotificationHub{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("notificationhubs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NotificationHubs that match those selectors.
func (c *notificationHubs) List(opts v1.ListOptions) (result *v1alpha1.NotificationHubList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NotificationHubList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("notificationhubs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested notificationHubs.
func (c *notificationHubs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("notificationhubs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a notificationHub and creates it.  Returns the server's representation of the notificationHub, and an error, if there is any.
func (c *notificationHubs) Create(notificationHub *v1alpha1.NotificationHub) (result *v1alpha1.NotificationHub, err error) {
	result = &v1alpha1.NotificationHub{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("notificationhubs").
		Body(notificationHub).
		Do().
		Into(result)
	return
}

// Update takes the representation of a notificationHub and updates it. Returns the server's representation of the notificationHub, and an error, if there is any.
func (c *notificationHubs) Update(notificationHub *v1alpha1.NotificationHub) (result *v1alpha1.NotificationHub, err error) {
	result = &v1alpha1.NotificationHub{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("notificationhubs").
		Name(notificationHub.Name).
		Body(notificationHub).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *notificationHubs) UpdateStatus(notificationHub *v1alpha1.NotificationHub) (result *v1alpha1.NotificationHub, err error) {
	result = &v1alpha1.NotificationHub{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("notificationhubs").
		Name(notificationHub.Name).
		SubResource("status").
		Body(notificationHub).
		Do().
		Into(result)
	return
}

// Delete takes name of the notificationHub and deletes it. Returns an error if one occurs.
func (c *notificationHubs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("notificationhubs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *notificationHubs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("notificationhubs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched notificationHub.
func (c *notificationHubs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NotificationHub, err error) {
	result = &v1alpha1.NotificationHub{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("notificationhubs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
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

// ElasticBeanstalkConfigurationTemplatesGetter has a method to return a ElasticBeanstalkConfigurationTemplateInterface.
// A group's client should implement this interface.
type ElasticBeanstalkConfigurationTemplatesGetter interface {
	ElasticBeanstalkConfigurationTemplates(namespace string) ElasticBeanstalkConfigurationTemplateInterface
}

// ElasticBeanstalkConfigurationTemplateInterface has methods to work with ElasticBeanstalkConfigurationTemplate resources.
type ElasticBeanstalkConfigurationTemplateInterface interface {
	Create(ctx context.Context, elasticBeanstalkConfigurationTemplate *v1alpha1.ElasticBeanstalkConfigurationTemplate, opts v1.CreateOptions) (*v1alpha1.ElasticBeanstalkConfigurationTemplate, error)
	Update(ctx context.Context, elasticBeanstalkConfigurationTemplate *v1alpha1.ElasticBeanstalkConfigurationTemplate, opts v1.UpdateOptions) (*v1alpha1.ElasticBeanstalkConfigurationTemplate, error)
	UpdateStatus(ctx context.Context, elasticBeanstalkConfigurationTemplate *v1alpha1.ElasticBeanstalkConfigurationTemplate, opts v1.UpdateOptions) (*v1alpha1.ElasticBeanstalkConfigurationTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ElasticBeanstalkConfigurationTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ElasticBeanstalkConfigurationTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ElasticBeanstalkConfigurationTemplate, err error)
	ElasticBeanstalkConfigurationTemplateExpansion
}

// elasticBeanstalkConfigurationTemplates implements ElasticBeanstalkConfigurationTemplateInterface
type elasticBeanstalkConfigurationTemplates struct {
	client rest.Interface
	ns     string
}

// newElasticBeanstalkConfigurationTemplates returns a ElasticBeanstalkConfigurationTemplates
func newElasticBeanstalkConfigurationTemplates(c *AwsV1alpha1Client, namespace string) *elasticBeanstalkConfigurationTemplates {
	return &elasticBeanstalkConfigurationTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the elasticBeanstalkConfigurationTemplate, and returns the corresponding elasticBeanstalkConfigurationTemplate object, and an error if there is any.
func (c *elasticBeanstalkConfigurationTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ElasticBeanstalkConfigurationTemplate, err error) {
	result = &v1alpha1.ElasticBeanstalkConfigurationTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elasticbeanstalkconfigurationtemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ElasticBeanstalkConfigurationTemplates that match those selectors.
func (c *elasticBeanstalkConfigurationTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ElasticBeanstalkConfigurationTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ElasticBeanstalkConfigurationTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elasticbeanstalkconfigurationtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested elasticBeanstalkConfigurationTemplates.
func (c *elasticBeanstalkConfigurationTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("elasticbeanstalkconfigurationtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a elasticBeanstalkConfigurationTemplate and creates it.  Returns the server's representation of the elasticBeanstalkConfigurationTemplate, and an error, if there is any.
func (c *elasticBeanstalkConfigurationTemplates) Create(ctx context.Context, elasticBeanstalkConfigurationTemplate *v1alpha1.ElasticBeanstalkConfigurationTemplate, opts v1.CreateOptions) (result *v1alpha1.ElasticBeanstalkConfigurationTemplate, err error) {
	result = &v1alpha1.ElasticBeanstalkConfigurationTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("elasticbeanstalkconfigurationtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(elasticBeanstalkConfigurationTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a elasticBeanstalkConfigurationTemplate and updates it. Returns the server's representation of the elasticBeanstalkConfigurationTemplate, and an error, if there is any.
func (c *elasticBeanstalkConfigurationTemplates) Update(ctx context.Context, elasticBeanstalkConfigurationTemplate *v1alpha1.ElasticBeanstalkConfigurationTemplate, opts v1.UpdateOptions) (result *v1alpha1.ElasticBeanstalkConfigurationTemplate, err error) {
	result = &v1alpha1.ElasticBeanstalkConfigurationTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elasticbeanstalkconfigurationtemplates").
		Name(elasticBeanstalkConfigurationTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(elasticBeanstalkConfigurationTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *elasticBeanstalkConfigurationTemplates) UpdateStatus(ctx context.Context, elasticBeanstalkConfigurationTemplate *v1alpha1.ElasticBeanstalkConfigurationTemplate, opts v1.UpdateOptions) (result *v1alpha1.ElasticBeanstalkConfigurationTemplate, err error) {
	result = &v1alpha1.ElasticBeanstalkConfigurationTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elasticbeanstalkconfigurationtemplates").
		Name(elasticBeanstalkConfigurationTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(elasticBeanstalkConfigurationTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the elasticBeanstalkConfigurationTemplate and deletes it. Returns an error if one occurs.
func (c *elasticBeanstalkConfigurationTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elasticbeanstalkconfigurationtemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *elasticBeanstalkConfigurationTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elasticbeanstalkconfigurationtemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched elasticBeanstalkConfigurationTemplate.
func (c *elasticBeanstalkConfigurationTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ElasticBeanstalkConfigurationTemplate, err error) {
	result = &v1alpha1.ElasticBeanstalkConfigurationTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("elasticbeanstalkconfigurationtemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

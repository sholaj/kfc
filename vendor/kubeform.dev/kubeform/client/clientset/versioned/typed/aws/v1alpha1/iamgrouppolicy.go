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

// IamGroupPoliciesGetter has a method to return a IamGroupPolicyInterface.
// A group's client should implement this interface.
type IamGroupPoliciesGetter interface {
	IamGroupPolicies(namespace string) IamGroupPolicyInterface
}

// IamGroupPolicyInterface has methods to work with IamGroupPolicy resources.
type IamGroupPolicyInterface interface {
	Create(ctx context.Context, iamGroupPolicy *v1alpha1.IamGroupPolicy, opts v1.CreateOptions) (*v1alpha1.IamGroupPolicy, error)
	Update(ctx context.Context, iamGroupPolicy *v1alpha1.IamGroupPolicy, opts v1.UpdateOptions) (*v1alpha1.IamGroupPolicy, error)
	UpdateStatus(ctx context.Context, iamGroupPolicy *v1alpha1.IamGroupPolicy, opts v1.UpdateOptions) (*v1alpha1.IamGroupPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.IamGroupPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.IamGroupPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IamGroupPolicy, err error)
	IamGroupPolicyExpansion
}

// iamGroupPolicies implements IamGroupPolicyInterface
type iamGroupPolicies struct {
	client rest.Interface
	ns     string
}

// newIamGroupPolicies returns a IamGroupPolicies
func newIamGroupPolicies(c *AwsV1alpha1Client, namespace string) *iamGroupPolicies {
	return &iamGroupPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iamGroupPolicy, and returns the corresponding iamGroupPolicy object, and an error if there is any.
func (c *iamGroupPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IamGroupPolicy, err error) {
	result = &v1alpha1.IamGroupPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamgrouppolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IamGroupPolicies that match those selectors.
func (c *iamGroupPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IamGroupPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IamGroupPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamgrouppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iamGroupPolicies.
func (c *iamGroupPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iamgrouppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a iamGroupPolicy and creates it.  Returns the server's representation of the iamGroupPolicy, and an error, if there is any.
func (c *iamGroupPolicies) Create(ctx context.Context, iamGroupPolicy *v1alpha1.IamGroupPolicy, opts v1.CreateOptions) (result *v1alpha1.IamGroupPolicy, err error) {
	result = &v1alpha1.IamGroupPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iamgrouppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iamGroupPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a iamGroupPolicy and updates it. Returns the server's representation of the iamGroupPolicy, and an error, if there is any.
func (c *iamGroupPolicies) Update(ctx context.Context, iamGroupPolicy *v1alpha1.IamGroupPolicy, opts v1.UpdateOptions) (result *v1alpha1.IamGroupPolicy, err error) {
	result = &v1alpha1.IamGroupPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamgrouppolicies").
		Name(iamGroupPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iamGroupPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *iamGroupPolicies) UpdateStatus(ctx context.Context, iamGroupPolicy *v1alpha1.IamGroupPolicy, opts v1.UpdateOptions) (result *v1alpha1.IamGroupPolicy, err error) {
	result = &v1alpha1.IamGroupPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamgrouppolicies").
		Name(iamGroupPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iamGroupPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the iamGroupPolicy and deletes it. Returns an error if one occurs.
func (c *iamGroupPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamgrouppolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iamGroupPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamgrouppolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched iamGroupPolicy.
func (c *iamGroupPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IamGroupPolicy, err error) {
	result = &v1alpha1.IamGroupPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iamgrouppolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

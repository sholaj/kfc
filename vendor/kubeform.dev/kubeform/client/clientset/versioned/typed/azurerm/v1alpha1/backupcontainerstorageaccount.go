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

// BackupContainerStorageAccountsGetter has a method to return a BackupContainerStorageAccountInterface.
// A group's client should implement this interface.
type BackupContainerStorageAccountsGetter interface {
	BackupContainerStorageAccounts(namespace string) BackupContainerStorageAccountInterface
}

// BackupContainerStorageAccountInterface has methods to work with BackupContainerStorageAccount resources.
type BackupContainerStorageAccountInterface interface {
	Create(ctx context.Context, backupContainerStorageAccount *v1alpha1.BackupContainerStorageAccount, opts v1.CreateOptions) (*v1alpha1.BackupContainerStorageAccount, error)
	Update(ctx context.Context, backupContainerStorageAccount *v1alpha1.BackupContainerStorageAccount, opts v1.UpdateOptions) (*v1alpha1.BackupContainerStorageAccount, error)
	UpdateStatus(ctx context.Context, backupContainerStorageAccount *v1alpha1.BackupContainerStorageAccount, opts v1.UpdateOptions) (*v1alpha1.BackupContainerStorageAccount, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BackupContainerStorageAccount, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BackupContainerStorageAccountList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackupContainerStorageAccount, err error)
	BackupContainerStorageAccountExpansion
}

// backupContainerStorageAccounts implements BackupContainerStorageAccountInterface
type backupContainerStorageAccounts struct {
	client rest.Interface
	ns     string
}

// newBackupContainerStorageAccounts returns a BackupContainerStorageAccounts
func newBackupContainerStorageAccounts(c *AzurermV1alpha1Client, namespace string) *backupContainerStorageAccounts {
	return &backupContainerStorageAccounts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backupContainerStorageAccount, and returns the corresponding backupContainerStorageAccount object, and an error if there is any.
func (c *backupContainerStorageAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BackupContainerStorageAccount, err error) {
	result = &v1alpha1.BackupContainerStorageAccount{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupcontainerstorageaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupContainerStorageAccounts that match those selectors.
func (c *backupContainerStorageAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackupContainerStorageAccountList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BackupContainerStorageAccountList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupcontainerstorageaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupContainerStorageAccounts.
func (c *backupContainerStorageAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backupcontainerstorageaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a backupContainerStorageAccount and creates it.  Returns the server's representation of the backupContainerStorageAccount, and an error, if there is any.
func (c *backupContainerStorageAccounts) Create(ctx context.Context, backupContainerStorageAccount *v1alpha1.BackupContainerStorageAccount, opts v1.CreateOptions) (result *v1alpha1.BackupContainerStorageAccount, err error) {
	result = &v1alpha1.BackupContainerStorageAccount{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backupcontainerstorageaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupContainerStorageAccount).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a backupContainerStorageAccount and updates it. Returns the server's representation of the backupContainerStorageAccount, and an error, if there is any.
func (c *backupContainerStorageAccounts) Update(ctx context.Context, backupContainerStorageAccount *v1alpha1.BackupContainerStorageAccount, opts v1.UpdateOptions) (result *v1alpha1.BackupContainerStorageAccount, err error) {
	result = &v1alpha1.BackupContainerStorageAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupcontainerstorageaccounts").
		Name(backupContainerStorageAccount.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupContainerStorageAccount).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *backupContainerStorageAccounts) UpdateStatus(ctx context.Context, backupContainerStorageAccount *v1alpha1.BackupContainerStorageAccount, opts v1.UpdateOptions) (result *v1alpha1.BackupContainerStorageAccount, err error) {
	result = &v1alpha1.BackupContainerStorageAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupcontainerstorageaccounts").
		Name(backupContainerStorageAccount.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupContainerStorageAccount).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the backupContainerStorageAccount and deletes it. Returns an error if one occurs.
func (c *backupContainerStorageAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupcontainerstorageaccounts").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupContainerStorageAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupcontainerstorageaccounts").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched backupContainerStorageAccount.
func (c *backupContainerStorageAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackupContainerStorageAccount, err error) {
	result = &v1alpha1.BackupContainerStorageAccount{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backupcontainerstorageaccounts").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
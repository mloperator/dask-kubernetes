// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/dask/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/apis/kubernetes.dask.org/v1"
	scheme "github.com/dask/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DaskWorkerGroupsGetter has a method to return a DaskWorkerGroupInterface.
// A group's client should implement this interface.
type DaskWorkerGroupsGetter interface {
	DaskWorkerGroups(namespace string) DaskWorkerGroupInterface
}

// DaskWorkerGroupInterface has methods to work with DaskWorkerGroup resources.
type DaskWorkerGroupInterface interface {
	Create(ctx context.Context, daskWorkerGroup *v1.DaskWorkerGroup, opts metav1.CreateOptions) (*v1.DaskWorkerGroup, error)
	Update(ctx context.Context, daskWorkerGroup *v1.DaskWorkerGroup, opts metav1.UpdateOptions) (*v1.DaskWorkerGroup, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DaskWorkerGroup, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DaskWorkerGroupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DaskWorkerGroup, err error)
	DaskWorkerGroupExpansion
}

// daskWorkerGroups implements DaskWorkerGroupInterface
type daskWorkerGroups struct {
	client rest.Interface
	ns     string
}

// newDaskWorkerGroups returns a DaskWorkerGroups
func newDaskWorkerGroups(c *KubernetesV1Client, namespace string) *daskWorkerGroups {
	return &daskWorkerGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the daskWorkerGroup, and returns the corresponding daskWorkerGroup object, and an error if there is any.
func (c *daskWorkerGroups) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DaskWorkerGroup, err error) {
	result = &v1.DaskWorkerGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("daskworkergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DaskWorkerGroups that match those selectors.
func (c *daskWorkerGroups) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DaskWorkerGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DaskWorkerGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("daskworkergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested daskWorkerGroups.
func (c *daskWorkerGroups) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("daskworkergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a daskWorkerGroup and creates it.  Returns the server's representation of the daskWorkerGroup, and an error, if there is any.
func (c *daskWorkerGroups) Create(ctx context.Context, daskWorkerGroup *v1.DaskWorkerGroup, opts metav1.CreateOptions) (result *v1.DaskWorkerGroup, err error) {
	result = &v1.DaskWorkerGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("daskworkergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(daskWorkerGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a daskWorkerGroup and updates it. Returns the server's representation of the daskWorkerGroup, and an error, if there is any.
func (c *daskWorkerGroups) Update(ctx context.Context, daskWorkerGroup *v1.DaskWorkerGroup, opts metav1.UpdateOptions) (result *v1.DaskWorkerGroup, err error) {
	result = &v1.DaskWorkerGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("daskworkergroups").
		Name(daskWorkerGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(daskWorkerGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the daskWorkerGroup and deletes it. Returns an error if one occurs.
func (c *daskWorkerGroups) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("daskworkergroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *daskWorkerGroups) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("daskworkergroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched daskWorkerGroup.
func (c *daskWorkerGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DaskWorkerGroup, err error) {
	result = &v1.DaskWorkerGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("daskworkergroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

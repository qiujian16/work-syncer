// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "open-cluster-management.io/api/client/work/clientset/versioned/scheme"
	v1 "open-cluster-management.io/api/work/v1"
)

// ManifestWorksGetter has a method to return a ManifestWorkInterface.
// A group's client should implement this interface.
type ManifestWorksGetter interface {
	ManifestWorks(namespace string) ManifestWorkInterface
}

// ManifestWorkInterface has methods to work with ManifestWork resources.
type ManifestWorkInterface interface {
	Create(ctx context.Context, manifestWork *v1.ManifestWork, opts metav1.CreateOptions) (*v1.ManifestWork, error)
	Update(ctx context.Context, manifestWork *v1.ManifestWork, opts metav1.UpdateOptions) (*v1.ManifestWork, error)
	UpdateStatus(ctx context.Context, manifestWork *v1.ManifestWork, opts metav1.UpdateOptions) (*v1.ManifestWork, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ManifestWork, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ManifestWorkList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ManifestWork, err error)
	ManifestWorkExpansion
}

// manifestWorks implements ManifestWorkInterface
type manifestWorks struct {
	client rest.Interface
	ns     string
}

// newManifestWorks returns a ManifestWorks
func newManifestWorks(c *WorkV1Client, namespace string) *manifestWorks {
	return &manifestWorks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the manifestWork, and returns the corresponding manifestWork object, and an error if there is any.
func (c *manifestWorks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ManifestWork, err error) {
	result = &v1.ManifestWork{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("manifestworks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ManifestWorks that match those selectors.
func (c *manifestWorks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ManifestWorkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ManifestWorkList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("manifestworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested manifestWorks.
func (c *manifestWorks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("manifestworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a manifestWork and creates it.  Returns the server's representation of the manifestWork, and an error, if there is any.
func (c *manifestWorks) Create(ctx context.Context, manifestWork *v1.ManifestWork, opts metav1.CreateOptions) (result *v1.ManifestWork, err error) {
	result = &v1.ManifestWork{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("manifestworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(manifestWork).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a manifestWork and updates it. Returns the server's representation of the manifestWork, and an error, if there is any.
func (c *manifestWorks) Update(ctx context.Context, manifestWork *v1.ManifestWork, opts metav1.UpdateOptions) (result *v1.ManifestWork, err error) {
	result = &v1.ManifestWork{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("manifestworks").
		Name(manifestWork.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(manifestWork).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *manifestWorks) UpdateStatus(ctx context.Context, manifestWork *v1.ManifestWork, opts metav1.UpdateOptions) (result *v1.ManifestWork, err error) {
	result = &v1.ManifestWork{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("manifestworks").
		Name(manifestWork.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(manifestWork).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the manifestWork and deletes it. Returns an error if one occurs.
func (c *manifestWorks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("manifestworks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *manifestWorks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("manifestworks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched manifestWork.
func (c *manifestWorks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ManifestWork, err error) {
	result = &v1.ManifestWork{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("manifestworks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

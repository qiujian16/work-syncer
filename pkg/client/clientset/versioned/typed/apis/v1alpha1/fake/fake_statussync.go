// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/qiujian16/work-syncer/pkg/apis/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStatusSyncs implements StatusSyncInterface
type FakeStatusSyncs struct {
	Fake *FakeWorkV1alpha1
	ns   string
}

var statussyncsResource = schema.GroupVersionResource{Group: "work.open-cluster-management.io", Version: "v1alpha1", Resource: "statussyncs"}

var statussyncsKind = schema.GroupVersionKind{Group: "work.open-cluster-management.io", Version: "v1alpha1", Kind: "StatusSync"}

// Get takes name of the statusSync, and returns the corresponding statusSync object, and an error if there is any.
func (c *FakeStatusSyncs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StatusSync, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(statussyncsResource, c.ns, name), &v1alpha1.StatusSync{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StatusSync), err
}

// List takes label and field selectors, and returns the list of StatusSyncs that match those selectors.
func (c *FakeStatusSyncs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StatusSyncList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(statussyncsResource, statussyncsKind, c.ns, opts), &v1alpha1.StatusSyncList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StatusSyncList{ListMeta: obj.(*v1alpha1.StatusSyncList).ListMeta}
	for _, item := range obj.(*v1alpha1.StatusSyncList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested statusSyncs.
func (c *FakeStatusSyncs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(statussyncsResource, c.ns, opts))

}

// Create takes the representation of a statusSync and creates it.  Returns the server's representation of the statusSync, and an error, if there is any.
func (c *FakeStatusSyncs) Create(ctx context.Context, statusSync *v1alpha1.StatusSync, opts v1.CreateOptions) (result *v1alpha1.StatusSync, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(statussyncsResource, c.ns, statusSync), &v1alpha1.StatusSync{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StatusSync), err
}

// Update takes the representation of a statusSync and updates it. Returns the server's representation of the statusSync, and an error, if there is any.
func (c *FakeStatusSyncs) Update(ctx context.Context, statusSync *v1alpha1.StatusSync, opts v1.UpdateOptions) (result *v1alpha1.StatusSync, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(statussyncsResource, c.ns, statusSync), &v1alpha1.StatusSync{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StatusSync), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStatusSyncs) UpdateStatus(ctx context.Context, statusSync *v1alpha1.StatusSync, opts v1.UpdateOptions) (*v1alpha1.StatusSync, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(statussyncsResource, "status", c.ns, statusSync), &v1alpha1.StatusSync{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StatusSync), err
}

// Delete takes name of the statusSync and deletes it. Returns an error if one occurs.
func (c *FakeStatusSyncs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(statussyncsResource, c.ns, name), &v1alpha1.StatusSync{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStatusSyncs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(statussyncsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StatusSyncList{})
	return err
}

// Patch applies the patch and returns the patched statusSync.
func (c *FakeStatusSyncs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StatusSync, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(statussyncsResource, c.ns, name, pt, data, subresources...), &v1alpha1.StatusSync{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StatusSync), err
}

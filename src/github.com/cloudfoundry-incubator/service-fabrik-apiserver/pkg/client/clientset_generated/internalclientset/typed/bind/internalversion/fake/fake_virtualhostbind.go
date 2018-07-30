//TODO copyright header
package fake

import (
	bind "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualhostbinds implements VirtualhostbindInterface
type FakeVirtualhostbinds struct {
	Fake *FakeBind
	ns   string
}

var virtualhostbindsResource = schema.GroupVersionResource{Group: "bind.servicefabrik.io", Version: "", Resource: "virtualhostbinds"}

var virtualhostbindsKind = schema.GroupVersionKind{Group: "bind.servicefabrik.io", Version: "", Kind: "Virtualhostbind"}

// Get takes name of the virtualhostbind, and returns the corresponding virtualhostbind object, and an error if there is any.
func (c *FakeVirtualhostbinds) Get(name string, options v1.GetOptions) (result *bind.Virtualhostbind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualhostbindsResource, c.ns, name), &bind.Virtualhostbind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bind.Virtualhostbind), err
}

// List takes label and field selectors, and returns the list of Virtualhostbinds that match those selectors.
func (c *FakeVirtualhostbinds) List(opts v1.ListOptions) (result *bind.VirtualhostbindList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualhostbindsResource, virtualhostbindsKind, c.ns, opts), &bind.VirtualhostbindList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &bind.VirtualhostbindList{}
	for _, item := range obj.(*bind.VirtualhostbindList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualhostbinds.
func (c *FakeVirtualhostbinds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualhostbindsResource, c.ns, opts))

}

// Create takes the representation of a virtualhostbind and creates it.  Returns the server's representation of the virtualhostbind, and an error, if there is any.
func (c *FakeVirtualhostbinds) Create(virtualhostbind *bind.Virtualhostbind) (result *bind.Virtualhostbind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualhostbindsResource, c.ns, virtualhostbind), &bind.Virtualhostbind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bind.Virtualhostbind), err
}

// Update takes the representation of a virtualhostbind and updates it. Returns the server's representation of the virtualhostbind, and an error, if there is any.
func (c *FakeVirtualhostbinds) Update(virtualhostbind *bind.Virtualhostbind) (result *bind.Virtualhostbind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualhostbindsResource, c.ns, virtualhostbind), &bind.Virtualhostbind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bind.Virtualhostbind), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualhostbinds) UpdateStatus(virtualhostbind *bind.Virtualhostbind) (*bind.Virtualhostbind, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualhostbindsResource, "status", c.ns, virtualhostbind), &bind.Virtualhostbind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bind.Virtualhostbind), err
}

// Delete takes name of the virtualhostbind and deletes it. Returns an error if one occurs.
func (c *FakeVirtualhostbinds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualhostbindsResource, c.ns, name), &bind.Virtualhostbind{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualhostbinds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualhostbindsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &bind.VirtualhostbindList{})
	return err
}

// Patch applies the patch and returns the patched virtualhostbind.
func (c *FakeVirtualhostbinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *bind.Virtualhostbind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualhostbindsResource, c.ns, name, data, subresources...), &bind.Virtualhostbind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bind.Virtualhostbind), err
}

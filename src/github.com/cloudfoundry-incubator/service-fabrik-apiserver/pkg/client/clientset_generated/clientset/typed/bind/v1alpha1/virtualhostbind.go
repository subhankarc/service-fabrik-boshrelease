//TODO copyright header
package v1alpha1

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind/v1alpha1"
	scheme "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualhostbindsGetter has a method to return a VirtualhostbindInterface.
// A group's client should implement this interface.
type VirtualhostbindsGetter interface {
	Virtualhostbinds(namespace string) VirtualhostbindInterface
}

// VirtualhostbindInterface has methods to work with Virtualhostbind resources.
type VirtualhostbindInterface interface {
	Create(*v1alpha1.Virtualhostbind) (*v1alpha1.Virtualhostbind, error)
	Update(*v1alpha1.Virtualhostbind) (*v1alpha1.Virtualhostbind, error)
	UpdateStatus(*v1alpha1.Virtualhostbind) (*v1alpha1.Virtualhostbind, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Virtualhostbind, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualhostbindList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Virtualhostbind, err error)
	VirtualhostbindExpansion
}

// virtualhostbinds implements VirtualhostbindInterface
type virtualhostbinds struct {
	client rest.Interface
	ns     string
}

// newVirtualhostbinds returns a Virtualhostbinds
func newVirtualhostbinds(c *BindV1alpha1Client, namespace string) *virtualhostbinds {
	return &virtualhostbinds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualhostbind, and returns the corresponding virtualhostbind object, and an error if there is any.
func (c *virtualhostbinds) Get(name string, options v1.GetOptions) (result *v1alpha1.Virtualhostbind, err error) {
	result = &v1alpha1.Virtualhostbind{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualhostbinds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Virtualhostbinds that match those selectors.
func (c *virtualhostbinds) List(opts v1.ListOptions) (result *v1alpha1.VirtualhostbindList, err error) {
	result = &v1alpha1.VirtualhostbindList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualhostbinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualhostbinds.
func (c *virtualhostbinds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualhostbinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a virtualhostbind and creates it.  Returns the server's representation of the virtualhostbind, and an error, if there is any.
func (c *virtualhostbinds) Create(virtualhostbind *v1alpha1.Virtualhostbind) (result *v1alpha1.Virtualhostbind, err error) {
	result = &v1alpha1.Virtualhostbind{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualhostbinds").
		Body(virtualhostbind).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualhostbind and updates it. Returns the server's representation of the virtualhostbind, and an error, if there is any.
func (c *virtualhostbinds) Update(virtualhostbind *v1alpha1.Virtualhostbind) (result *v1alpha1.Virtualhostbind, err error) {
	result = &v1alpha1.Virtualhostbind{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualhostbinds").
		Name(virtualhostbind.Name).
		Body(virtualhostbind).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualhostbinds) UpdateStatus(virtualhostbind *v1alpha1.Virtualhostbind) (result *v1alpha1.Virtualhostbind, err error) {
	result = &v1alpha1.Virtualhostbind{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualhostbinds").
		Name(virtualhostbind.Name).
		SubResource("status").
		Body(virtualhostbind).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualhostbind and deletes it. Returns an error if one occurs.
func (c *virtualhostbinds) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualhostbinds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualhostbinds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualhostbinds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualhostbind.
func (c *virtualhostbinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Virtualhostbind, err error) {
	result = &v1alpha1.Virtualhostbind{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualhostbinds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
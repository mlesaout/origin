package internalversion

import (
	api "github.com/openshift/origin/pkg/build/api"
	scheme "github.com/openshift/origin/pkg/build/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BuildConfigsGetter has a method to return a BuildConfigInterface.
// A group's client should implement this interface.
type BuildConfigsGetter interface {
	BuildConfigs(namespace string) BuildConfigInterface
}

// BuildConfigInterface has methods to work with BuildConfig resources.
type BuildConfigInterface interface {
	Create(*api.BuildConfig) (*api.BuildConfig, error)
	Update(*api.BuildConfig) (*api.BuildConfig, error)
	UpdateStatus(*api.BuildConfig) (*api.BuildConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*api.BuildConfig, error)
	List(opts v1.ListOptions) (*api.BuildConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *api.BuildConfig, err error)
	BuildConfigExpansion
}

// buildConfigs implements BuildConfigInterface
type buildConfigs struct {
	client rest.Interface
	ns     string
}

// newBuildConfigs returns a BuildConfigs
func newBuildConfigs(c *BuildClient, namespace string) *buildConfigs {
	return &buildConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a buildConfig and creates it.  Returns the server's representation of the buildConfig, and an error, if there is any.
func (c *buildConfigs) Create(buildConfig *api.BuildConfig) (result *api.BuildConfig, err error) {
	result = &api.BuildConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("buildconfigs").
		Body(buildConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a buildConfig and updates it. Returns the server's representation of the buildConfig, and an error, if there is any.
func (c *buildConfigs) Update(buildConfig *api.BuildConfig) (result *api.BuildConfig, err error) {
	result = &api.BuildConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("buildconfigs").
		Name(buildConfig.Name).
		Body(buildConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *buildConfigs) UpdateStatus(buildConfig *api.BuildConfig) (result *api.BuildConfig, err error) {
	result = &api.BuildConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("buildconfigs").
		Name(buildConfig.Name).
		SubResource("status").
		Body(buildConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the buildConfig and deletes it. Returns an error if one occurs.
func (c *buildConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("buildconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *buildConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("buildconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the buildConfig, and returns the corresponding buildConfig object, and an error if there is any.
func (c *buildConfigs) Get(name string, options v1.GetOptions) (result *api.BuildConfig, err error) {
	result = &api.BuildConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("buildconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BuildConfigs that match those selectors.
func (c *buildConfigs) List(opts v1.ListOptions) (result *api.BuildConfigList, err error) {
	result = &api.BuildConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("buildconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested buildConfigs.
func (c *buildConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("buildconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched buildConfig.
func (c *buildConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *api.BuildConfig, err error) {
	result = &api.BuildConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("buildconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

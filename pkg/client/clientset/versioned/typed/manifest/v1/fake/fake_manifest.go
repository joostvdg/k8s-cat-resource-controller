/*
Copyright The Kubernetes Authors.

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

package fake

import (
	manifestv1 "github.com/joostvdg/k8s-cat-resource-controller/pkg/apis/manifest/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManifests implements ManifestInterface
type FakeManifests struct {
	Fake *FakeManifestV1
	ns   string
}

var manifestsResource = schema.GroupVersionResource{Group: "manifest", Version: "v1", Resource: "manifests"}

var manifestsKind = schema.GroupVersionKind{Group: "manifest", Version: "v1", Kind: "Manifest"}

// Get takes name of the manifest, and returns the corresponding manifest object, and an error if there is any.
func (c *FakeManifests) Get(name string, options v1.GetOptions) (result *manifestv1.Manifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(manifestsResource, c.ns, name), &manifestv1.Manifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*manifestv1.Manifest), err
}

// List takes label and field selectors, and returns the list of Manifests that match those selectors.
func (c *FakeManifests) List(opts v1.ListOptions) (result *manifestv1.ManifestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(manifestsResource, manifestsKind, c.ns, opts), &manifestv1.ManifestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &manifestv1.ManifestList{ListMeta: obj.(*manifestv1.ManifestList).ListMeta}
	for _, item := range obj.(*manifestv1.ManifestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested manifests.
func (c *FakeManifests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(manifestsResource, c.ns, opts))

}

// Create takes the representation of a manifest and creates it.  Returns the server's representation of the manifest, and an error, if there is any.
func (c *FakeManifests) Create(manifest *manifestv1.Manifest) (result *manifestv1.Manifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(manifestsResource, c.ns, manifest), &manifestv1.Manifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*manifestv1.Manifest), err
}

// Update takes the representation of a manifest and updates it. Returns the server's representation of the manifest, and an error, if there is any.
func (c *FakeManifests) Update(manifest *manifestv1.Manifest) (result *manifestv1.Manifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(manifestsResource, c.ns, manifest), &manifestv1.Manifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*manifestv1.Manifest), err
}

// Delete takes name of the manifest and deletes it. Returns an error if one occurs.
func (c *FakeManifests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(manifestsResource, c.ns, name), &manifestv1.Manifest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManifests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(manifestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &manifestv1.ManifestList{})
	return err
}

// Patch applies the patch and returns the patched manifest.
func (c *FakeManifests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *manifestv1.Manifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(manifestsResource, c.ns, name, pt, data, subresources...), &manifestv1.Manifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*manifestv1.Manifest), err
}

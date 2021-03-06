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
	v1alpha1 "github.com/awslabs/aws-eks-cluster-controller/pkg/apis/cluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeGroups implements NodeGroupInterface
type FakeNodeGroups struct {
	Fake *FakeClusterV1alpha1
	ns   string
}

var nodegroupsResource = schema.GroupVersionResource{Group: "cluster.eks.amazonaws.com", Version: "v1alpha1", Resource: "nodegroups"}

var nodegroupsKind = schema.GroupVersionKind{Group: "cluster.eks.amazonaws.com", Version: "v1alpha1", Kind: "NodeGroup"}

// Get takes name of the nodeGroup, and returns the corresponding nodeGroup object, and an error if there is any.
func (c *FakeNodeGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.NodeGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodegroupsResource, c.ns, name), &v1alpha1.NodeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeGroup), err
}

// List takes label and field selectors, and returns the list of NodeGroups that match those selectors.
func (c *FakeNodeGroups) List(opts v1.ListOptions) (result *v1alpha1.NodeGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodegroupsResource, nodegroupsKind, c.ns, opts), &v1alpha1.NodeGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeGroupList{ListMeta: obj.(*v1alpha1.NodeGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodeGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeGroups.
func (c *FakeNodeGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodegroupsResource, c.ns, opts))

}

// Create takes the representation of a nodeGroup and creates it.  Returns the server's representation of the nodeGroup, and an error, if there is any.
func (c *FakeNodeGroups) Create(nodeGroup *v1alpha1.NodeGroup) (result *v1alpha1.NodeGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodegroupsResource, c.ns, nodeGroup), &v1alpha1.NodeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeGroup), err
}

// Update takes the representation of a nodeGroup and updates it. Returns the server's representation of the nodeGroup, and an error, if there is any.
func (c *FakeNodeGroups) Update(nodeGroup *v1alpha1.NodeGroup) (result *v1alpha1.NodeGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodegroupsResource, c.ns, nodeGroup), &v1alpha1.NodeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeGroups) UpdateStatus(nodeGroup *v1alpha1.NodeGroup) (*v1alpha1.NodeGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nodegroupsResource, "status", c.ns, nodeGroup), &v1alpha1.NodeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeGroup), err
}

// Delete takes name of the nodeGroup and deletes it. Returns an error if one occurs.
func (c *FakeNodeGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nodegroupsResource, c.ns, name), &v1alpha1.NodeGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodegroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeGroupList{})
	return err
}

// Patch applies the patch and returns the patched nodeGroup.
func (c *FakeNodeGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodegroupsResource, c.ns, name, data, subresources...), &v1alpha1.NodeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeGroup), err
}

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

// FakeEKSs implements EKSInterface
type FakeEKSs struct {
	Fake *FakeClusterV1alpha1
	ns   string
}

var ekssResource = schema.GroupVersionResource{Group: "cluster.eks.amazonaws.com", Version: "v1alpha1", Resource: "ekss"}

var ekssKind = schema.GroupVersionKind{Group: "cluster.eks.amazonaws.com", Version: "v1alpha1", Kind: "EKS"}

// Get takes name of the eKS, and returns the corresponding eKS object, and an error if there is any.
func (c *FakeEKSs) Get(name string, options v1.GetOptions) (result *v1alpha1.EKS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ekssResource, c.ns, name), &v1alpha1.EKS{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EKS), err
}

// List takes label and field selectors, and returns the list of EKSs that match those selectors.
func (c *FakeEKSs) List(opts v1.ListOptions) (result *v1alpha1.EKSList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ekssResource, ekssKind, c.ns, opts), &v1alpha1.EKSList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EKSList{ListMeta: obj.(*v1alpha1.EKSList).ListMeta}
	for _, item := range obj.(*v1alpha1.EKSList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eKSs.
func (c *FakeEKSs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ekssResource, c.ns, opts))

}

// Create takes the representation of a eKS and creates it.  Returns the server's representation of the eKS, and an error, if there is any.
func (c *FakeEKSs) Create(eKS *v1alpha1.EKS) (result *v1alpha1.EKS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ekssResource, c.ns, eKS), &v1alpha1.EKS{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EKS), err
}

// Update takes the representation of a eKS and updates it. Returns the server's representation of the eKS, and an error, if there is any.
func (c *FakeEKSs) Update(eKS *v1alpha1.EKS) (result *v1alpha1.EKS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ekssResource, c.ns, eKS), &v1alpha1.EKS{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EKS), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEKSs) UpdateStatus(eKS *v1alpha1.EKS) (*v1alpha1.EKS, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ekssResource, "status", c.ns, eKS), &v1alpha1.EKS{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EKS), err
}

// Delete takes name of the eKS and deletes it. Returns an error if one occurs.
func (c *FakeEKSs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ekssResource, c.ns, name), &v1alpha1.EKS{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEKSs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ekssResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EKSList{})
	return err
}

// Patch applies the patch and returns the patched eKS.
func (c *FakeEKSs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EKS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ekssResource, c.ns, name, pt, data, subresources...), &v1alpha1.EKS{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EKS), err
}

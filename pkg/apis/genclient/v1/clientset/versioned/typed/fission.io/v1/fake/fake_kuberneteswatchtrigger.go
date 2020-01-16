/*
Copyright The Fission Authors.

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
	fissioniov1 "github.com/fission/fission/pkg/apis/fission.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubernetesWatchTriggers implements KubernetesWatchTriggerInterface
type FakeKubernetesWatchTriggers struct {
	Fake *FakeV1V1
	ns   string
}

var kuberneteswatchtriggersResource = schema.GroupVersionResource{Group: "fission.io", Version: "v1", Resource: "kuberneteswatchtriggers"}

var kuberneteswatchtriggersKind = schema.GroupVersionKind{Group: "fission.io", Version: "v1", Kind: "KubernetesWatchTrigger"}

// Get takes name of the _kubernetesWatchTrigger, and returns the corresponding kubernetesWatchTrigger object, and an error if there is any.
func (c *FakeKubernetesWatchTriggers) Get(name string, options v1.GetOptions) (result *fissioniov1.KubernetesWatchTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kuberneteswatchtriggersResource, c.ns, name), &fissioniov1.KubernetesWatchTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fissioniov1.KubernetesWatchTrigger), err
}

// List takes label and field selectors, and returns the list of KubernetesWatchTriggers that match those selectors.
func (c *FakeKubernetesWatchTriggers) List(opts v1.ListOptions) (result *fissioniov1.KubernetesWatchTriggerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kuberneteswatchtriggersResource, kuberneteswatchtriggersKind, c.ns, opts), &fissioniov1.KubernetesWatchTriggerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &fissioniov1.KubernetesWatchTriggerList{ListMeta: obj.(*fissioniov1.KubernetesWatchTriggerList).ListMeta}
	for _, item := range obj.(*fissioniov1.KubernetesWatchTriggerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubernetesWatchTriggers.
func (c *FakeKubernetesWatchTriggers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kuberneteswatchtriggersResource, c.ns, opts))

}

// Create takes the representation of a _kubernetesWatchTrigger and creates it.  Returns the server's representation of the kubernetesWatchTrigger, and an error, if there is any.
func (c *FakeKubernetesWatchTriggers) Create(_kubernetesWatchTrigger *fissioniov1.KubernetesWatchTrigger) (result *fissioniov1.KubernetesWatchTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kuberneteswatchtriggersResource, c.ns, _kubernetesWatchTrigger), &fissioniov1.KubernetesWatchTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fissioniov1.KubernetesWatchTrigger), err
}

// Update takes the representation of a _kubernetesWatchTrigger and updates it. Returns the server's representation of the kubernetesWatchTrigger, and an error, if there is any.
func (c *FakeKubernetesWatchTriggers) Update(_kubernetesWatchTrigger *fissioniov1.KubernetesWatchTrigger) (result *fissioniov1.KubernetesWatchTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kuberneteswatchtriggersResource, c.ns, _kubernetesWatchTrigger), &fissioniov1.KubernetesWatchTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fissioniov1.KubernetesWatchTrigger), err
}

// Delete takes name of the _kubernetesWatchTrigger and deletes it. Returns an error if one occurs.
func (c *FakeKubernetesWatchTriggers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kuberneteswatchtriggersResource, c.ns, name), &fissioniov1.KubernetesWatchTrigger{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubernetesWatchTriggers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kuberneteswatchtriggersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &fissioniov1.KubernetesWatchTriggerList{})
	return err
}

// Patch applies the patch and returns the patched kubernetesWatchTrigger.
func (c *FakeKubernetesWatchTriggers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *fissioniov1.KubernetesWatchTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kuberneteswatchtriggersResource, c.ns, name, pt, data, subresources...), &fissioniov1.KubernetesWatchTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fissioniov1.KubernetesWatchTrigger), err
}

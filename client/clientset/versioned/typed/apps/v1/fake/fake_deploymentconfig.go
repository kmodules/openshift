/*
Copyright AppsCode Inc. and Contributors

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
	"context"
	json "encoding/json"
	"fmt"

	v1 "kmodules.xyz/openshift/apis/apps/v1"
	appsv1 "kmodules.xyz/openshift/client/applyconfiguration/apps/v1"

	v1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeploymentConfigs implements DeploymentConfigInterface
type FakeDeploymentConfigs struct {
	Fake *FakeAppsV1
	ns   string
}

var deploymentconfigsResource = v1.SchemeGroupVersion.WithResource("deploymentconfigs")

var deploymentconfigsKind = v1.SchemeGroupVersion.WithKind("DeploymentConfig")

// Get takes name of the deploymentConfig, and returns the corresponding deploymentConfig object, and an error if there is any.
func (c *FakeDeploymentConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deploymentconfigsResource, c.ns, name), &v1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DeploymentConfig), err
}

// List takes label and field selectors, and returns the list of DeploymentConfigs that match those selectors.
func (c *FakeDeploymentConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DeploymentConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deploymentconfigsResource, deploymentconfigsKind, c.ns, opts), &v1.DeploymentConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.DeploymentConfigList{ListMeta: obj.(*v1.DeploymentConfigList).ListMeta}
	for _, item := range obj.(*v1.DeploymentConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deploymentConfigs.
func (c *FakeDeploymentConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deploymentconfigsResource, c.ns, opts))

}

// Create takes the representation of a deploymentConfig and creates it.  Returns the server's representation of the deploymentConfig, and an error, if there is any.
func (c *FakeDeploymentConfigs) Create(ctx context.Context, deploymentConfig *v1.DeploymentConfig, opts metav1.CreateOptions) (result *v1.DeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deploymentconfigsResource, c.ns, deploymentConfig), &v1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DeploymentConfig), err
}

// Update takes the representation of a deploymentConfig and updates it. Returns the server's representation of the deploymentConfig, and an error, if there is any.
func (c *FakeDeploymentConfigs) Update(ctx context.Context, deploymentConfig *v1.DeploymentConfig, opts metav1.UpdateOptions) (result *v1.DeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deploymentconfigsResource, c.ns, deploymentConfig), &v1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DeploymentConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeploymentConfigs) UpdateStatus(ctx context.Context, deploymentConfig *v1.DeploymentConfig, opts metav1.UpdateOptions) (*v1.DeploymentConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deploymentconfigsResource, "status", c.ns, deploymentConfig), &v1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DeploymentConfig), err
}

// Delete takes name of the deploymentConfig and deletes it. Returns an error if one occurs.
func (c *FakeDeploymentConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(deploymentconfigsResource, c.ns, name, opts), &v1.DeploymentConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeploymentConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deploymentconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.DeploymentConfigList{})
	return err
}

// Patch applies the patch and returns the patched deploymentConfig.
func (c *FakeDeploymentConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentconfigsResource, c.ns, name, pt, data, subresources...), &v1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DeploymentConfig), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied deploymentConfig.
func (c *FakeDeploymentConfigs) Apply(ctx context.Context, deploymentConfig *appsv1.DeploymentConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.DeploymentConfig, err error) {
	if deploymentConfig == nil {
		return nil, fmt.Errorf("deploymentConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(deploymentConfig)
	if err != nil {
		return nil, err
	}
	name := deploymentConfig.Name
	if name == nil {
		return nil, fmt.Errorf("deploymentConfig.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentconfigsResource, c.ns, *name, types.ApplyPatchType, data), &v1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DeploymentConfig), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeDeploymentConfigs) ApplyStatus(ctx context.Context, deploymentConfig *appsv1.DeploymentConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.DeploymentConfig, err error) {
	if deploymentConfig == nil {
		return nil, fmt.Errorf("deploymentConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(deploymentConfig)
	if err != nil {
		return nil, err
	}
	name := deploymentConfig.Name
	if name == nil {
		return nil, fmt.Errorf("deploymentConfig.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentconfigsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DeploymentConfig), err
}

// Instantiate takes the representation of a deploymentRequest and creates it.  Returns the server's representation of the deploymentConfig, and an error, if there is any.
func (c *FakeDeploymentConfigs) Instantiate(ctx context.Context, deploymentConfigName string, deploymentRequest *v1.DeploymentRequest, opts metav1.CreateOptions) (result *v1.DeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateSubresourceAction(deploymentconfigsResource, deploymentConfigName, "instantiate", c.ns, deploymentRequest), &v1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DeploymentConfig), err
}

// Rollback takes the representation of a deploymentConfigRollback and creates it.  Returns the server's representation of the deploymentConfig, and an error, if there is any.
func (c *FakeDeploymentConfigs) Rollback(ctx context.Context, deploymentConfigName string, deploymentConfigRollback *v1.DeploymentConfigRollback, opts metav1.CreateOptions) (result *v1.DeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateSubresourceAction(deploymentconfigsResource, deploymentConfigName, "rollback", c.ns, deploymentConfigRollback), &v1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DeploymentConfig), err
}

// GetScale takes name of the deploymentConfig, and returns the corresponding scale object, and an error if there is any.
func (c *FakeDeploymentConfigs) GetScale(ctx context.Context, deploymentConfigName string, options metav1.GetOptions) (result *v1beta1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceAction(deploymentconfigsResource, c.ns, "scale", deploymentConfigName), &v1beta1.Scale{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeDeploymentConfigs) UpdateScale(ctx context.Context, deploymentConfigName string, scale *v1beta1.Scale, opts metav1.UpdateOptions) (result *v1beta1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deploymentconfigsResource, "scale", c.ns, scale), &v1beta1.Scale{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Scale), err
}

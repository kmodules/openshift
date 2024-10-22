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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "kmodules.xyz/openshift/apis/apps/v1"

	corev1 "k8s.io/api/core/v1"
)

// DeploymentStrategyApplyConfiguration represents an declarative configuration of the DeploymentStrategy type for use
// with apply.
type DeploymentStrategyApplyConfiguration struct {
	Type                  *v1.DeploymentStrategyType                          `json:"type,omitempty"`
	CustomParams          *CustomDeploymentStrategyParamsApplyConfiguration   `json:"customParams,omitempty"`
	RecreateParams        *RecreateDeploymentStrategyParamsApplyConfiguration `json:"recreateParams,omitempty"`
	RollingParams         *RollingDeploymentStrategyParamsApplyConfiguration  `json:"rollingParams,omitempty"`
	Resources             *corev1.ResourceRequirements                        `json:"resources,omitempty"`
	Labels                map[string]string                                   `json:"labels,omitempty"`
	Annotations           map[string]string                                   `json:"annotations,omitempty"`
	ActiveDeadlineSeconds *int64                                              `json:"activeDeadlineSeconds,omitempty"`
}

// DeploymentStrategyApplyConfiguration constructs an declarative configuration of the DeploymentStrategy type for use with
// apply.
func DeploymentStrategy() *DeploymentStrategyApplyConfiguration {
	return &DeploymentStrategyApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *DeploymentStrategyApplyConfiguration) WithType(value v1.DeploymentStrategyType) *DeploymentStrategyApplyConfiguration {
	b.Type = &value
	return b
}

// WithCustomParams sets the CustomParams field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CustomParams field is set to the value of the last call.
func (b *DeploymentStrategyApplyConfiguration) WithCustomParams(value *CustomDeploymentStrategyParamsApplyConfiguration) *DeploymentStrategyApplyConfiguration {
	b.CustomParams = value
	return b
}

// WithRecreateParams sets the RecreateParams field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RecreateParams field is set to the value of the last call.
func (b *DeploymentStrategyApplyConfiguration) WithRecreateParams(value *RecreateDeploymentStrategyParamsApplyConfiguration) *DeploymentStrategyApplyConfiguration {
	b.RecreateParams = value
	return b
}

// WithRollingParams sets the RollingParams field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RollingParams field is set to the value of the last call.
func (b *DeploymentStrategyApplyConfiguration) WithRollingParams(value *RollingDeploymentStrategyParamsApplyConfiguration) *DeploymentStrategyApplyConfiguration {
	b.RollingParams = value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *DeploymentStrategyApplyConfiguration) WithResources(value corev1.ResourceRequirements) *DeploymentStrategyApplyConfiguration {
	b.Resources = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *DeploymentStrategyApplyConfiguration) WithLabels(entries map[string]string) *DeploymentStrategyApplyConfiguration {
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *DeploymentStrategyApplyConfiguration) WithAnnotations(entries map[string]string) *DeploymentStrategyApplyConfiguration {
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithActiveDeadlineSeconds sets the ActiveDeadlineSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ActiveDeadlineSeconds field is set to the value of the last call.
func (b *DeploymentStrategyApplyConfiguration) WithActiveDeadlineSeconds(value int64) *DeploymentStrategyApplyConfiguration {
	b.ActiveDeadlineSeconds = &value
	return b
}
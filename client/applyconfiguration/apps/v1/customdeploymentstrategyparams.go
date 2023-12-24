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
	v1 "k8s.io/api/core/v1"
)

// CustomDeploymentStrategyParamsApplyConfiguration represents an declarative configuration of the CustomDeploymentStrategyParams type for use
// with apply.
type CustomDeploymentStrategyParamsApplyConfiguration struct {
	Image       *string     `json:"image,omitempty"`
	Environment []v1.EnvVar `json:"environment,omitempty"`
	Command     []string    `json:"command,omitempty"`
}

// CustomDeploymentStrategyParamsApplyConfiguration constructs an declarative configuration of the CustomDeploymentStrategyParams type for use with
// apply.
func CustomDeploymentStrategyParams() *CustomDeploymentStrategyParamsApplyConfiguration {
	return &CustomDeploymentStrategyParamsApplyConfiguration{}
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *CustomDeploymentStrategyParamsApplyConfiguration) WithImage(value string) *CustomDeploymentStrategyParamsApplyConfiguration {
	b.Image = &value
	return b
}

// WithEnvironment adds the given value to the Environment field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Environment field.
func (b *CustomDeploymentStrategyParamsApplyConfiguration) WithEnvironment(values ...v1.EnvVar) *CustomDeploymentStrategyParamsApplyConfiguration {
	for i := range values {
		b.Environment = append(b.Environment, values[i])
	}
	return b
}

// WithCommand adds the given value to the Command field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Command field.
func (b *CustomDeploymentStrategyParamsApplyConfiguration) WithCommand(values ...string) *CustomDeploymentStrategyParamsApplyConfiguration {
	for i := range values {
		b.Command = append(b.Command, values[i])
	}
	return b
}

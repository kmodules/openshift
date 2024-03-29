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

// ExecNewPodHookApplyConfiguration represents an declarative configuration of the ExecNewPodHook type for use
// with apply.
type ExecNewPodHookApplyConfiguration struct {
	Command       []string    `json:"command,omitempty"`
	Env           []v1.EnvVar `json:"env,omitempty"`
	ContainerName *string     `json:"containerName,omitempty"`
	Volumes       []string    `json:"volumes,omitempty"`
}

// ExecNewPodHookApplyConfiguration constructs an declarative configuration of the ExecNewPodHook type for use with
// apply.
func ExecNewPodHook() *ExecNewPodHookApplyConfiguration {
	return &ExecNewPodHookApplyConfiguration{}
}

// WithCommand adds the given value to the Command field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Command field.
func (b *ExecNewPodHookApplyConfiguration) WithCommand(values ...string) *ExecNewPodHookApplyConfiguration {
	for i := range values {
		b.Command = append(b.Command, values[i])
	}
	return b
}

// WithEnv adds the given value to the Env field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Env field.
func (b *ExecNewPodHookApplyConfiguration) WithEnv(values ...v1.EnvVar) *ExecNewPodHookApplyConfiguration {
	for i := range values {
		b.Env = append(b.Env, values[i])
	}
	return b
}

// WithContainerName sets the ContainerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContainerName field is set to the value of the last call.
func (b *ExecNewPodHookApplyConfiguration) WithContainerName(value string) *ExecNewPodHookApplyConfiguration {
	b.ContainerName = &value
	return b
}

// WithVolumes adds the given value to the Volumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Volumes field.
func (b *ExecNewPodHookApplyConfiguration) WithVolumes(values ...string) *ExecNewPodHookApplyConfiguration {
	for i := range values {
		b.Volumes = append(b.Volumes, values[i])
	}
	return b
}

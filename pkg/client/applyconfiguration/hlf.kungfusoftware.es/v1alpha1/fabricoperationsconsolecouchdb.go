/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// FabricOperationsConsoleCouchDBApplyConfiguration represents a declarative configuration of the FabricOperationsConsoleCouchDB type for use
// with apply.
type FabricOperationsConsoleCouchDBApplyConfiguration struct {
	Image            *string                    `json:"image,omitempty"`
	Tag              *string                    `json:"tag,omitempty"`
	Username         *string                    `json:"username,omitempty"`
	Password         *string                    `json:"password,omitempty"`
	Storage          *StorageApplyConfiguration `json:"storage,omitempty"`
	Resources        *v1.ResourceRequirements   `json:"resources,omitempty"`
	ImagePullSecrets []v1.LocalObjectReference  `json:"imagePullSecrets,omitempty"`
	Affinity         *v1.Affinity               `json:"affinity,omitempty"`
	Tolerations      []v1.Toleration            `json:"tolerations,omitempty"`
	ImagePullPolicy  *v1.PullPolicy             `json:"imagePullPolicy,omitempty"`
}

// FabricOperationsConsoleCouchDBApplyConfiguration constructs a declarative configuration of the FabricOperationsConsoleCouchDB type for use with
// apply.
func FabricOperationsConsoleCouchDB() *FabricOperationsConsoleCouchDBApplyConfiguration {
	return &FabricOperationsConsoleCouchDBApplyConfiguration{}
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *FabricOperationsConsoleCouchDBApplyConfiguration) WithImage(value string) *FabricOperationsConsoleCouchDBApplyConfiguration {
	b.Image = &value
	return b
}

// WithTag sets the Tag field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tag field is set to the value of the last call.
func (b *FabricOperationsConsoleCouchDBApplyConfiguration) WithTag(value string) *FabricOperationsConsoleCouchDBApplyConfiguration {
	b.Tag = &value
	return b
}

// WithUsername sets the Username field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Username field is set to the value of the last call.
func (b *FabricOperationsConsoleCouchDBApplyConfiguration) WithUsername(value string) *FabricOperationsConsoleCouchDBApplyConfiguration {
	b.Username = &value
	return b
}

// WithPassword sets the Password field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Password field is set to the value of the last call.
func (b *FabricOperationsConsoleCouchDBApplyConfiguration) WithPassword(value string) *FabricOperationsConsoleCouchDBApplyConfiguration {
	b.Password = &value
	return b
}

// WithStorage sets the Storage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Storage field is set to the value of the last call.
func (b *FabricOperationsConsoleCouchDBApplyConfiguration) WithStorage(value *StorageApplyConfiguration) *FabricOperationsConsoleCouchDBApplyConfiguration {
	b.Storage = value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *FabricOperationsConsoleCouchDBApplyConfiguration) WithResources(value v1.ResourceRequirements) *FabricOperationsConsoleCouchDBApplyConfiguration {
	b.Resources = &value
	return b
}

// WithImagePullSecrets adds the given value to the ImagePullSecrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImagePullSecrets field.
func (b *FabricOperationsConsoleCouchDBApplyConfiguration) WithImagePullSecrets(values ...v1.LocalObjectReference) *FabricOperationsConsoleCouchDBApplyConfiguration {
	for i := range values {
		b.ImagePullSecrets = append(b.ImagePullSecrets, values[i])
	}
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *FabricOperationsConsoleCouchDBApplyConfiguration) WithAffinity(value v1.Affinity) *FabricOperationsConsoleCouchDBApplyConfiguration {
	b.Affinity = &value
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *FabricOperationsConsoleCouchDBApplyConfiguration) WithTolerations(values ...v1.Toleration) *FabricOperationsConsoleCouchDBApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}

// WithImagePullPolicy sets the ImagePullPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImagePullPolicy field is set to the value of the last call.
func (b *FabricOperationsConsoleCouchDBApplyConfiguration) WithImagePullPolicy(value v1.PullPolicy) *FabricOperationsConsoleCouchDBApplyConfiguration {
	b.ImagePullPolicy = &value
	return b
}

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package v1alpha1

import (
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	deploymentDirectorStorage = builders.NewApiResource( // Resource status endpoint
		deployment.InternalDirector,
		DirectorSchemeFns{},
		func() runtime.Object { return &Director{} },     // Register versioned resource
		func() runtime.Object { return &DirectorList{} }, // Register versioned resource list
		&DirectorStrategy{builders.StorageStrategySingleton},
	)
	deploymentDockerStorage = builders.NewApiResource( // Resource status endpoint
		deployment.InternalDocker,
		DockerSchemeFns{},
		func() runtime.Object { return &Docker{} },     // Register versioned resource
		func() runtime.Object { return &DockerList{} }, // Register versioned resource list
		&DockerStrategy{builders.StorageStrategySingleton},
	)
	deploymentVirtualhostStorage = builders.NewApiResource( // Resource status endpoint
		deployment.InternalVirtualhost,
		VirtualhostSchemeFns{},
		func() runtime.Object { return &Virtualhost{} },     // Register versioned resource
		func() runtime.Object { return &VirtualhostList{} }, // Register versioned resource list
		&VirtualhostStrategy{builders.StorageStrategySingleton},
	)
	ApiVersion = builders.NewApiVersion("deployment.servicefabrik.io", "v1alpha1").WithResources(
		deploymentDirectorStorage,
		builders.NewApiResource( // Resource status endpoint
			deployment.InternalDirectorStatus,
			DirectorSchemeFns{},
			func() runtime.Object { return &Director{} },     // Register versioned resource
			func() runtime.Object { return &DirectorList{} }, // Register versioned resource list
			&DirectorStatusStrategy{builders.StatusStorageStrategySingleton},
		), deploymentDockerStorage,
		builders.NewApiResource( // Resource status endpoint
			deployment.InternalDockerStatus,
			DockerSchemeFns{},
			func() runtime.Object { return &Docker{} },     // Register versioned resource
			func() runtime.Object { return &DockerList{} }, // Register versioned resource list
			&DockerStatusStrategy{builders.StatusStorageStrategySingleton},
		), deploymentVirtualhostStorage,
		builders.NewApiResource( // Resource status endpoint
			deployment.InternalVirtualhostStatus,
			VirtualhostSchemeFns{},
			func() runtime.Object { return &Virtualhost{} },     // Register versioned resource
			func() runtime.Object { return &VirtualhostList{} }, // Register versioned resource list
			&VirtualhostStatusStrategy{builders.StatusStorageStrategySingleton},
		))

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

//
// Director Functions and Structs
//
// +k8s:deepcopy-gen=false
type DirectorSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type DirectorStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type DirectorStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DirectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Director `json:"items"`
}

//
// Docker Functions and Structs
//
// +k8s:deepcopy-gen=false
type DockerSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type DockerStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type DockerStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DockerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Docker `json:"items"`
}

//
// Virtualhost Functions and Structs
//
// +k8s:deepcopy-gen=false
type VirtualhostSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type VirtualhostStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type VirtualhostStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualhostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Virtualhost `json:"items"`
}

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
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	bindDirectorBindStorage = builders.NewApiResource( // Resource status endpoint
		bind.InternalDirectorBind,
		DirectorBindSchemeFns{},
		func() runtime.Object { return &DirectorBind{} },     // Register versioned resource
		func() runtime.Object { return &DirectorBindList{} }, // Register versioned resource list
		&DirectorBindStrategy{builders.StorageStrategySingleton},
	)
	bindDockerBindStorage = builders.NewApiResource( // Resource status endpoint
		bind.InternalDockerBind,
		DockerBindSchemeFns{},
		func() runtime.Object { return &DockerBind{} },     // Register versioned resource
		func() runtime.Object { return &DockerBindList{} }, // Register versioned resource list
		&DockerBindStrategy{builders.StorageStrategySingleton},
	)
	ApiVersion = builders.NewApiVersion("bind.servicefabrik.io", "v1alpha1").WithResources(
		bindDirectorBindStorage,
		builders.NewApiResource( // Resource status endpoint
			bind.InternalDirectorBindStatus,
			DirectorBindSchemeFns{},
			func() runtime.Object { return &DirectorBind{} },     // Register versioned resource
			func() runtime.Object { return &DirectorBindList{} }, // Register versioned resource list
			&DirectorBindStatusStrategy{builders.StatusStorageStrategySingleton},
		), bindDockerBindStorage,
		builders.NewApiResource( // Resource status endpoint
			bind.InternalDockerBindStatus,
			DockerBindSchemeFns{},
			func() runtime.Object { return &DockerBind{} },     // Register versioned resource
			func() runtime.Object { return &DockerBindList{} }, // Register versioned resource list
			&DockerBindStatusStrategy{builders.StatusStorageStrategySingleton},
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
// DirectorBind Functions and Structs
//
// +k8s:deepcopy-gen=false
type DirectorBindSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type DirectorBindStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type DirectorBindStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DirectorBindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DirectorBind `json:"items"`
}

//
// DockerBind Functions and Structs
//
// +k8s:deepcopy-gen=false
type DockerBindSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type DockerBindStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type DockerBindStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DockerBindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DockerBind `json:"items"`
}

// Copyright 2021 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// KogitoRuntimeInterface ...
// +kubebuilder:object:generate=false
type KogitoRuntimeInterface interface {
	KogitoService
	// GetSpec gets the Kogito Service specification structure.
	GetRuntimeSpec() KogitoRuntimeSpecInterface
	// GetStatus gets the Kogito Service Status structure.
	GetRuntimeStatus() KogitoRuntimeStatusInterface
}

// KogitoRuntimeListInterface ...
// +kubebuilder:object:generate=false
type KogitoRuntimeListInterface interface {
	runtime.Object
	// GetItems gets all items
	GetItems() []KogitoRuntimeInterface
}

// KogitoRuntimeSpecInterface ...
// +kubebuilder:object:generate=false
type KogitoRuntimeSpecInterface interface {
	KogitoServiceSpecInterface
	IsEnableIstio() bool
	SetEnableIstio(enableIstio bool)
}

// KogitoRuntimeStatusInterface ...
// +kubebuilder:object:generate=false
type KogitoRuntimeStatusInterface interface {
	KogitoServiceStatusInterface
}

// KogitoRuntimeHandler ...
// +kubebuilder:object:generate=false
type KogitoRuntimeHandler interface {
	FetchKogitoRuntimeInstance(key types.NamespacedName) (KogitoRuntimeInterface, error)
	FetchAllKogitoRuntimeInstances(namespace string) (KogitoRuntimeListInterface, error)
}

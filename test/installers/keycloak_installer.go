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

package installers

import (
	keycloak "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	"github.com/kiegroup/kogito-operator/core/client/kubernetes"
	"github.com/kiegroup/kogito-operator/test/framework"
)

var (
	// keycloakOlmNamespacedInstaller installs Keycloak in the namespace using OLM
	keycloakOlmNamespacedInstaller = OlmNamespacedServiceInstaller{
		subscriptionName:                  "keycloak-operator",
		channel:                           "alpha",
		catalog:                           framework.CommunityCatalog,
		installationTimeoutInMinutes:      10,
		getAllNamespacedOlmCrsInNamespace: getKeycloakCrsInNamespace,
	}
)

// GetKeycloakInstaller returns Keycloak installer
func GetKeycloakInstaller() ServiceInstaller {
	return &keycloakOlmNamespacedInstaller
}

func getKeycloakCrsInNamespace(namespace string) ([]kubernetes.ResourceObject, error) {
	crs := []kubernetes.ResourceObject{}

	keycloaks := &keycloak.KeycloakList{}
	if err := framework.GetObjectsInNamespace(namespace, keycloaks); err != nil {
		return nil, err
	}
	for i := range keycloaks.Items {
		crs = append(crs, &keycloaks.Items[i])
	}

	keycloakClients := &keycloak.KeycloakClientList{}
	if err := framework.GetObjectsInNamespace(namespace, keycloakClients); err != nil {
		return nil, err
	}
	for i := range keycloakClients.Items {
		crs = append(crs, &keycloakClients.Items[i])
	}

	keycloakUsers := &keycloak.KeycloakUserList{}
	if err := framework.GetObjectsInNamespace(namespace, keycloakUsers); err != nil {
		return nil, err
	}
	for i := range keycloakUsers.Items {
		crs = append(crs, &keycloakUsers.Items[i])
	}

	keycloakRealms := &keycloak.KeycloakRealmList{}
	if err := framework.GetObjectsInNamespace(namespace, keycloakRealms); err != nil {
		return nil, err
	}
	for i := range keycloakRealms.Items {
		crs = append(crs, &keycloakRealms.Items[i])
	}

	return crs, nil
}

// Copyright 2020 Red Hat, Inc. and/or its affiliates
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

package infrastructure

import (
	"github.com/kiegroup/kogito-cloud-operator/community-kogito-operator/core/client/kubernetes"
	"github.com/kiegroup/kogito-cloud-operator/community-kogito-operator/core/operator"
	mongodb "github.com/mongodb/mongodb-kubernetes-operator/pkg/apis/mongodb/v1"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MongoDBOperatorName is the MongoDB Operator default name
	MongoDBOperatorName = "mongodb-kubernetes-operator"

	// DefaultMongoDBAuthDatabase is the default authentication database in MongoDB
	DefaultMongoDBAuthDatabase = "admin"
	// DefaultMongoDBPasswordSecretRef is the default key for the secret reference in MongoDB
	DefaultMongoDBPasswordSecretRef = "password"

	// MongoDBKind refers to MongoDB Kind
	MongoDBKind = "MongoDB"

	// MongoDBAppSecretAuthDatabaseKey is the secret authentication database key set in the linked secret for an application
	MongoDBAppSecretAuthDatabaseKey = "auth-database"
	// MongoDBAppSecretDatabaseKey is the secret database key set in the linked secret for an application
	MongoDBAppSecretDatabaseKey = "database"
	// MongoDBAppSecretUsernameKey is the secret username key set in the linked secret for an application
	MongoDBAppSecretUsernameKey = "username"
	// MongoDBAppSecretPasswordKey is the secret password key set in the linked secret for an application
	MongoDBAppSecretPasswordKey = "password"
)

var (
	// MongoDBAPIVersion refers to MongoDB APIVersion
	MongoDBAPIVersion = mongodb.SchemeGroupVersion.String()

	mongoDBServerGroup = mongodb.SchemeGroupVersion.Group
)

// MongoDBHandler ...
type MongoDBHandler interface {
	IsMongoDBAvailable() bool
	IsMongoDBOperatorAvailable(namespace string) (bool, error)
}

type mongoDBHandler struct {
	*operator.Context
}

// NewMongoDBHandler ...
func NewMongoDBHandler(context *operator.Context) MongoDBHandler {
	return &mongoDBHandler{
		context,
	}
}

// IsMongoDBAvailable checks if MongoDB CRD is available in the cluster
func (m *mongoDBHandler) IsMongoDBAvailable() bool {
	return m.Client.HasServerGroup(mongoDBServerGroup)
}

// IsMongoDBOperatorAvailable verify if MongoDB Operator is running in the given namespace and the CRD is available
func (m *mongoDBHandler) IsMongoDBOperatorAvailable(namespace string) (bool, error) {
	m.Log.Debug("Checking if MongoDB Operator is available in the namespace", "namespace", namespace)
	// first check for CRD
	if m.IsMongoDBAvailable() {
		m.Log.Debug("MongoDB CRDs available. Checking if MongoDB Operator is deployed in the namespace", "namespace", namespace)
		// then check if there's an MongoDB Operator deployed
		deployment := &v1.Deployment{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: MongoDBOperatorName}}
		exists := false
		var err error
		if exists, err = kubernetes.ResourceC(m.Client).Fetch(deployment); err != nil {
			return false, nil
		}
		if exists {
			m.Log.Debug("MongoDB Operator is available in the namespace", "namespace", namespace)
			return true, nil
		}
	} else {
		m.Log.Debug("Couldn't find MongoDB CRDs")
	}
	m.Log.Debug("Looks like MongoDB Operator is not available in the namespace", "namespace", namespace)
	return false, nil
}

// +build !ignore_autogenerated

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
// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kiegroup/kogito-operator/api/v1beta1.Builds":                        schema_kiegroup_kogito_operator_api_v1beta1_Builds(ref),
		"github.com/kiegroup/kogito-operator/api/v1beta1.GitSource":                     schema_kiegroup_kogito_operator_api_v1beta1_GitSource(ref),
		"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoBuild":                   schema_kiegroup_kogito_operator_api_v1beta1_KogitoBuild(ref),
		"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoBuildStatus":             schema_kiegroup_kogito_operator_api_v1beta1_KogitoBuildStatus(ref),
		"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoInfra":                   schema_kiegroup_kogito_operator_api_v1beta1_KogitoInfra(ref),
		"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoInfraSpec":               schema_kiegroup_kogito_operator_api_v1beta1_KogitoInfraSpec(ref),
		"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoInfraStatus":             schema_kiegroup_kogito_operator_api_v1beta1_KogitoInfraStatus(ref),
		"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoRuntime":                 schema_kiegroup_kogito_operator_api_v1beta1_KogitoRuntime(ref),
		"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoSupportingService":       schema_kiegroup_kogito_operator_api_v1beta1_KogitoSupportingService(ref),
		"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoSupportingServiceSpec":   schema_kiegroup_kogito_operator_api_v1beta1_KogitoSupportingServiceSpec(ref),
		"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoSupportingServiceStatus": schema_kiegroup_kogito_operator_api_v1beta1_KogitoSupportingServiceStatus(ref),
		"github.com/kiegroup/kogito-operator/api/v1beta1.WebHookSecret":                 schema_kiegroup_kogito_operator_api_v1beta1_WebHookSecret(ref),
	}
}

func schema_kiegroup_kogito_operator_api_v1beta1_Builds(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Builds ...",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"new": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Builds are being created.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"pending": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Builds are about to start running.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"running": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Builds are running.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"complete": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Builds have executed and succeeded.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"failed": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Builds have executed and failed.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"error": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Builds have been prevented from executing by an error.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"cancelled": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Builds have been stopped from executing.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func schema_kiegroup_kogito_operator_api_v1beta1_GitSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GitSource Git coordinates to locate the source code to build.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"uri": {
						SchemaProps: spec.SchemaProps{
							Description: "Git URI for the s2i source.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"reference": {
						SchemaProps: spec.SchemaProps{
							Description: "Branch to use in the Git repository.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"contextDir": {
						SchemaProps: spec.SchemaProps{
							Description: "Context/subdirectory where the code is located, relative to the repo root.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"uri"},
			},
		},
	}
}

func schema_kiegroup_kogito_operator_api_v1beta1_KogitoBuild(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoBuild handles how to build a custom Kogito service in a Kubernetes/OpenShift cluster.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kiegroup/kogito-operator/api/v1beta1.KogitoBuildSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kiegroup/kogito-operator/api/v1beta1.KogitoBuildStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoBuildSpec", "github.com/kiegroup/kogito-operator/api/v1beta1.KogitoBuildStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_kiegroup_kogito_operator_api_v1beta1_KogitoBuildStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoBuildStatus defines the observed state of KogitoBuild.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "History of conditions for the resource",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.Condition"),
									},
								},
							},
						},
					},
					"latestBuild": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"builds": {
						SchemaProps: spec.SchemaProps{
							Description: "History of builds",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kiegroup/kogito-operator/api/v1beta1.Builds"),
						},
					},
				},
				Required: []string{"conditions", "builds"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-operator/api/v1beta1.Builds", "k8s.io/apimachinery/pkg/apis/meta/v1.Condition"},
	}
}

func schema_kiegroup_kogito_operator_api_v1beta1_KogitoInfra(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoInfra is the resource to bind a Custom Resource (CR) not managed by Kogito Operator to a given deployed Kogito service. It holds the reference of a CR managed by another operator such as Strimzi. For example: one can create a Kafka CR via Strimzi and link this resource using KogitoInfra to a given Kogito service (custom or supporting, such as Data Index). Please refer to the Kogito Operator documentation (https://docs.jboss.org/kogito/release/latest/html_single/) for more information.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kiegroup/kogito-operator/api/v1beta1.KogitoInfraSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kiegroup/kogito-operator/api/v1beta1.KogitoInfraStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoInfraSpec", "github.com/kiegroup/kogito-operator/api/v1beta1.KogitoInfraStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_kiegroup_kogito_operator_api_v1beta1_KogitoInfraSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoInfraSpec defines the desired state of KogitoInfra.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"resource": {
						SchemaProps: spec.SchemaProps{
							Description: "Resource for the service. Example: Infinispan/Kafka/Keycloak.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kiegroup/kogito-operator/api/v1beta1.Resource"),
						},
					},
					"infraProperties": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-map-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Optional properties which would be needed to setup correct runtime/service configuration, based on the resource type. For example, MongoDB will require `username` and `database` as properties for a correct setup, else it will fail",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"resource"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-operator/api/v1beta1.Resource"},
	}
}

func schema_kiegroup_kogito_operator_api_v1beta1_KogitoInfraStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoInfraStatus defines the observed state of KogitoInfra.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "History of conditions for the resource",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.Condition"),
									},
								},
							},
						},
					},
					"runtimeProperties": {
						SchemaProps: spec.SchemaProps{
							Description: "Runtime variables extracted from the linked resource that will be added to the deployed Kogito service.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/kiegroup/kogito-operator/api/v1beta1.RuntimeProperties"),
									},
								},
							},
						},
					},
					"volumes": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "List of volumes that should be added to the services bound to this infra instance",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/kiegroup/kogito-operator/api/v1beta1.KogitoInfraVolume"),
									},
								},
							},
						},
					},
				},
				Required: []string{"conditions"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoInfraVolume", "github.com/kiegroup/kogito-operator/api/v1beta1.RuntimeProperties", "k8s.io/apimachinery/pkg/apis/meta/v1.Condition"},
	}
}

func schema_kiegroup_kogito_operator_api_v1beta1_KogitoRuntime(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoRuntime is a custom Kogito service.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kiegroup/kogito-operator/api/v1beta1.KogitoRuntimeSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kiegroup/kogito-operator/api/v1beta1.KogitoRuntimeStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoRuntimeSpec", "github.com/kiegroup/kogito-operator/api/v1beta1.KogitoRuntimeStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_kiegroup_kogito_operator_api_v1beta1_KogitoSupportingService(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoSupportingService deploys the Supporting service in the given namespace.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kiegroup/kogito-operator/api/v1beta1.KogitoSupportingServiceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kiegroup/kogito-operator/api/v1beta1.KogitoSupportingServiceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoSupportingServiceSpec", "github.com/kiegroup/kogito-operator/api/v1beta1.KogitoSupportingServiceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_kiegroup_kogito_operator_api_v1beta1_KogitoSupportingServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoSupportingServiceSpec defines the desired state of KogitoSupportingService.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Number of replicas that the service will have deployed in the cluster. Default value: 1.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"env": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Environment variables to be added to the runtime container. Keys must be a C_IDENTIFIER.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image definition for the service. Example: \"quay.io/kiegroup/kogito-service:latest\". On OpenShift an ImageStream will be created in the current namespace pointing to the given image.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"insecureImageRegistry": {
						SchemaProps: spec.SchemaProps{
							Description: "A flag indicating that image streams created by Kogito Operator should be configured to allow pulling from insecure registries. Usable just on OpenShift. Defaults to 'false'.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Defined compute resource requirements for the deployed service.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"deploymentLabels": {
						SchemaProps: spec.SchemaProps{
							Description: "Additional labels to be added to the Deployment and Pods managed by the operator.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"serviceLabels": {
						SchemaProps: spec.SchemaProps{
							Description: "Additional labels to be added to the Service managed by the operator.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"propertiesConfigMap": {
						SchemaProps: spec.SchemaProps{
							Description: "Custom ConfigMap with application.properties file to be mounted for the Kogito service. The ConfigMap must be created in the same namespace. Use this property if you need custom properties to be mounted before the application deployment. If left empty, one will be created for you. Later it can be updated to add any custom properties to apply to the service.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"infra": {
						SchemaProps: spec.SchemaProps{
							Description: "Infra provides list of dependent KogitoInfra objects.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"monitoring": {
						SchemaProps: spec.SchemaProps{
							Description: "Create Service monitor instance to connect with Monitoring service",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kiegroup/kogito-operator/api/v1beta1.Monitoring"),
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Description: "Application properties that will be set to the service. For example 'MY_VAR: my_value'.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"probes": {
						SchemaProps: spec.SchemaProps{
							Description: "Configure liveness, readiness and startup probes for containers",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kiegroup/kogito-operator/api/v1beta1.KogitoProbe"),
						},
					},
					"trustStore": {
						SchemaProps: spec.SchemaProps{
							Description: "Custom TrustStore that will be used by this service to make calls to TLS endpoints",
							Ref:         ref("github.com/kiegroup/kogito-operator/api/v1beta1.TLSKeyStore"),
						},
					},
					"serviceType": {
						SchemaProps: spec.SchemaProps{
							Description: "Defines the type for the supporting service, eg: DataIndex, JobsService Default value: JobsService",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"serviceType"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoProbe", "github.com/kiegroup/kogito-operator/api/v1beta1.Monitoring", "github.com/kiegroup/kogito-operator/api/v1beta1.TLSKeyStore", "k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_kiegroup_kogito_operator_api_v1beta1_KogitoSupportingServiceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KogitoSupportingServiceStatus defines the observed state of KogitoSupportingService.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "History of conditions for the resource",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.Condition"),
									},
								},
							},
						},
					},
					"deploymentConditions": {
						SchemaProps: spec.SchemaProps{
							Description: "General conditions for the Kogito Service deployment.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/apps/v1.DeploymentCondition"),
									},
								},
							},
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image is the resolved image for this service.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"externalURI": {
						SchemaProps: spec.SchemaProps{
							Description: "URI is where the service is exposed.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"cloudEvents": {
						SchemaProps: spec.SchemaProps{
							Description: "Describes the CloudEvents that this instance can consume or produce",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kiegroup/kogito-operator/api/v1beta1.KogitoCloudEventsStatus"),
						},
					},
				},
				Required: []string{"conditions"},
			},
		},
		Dependencies: []string{
			"github.com/kiegroup/kogito-operator/api/v1beta1.KogitoCloudEventsStatus", "k8s.io/api/apps/v1.DeploymentCondition", "k8s.io/apimachinery/pkg/apis/meta/v1.Condition"},
	}
}

func schema_kiegroup_kogito_operator_api_v1beta1_WebHookSecret(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WebHookSecret Secret to use for a given webHook.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "WebHook type, either GitHub or Generic.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret value for webHook",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

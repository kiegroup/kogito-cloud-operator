// Copyright 2019 Red Hat, Inc. and/or its affiliates
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

package resource

import (
	"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1"
	"github.com/kiegroup/kogito-cloud-operator/pkg/client"
	"github.com/kiegroup/kogito-cloud-operator/pkg/client/meta"
	"github.com/kiegroup/kogito-cloud-operator/pkg/infrastructure"
	"github.com/kiegroup/kogito-cloud-operator/pkg/util"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"path"
	"strconv"
)

func newStatefulset(instance *v1alpha1.KogitoDataIndex, secret *corev1.Secret, externalURI string, cli *client.Client) (*appsv1.StatefulSet, error) {
	// define the http port

	httpPort := defineDataIndexHTTPPort(instance)
	log.Debugf("Method defineDataIndexHTTPPort(instance) returned http port number %i", httpPort)

	// create a standard probe
	probe := defaultProbe
	probe.Handler.TCPSocket = &corev1.TCPSocketAction{Port: intstr.IntOrString{IntVal: httpPort}}
	// environment variables
	removeManagedEnvVars(instance)
	// set KOGITO_DATA_INDEX_HTTP_PORT env
	hasEnv := false
	for envName := range instance.Spec.Env {
		if envName == dataIndexEnvKeyHTTPPort {
			hasEnv = true
		}
	}
	if len(instance.Spec.Env) == 0 || !hasEnv {
		envMap := make(map[string]string)
		envMap[dataIndexEnvKeyHTTPPort] = util.FormatInt32ToString(httpPort)
		log.Debugf("Has no %s Env, adding it to spec %s", dataIndexEnvKeyHTTPPort, envMap)
		instance.Spec.Env = util.AppendStringMap(instance.Spec.Env, envMap)
	}
	// from cr
	envs := instance.Spec.Env
	envs = util.AppendStringMap(envs, fromInfinispanToStringMap(instance.Spec.Infinispan))
	envs = util.AppendStringMap(envs, fromKafkaToStringMap(externalURI))

	if instance.Spec.Replicas == 0 {
		instance.Spec.Replicas = defaultReplicas
	}
	if len(instance.Spec.Image) == 0 {
		instance.Spec.Image = DefaultImage
	}

	statefulset := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name,
			Namespace: instance.Namespace,
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: &instance.Spec.Replicas,
			Selector: &metav1.LabelSelector{},
			UpdateStrategy: appsv1.StatefulSetUpdateStrategy{
				Type: appsv1.RollingUpdateStatefulSetStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            instance.Name,
							Image:           instance.Spec.Image,
							Env:             util.FromMapToEnvVar(envs),
							Resources:       extractResources(instance),
							ImagePullPolicy: corev1.PullAlways,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: httpPort,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							LivenessProbe:  probe,
							ReadinessProbe: probe,
						},
					},
				},
			},
		},
	}

	if err := mountProtoBufConfigMaps(statefulset, cli); err != nil {
		return nil, err
	}
	setInfinispanCredentialsSecret(instance.Spec.Infinispan, secret, &statefulset.Spec.Template.Spec.Containers[0])
	meta.SetGroupVersionKind(&statefulset.TypeMeta, meta.KindStatefulSet)
	addDefaultMetadata(&statefulset.ObjectMeta, instance)
	addDefaultMetadata(&statefulset.Spec.Template.ObjectMeta, instance)
	statefulset.Spec.Selector.MatchLabels = statefulset.Labels

	log.Infof("Stateful set ports ", statefulset.Spec.Template.Spec.Containers[0].Ports)
	return statefulset, nil
}

// mountProtoBufConfigMaps mounts protobuf configMaps from KogitoApps into the given stateful set
func mountProtoBufConfigMaps(statefulset *appsv1.StatefulSet, cli *client.Client) (err error) {
	var cms *corev1.ConfigMapList
	if cms, err = infrastructure.GetProtoBufConfigMaps(statefulset.Namespace, cli); err != nil {
		return err
	}

	for _, cm := range cms.Items {
		statefulset.Spec.Template.Spec.Volumes =
			append(statefulset.Spec.Template.Spec.Volumes, corev1.Volume{
				Name: cm.Name,
				VolumeSource: corev1.VolumeSource{
					ConfigMap: &corev1.ConfigMapVolumeSource{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: cm.Name,
						},
					},
				},
			})
		statefulset.Spec.Template.Spec.Containers[0].VolumeMounts =
			append(statefulset.Spec.Template.Spec.Containers[0].VolumeMounts, corev1.VolumeMount{Name: cm.Name, MountPath: path.Join(defaultProtobufMountPath, cm.Labels["app"])})
	}
	if len(statefulset.Spec.Template.Spec.Volumes) > 0 {
		for k, v := range protoBufEnvs {
			util.SetEnvVar(k, v, &statefulset.Spec.Template.Spec.Containers[0])
		}
	} else {
		for _, v := range protoBufKeys {
			util.SetEnvVar(v, "", &statefulset.Spec.Template.Spec.Containers[0])
		}
	}

	return nil
}

// defineDataIndexHTTPPort will define which port the dataindex should be listening to. To set it use httpPort cr parameter.
// defaults to 8080
func defineDataIndexHTTPPort(instance *v1alpha1.KogitoDataIndex) int32 {
	// first check if the env KOGITO_DATA_INDEX_HTTP_PORT is set
	for env, value := range instance.Spec.Env {
		if env == "KOGITO_DATA_INDEX_HTTP_PORT" {
			parsedPort, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				log.Warnf("Failed to parse provided port via Env %s, using the default.", value)
				return defaultExposedPort
			}
			return int32(parsedPort)
		}
	}

	// port should be greater than 0
	if instance.Spec.HTTPPort < 1 {
		log.Debugf("HTTPPort not set, returning default http port.")
		return defaultExposedPort
	} else {
		log.Debugf("HTTPPort is set, returning port number %i", int(instance.Spec.HTTPPort))
		return instance.Spec.HTTPPort
	}
}

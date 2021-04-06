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

package framework

import (
	"fmt"
	"github.com/kiegroup/kogito-operator/core/infrastructure/kafka/v1beta2"

	"github.com/kiegroup/kogito-operator/core/client/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeployKafkaInstance deploys an instance of Kafka
func DeployKafkaInstance(namespace string, kafka *v1beta2.Kafka) error {
	GetLogger(namespace).Info("Creating Kafka instance %s.", "name", kafka.Name)

	if err := kubernetes.ResourceC(kubeClient).Create(kafka); err != nil {
		return fmt.Errorf("Error while creating Kafka: %v ", err)
	}

	return nil
}

// DeployKafkaTopic deploys a Kafka topic
func DeployKafkaTopic(namespace, kafkaTopicName, kafkaInstanceName string) error {
	GetLogger(namespace).Info("Creating Kafka", "topic", kafkaTopicName, "instanceName", kafkaInstanceName)

	kafkaTopic := &v1beta2.KafkaTopic{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      kafkaTopicName,
			Labels:    map[string]string{"strimzi.io/cluster": kafkaInstanceName},
		},
		Spec: v1beta2.KafkaTopicSpec{
			Replicas:   1,
			Partitions: 1,
		},
	}

	if err := kubernetes.ResourceC(kubeClient).Create(kafkaTopic); err != nil {
		return fmt.Errorf("Error while creating Kafka Topic: %v ", err)
	}

	return nil
}

// ScaleKafkaInstanceDown scales a Kafka instance down by killing its pod temporarily
func ScaleKafkaInstanceDown(namespace, kafkaInstanceName string) error {
	GetLogger(namespace).Info("Scaling Kafka Instance down", "instance name", kafkaInstanceName)
	pods, err := GetPodsWithLabels(namespace, map[string]string{"strimzi.io/name": kafkaInstanceName + "-kafka"})
	if err != nil {
		return err
	} else if len(pods.Items) != 1 {
		return fmt.Errorf("Kafka instance should have just one kafka pod running")
	}
	if err = DeleteObject(&pods.Items[0]); err != nil {
		return fmt.Errorf("Error scaling Kafka instance down by deleting a kafka pod. The nested error is: %v", err)
	}

	return nil
}

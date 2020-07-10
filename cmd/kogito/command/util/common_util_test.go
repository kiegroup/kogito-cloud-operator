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

package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ValidateKeyPair_CorrectInput(t *testing.T) {
	keyValuePair := []string{
		"VAR1=key1",
		"VAR2=key2",
	}

	err := CheckKeyPair(keyValuePair)
	assert.NoError(t, err)
}

func Test_ValidateKeyPair_InCorrectInput(t *testing.T) {
	keyValuePair := []string{
		"VAR1=key1",
		"VAR2",
	}

	err := CheckKeyPair(keyValuePair)
	assert.Error(t, err)
}

func Test_ValidateSecretEnvVar_CorrectInput(t *testing.T) {
	keyValuePair := []string{
		"VAR1=secretName1#secretKey1",
		"VAR2=secretName2#secretKey2",
	}

	err := CheckSecretKeyPair(keyValuePair)
	assert.NoError(t, err)
}

func Test_ValidateSecretEnvVar_InCorrectInput(t *testing.T) {
	keyValuePair := []string{
		"VAR1=secretName1#secretKey1",
		"VAR2=secretName2@secretKey2",
	}

	err := CheckSecretKeyPair(keyValuePair)
	assert.Error(t, err)
}

func Test_ValidateImageTag_CorrectInput(t *testing.T) {
	image := "quay.io/kiegroup/data_index:1.0"
	err := CheckImageTag(image)
	assert.NoError(t, err)
}

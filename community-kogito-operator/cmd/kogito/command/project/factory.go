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

package project

import (
	"github.com/kiegroup/kogito-cloud-operator/community-kogito-operator/cmd/kogito/command/context"
	"github.com/spf13/cobra"
)

// BuildCommands creates the commands available in this package
func BuildCommands(ctx *context.CommandContext, rootCommand *cobra.Command) {
	initDeleteProjectCommand(ctx, rootCommand)
	initNewProjectCommand(ctx, rootCommand)
	initUseProjectCommand(ctx, rootCommand)
	initDisplayProjectCommand(ctx, rootCommand)
}

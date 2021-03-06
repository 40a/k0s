/*
Copyright 2021 Mirantis, Inc.

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
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	installServerCmd = &cobra.Command{
		Use:   "server",
		Short: "Helper command for setting up k0s as server node on a brand-new system. Must be run as root (or with sudo)",
		Example: `All default values of server command will be passed to the service stub unless overriden. 

With server subcommand you can setup a single node cluster by running:

	k0s install server --enable-worker
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			flagsAndVals := []string{"server"}
			cmd.Flags().Visit(func(f *pflag.Flag) {
				flagsAndVals = append(flagsAndVals, fmt.Sprintf("--%s=%s", f.Name, f.Value.String()))
			})

			return setup("server", flagsAndVals)
		},
	}
)

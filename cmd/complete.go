/*
Copyright The Helm Authors.

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
	"io"

	v2 "github.com/helm/helm-2to3/pkg/v2"
	"github.com/spf13/cobra"
)

func newCompleteCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:                "__complete",
		Hidden:             true,
		DisableFlagParsing: true,
		RunE:               runComplete,
	}

	return cmd
}

func runComplete(cmd *cobra.Command, args []string) error {
	for _, arg := range args {
		if arg == "convert" {
			retrieveOptions := v2.RetrieveOptions{
				// TODO support these overrides by parsing flags and setting these values
				//				TillerNamespace:  convertOptions.TillerNamespace,
				//				TillerLabel:      convertOptions.TillerLabel,
				//				TillerOutCluster: convertOptions.TillerOutCluster,
				//				StorageType:      convertOptions.StorageType,
			}
			v2Releases, err := v2.GetReleaseVersions(retrieveOptions)
			if err != nil {
				return err
			}
			for _, release := range v2Releases {
				fmt.Println(release.GetName())
			}
			return nil
		}
	}
	return nil
}

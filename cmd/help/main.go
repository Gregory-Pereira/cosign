// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/sigstore/cosign/v2/cmd/cosign/cli"
	"github.com/sigstore/cosign/v2/cmd/cosign/cli/templates"
	errors "github.com/sigstore/cosign/v2/cmd/cosign/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func main() {
	var dir string
	root := &cobra.Command{
		Use:          "gendoc",
		Short:        "Generate cosign's help docs",
		SilenceUsage: true,
		Args:         cobra.NoArgs,
		RunE: func(*cobra.Command, []string) error {
			err := errors.GenerateExitCodeDocs(dir)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			return doc.GenMarkdownTree(cli.New(), dir)
		},
	}
	root.Flags().StringVarP(&dir, "dir", "d", "doc", "Path to directory in which to generate docs")

	templates.SetCustomUsageFunc(root)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
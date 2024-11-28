// Copyright 2024 Shiqinfeng &lt;150627601@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/shiqinfeng1/gomono-layout/internal/gomonoctl/cmd/project"
)

var rootCmd = &cobra.Command{
	Use:   "gomonoctl",
	Short: "gomonoctl: An elegant toolkit for Go microservices.",
	Long:  `gomonoctl: An elegant toolkit for Go microservices.`,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

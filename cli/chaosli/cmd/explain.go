// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025 Datadog, Inc.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "explains disruption config",
	Long:  `translates the yaml of the disruption configuration to english.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		explanation(path)
	},
}

func explanation(path string) {
	disruption := ReadUnmarshalValidate(path)

	expl := disruption.Spec.Explain()

	for _, e := range expl {
		fmt.Println(e)
	}
}

func init() {
	explainCmd.Flags().String("path", "", "The path to the disruption file to be explained.")

	err := explainCmd.MarkFlagRequired("path")
	if err != nil {
		return
	}
}

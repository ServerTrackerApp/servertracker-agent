/*
 * Copyright (c) 2024 Luca Fröschke
 */

package cmd

import (
	"github.com/spf13/cobra"
	"go.servertracker.net/agent/log"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the agent",
	Run: func(cmd *cobra.Command, args []string) {
		log.Log("Starting agent", log.INFO)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

/*
 * Copyright (c) 2024 Luca Fröschke
 */

package cmd

import (
	"github.com/spf13/cobra"
	"go.servertracker.net/agent/config"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "servertracker",
	Short: "ServerTracker agent for collecting system metrics",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config.Init()
}

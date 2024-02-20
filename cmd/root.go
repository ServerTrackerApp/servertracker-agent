/*
 * Copyright (c) 2024 Luca Fröschke
 */

package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
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
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath("/etc/servertracker/")
	viper.AddConfigPath(home)
	viper.SetConfigName("agent")
	viper.SetConfigType("yml")

	viper.SetDefault("server_id", "")
	viper.SetDefault("auth_key", "")
	viper.SetDefault("interval", 60)

	viper.SetDefault("plugin_location", "/etc/servertracker/plugins")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	if _, err := os.Stat(viper.ConfigFileUsed()); os.IsNotExist(err) {
		fmt.Println("Config file not found, creating it...")
		fmt.Println("Config file created at:", viper.ConfigFileUsed())
		err := viper.SafeWriteConfig()
		if err != nil {
			fmt.Println("Error creating config file:", err)
		}
	}
}

/*
 * Copyright (c) 2024 Luca Fröschke
 */

package config

import "github.com/spf13/viper"

const (
	DefaultLocation = "/etc/servertracker/agent.yml"
	PluginLocation  = "/etc/servertracker/plugins"
)

type Config struct {
	ServerId       string `yaml:"server_id"`
	AuthKey        string `yaml:"auth_key"`
	Interval       int    `yaml:"interval"`
	PluginLocation string `yaml:"plugin_location"`
}

func NewConfig() *Config {
	return &Config{
		Interval:       60,
		PluginLocation: PluginLocation,
	}
}

func FromViper() *Config {
	return &Config{
		ServerId:       viper.GetString("server_id"),
		AuthKey:        viper.GetString("auth_key"),
		Interval:       viper.GetInt("interval"),
		PluginLocation: viper.GetString("plugin_location"),
	}
}

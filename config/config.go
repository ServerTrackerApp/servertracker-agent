/*
 * Copyright (c) 2024 Luca Fröschke
 */

package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"sync"
)

var (
	config *Config
	once   sync.Once
)

const (
	Location       = "/etc/servertracker/agent.yml"
	PluginLocation = "/etc/servertracker/plugins"
)

type Config struct {
	ServerId       string                    `yaml:"server_id"`
	AuthKey        string                    `yaml:"auth_key"`
	SyncInterval   int                       `yaml:"sync_interval"`
	PluginLocation string                    `yaml:"plugin_location"`
	PluginSettings map[string]PluginSettings `yaml:"plugin_settings"`
}

type PluginSettings struct {
	PluginName string
	Settings   map[string]interface{}
}

func newConfig() *Config {
	return &Config{
		ServerId:       "",
		AuthKey:        "",
		SyncInterval:   60,
		PluginLocation: PluginLocation,
		PluginSettings: make(map[string]PluginSettings),
	}
}

func Init() {
	once.Do(func() {
		filePath := Location
		dirPath := filepath.Dir(filePath)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			err = os.MkdirAll(dirPath, 0755)
			if err != nil {
				panic(err)
			}
		}
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			defaultConfig := newConfig()
			configBytes, _ := yaml.Marshal(defaultConfig)
			err = os.WriteFile(filePath, configBytes, 0644)
			if err != nil {
				panic(err)
			}
			config = defaultConfig
		} else {
			configBytes, err := os.ReadFile(filePath)
			if err != nil {
				panic(err)
			}
			err = yaml.Unmarshal(configBytes, &config)
			if err != nil {
				panic(err)
			}
		}
	})
}

func GetConfig() *Config {
	return config
}

func SetConfig(c *Config) {
	config = c
}

func Save() {
	filePath := Location
	configBytes, _ := yaml.Marshal(config)
	err := os.WriteFile(filePath, configBytes, 0644)
	if err != nil {
		panic(err)
	}
}

func GetPluginSettings(pluginName string) (PluginSettings, bool) {
	settings, ok := config.PluginSettings[pluginName]
	return settings, ok
}

func SetPluginSettings(pluginName string, settings PluginSettings) {
	config.PluginSettings[pluginName] = settings
}

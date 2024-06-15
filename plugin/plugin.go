/*
 * Copyright (c) 2024 Luca Fröschke
 */

package plugin

import (
	"fmt"
	"go.servertracker.net/agent/config"
	"go.servertracker.net/agent/datastore"
	"go.servertracker.net/agent/log"
	"os"
	"path/filepath"
	"plugin"
	"sync"
)

var (
	Plugins = make(map[string]Plugin)
)

type Plugin interface {
	Name() string
	Init(settings config.PluginSettings, ds *datastore.DataStore) error
	Run() error
	Exit() error
}

func LoadPlugins() {
	pluginLocation := config.GetConfig().PluginLocation

	if _, err := os.Stat(pluginLocation); os.IsNotExist(err) {
		err = os.MkdirAll(pluginLocation, 0755)
		if err != nil {
			log.Log(fmt.Sprintf("Failed to create plugin directory: %s", err), log.ERROR)
			return
		}
	}

	files, err := os.ReadDir(pluginLocation)
	if err != nil {
		log.Log(fmt.Sprintf("Failed to read plugin directory: %s", err), log.ERROR)
		return
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".so" {
			continue
		}

		pluginPath := filepath.Join(pluginLocation, file.Name())
		p, err := plugin.Open(pluginPath)
		if err != nil {
			log.Log(fmt.Sprintf("Failed to open plugin: %s", err), log.ERROR)
			continue
		}

		symPlugin, err := p.Lookup("Plugin")
		if err != nil {
			log.Log(fmt.Sprintf("Failed to lookup symbol: %s", err), log.ERROR)
			continue
		}

		var plug Plugin
		plug, ok := symPlugin.(Plugin)
		if !ok {
			log.Log("Plugin does not implement the Plugin interface", log.ERROR)
			continue
		}

		settings, _ := config.GetPluginSettings(plug.Name())
		err = plug.Init(settings, datastore.GetInstance())
		if err != nil {
			log.Log(fmt.Sprintf("Failed to initialize plugin: %s", err), log.ERROR)
			continue
		}

		Plugins[plug.Name()] = plug
	}
}

func StartPlugins() {
	var wg sync.WaitGroup

	for _, plug := range Plugins {
		log.Log(fmt.Sprintf("Starting plugin: %s", plug.Name()), log.INFO)
		wg.Add(1)
		go func(plug Plugin) {
			defer wg.Done()
			err := plug.Run()
			if err != nil {
				log.Log(fmt.Sprintf("Failed to run plugin: %s", err), log.ERROR)
			}
		}(plug)
	}

	wg.Wait()
}

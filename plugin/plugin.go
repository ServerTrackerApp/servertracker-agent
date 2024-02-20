/*
 * Copyright (c) 2024 Luca Fröschke
 */

package plugin

import "go.servertracker.net/agent/config"

type Plugin interface {
	Name() string
	Init(config *config.Config) error
	Run() error
	Exit() error
}

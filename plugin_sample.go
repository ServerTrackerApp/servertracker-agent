//go:build ignore

/*
 * Copyright (c) 2024 Luca Fröschke
 */

package main

type SamplePlugin struct{}

func (p *SamplePlugin) Name() string {
	return "servertracker.sample_plugin"

}

func (p *SamplePlugin) Init() error {
	fmt.Println("Initializing SamplePlugin")
	return nil
}

func (p *SamplePlugin) Run() error {
	fmt.Println("Running SamplePlugin")
	return nil
}

func (p *SamplePlugin) Exit() error {
	fmt.Println("Exiting SamplePlugin")
	return nil
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

type config struct {
	Module struct {
		Enabled      bool
		moduleConfig `mapstructure:",squash"`
	}
}

type moduleConfig struct {
	Token string
}

var (
	cfgFile string
	C       config
)

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		mytest()
	},
}

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

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	testCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

func mytest() {
	viper.SetConfigName(".foamloon")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
	}

	if C.Module.Enabled {
		// Do something with C.Module.Token
		fmt.Println("Module enabled. Token:", C.Module.Token)
	} else {
		fmt.Println("Module disabled.")
	}

	// if err := rootCmd.Execute(); err != nil {
	// 	fmt.Println(err)
	// }
}

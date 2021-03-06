package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cfgFileName string = ".orion"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "orion",
	Short: "Orion is an opinionated tool for configuring cloud providers with terraform templates.",
	Long: `


	 _____                                 
	/\  __'\         __                    
	\ \ \/\ \  _ __ /\_\    ___     ___    
	 \ \ \ \ \/\''__\/\ \  / __'\ /' _ '\  
	  \ \ \_\ \ \ \/ \ \ \/\ \L\ \/\ \/\ \ 
	   \ \_____\ \_\  \ \_\ \____/\ \_\ \_\
	    \/_____/\/_/   \/_/\/___/  \/_/\/_/
										   
										   	
                                                  
This project is an attempt to combine and share best practices when building production 
ready cloud native managed service solutions. Orion's infrastructure turn-key starter 
templates are based on real world engagements with enterprise customers.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.orion)")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".orion" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(cfgFileName)
	}

	if err := viper.ReadInConfig(); err != nil {
		// TODO: Enable this if we start using viper
		// fmt.Println("Can't read config:", err)
		// os.Exit(1)
	}
}

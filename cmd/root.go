/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/jiwoniy/otmk-kipris-collector/nice"
	"github.com/jiwoniy/otmk-kipris-collector/app"
	"github.com/jiwoniy/otmk-kipris-collector/kipris/collector"
	"github.com/jiwoniy/otmk-kipris-collector/kipris/types"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "otmk-kipris-collector",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var appCfg types.ApplicationConfig

		cfgData, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(cfgData, &appCfg); err != nil {
			panic(err)
		}

		var config types.CollectorConfig
		mode := "dev"
		if len(args) < 1 {
			config = appCfg.Dev
		} else {
			env := args[0]
			switch env {
			case "prod":
				config = appCfg.Prod
				mode = "prod"
			case "dev":
				config = appCfg.Dev
			case "test":
				config = appCfg.Test
				mode = "test"
			default:
				config = appCfg.Dev
			}
		}

		collectorInstance, err := collector.New(config)

		if err != nil {
			panic(err)
		}

		application := app.NewApplication(collectorInstance)
		restConfig := types.RestConfig{
			ListenAddr: config.ListenAddr,
		}
		app.StartApplication(application, mode, restConfig)

		// nice import
		// db, err := gorm.Open("mysql", "nice_admin:Nice0518!@@(61.97.187.142:3306)/nice?charset=utf8&parseTime=True&loc=Local")
		// if err != nil {
		// 	panic(err)
		// }
		// defer db.Close()
		// client := nice.NewKeeper(db)
		// folderPath := "./nice/csv"
		// client.ImportNiceCsv(folderPath, db)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "collector_config.json", "config file (default is $HOME/.collector_config.json)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
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

		// Search config in home directory with name ".kipris-collector" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".otmk-kipris-collector")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

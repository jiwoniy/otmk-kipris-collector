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
package main

import (
	"github.com/jiwoniy/otmk-kipris-collector/app"
	"github.com/jiwoniy/otmk-kipris-collector/kipris/collector"
	"github.com/jiwoniy/otmk-kipris-collector/kipris/types"
)

func main() {
	// cmd.Execute()

	collectorInstance, err := collector.New()

	if err != nil {
		panic(err)
	}

	application := app.NewApplication(collectorInstance)
	restConfig := types.RestConfig{
		ListenAddr: ":8080",
	}
	app.StartApplication(application, restConfig)
}

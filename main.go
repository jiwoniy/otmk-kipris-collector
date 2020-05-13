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
	"github.com/jiwoniy/otmk-kipris-collector/collector"
	"github.com/jiwoniy/otmk-kipris-collector/types"
)

func main() {
	// cmd.Execute()

	config := types.CollectorConfig{
		Endpoint:     "http://plus.kipris.or.kr/openapi/rest",
		AccessKey:    "=JbKg6deF5WolYTZcZkypzgLBbSVbjZC6VEgfccaQyw=",
		DbType:       "sqlite3",
		DbConnString: "./test.db",
		// DbType:       "mysql",
		// DbConnString: "kipris_server:OnthemarkKipris0507!@@(61.97.187.142:3306)/kipris?charset=utf8&parseTime=True&loc=Local",
	}

	collectorInstance, err := collector.NewCollector(config)
	if err != nil {
		panic(err)
	}

	application := app.NewApplication(collectorInstance)
	restConfig := types.RestConfig{
		ListenAddr: ":8080",
	}
	app.StartApplication(application, restConfig)
}

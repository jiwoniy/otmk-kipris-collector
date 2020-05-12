package app

import (
	"testing"

	"github.com/jiwoniy/otmk-kipris-collector/collector"
	"github.com/jiwoniy/otmk-kipris-collector/types"
)

// func TestRestApp(t *testing.T) {
// 	config := types.CollectorConfig{
// 		Endpoint:     "http://plus.kipris.or.kr/openapi/rest",
// 		AccessKey:    "=JbKg6deF5WolYTZcZkypzgLBbSVbjZC6VEgfccaQyw=",
// 		DbType:       "mysql",
// 		DbConnString: "kipris_server:OnthemarkKipris0507!@@(61.97.187.142:3306)/kipris?charset=utf8&parseTime=True&loc=Local",
// 	}

// 	collectorInstance, err := collector.NewCollector(config)
// 	if err != nil {
// 		panic(err)
// 	}

// 	app := NewApplication(collectorInstance)
// 	restConfig := types.RestConfig{
// 		ListenAddr: ":8082",
// 	}
// 	StartApplication(app, restConfig)
// }

func TestRestRouter(t *testing.T) {
	config := types.CollectorConfig{
		Endpoint:     "http://plus.kipris.or.kr/openapi/rest",
		AccessKey:    "=JbKg6deF5WolYTZcZkypzgLBbSVbjZC6VEgfccaQyw=",
		DbType:       "sqlite3",
		DbConnString: ":memory:",
	}

	collectorInstance, err := collector.NewCollector(config)
	if err != nil {
		panic(err)
	}

	app := NewApplication(collectorInstance)

	getMethods, _ := app.collector.GetMethods()
	for _, restMethod := range getMethods {
		// router.GET(restMethod.Path, restHandler(restMethod.Handler))
		// restHandler(restMethod.Handler)
		// restMethod.Handler()
		// fmt.Println(restMethod)
		// r.GET(restMethod.Path, restHandler(restMethod.Handler))
	}

}

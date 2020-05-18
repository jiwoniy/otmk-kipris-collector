package collector

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jiwoniy/otmk-kipris-collector/kipris/model"
	"github.com/jiwoniy/otmk-kipris-collector/kipris/types"
)

type testcases struct {
	input          string
	success        bool
	data           interface{}
	responseStatus model.KiprisResponseStatus
}

func TestCollectorRest(t *testing.T) {
	config := types.CollectorConfig{
		KiprisConfig: types.KiprisConfig{
			Endpoint:  "http://plus.kipris.or.kr/openapi/rest",
			AccessKey: "I0Jnw4w6/UpQSp1zHPsIDSztV9=hgVUNI6IANH3bCEw=", // onthe mark key
		},
		DbConfig: types.DbConfig{
			DbType: "sqlite3",
			// DbConnString: "./rest.db",
			DbConnString: ":memory:",
			// DbType:       "mysql",
			// DbConnString: "kipris_server:OnthemarkKipris0507!@@(61.97.187.142:3306)/kipris?charset=utf8&parseTime=True&loc=Local",
		},

		// DbConnString: "./test.db",
	}

	collector, err := NewCollector(config)
	if err != nil {
		t.Error(err)
	}

	param := types.TaskParameters{
		Year: "2020",
	}
	err = collector.CreateTask(param)
	err = collector.CreateTask(param)
	// err = collector.CreateTask(param)
	if err != nil {
		t.Error(err)
	}

	// pagination, _ := collector.GetTaskList("1", "20")

	// s := collector.GetStorage()
	// db := s.GetDB()

	// pagination, _ := s.GetTaskApplicationNumberList(db, 1, 1, 5)
	// fmt.Println(pagination)
	// fmt.Println(pagination.Data)

	// a, _ := json.Marshal(pagination)
	// fmt.Println(string(a))

	data, _ := collector.GetTaskById(10)

	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

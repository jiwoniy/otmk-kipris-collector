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
		Endpoint:     "http://plus.kipris.or.kr/openapi/rest",
		AccessKey:    "I0Jnw4w6/UpQSp1zHPsIDSztV9=hgVUNI6IANH3bCEw=", // onthe mark key
		DbType:       "sqlite3",
		DbConnString: ":memory:",
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
	err = collector.CreateTask(param)
	if err != nil {
		t.Error(err)
	}

	// pagination, _ := collector.GetTaskList("1", "20")

	// a, _ := json.Marshal(pagination)
	// fmt.Println(string(a))

	data, _ := collector.GetTaskById(10)

	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

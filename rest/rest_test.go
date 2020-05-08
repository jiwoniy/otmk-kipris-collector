package rest

import (
	"testing"

	"github.com/jiwoniy/otmk-kipris-collector/query"
	"github.com/jiwoniy/otmk-kipris-collector/types"
)

func TestRest(t *testing.T) {
	queryConfig := types.QueryConfig{
		DbType:       "mysql",
		DbConnString: "kipris_server:OnthemarkKipris0507!@@(61.97.187.142:3306)/kipris?charset=utf8&parseTime=True&loc=Local",
	}

	queryApp, err := query.NewApp(queryConfig)
	if err != nil {
		t.Error(err)
	}

	config := types.RestConfig{
		ListenAddr: ":8082",
	}

	StartApplication(queryApp, config)
}

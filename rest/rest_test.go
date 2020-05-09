package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jiwoniy/otmk-kipris-collector/query"
	"github.com/jiwoniy/otmk-kipris-collector/types"
	"github.com/stretchr/testify/suite"
	"gopkg.in/go-playground/assert.v1"
)

type RestTestSuite struct {
	suite.Suite
	queryApp types.Query
}

func (suite *RestTestSuite) SetupTest() {
	queryConfig := types.QueryConfig{
		DbType:       "mysql",
		DbConnString: "kipris_server:OnthemarkKipris0507!@@(61.97.187.142:3306)/kipris?charset=utf8&parseTime=True&loc=Local",
	}

	queryApp, err := query.NewApp(queryConfig)
	if err != nil {
		suite.Error(err)
	}

	suite.queryApp = queryApp

	// config := types.RestConfig{
	// 	ListenAddr: ":8082",
	// }
	// suite.config = config
}

// func TestRest(t *testing.T) {
// 	StartApplication(suite.queryApp, suite.config)
// }

func (suite *RestTestSuite) TestPingRoute() {
	router := setupRouter(suite.queryApp)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Equal(suite.T(), "pong", w.Body.String())
}

// func (suite *RestTestSuite) TestPing2Route() {
// 	router := setupRouter(suite.queryApp)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/ping2", nil)
// 	router.ServeHTTP(w, req)

// 	// b, _ := json.MarshalIndent(w.Body, "", " ")
// 	fmt.Println(w.Body.String())
// 	// fmt.Println(string(b))

// 	assert.Equal(suite.T(), "application/json; charset=utf-8", w.Header().Get("Content-Type"))
// 	assert.Equal(suite.T(), http.StatusOK, w.Code)
// 	// assert.Equal(t, "pong", w.Body.String())
// }

func TestRestSuite(t *testing.T) {
	suite.Run(t, new(RestTestSuite))
}

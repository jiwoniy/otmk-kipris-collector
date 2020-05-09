package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jiwoniy/otmk-kipris-collector/types"
)

// type ResultResponse struct {
// 	Success bool        `json:"success"`
// 	Data    interface{} `json:"data,omitempty"`
// 	Error   string      `json:"error,omitempty"`
// }

// func writeResponse(c *gin.Context, data interface{}) {
// 	var response = ResultResponse{
// 		Success: true,
// 		Data:    data,
// 	}

// 	c.JSON(http.StatusOK, response)
// }

func setupRouter(app types.Query) *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// r.GET("/ping2", func(c *gin.Context) {
	// 	data := app.GetApplicationNumber("4020200000001")
	// 	c.JSON(http.StatusOK, data)
	// })

	return r
}

func StartApplication(app types.Query, config types.RestConfig) {
	r := setupRouter(app)
	r.Run(config.ListenAddr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

package rest

import (
	"fmt"

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

func StartApplication(app types.Query, config types.RestConfig) {
	r := gin.Default()

	fmt.Println(app)

	r.GET("ping", app.GetApplicationNumber("dd"))

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// CQRS
	// command
	// craw application number
	// status

	// query
	// get application number

	r.Run(config.ListenAddr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

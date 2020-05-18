package app

import (
	"github.com/gin-gonic/gin"

	"github.com/jiwoniy/otmk-kipris-collector/kipris/types"
)

func restHandler(fn func(ctx *gin.Context)) gin.HandlerFunc {
	return fn
}

func setupRouter(app types.RestClient) *gin.Engine {
	r := gin.Default()
	r.Use(cORSMiddleware())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "no route",
		})
	})

	getMethods, _ := app.GetMethods()
	postMethods, _ := app.PostMethods()
	for _, restMethod := range getMethods {
		r.GET(restMethod.Path, restHandler(restMethod.Handler))
	}
	for _, restMethod := range postMethods {
		r.POST(restMethod.Path, restHandler(restMethod.Handler))
	}

	return r
}

func StartApplication(app *Application, mode string, config types.RestConfig) {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// - using env:   export GIN_MODE=release
	// - using code:  gin.SetMode(gin.ReleaseMode)
	switch mode {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	r := setupRouter(app.collector)

	r.Run(config.ListenAddr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func cORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET,POST")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

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

func StartApplication(app *Application, config types.RestConfig) {
	// - using env:   export GIN_MODE=release
	// - using code:  gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.TestMode)
	r := setupRouter(app.collector)

	// srv := &http.Server{
	// 	Addr:    config.ListenAddr,
	// 	Handler: r,
	// }

	// go func() {
	// 	// service connections
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// // Wait for interrupt signal to gracefully shutdown the server with
	// // a timeout of 5 seconds.
	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// log.Println("Shutdown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }
	// log.Println("Server exiting")

	r.Run(config.ListenAddr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

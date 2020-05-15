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

	// r.Use(cors.Default())
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:8080"},
	// 	AllowMethods:     []string{"GET", "POST"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "http://localhost:8080"
	// 	},
	// 	// MaxAge: 12 * time.Hour,
	// }))

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

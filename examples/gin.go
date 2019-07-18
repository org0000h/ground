package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Creates an instance without any middleware by default
	framework := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	framework.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	framework.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	//framework.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// Authorization group
	// authorized := framework.Group("/", AuthRequired())
	// exactly the same as:
	//authorized := framework.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//	authorized.POST("/read", readEndpoint)
	//
	//	// nested group
	//	testing := authorized.Group("testing")
	//	testing.GET("/analytics", analyticsEndpoint)
	//}

	// Listen and serve on 0.0.0.0:8080
	framework.Run(":8080")
}

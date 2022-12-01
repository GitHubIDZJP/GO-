package ginModule

import "fmt"

// 中间件
func ginNew() {

	fmt.Println("中间件和使用HTTP方法")
	// 使用默认中间件创建 gin 路由器:
	// logger and recovery (crash-free) middleware(使用默认中间件创建 gin 路由器)
	/*
		//使用HTTP方法
		router := gin.Default()

		router.GET("/someGet", getting)
		router.POST("/somePost", posting)
		router.PUT("/somePut", putting)
		router.DELETE("/someDelete", deleting)
		router.PATCH("/somePatch", patching)
		router.HEAD("/someHead", head)
		router.OPTIONS("/someOptions", options)

	*/

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	//router.Run()

	// 默认情况下创建没有任何中间件的路由器
	//r := gin.New()

	/* 全球的中间件
	日志记录器中间件将把日志写入gin.DefaultWriter即使你设置了GIN_MODE=release
	默认为gin。DefaultWriter = os.Stdout
	*/
	//	r.Use(gin.Logger())

	// 恢复中间件从任何恐慌中恢复，如果有，则写入 500.
	//r.Use(gin.Recovery())

	// 每个路由中间件，您可以根据需要添加任意数量的中间件.
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 授权组
	// authorized := r.Group("/", AuthRequired())
	// 完全一样:
	//authorized := r.Group("/")
	// 每个组中间件！在这种情况下，我们使用自定义创建的
	// AuthRequired() 中间件仅在“授权”组中.
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//	authorized.POST("/read", readEndpoint)

	// 嵌套组
	//	testing := authorized.Group("testing")
	//	testing.GET("/analytics", analyticsEndpoint)
	//}

	// Listen and serve on 0.0.0.0:8080
	//r.Run(":8080")
}

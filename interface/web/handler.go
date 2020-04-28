package web


// Router returns a http router with endpoints set
func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/users/:id", getUser)
	router.GET("/users", getUsers)
	router.POST("/users", addUser)
	router.GET("/users-match/:id", getUserMatch)

	return router
}

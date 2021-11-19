package application

func mapUrls() {
	router.GET("/api/employee", handler.Get)
	router.POST("/api/employee", handler.Create)
	router.GET("/api/employee/:employee_id", handler.GetByID)
	router.POST("/api/employee/:employee_id", handler.UploadIcon)
	router.PATCH("/api/employee/:employee_id", handler.Update)
	router.GET("/api/employee/search/s", handler.Search)
	router.DELETE("/api/employee/:employee_id", handler.DeleteIcon)
}

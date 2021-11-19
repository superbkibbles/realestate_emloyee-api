package application

import (
	"github.com/gin-gonic/gin"
	"github.com/superbkibbles/realestate_employee-api/src/clients/elasticsearch"
	"github.com/superbkibbles/realestate_employee-api/src/http"
	"github.com/superbkibbles/realestate_employee-api/src/repository/db"
	"github.com/superbkibbles/realestate_employee-api/src/services/employeeService"
)

var (
	router  = gin.Default()
	handler http.EmployeeHandler
)

func StartApp() {
	elasticsearch.Client.Init()
	handler = http.NewComplexHandler(employeeService.NewComplexService(db.NewDbRepository()))
	mapUrls()
	router.Static("assets", "clients/visuals")

	router.Run(":3050")
}

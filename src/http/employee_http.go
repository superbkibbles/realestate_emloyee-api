package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/superbkibbles/bookstore_utils-go/rest_errors"
	"github.com/superbkibbles/realestate_employee-api/src/domain/employee"
	"github.com/superbkibbles/realestate_employee-api/src/domain/query"
	"github.com/superbkibbles/realestate_employee-api/src/domain/update"
	"github.com/superbkibbles/realestate_employee-api/src/services/employeeService"
)

type EmployeeHandler interface {
	Get(*gin.Context)
	Create(*gin.Context)
	GetByID(*gin.Context)
	UploadIcon(*gin.Context)
	Update(*gin.Context)
	Search(*gin.Context)
	DeleteIcon(*gin.Context)
}

type employeeHandler struct {
	srv employeeService.EmployeeService
}

func NewComplexHandler(srv employeeService.EmployeeService) EmployeeHandler {
	return &employeeHandler{
		srv: srv,
	}
}

func (eh *employeeHandler) Get(c *gin.Context) {
	employees, err := eh.srv.Get()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, employees)
}

func (eh *employeeHandler) Create(c *gin.Context) {
	var e employee.Employee
	if err := c.ShouldBindJSON(&e); err != nil {
		restErr := rest_errors.NewBadRequestErr("Invalid JSON body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	if err := eh.srv.Save(&e); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, e)
}

func (eh *employeeHandler) GetByID(c *gin.Context) {
	employeeID := strings.TrimSpace(c.Param("employee_id"))

	e, err := eh.srv.GetByID(employeeID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusFound, e)
}

func (eh *employeeHandler) UploadIcon(c *gin.Context) {
	employeeID := strings.TrimSpace(c.Param("employee_id"))

	file, err := c.FormFile("photo")
	if err != nil {
		restErr := rest_errors.NewBadRequestErr("Bad Request")
		c.JSON(restErr.Status(), restErr)
		return
	}

	e, uploadErr := eh.srv.UploadIcon(employeeID, file)
	if uploadErr != nil {
		c.JSON(uploadErr.Status(), uploadErr)
		return
	}

	c.JSON(http.StatusOK, e)
}

func (eh *employeeHandler) Update(c *gin.Context) {
	employeeID := strings.TrimSpace(c.Param("employee_id"))
	var esUpdate update.EsUpdate
	if err := c.ShouldBindJSON(&esUpdate); err != nil {
		restErr := rest_errors.NewBadRequestErr("Invalid JSON body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	e, err := eh.srv.Update(employeeID, esUpdate)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, e)
}

func (eh *employeeHandler) Search(c *gin.Context) {
	var q query.EsQuery
	if err := c.ShouldBindJSON(&q); err != nil {
		restErr := rest_errors.NewBadRequestErr("Invalid JSON body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	e, err := eh.srv.Search(q)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusFound, e)
}

func (eh *employeeHandler) DeleteIcon(c *gin.Context) {
	employeeID := strings.TrimSpace(c.Param("employee_id"))

	if err := eh.srv.DeleteIcon(employeeID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.String(http.StatusOK, "Deleted")
}

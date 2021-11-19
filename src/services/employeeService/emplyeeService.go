package employeeService

import (
	"mime/multipart"
	"strings"

	"github.com/superbkibbles/bookstore_utils-go/rest_errors"
	"github.com/superbkibbles/realestate_employee-api/src/domain/employee"
	"github.com/superbkibbles/realestate_employee-api/src/domain/query"
	"github.com/superbkibbles/realestate_employee-api/src/domain/update"
	"github.com/superbkibbles/realestate_employee-api/src/repository/db"
	"github.com/superbkibbles/realestate_employee-api/src/utils/date_utils"
	"github.com/superbkibbles/realestate_employee-api/src/utils/file_utils"
)

type EmployeeService interface {
	Get() (employee.Employees, rest_errors.RestErr)
	Save(complex *employee.Employee) rest_errors.RestErr
	GetByID(complexID string) (*employee.Employee, rest_errors.RestErr)
	UploadIcon(id string, fileHeader *multipart.FileHeader) (*employee.Employee, rest_errors.RestErr)
	Update(id string, updateRequest update.EsUpdate) (*employee.Employee, rest_errors.RestErr)
	Search(updateRequest query.EsQuery) (employee.Employees, rest_errors.RestErr)
	DeleteIcon(agencyID string) rest_errors.RestErr
}

type employeeService struct {
	dbRepo db.DbRepository
}

func NewComplexService(dbRepo db.DbRepository) EmployeeService {
	return &employeeService{
		dbRepo: dbRepo,
	}
}

func (srv *employeeService) Get() (employee.Employees, rest_errors.RestErr) {
	return srv.dbRepo.Get()
}

func (srv *employeeService) Save(e *employee.Employee) rest_errors.RestErr {
	e.Status = employee.STATUS_ACTIVE
	e.DateCreated = date_utils.GetNowDBFromat()
	return srv.dbRepo.Save(e)
}

func (cs *employeeService) GetByID(complexID string) (*employee.Employee, rest_errors.RestErr) {
	return cs.dbRepo.GetByID(complexID)
}

func (srv *employeeService) UploadIcon(id string, fileHeader *multipart.FileHeader) (*employee.Employee, rest_errors.RestErr) {
	complex, err := srv.GetByID(id)
	if err != nil {
		return nil, err
	}
	file, fErr := fileHeader.Open()
	if fErr != nil {
		return nil, rest_errors.NewInternalServerErr("Error while trying to open the file", nil)
	}
	filePath, err := file_utils.SaveFile(fileHeader, file)
	if err != nil {
		return nil, err
	}
	complex.Photo = "http://localhost:3050/assets/" + filePath

	srv.dbRepo.UploadIcon(complex, id)
	return complex, nil
}

func (srv *employeeService) Update(id string, updateRequest update.EsUpdate) (*employee.Employee, rest_errors.RestErr) {
	return srv.dbRepo.Update(id, updateRequest)
}

func (srv *employeeService) Search(updateRequest query.EsQuery) (employee.Employees, rest_errors.RestErr) {
	return srv.dbRepo.Search(updateRequest)
}

func (srv *employeeService) DeleteIcon(agencyID string) rest_errors.RestErr {
	agency, err := srv.GetByID(agencyID)
	if err != nil {
		return err
	}

	splittedPath := strings.Split(agency.Photo, "/")
	fileName := splittedPath[len(splittedPath)-1]

	file_utils.DeleteFile(fileName)

	agency.Photo = ""
	srv.dbRepo.UploadIcon(agency, agencyID)
	return nil
}

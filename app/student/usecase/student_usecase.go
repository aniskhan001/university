package usecase

import (
	"university/app/student/repo"

	deptRepo "university/app/department/repo"

	"github.com/labstack/echo/v4"
)

type StudentUsecase interface {
	Get(echo.Context, uint) (*repo.StudentResp, error)
	GetDetails(echo.Context, uint) (*repo.StudentDetailsResp, error)
	List(echo.Context) ([]repo.StudentResp, error)
	ListByDept(echo.Context, uint) ([]repo.StudentResp, error)
	Insert(echo.Context) (*repo.StudentResp, error)
	InsertMany(echo.Context) ([]repo.StudentResp, error)
	Edit(echo.Context, uint) (*repo.StudentResp, error)
}

type studentUsecase struct {
	repo     repo.StudentRepository
	deptRepo deptRepo.DeptRepository
}

func NewStudentUsecase(repo repo.StudentRepository, deptRepo deptRepo.DeptRepository) StudentUsecase {
	return &studentUsecase{
		repo:     repo,
		deptRepo: deptRepo,
	}
}

func (du *studentUsecase) Get(c echo.Context, id uint) (*repo.StudentResp, error) {
	dept, err := du.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return repo.ToStudentResponse(dept), nil
}

func (du *studentUsecase) GetDetails(c echo.Context, id uint) (*repo.StudentDetailsResp, error) {
	student, err := du.repo.Get(id)
	if err != nil {
		return nil, err
	}

	dept, err := du.deptRepo.Get(student.Department)
	if err != nil {
		return nil, err
	}

	return repo.StudentDetailsResponse(student, dept), nil
}

func (du *studentUsecase) List(c echo.Context) ([]repo.StudentResp, error) {
	dept, err := du.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return repo.ToStudentsResponse(dept), nil
}

func (du *studentUsecase) ListByDept(c echo.Context, deptID uint) ([]repo.StudentResp, error) {
	dept, err := du.repo.GetAllFromDept(deptID)
	if err != nil {
		return nil, err
	}

	return repo.ToStudentsResponse(dept), nil
}

func (du *studentUsecase) Insert(c echo.Context) (*repo.StudentResp, error) {
	// reading data from request
	var deptResp repo.StudentResp
	err := c.Bind(&deptResp)
	if err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	dept, err := du.repo.Insert(repo.ToStudentModel(&deptResp))
	if err != nil {
		return nil, err
	}

	return repo.ToStudentResponse(dept), nil
}

func (du *studentUsecase) InsertMany(c echo.Context) ([]repo.StudentResp, error) {
	// reading data from request
	var deptsResp []repo.StudentResp
	err := c.Bind(&deptsResp)
	if err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	depts, err := du.repo.InsertMany(repo.ToStudentsModel(deptsResp))
	if err != nil {
		return nil, err
	}

	return repo.ToStudentsResponse(depts), nil
}

func (du *studentUsecase) Edit(c echo.Context, id uint) (*repo.StudentResp, error) {
	// reading data from request
	var deptResp repo.StudentResp
	err := c.Bind(&deptResp)
	if err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	dept, err := du.repo.Edit(id, repo.ToStudentModel(&deptResp))
	if err != nil {
		return nil, err
	}

	return repo.ToStudentResponse(dept), nil
}

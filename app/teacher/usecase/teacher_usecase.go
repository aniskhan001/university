package usecase

import (
	"university/app/teacher/repo"

	deptRepo "university/app/department/repo"

	"github.com/labstack/echo/v4"
)

type TeacherUsecase interface {
	Get(echo.Context, uint) (*repo.TeacherResp, error)
	GetDetails(echo.Context, uint) (*repo.TeacherDetailsResp, error)
	List(echo.Context) ([]repo.TeacherResp, error)
	ListByDept(echo.Context, uint) ([]repo.TeacherResp, error)
	Insert(echo.Context) (*repo.TeacherResp, error)
	InsertMany(echo.Context) ([]repo.TeacherResp, error)
	Edit(echo.Context, uint) (*repo.TeacherResp, error)
}

type teacherUsecase struct {
	repo     repo.TeacherRepository
	deptRepo deptRepo.DeptRepository
}

func NewTeacherUsecase(repo repo.TeacherRepository, deptRepo deptRepo.DeptRepository) TeacherUsecase {
	return &teacherUsecase{
		repo:     repo,
		deptRepo: deptRepo,
	}
}

func (du *teacherUsecase) Get(c echo.Context, id uint) (*repo.TeacherResp, error) {
	dept, err := du.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return repo.ToTeacherResponse(dept), nil
}

func (du *teacherUsecase) GetDetails(c echo.Context, id uint) (*repo.TeacherDetailsResp, error) {
	teacher, err := du.repo.Get(id)
	if err != nil {
		return nil, err
	}

	dept, err := du.deptRepo.Get(teacher.Department)
	if err != nil {
		return nil, err
	}

	return repo.TeacherDetailsResponse(teacher, dept), nil
}

func (du *teacherUsecase) List(c echo.Context) ([]repo.TeacherResp, error) {
	dept, err := du.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return repo.ToTeachersResponse(dept), nil
}

func (du *teacherUsecase) ListByDept(c echo.Context, deptID uint) ([]repo.TeacherResp, error) {
	dept, err := du.repo.GetAllFromDept(deptID)
	if err != nil {
		return nil, err
	}

	return repo.ToTeachersResponse(dept), nil
}

func (du *teacherUsecase) Insert(c echo.Context) (*repo.TeacherResp, error) {
	// reading data from request
	var deptResp repo.TeacherResp
	err := c.Bind(&deptResp)
	if err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	dept, err := du.repo.Insert(repo.ToTeacherModel(&deptResp))
	if err != nil {
		return nil, err
	}

	return repo.ToTeacherResponse(dept), nil
}

func (du *teacherUsecase) InsertMany(c echo.Context) ([]repo.TeacherResp, error) {
	// reading data from request
	var deptsResp []repo.TeacherResp
	err := c.Bind(&deptsResp)
	if err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	depts, err := du.repo.InsertMany(repo.ToTeachersModel(deptsResp))
	if err != nil {
		return nil, err
	}

	return repo.ToTeachersResponse(depts), nil
}

func (du *teacherUsecase) Edit(c echo.Context, id uint) (*repo.TeacherResp, error) {
	// reading data from request
	var deptResp repo.TeacherResp
	err := c.Bind(&deptResp)
	if err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	dept, err := du.repo.Edit(id, repo.ToTeacherModel(&deptResp))
	if err != nil {
		return nil, err
	}

	return repo.ToTeacherResponse(dept), nil
}

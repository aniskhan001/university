package usecase

import (
	"university/app/department/repo"

	"github.com/labstack/echo/v4"
)

type DeptUsecase interface {
	Get(echo.Context, uint) (*repo.DeptResp, error)
	List(echo.Context) ([]repo.DeptResp, error)
	Insert(echo.Context) (*repo.DeptResp, error)
	InsertMany(echo.Context) ([]repo.DeptResp, error)
	Edit(echo.Context, uint) (*repo.DeptResp, error)
}

type deptUsecase struct {
	repo repo.DeptRepository
}

func NewDeptUsecase(repo repo.DeptRepository) DeptUsecase {
	return &deptUsecase{
		repo: repo,
	}
}

func (du *deptUsecase) Get(c echo.Context, id uint) (*repo.DeptResp, error) {
	dept, err := du.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return repo.ToDeptResponse(dept), nil
}

func (du *deptUsecase) List(c echo.Context) ([]repo.DeptResp, error) {
	dept, err := du.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return repo.ToDeptsResponse(dept), nil
}

func (du *deptUsecase) Insert(c echo.Context) (*repo.DeptResp, error) {
	// reading data from request
	var deptResp repo.DeptResp
	err := c.Bind(&deptResp)
	if err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	dept, err := du.repo.Insert(repo.ToDeptModel(&deptResp))
	if err != nil {
		return nil, err
	}

	return repo.ToDeptResponse(dept), nil
}

func (du *deptUsecase) InsertMany(c echo.Context) ([]repo.DeptResp, error) {
	// reading data from request
	var deptsResp []repo.DeptResp
	err := c.Bind(&deptsResp)
	if err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	depts, err := du.repo.InsertMany(repo.ToDeptsModel(deptsResp))
	if err != nil {
		return nil, err
	}

	return repo.ToDeptsResponse(depts), nil
}

func (du *deptUsecase) Edit(c echo.Context, id uint) (*repo.DeptResp, error) {
	// reading data from request
	var deptResp repo.DeptResp
	err := c.Bind(&deptResp)
	if err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	dept, err := du.repo.Edit(id, repo.ToDeptModel(&deptResp))
	if err != nil {
		return nil, err
	}

	return repo.ToDeptResponse(dept), nil
}

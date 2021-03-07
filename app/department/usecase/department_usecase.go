package usecase

import (
	"university/app/department/repository"

	"github.com/labstack/echo/v4"
)

type DeptUsecase interface {
	Get(c echo.Context, id uint) (*DeptResp, error)
	List(c echo.Context) ([]DeptResp, error)
}

type deptUsecase struct {
	repo repository.DeptRepository
}

func NewDeptUsecase(repo repository.DeptRepository) DeptUsecase {
	return &deptUsecase{
		repo: repo,
	}
}

func (du *deptUsecase) Get(c echo.Context, id uint) (*DeptResp, error) {
	dept, err := du.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return toDeptResp(dept), nil
}

func (du *deptUsecase) List(c echo.Context) ([]DeptResp, error) {
	dept, err := du.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return toDeptResps(dept), nil
}

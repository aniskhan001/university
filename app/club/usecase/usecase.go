package usecase

import (
	"university/app/club/repo"

	"github.com/labstack/echo/v4"
)

type Usecase interface {
	Get(echo.Context, uint) (*repo.Response, error)
	List(echo.Context) ([]repo.Response, error)
	Insert(echo.Context) (*repo.Response, error)
	Edit(echo.Context, uint) (*repo.Response, error)
}

type usecase struct {
	repo repo.Repo
}

func Init(repo repo.Repo) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Get(c echo.Context, id uint) (*repo.Response, error) {
	res, err := uc.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return repo.ToResponse(res), nil
}

func (uc *usecase) List(c echo.Context) ([]repo.Response, error) {
	res, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return repo.ToListResponse(res), nil
}

func (uc *usecase) Insert(c echo.Context) (*repo.Response, error) {
	// reading data from request
	var resp repo.Response
	if err := c.Bind(&resp); err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	res, err := uc.repo.Insert(repo.ToModel(&resp))
	if err != nil {
		return nil, err
	}

	return repo.ToResponse(res), nil
}

func (uc *usecase) Edit(c echo.Context, id uint) (*repo.Response, error) {
	// reading data from request
	var resp repo.Response
	if err := c.Bind(&resp); err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	res, err := uc.repo.Edit(id, repo.ToModel(&resp))
	if err != nil {
		return nil, err
	}

	return repo.ToResponse(res), nil
}

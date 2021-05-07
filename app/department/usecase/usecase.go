package usecase

import (
	"university/app/department/repo"

	"github.com/labstack/echo/v4"
)

type Usecase interface {
	Get(echo.Context, uint) (*repo.Presenter, error)
	List(echo.Context) ([]repo.Presenter, error)
	Insert(echo.Context) (*repo.Presenter, error)
	InsertMany(echo.Context) ([]repo.Presenter, error)
	Edit(echo.Context, uint) (*repo.Presenter, error)
}

type usecase struct {
	repo repo.Repo
}

func Init(repo repo.Repo) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Get(c echo.Context, id uint) (*repo.Presenter, error) {
	res, err := uc.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return repo.ToPresenter(res), nil
}

func (uc *usecase) List(c echo.Context) ([]repo.Presenter, error) {
	res, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return repo.ToPresenterList(res), nil
}

func (uc *usecase) Insert(c echo.Context) (*repo.Presenter, error) {
	// reading data from request
	data := repo.Presenter{}
	if err := c.Bind(&data); err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	res, err := uc.repo.Insert(repo.ToModel(&data))
	if err != nil {
		return nil, err
	}

	return repo.ToPresenter(res), nil
}

func (uc *usecase) InsertMany(c echo.Context) ([]repo.Presenter, error) {
	// reading data from request
	data := []repo.Presenter{}
	if err := c.Bind(&data); err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	res, err := uc.repo.InsertMany(repo.ToModelList(data))
	if err != nil {
		return nil, err
	}

	return repo.ToPresenterList(res), nil
}

func (uc *usecase) Edit(c echo.Context, id uint) (*repo.Presenter, error) {
	// reading data from request
	data := repo.Presenter{}
	if err := c.Bind(&data); err != nil {
		return nil, err
	}

	// passing model to repository instead of binded data
	res, err := uc.repo.Edit(id, repo.ToModel(&data))
	if err != nil {
		return nil, err
	}

	return repo.ToPresenter(res), nil
}

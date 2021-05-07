package usecase

import (
	"university/app/student/repo"

	deptRepo "university/app/department/repo"

	"github.com/labstack/echo/v4"
)

type Usecase interface {
	Get(echo.Context, uint) (*repo.Presenter, error)
	GetDetails(echo.Context, uint) (*repo.DetailPresenter, error)
	List(echo.Context) ([]repo.Presenter, error)
	ListByDept(echo.Context, uint) ([]repo.Presenter, error)
	Insert(echo.Context) (*repo.Presenter, error)
	InsertMany(echo.Context) ([]repo.Presenter, error)
	Edit(echo.Context, uint) (*repo.Presenter, error)
}

type usecase struct {
	repo     repo.Repo
	deptRepo deptRepo.Repo
}

func Init(repo repo.Repo, deptRepo deptRepo.Repo) Usecase {
	return &usecase{
		repo:     repo,
		deptRepo: deptRepo,
	}
}

func (uc *usecase) Get(c echo.Context, id uint) (*repo.Presenter, error) {
	res, err := uc.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return repo.ToPresenter(res), nil
}

func (uc *usecase) GetDetails(c echo.Context, id uint) (*repo.DetailPresenter, error) {
	res, err := uc.repo.Get(id)
	if err != nil {
		return nil, err
	}

	deptRes, err := uc.deptRepo.Get(res.Department)
	if err != nil {
		return nil, err
	}

	return repo.ToDetailPresenter(res, deptRes), nil
}

func (uc *usecase) List(c echo.Context) ([]repo.Presenter, error) {
	res, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return repo.ToPresenterList(res), nil
}

func (uc *usecase) ListByDept(c echo.Context, deptID uint) ([]repo.Presenter, error) {
	res, err := uc.repo.GetAllFromDept(deptID)
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

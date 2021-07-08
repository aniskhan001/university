package repo

import (
	"university/model"
)

type Repo interface {
	Get(uint) (*model.Club, error)
	GetAll() ([]model.Club, error)
	Insert(*model.Club) (*model.Club, error)
	Edit(uint, *model.Club) (*model.Club, error)
}

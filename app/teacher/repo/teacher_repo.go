package repo

import (
	"university/app/teacher/model"
	"university/infrastructure/db"

	"gorm.io/gorm"
)

type teacherRepository struct {
	*gorm.DB
}

type TeacherRepository interface {
	Get(uint) (*model.Teacher, error)
	Insert(*model.Teacher) (*model.Teacher, error)
	InsertMany([]model.Teacher) ([]model.Teacher, error)
	GetAll() ([]model.Teacher, error)
	GetAllFromDept(uint) ([]model.Teacher, error)
	Edit(uint, *model.Teacher) (*model.Teacher, error)
}

// NewTeacherRepository will create an object that represent the DeptRepository interface
func NewTeacherRepository(DB *gorm.DB) TeacherRepository {
	return &teacherRepository{
		DB: db.Get().DB,
	}
}

func (db *teacherRepository) Get(id uint) (*model.Teacher, error) {
	var teacher model.Teacher

	if res := db.Select("id", "name").Where("id = ?", id).First(&teacher); res.Error != nil {
		return nil, res.Error
	}

	return &teacher, nil
}

func (db *teacherRepository) GetAll() ([]model.Teacher, error) {
	var teachers []model.Teacher

	if res := db.Select("id", "name").Find(&teachers); res.Error != nil {
		return nil, res.Error
	}

	return teachers, nil
}

func (db *teacherRepository) GetAllFromDept(deptID uint) ([]model.Teacher, error) {
	var teachers []model.Teacher

	if res := db.Select("id", "name").Where("department = ?", deptID).Find(&teachers); res.Error != nil {
		return nil, res.Error
	}

	return teachers, nil
}

func (db *teacherRepository) Insert(teacher *model.Teacher) (*model.Teacher, error) {
	if res := db.Create(teacher); res.Error != nil {
		return nil, res.Error
	}

	return teacher, nil
}

func (db *teacherRepository) InsertMany(teachers []model.Teacher) ([]model.Teacher, error) {
	if res := db.Create(teachers); res.Error != nil {
		return nil, res.Error
	}

	return teachers, nil
}

func (db *teacherRepository) Edit(id uint, teacher *model.Teacher) (*model.Teacher, error) {
	res := db.Model(model.Teacher{}).Where("id = ?", id).Updates(*teacher)
	if res.Error != nil {
		return nil, res.Error
	}

	return db.Get(id)
}

package repo

import (
	"university/infrastructure/db"
	"university/model"

	"gorm.io/gorm"
)

type studentRepository struct {
	*gorm.DB
}

type StudentRepository interface {
	Get(uint) (*model.Student, error)
	Insert(*model.Student) (*model.Student, error)
	InsertMany([]model.Student) ([]model.Student, error)
	GetAll() ([]model.Student, error)
	GetAllFromDept(uint) ([]model.Student, error)
	Edit(uint, *model.Student) (*model.Student, error)
}

// NewStudentRepository will create an object that represent the DeptRepository interface
func NewStudentRepository(DB *gorm.DB) StudentRepository {
	return &studentRepository{
		DB: db.Get().DB,
	}
}

func (db *studentRepository) Get(id uint) (*model.Student, error) {
	var student model.Student

	if res := db.Select("id", "name", "department").Where("id = ?", id).First(&student); res.Error != nil {
		return nil, res.Error
	}

	return &student, nil
}

func (db *studentRepository) GetAll() ([]model.Student, error) {
	var students []model.Student

	if res := db.Select("id", "name").Find(&students); res.Error != nil {
		return nil, res.Error
	}

	return students, nil
}

func (db *studentRepository) GetAllFromDept(deptID uint) ([]model.Student, error) {
	var students []model.Student

	if res := db.Select("id", "name").Where("department = ?", deptID).Find(&students); res.Error != nil {
		return nil, res.Error
	}

	return students, nil
}

func (db *studentRepository) Insert(student *model.Student) (*model.Student, error) {
	if res := db.Create(student); res.Error != nil {
		return nil, res.Error
	}

	return student, nil
}

func (db *studentRepository) InsertMany(students []model.Student) ([]model.Student, error) {
	if res := db.Create(students); res.Error != nil {
		return nil, res.Error
	}

	return students, nil
}

func (db *studentRepository) Edit(id uint, student *model.Student) (*model.Student, error) {
	res := db.Model(model.Student{}).Where("id = ?", id).Updates(*student)
	if res.Error != nil {
		return nil, res.Error
	}

	return db.Get(id)
}

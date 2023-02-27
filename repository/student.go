package repository

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID          uint64    `json:"id,omitempty" gorm:"primaryKey"`
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	Email       string    `json:"email,omitempty" gorm:"unique"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return Repo{db: db}
}

func (a *Repo) Create(student *Student) error {
	return a.db.Create(student).Error
}

func (a *Repo) FindOne(id uint64) (*Student, error) {
	student := Student{}
	res := a.db.First(&student, Student{ID: id})
	if res.Error != nil {
		return nil, res.Error
	}
	return &student, nil
}

func (a *Repo) FindAll() ([]Student, error) {
	students := make([]Student, 0)
	return students, a.db.Find(&students).Error
}

func (a *Repo) Update(upd *Student) (*Student, error) {
	student, err := a.FindOne(upd.ID)
	if err != nil {
		return nil, err
	}
	student.FirstName = upd.FirstName
	student.LastName = upd.LastName
	student.Email = upd.Email
	student.PhoneNumber = upd.PhoneNumber
	student.UpdatedAt = time.Now()
	a.db.Save(student)
	return student, nil
}

func (a *Repo) Delete(id uint64) error {
	student, err := a.FindOne(id)
	if err != nil {
		return err
	}
	a.db.Delete(student)
	return nil
}

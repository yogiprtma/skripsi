package subject

import "gorm.io/gorm"

type Repository interface {
	Save(subject Subject) (Subject, error)
	FindAll() ([]Subject, error)
	FindByID(ID int) (Subject, error)
	Update(subject Subject) (Subject, error)
	DeleteByID(ID int) (Subject, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(subject Subject) (Subject, error) {
	err := r.db.Create(&subject).Error
	if err != nil {
		return subject, err
	}

	return subject, nil
}

func (r *repository) FindAll() ([]Subject, error) {
	var subjects []Subject

	err := r.db.Find(&subjects).Error
	if err != nil {
		return subjects, err
	}

	return subjects, nil
}

func (r *repository) FindByID(ID int) (Subject, error) {
	var subject Subject

	err := r.db.Where("id = ?", ID).Find(&subject).Error
	if err != nil {
		return subject, err
	}

	return subject, nil
}

func (r *repository) Update(subject Subject) (Subject, error) {
	err := r.db.Save(&subject).Error
	if err != nil {
		return subject, err
	}

	return subject, nil
}

func (r *repository) DeleteByID(ID int) (Subject, error) {
	var subject Subject

	err := r.db.Where("id = ?", ID).Delete(&subject).Error
	if err != nil {
		return subject, err
	}

	return subject, nil
}

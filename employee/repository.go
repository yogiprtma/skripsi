package employee

import "gorm.io/gorm"

type Repository interface {
	Save(employee Employee) (Employee, error)
	FindAll() ([]Employee, error)
	Update(ID int, name, nip, role string) (Employee, error)
	FindByID(ID int) (Employee, error)
	UpdatePassword(ID int, password string) (Employee, error)
	DeleteByID(ID int) (Employee, error)
	FindByNIP(nip string) (Employee, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(employee Employee) (Employee, error) {
	err := r.db.Create(&employee).Error
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) FindAll() ([]Employee, error) {
	var employees []Employee

	err := r.db.Not("nip = ?", "admin").Find(&employees).Error
	if err != nil {
		return employees, err
	}

	return employees, nil
}

func (r *repository) Update(ID int, name, nip, role string) (Employee, error) {
	var employee Employee
	err := r.db.Model(&employee).Where("id = ?", ID).Updates(Employee{Name: name, Role: role, Nip: nip}).Error
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) FindByID(ID int) (Employee, error) {
	var employee Employee

	err := r.db.Where("id = ?", ID).Find(&employee).Error
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) UpdatePassword(ID int, password string) (Employee, error) {
	var employee Employee
	err := r.db.Model(&employee).Where("id = ?", ID).Update("password_hash", password).Error
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) DeleteByID(ID int) (Employee, error) {
	var employee Employee

	err := r.db.Where("id = ?", ID).Delete(&employee).Error
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) FindByNIP(nip string) (Employee, error) {
	var employee Employee

	err := r.db.Where("nip = ?", nip).Find(&employee).Error
	if err != nil {
		return employee, err
	}

	return employee, nil
}

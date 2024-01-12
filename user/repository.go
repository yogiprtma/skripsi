package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByNPM(NPM string) (User, error)
	FindAll() ([]User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
	UpdatePassword(ID int, password string) (User, error)
	DeleteByID(ID int) (User, error)
	FindByTwoDepartments(first string, second string) ([]User, error)
	FindByDepartment(department string) ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByNPM(NPM string) (User, error) {
	var user User

	err := r.db.Where("npm = ?", NPM).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Delete(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdatePassword(ID int, password string) (User, error) {
	var user User
	err := r.db.Model(&user).Where("id = ?", ID).Update("password_hash", password).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByTwoDepartments(first string, second string) ([]User, error) {
	var users []User

	err := r.db.Where("department IN ?", []string{first, second}).Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) FindByDepartment(department string) ([]User, error) {
	var users []User

	err := r.db.Where("department = ?", department).Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

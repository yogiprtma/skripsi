package user

import validation "github.com/go-ozzo/ozzo-validation"

type LoginInput struct {
	NPM           string `form:"npm" binding:"required"`
	Password      string `form:"password" binding:"required"`
	Error         error
	ErrorPassword error
}

func (form LoginInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.NPM, validation.Required.Error("Input tidak boleh kosong.")),
		validation.Field(&form.Password, validation.Required.Error("Input tidak boleh kosong.")),
	)
}

type FormCreateUserInput struct {
	Name       string `form:"name" binding:"required"`
	NPM        string `form:"npm" binding:"required"`
	Department string `form:"department" binding:"required,oneof='Informatika' 'Teknik Mesin' 'Teknik Sipil' 'Teknik Elektro' 'Arsitektur' 'Sistem Informasi'"`
	Password   string `form:"password" binding:"required,min=6,max=14"`
	Error      error
}

func (form FormCreateUserInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("Nama tidak boleh kosong.")),
		validation.Field(&form.NPM, validation.Required.Error("NPM tidak boleh kosong.")),
		validation.Field(&form.Department, validation.Required.Error("Prodi tidak boleh kosong."),
			validation.In("Informatika", "Teknik Mesin", "Teknik Sipil", "Teknik Elektro", "Arsitektur", "Sistem Informasi").Error("Prodi harus valid.")),
		validation.Field(&form.Password, validation.Required.Error("Password tidak boleh kosong."),
			validation.Length(6, 14).Error("Password harus diantara 6 sampai 14 karakter.")),
	)
}

type FormUpdateUserInput struct {
	ID         int
	Name       string `form:"name" binding:"required"`
	NPM        string `form:"npm" binding:"required"`
	Department string `form:"department" binding:"required,oneof='Informatika' 'Teknik Mesin' 'Teknik Sipil' 'Teknik Elektro' 'Arsitektur' 'Sistem Informasi'"`
	Error      error
}

func (form FormUpdateUserInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("Nama tidak boleh kosong.")),
		validation.Field(&form.NPM, validation.Required.Error("NPM tidak boleh kosong.")),
		validation.Field(&form.Department, validation.Required.Error("Prodi tidak boleh kosong."),
			validation.In("Informatika", "Teknik Mesin", "Teknik Sipil", "Teknik Elektro", "Arsitektur", "Sistem Informasi").Error("Prodi harus valid.")),
	)
}

type FormUpdatePasswordUserInput struct {
	ID            int
	Password      string `form:"password1" binding:"required,min=6,max=14"`
	NewPassword   string `form:"password2" binding:"required,min=6,max=14"`
	Error         error
	ErrorPassword string
}

func (form FormUpdatePasswordUserInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Password, validation.Required.Error("Password tidak boleh kosong."),
			validation.Length(6, 14).Error("Password harus diantara 6 sampai 14 karakter.")),
		validation.Field(&form.NewPassword, validation.Required.Error("Password tidak boleh kosong."),
			validation.Length(6, 14).Error("Password harus diantara 6 sampai 14 karakter.")),
	)
}

type FormFindPublicAndPrivateKeyUserInput struct {
	NPM          string `form:"npm" binding:"required"`
	Error        error
	ErrorFindNPM error
	ErrorKey     error
}

func (form FormFindPublicAndPrivateKeyUserInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.NPM, validation.Required.Error("NPM tidak boleh kosong.")),
	)
}

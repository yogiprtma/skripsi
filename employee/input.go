package employee

import validation "github.com/go-ozzo/ozzo-validation"

type FormCreateEmployeeInput struct {
	Name     string `form:"name" binding:"required"`
	NIP      string `form:"nip" binding:"required"`
	Role     string `form:"role" binding:"required,oneof='Wakil Dekan Bidang Akademik' 'Koordinator Prodi Informatika' 'Koordinator Prodi Teknik Sipil' 'Koordinator Prodi Teknik Elektro' 'Koordinator Prodi Teknik Mesin' 'Koordinator Prodi Arsitektur' 'Koordinator Prodi Sistem Informasi' 'Karyawan Akademik Bagian Informatika dan Sistem Informasi' 'Karyawan Akademik Bagian Teknik Sipil dan Arsitektur' 'Karyawan Akademik Bagian Teknik Mesin dan Teknik Elektro'"`
	Password string `form:"password" binding:"required,min=6,max=14"`
	Error    error
}

func (form FormCreateEmployeeInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("Nama tidak boleh kosong.")),
		validation.Field(&form.NIP, validation.Required.Error("NIP tidak boleh kosong.")),
		validation.Field(&form.Role, validation.Required.Error("Jabatan tidak boleh kosong."),
			validation.In("Wakil Dekan Bidang Akademik", "Koordinator Prodi Informatika", "Koordinator Prodi Teknik Sipil", "Koordinator Prodi Teknik Elektro", "Koordinator Prodi Teknik Mesin", "Koordinator Prodi Arsitektur", "Koordinator Prodi Sistem Informasi", "Karyawan Akademik Bagian Informatika dan Sistem Informasi", "Karyawan Akademik Bagian Teknik Sipil dan Arsitektur", "Karyawan Akademik Bagian Teknik Mesin dan Teknik Elektro").Error("Jabatan harus valid.")),
		validation.Field(&form.Password, validation.Required.Error("Password tidak boleh kosong."),
			validation.Length(6, 14).Error("Password harus diantara 6 sampai 14 karakter.")),
	)
}

type FormUpdateEmployeeInput struct {
	ID    int
	Name  string `form:"name" binding:"required"`
	NIP   string `form:"nip" binding:"required"`
	Role  string `form:"role" binding:"required,oneof='Wakil Dekan Bidang Akademik' 'Koordinator Prodi Informatika' 'Koordinator Prodi Teknik Sipil' 'Koordinator Prodi Teknik Elektro' 'Koordinator Prodi Teknik Mesin' 'Koordinator Prodi Arsitektur' 'Koordinator Prodi Sistem Informasi' 'Karyawan Akademik Bagian Informatika dan Sistem Informasi' 'Karyawan Akademik Bagian Teknik Sipil dan Arsitektur' 'Karyawan Akademik Bagian Teknik Mesin dan Teknik Elektro'"`
	Error error
}

func (form FormUpdateEmployeeInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("Nama tidak boleh kosong.")),
		validation.Field(&form.NIP, validation.Required.Error("NIP tidak boleh kosong.")),
		validation.Field(&form.Role, validation.Required.Error("Jabatan tidak boleh kosong."),
			validation.In("Wakil Dekan Bidang Akademik", "Koordinator Prodi Informatika", "Koordinator Prodi Teknik Sipil", "Koordinator Prodi Teknik Elektro", "Koordinator Prodi Teknik Mesin", "Koordinator Prodi Arsitektur", "Koordinator Prodi Sistem Informasi", "Karyawan Akademik Bagian Informatika dan Sistem Informasi", "Karyawan Akademik Bagian Teknik Sipil dan Arsitektur", "Karyawan Akademik Bagian Teknik Mesin dan Teknik Elektro").Error("Jabatan harus valid.")),
	)
}

type FormUpdatePasswordEmployeeInput struct {
	ID            int
	Password      string `form:"password1" binding:"required,min=6,max=14"`
	NewPassword   string `form:"password2" binding:"required,min=6,max=14"`
	Error         error
	ErrorPassword string
}

func (form FormUpdatePasswordEmployeeInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Password, validation.Required.Error("Password tidak boleh kosong."),
			validation.Length(6, 14).Error("Password harus diantara 6 sampai 14 karakter.")),
		validation.Field(&form.NewPassword, validation.Required.Error("Password tidak boleh kosong."),
			validation.Length(6, 14).Error("Password harus diantara 6 sampai 14 karakter.")),
	)
}

type LoginInput struct {
	Nip           string `form:"nip" binding:"required"`
	Password      string `form:"password" binding:"required"`
	Error         error
	ErrorPassword error
}

func (form LoginInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Nip, validation.Required.Error("Input tidak boleh kosong.")),
		validation.Field(&form.Password, validation.Required.Error("Input tidak boleh kosong.")),
	)
}

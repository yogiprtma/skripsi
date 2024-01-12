package subject

import validation "github.com/go-ozzo/ozzo-validation"

type FormCreateSubjectInput struct {
	Name  string `form:"name" binding:"required"`
	Error error
}

func (form FormCreateSubjectInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("Perihal tidak boleh kosong.")),
	)
}

type FormUpdateSubjectInput struct {
	ID    int
	Name  string `form:"name" binding:"required"`
	Error error
}

func (form FormUpdateSubjectInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("Perihal tidak boleh kosong.")),
	)
}

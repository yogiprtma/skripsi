package subject

type UserDataFormatter struct {
	ID   int
	Name string
	No   int
}

func FormatDataSubject(subject Subject) UserDataFormatter {
	formatter := UserDataFormatter{
		ID:   subject.ID,
		Name: subject.Name,
	}

	return formatter
}

func FormatDataSubjects(subjects []Subject) []UserDataFormatter {
	subjectsFormatter := []UserDataFormatter{}

	for index, subject := range subjects {
		subjectFormatter := FormatDataSubject(subject)
		subjectFormatter.No = index + 1

		subjectsFormatter = append(subjectsFormatter, subjectFormatter)
	}

	return subjectsFormatter
}

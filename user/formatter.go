package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	NPM        string `json:"npm"`
	Department string `json:"department"`
}

func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		NPM:        user.NPM,
		Department: user.Department,
	}

	return formatter
}

type UserDataFormatter struct {
	ID         int
	Name       string
	NPM        string
	Department string
	PublicKey  string
	PrivateKey string
	No         int
}

func FormatDataUser(user User) UserDataFormatter {
	UserFormatter := UserDataFormatter{}
	UserFormatter.ID = user.ID
	UserFormatter.Name = user.Name
	UserFormatter.NPM = user.NPM
	UserFormatter.Department = user.Department
	UserFormatter.PrivateKey = user.PrivateKey
	UserFormatter.PublicKey = user.PublicKey

	return UserFormatter
}

func FormatDataUsers(users []User) []UserDataFormatter {
	usersFormatter := []UserDataFormatter{}

	for index, user := range users {
		userFormatter := FormatDataUser(user)
		userFormatter.No = index + 1
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}

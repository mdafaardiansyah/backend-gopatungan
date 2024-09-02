package user

type UserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Job   string `json:"job"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Job:   user.Job,
		Email: user.Email,
		Token: token,
	}
	return formatter
}

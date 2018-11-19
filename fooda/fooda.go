package fooda

var chat Chat

func init() {
	chat = Hipchat{}
}

type Chat interface {
	Notify(string, string) error
}

func Remind(handle string) {
	chat.Notify(handle, "Your forgot fooda!")
}

func Forgetters() ([]User, error) {
	orders, err := TodaysOrders()

	if err != nil {
		return nil, err
	}

	users, err := AllUsers()

	missing := []User{}

User:
	for _, user := range users {
		for _, order := range orders {
			if user.Email == order.User.Email {
				continue User
			}
		}

		missing = append(missing, user)
	}

	return missing, nil
}

package fooda

func GetUser(email string) (User, error) {
	user := User{}

	err := db.Get(&user, `select * from users where email=$1`, email)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func AllUsers() ([]User, error) {
	rows, err := db.Queryx(`select * from users`)

	if err != nil {
		return nil, err
	}

	users := []User{}

	for rows.Next() {
		user := User{}
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

type User struct {
	ID     int
	Email  string
	Handle string
}

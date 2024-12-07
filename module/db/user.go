package db

type User struct {
	ID   int
	Name string
	Age  int
}

func (db *database) InsertUser(name string, age int) {
	db.Exec(`INSERT INTO user (name, age) VALUES (?, ?)`, name, age)
}

func (db *database) GetUser(name string) (string, int) {
	var age int
	db.QueryRow(`SELECT * FROM user WHERE name = ?`, name).Scan(&name, &age)
	return name, age
}

func (db *database) UpdateUser(name string, age int) {
	db.Exec(`UPDATE user SET age = ? WHERE name = ?`, age, name)
}

func (db *database) DeleteUser(name string) {
	db.Exec(`DELETE FROM user WHERE name = ?`, name)
}

func (db *database) GetAllUsers() ([]*User, error) {
	rows, err := db.Query(`SELECT id, name, age FROM user`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name, &user.Age)
		users = append(users, &user)
	}
	return users, nil
}

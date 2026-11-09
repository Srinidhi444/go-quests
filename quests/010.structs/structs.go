package structs

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func NewUser(id int, name, email string, age int) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
		Age:   age,
	}
}

func (u *User) IsAdult() bool {
	if u.Age >= 18 {
		return true
	}
	return false
}

func (u User) DisplayName() string {
	s := u.Name + " <" + u.Email + ">"
	return s
}

func (u *User) UpdateEmail(email string) {
	// TODO: implement
	// Read README.md for the instructions
	u.Email = email
}

func (u *User) Birthday() {
	// TODO: implement
	// Read README.md for the instructions
	u.Age++
}

func CloneUser(u User) User {
	return User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Age:   u.Age,
	}
}

package structs

type User struct {
	ID   int
	Name string
}

var inMemoryUsers []User

func GenerateUser(userName string) (user User) {
	lastID := 0
	if len(inMemoryUsers) > 0 {
		lastID = inMemoryUsers[len(inMemoryUsers)-1].ID
	}

	user = User{}
	user.ID = lastID + 1
	user.Name = userName

	inMemoryUsers = append(inMemoryUsers, user)

	return
}

func GetUsers() []User {
	return inMemoryUsers
}

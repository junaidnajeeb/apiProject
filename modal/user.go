package modal

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

var UserList = make(map[uint64]User)

// UserAdd ss
func UserAdd(newUser User) {
	UserList[newUser.ID] = newUser
}

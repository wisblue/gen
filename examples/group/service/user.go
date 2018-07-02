package service

// User is a user
type User struct {
	UserID int    // UserID is the User ID
	Name   string // Name is the name
}

var sizeuser int
var mapusers = map[int]*User{}

// UserServices #path:"/user/"#
type UserServices struct{}

// CreateUser #route:"PUT /"#
func (*UserServices) Create(user *User) (userID int, err error) {
	sizeitems++
	user.UserID = sizeitems
	mapusers[sizeitems] = user
	return sizeitems, nil
}

// DeleteUser #route:"DELETE /{userID}"#
func (UserServices) Delete(userID int) (err error) {
	delete(mapusers, userID)
	return nil
}

// GetUser #route:"GET /{userID}"#
func (UserServices) Get(userID int) (user *User, err error) {
	return mapusers[userID], nil
}

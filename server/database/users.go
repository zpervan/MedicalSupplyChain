package database

// Representation of a user in the database.
type User struct {
	Id         string
	Username   string
	Password   string
	Name       string
	Surname    string
	Role       string
	Department string
}

// Create a new user in the database if the user does not exist.
//
// Returns an error if the creation was not successful or the user already exists.
func (d Database) CreateUser(newUser User) error {
	return nil
}

// Fetch all users from the database.
//
// Returns an error if the users cannot be fetched, i.e. database is not reachable.
func (d Database) FetchUsers() error {
	return nil
}

// Update the user with the specified username.
//
// Returns an error if the update was not successful.
func (d Database) UpdateUser(username string) error {
	return nil
}

// Delete the user with the specified username.
//
// Returns an error if the update was not successful.
func (d Database) DeleteUser(username string) error {
	return nil
}
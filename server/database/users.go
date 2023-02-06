package database

import (
	"errors"
	ms "server/proto"
	"time"
)

// Create a new user in the database if the user does not exist.
//
// Returns an error if the creation was not successful or the user already exists.
func (d Database) CreateUser(newUser *ms.User) error {
	if newUser == nil {
		return errors.New("cannot create user. reason: passed user data is empty")
	}

	_, err := d.handler.Exec("INSERT INTO user_account(username, password, name, surname, department, role, created_on) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		newUser.Username, newUser.Password, newUser.Name, newUser.Surname, newUser.Department, newUser.Role, time.Now().Format(time.RFC1123))

	if err != nil {
		d.logger.Error("cannot create user. reason: " + err.Error())
		return err
	}

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
func (d Database) UpdateUser(user *ms.User) error {
	return nil
}

// Delete the user with the specified username.
//
// Returns an error if the update was not successful.
func (d Database) DeleteUser(username string) error {
	return nil
}

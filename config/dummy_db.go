package config

import (
	"sampleapp/crypto"

	"errors"
)

func NewDummyUser(username, email string) *DummyUserModel {
	return &DummyUserModel{
		Username: username,
		Email: email,
	}
}

type DummyUserModel struct {
	Username string
	Password string
	Email string
	authenticated bool
}

func (u *DummyUserModel) SetPassword(password string) error {
	hash, err := crypto.PasswordEncrypt(password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

func (u *DummyUserModel) Authenticate() {
	u.authenticated = true
}

type DummyDatabase struct {
	database map[string]interface{}
}

var store DummyDatabase

func init() {
	store.database = map[string]interface{}{}
}

func DummyDB() *DummyDatabase {
	return &store
}

func (db *DummyDatabase) Exists(username string) bool {
	_, r := db.database[username]
	return r
}

func (db *DummyDatabase) SaveUser(username, email, password string) error {
	if db.Exists(username) {
		return errors.New("user \"" + username + "\" already exists")
	}

	user := NewDummyUser(username, email)
	if err := user.SetPassword(password); err != nil {
		return err
	}
	db.database[username] = user
	return nil
}

func (db *DummyDatabase) GetUser(username, password string) (*DummyUserModel, error) {
	buffer, exists := db.database[username]
	if !exists {
		return nil, errors.New("user \"" + username + "\" doesn't exists")
	}

	user := buffer.(*DummyUserModel)
	if err := crypto.CompareHashAndPassword(user.Password, password); err != nil {
		return nil, errors.New("user \"" + username + "\" doesn't exists")
	}

	return user, nil
}

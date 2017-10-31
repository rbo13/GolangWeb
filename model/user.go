package model

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"time"
)

const passwordSalt = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890abcdefghijklmnopqrstuvwxyz)(*&^%$#@!~!@#$%^&*()-=_+"

type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

func Login(email, password string) (*User, error) {
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	row := db.QueryRow(`
		SELECT id, email, firstname, lastname
		FROM user
		WHERE email = ? 
		  AND password = ?`, email, pwd)
	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}
	t := time.Now()
	_, err = db.Exec(`
		UPDATE user
		SET lastlogin = ?
		WHERE id = ?`, t, result.ID)
	if err != nil {
		log.Printf("Failed to update login time for user %v to %v: %v", result.Email, t, err)
	}
	return result, nil
}

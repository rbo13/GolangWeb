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

// User ...
type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

// Login handler
func Login(email, password string) (*User, error) {
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	stmtOut, err := db.Prepare(`SELECT id, email, firstname, lastname
		FROM user
		WHERE email = ?
		AND password = ?`)
	err = stmtOut.QueryRow(email, pwd).Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName)

	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}

	t := time.Now()
	stmt, updateErr := db.Prepare(`UPDATE user SET lastlogin = ? WHERE id = ?`)

	if updateErr != nil {
		log.Printf("Failed to update login time for user: %v to %v: %v", result.Email, t, updateErr)
	}

	stmt.Exec(t, result.ID)

	return result, nil
}

// Signup handler
// func Signup(email, password string) {
// 	hasher := sha512.New()
// 	hasher.Write([]byte(passwordSalt))
// 	hasher.Write([]byte(email))
// 	hasher.Write([]byte(password))
// 	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

// 	stmt, dberr := db.Prepare(`INSERT user SET email = ?, password = ?`)

// 	if dberr != nil {
// 		fmt.Println(dberr)
// 		return
// 	}

// 	res, insertError := stmt.Exec(email, pwd)

// 	if insertError != nil {
// 		fmt.Println(insertError)
// 	}

// 	log.Println("Successfully added new user: ", res)
// }

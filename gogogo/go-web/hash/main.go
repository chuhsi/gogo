package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/*
	密码哈希值
 */

func HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "123"
	hash, _ := HashedPassword(password)
	fmt.Println("password", password)
	fmt.Println("hashed password", hash)

	b := CheckPasswordHash(password, hash)
	fmt.Println("match", b)
}

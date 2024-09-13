package main

import (
	"errors"
	"fmt"
)

type Services struct{}

// init repo
var repo = initRepository()

func initServices() *Services {
	return &Services{}
}

// LOGIN SERVICE
func (x *Services) login(userInfo UserBasicInfo) (string, error) {
	//check the user infos is okay or lack?
	if !isValidEmail(userInfo.Email) {
		return "", errors.New("invalid e-mail")
	}
	//MAKE DB TRANSACTIONS
	userInfoDB, err := repo.login(&userInfo)
	if err != nil {
		fmt.Println("Error service", err)
		return "", err
	}
	//IS PASSWORD CORRECT?
	if !isPasswordCorrect(userInfoDB.Password, userInfo.Password) {
		return "", errors.New("invalid email or password ")
	}

	//CREATE JWT
	token, err := createJWT(userInfoDB.ID, userInfoDB.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// SIGNUP SERVICE
func (x *Services) signup(userInfo UserBasicInfo) error {
	//CHECK USER'S INFO
	if !isValidEmail(userInfo.Email) {
		return errors.New("invalid e-mail")
	} else if !isValidPassword(userInfo.Password) {
		return errors.New("invalid password")
	} else if !isValidName(userInfo.Name) {
		return errors.New("invalid name")
	} else if !isValidName(userInfo.Surname) {
		return errors.New("invalid surname")
	}
	var err error
	//HASH PASSWORD
	userInfo.Password, err = hashPassword(userInfo.Password)
	if err != nil {
		return err
	}

	//MAKE DB TRANSACTIONS
	err = repo.signup(&userInfo)
	if err != nil {
		return err
	}

	//RETURN AN ANSWER
	return nil
}

// EDIT SERVICE
func (x *Services) edit(userInfo UserBasicInfo, token string) error {
	//parse jwt and check email is equal or not?
	userID, err := parseJWT(token)
	if err != nil {
		return errors.New("something went wrong,parsejwt")
	}
	//check email
	userInfoDB, err := repo.getUserInfoDB(userID)
	if err != nil {
		fmt.Println("Error get user info")
		return err
	}

	//is user-id equal?
	if userInfoDB.ID != userID {
		return errors.New("user id and jwt user id isn't equal")
	}
	
	//check is email registered for anyother people or not
	err = repo.checkEmailIsUnique(userInfo.Email, userID)
	if err != nil {
		return err
	}
	//check credentials
	if !isValidEmail(userInfo.Email) {
		return errors.New("invalid e-mail")
	} else if !isValidName(userInfo.Name) {
		return errors.New("invalid name")
	} else if !isValidName(userInfo.Surname) {
		return errors.New("invalid surname")
	}

	
	//SAVE NEW INFOS
	err = repo.edit(userInfoDB.ID, userInfo)
	if err != nil {
		return err
	}

	//
	return nil
}

// CHANGE PASSWORD SERVICE
func (x *Services) changePassword(newPassword string, token string) error {
	//parse jwt and check email is equal or not?
	userID, err := parseJWT(token)
	if err != nil {
		return errors.New("something went wrong,parsejwt")
	}

	//hash password
	newPassword, err = hashPassword(newPassword)
	if err != nil {
		return errors.New("something went wrong,hash password")
	}

	//SAVE NEW INFOS
	err = repo.changePassword(newPassword, userID)
	if err != nil {
		return err
	}

	//
	return nil
}

// DELETE SERVICE
func (x *Services) delete(token string) error {
	userID, err := parseJWT(token)
	if err != nil {
		return err
	}
	err = repo.delete(userID)
	if err != nil {
		return err
	}
	return nil
}

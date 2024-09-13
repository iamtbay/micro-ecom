package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repository struct {
}

func initRepository() *Repository {
	return &Repository{}
}

// LOGIN
func (x *Repository) login(userInfo *UserBasicInfo) (UserInfoDB, error) {
	query := `SELECT * FROM users WHERE email=$1`
	var userInfoDB UserInfoDB
	err := conn.QueryRow(context.Background(), query, userInfo.Email).Scan(
		&userInfoDB.ID,
		&userInfoDB.Name,
		&userInfoDB.Surname,
		&userInfoDB.Email,
		&userInfoDB.Password,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return UserInfoDB{}, errors.New("invalid email or password(pass)")
		}
		return UserInfoDB{}, err
	}

	return userInfoDB, nil

}

// SIGN UP
func (x *Repository) signup(userInfo *UserBasicInfo) error {
	query := `INSERT INTO users(name,surname,email,password) VALUES($1, $2, $3, $4)`
	_, err := conn.Exec(context.Background(), query, userInfo.Name, userInfo.Surname, userInfo.Email, userInfo.Password)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return errors.New("email is already in use, please try another")
			}
		}
		return err
	}
	fmt.Println("succesful")
	return nil
}

// EDIT
func (x *Repository) edit(userID int32, newUserInfo UserBasicInfo) error {
	query := `UPDATE users
			SET 
				name=$1, surname=$2, email=$3
			WHERE user_id=$4`
	_, err := conn.Exec(context.Background(), query,
		newUserInfo.Name,
		newUserInfo.Surname,
		newUserInfo.Email,
		userID,
	)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("something went wrong")
	}
	return nil
}

// CHANGE PASSWORD
func (x *Repository) changePassword(newPassword string, userID int32) error {
	query := `UPDATE users
			SET password=$1
			WHERE user_id=$2`
	_, err := conn.Exec(context.Background(), query,
		newPassword,
		userID,
	)
	if err != nil {
		return errors.New("something went wrong")
	}
	return nil
}

// DELETE USER
func (x *Repository) delete(userID int32) error {
	query := `UPDATE users SET name='', surname='', email='', password='' WHERE user_id=$1`
	_, err := conn.Exec(context.Background(), query, userID)
	if err != nil {
		return err
	}
	return nil
}

// QUERY METHODS

func (x *Repository) getUserInfoDB(userID int32) (UserInfoDB, error) {
	query := `SELECT * FROM users WHERE user_id=$1`
	fmt.Println("user id is", userID)
	var userInfo UserInfoDB
	err := conn.QueryRow(context.Background(), query, userID).Scan(
		&userInfo.ID,
		&userInfo.Name,
		&userInfo.Surname,
		&userInfo.Email,
		&userInfo.Password,
	)
	if err != nil {
		fmt.Println(err)
		return userInfo, err
	}
	return userInfo, nil
}

// check e-mail to edit user email
func (x *Repository) checkEmailIsUnique(email string, userID int32) error {
	var count int
	var dbUserID int32
	query := `SELECT user_id, count(*) over () as total_count
		FROM users WHERE email=$1`
	err := conn.QueryRow(context.Background(), query, email).Scan(&dbUserID, &count)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	if dbUserID != userID {
		fmt.Println(dbUserID, userID)
		return errors.New("error isn't equal user ids")
	}
	return nil
}

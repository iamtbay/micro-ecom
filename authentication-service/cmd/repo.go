package main

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repository struct {
}

func initRepository() *Repository {
	return &Repository{}
}

// !
// CHECK USER
func (x *Repository) checkUser(userID uuid.UUID) (UserInfoDB, error) {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//query
	query := `SELECT user_id,name,surname,email FROM users WHERE user_id=$1`
	var userInfoDB UserInfoDB
	err := conn.QueryRow(ctx, query, userID).Scan(
		&userInfoDB.ID,
		&userInfoDB.Name,
		&userInfoDB.Surname,
		&userInfoDB.Email,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return UserInfoDB{}, errors.New("user couldn't verify")
		}
		return UserInfoDB{}, err
	}

	return userInfoDB, nil

}

// !
// LOGIN
func (x *Repository) login(userInfo *UserBasicInfo) (UserInfoDB, error) {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `SELECT * FROM users WHERE email=$1`
	var userInfoDB UserInfoDB
	err := conn.QueryRow(ctx, query, userInfo.Email).Scan(
		&userInfoDB.ID,
		&userInfoDB.Name,
		&userInfoDB.Surname,
		&userInfoDB.Email,
		&userInfoDB.Password,
		&userInfoDB.IsAdmin,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return UserInfoDB{}, errors.New("invalid email or password")
		}
		return UserInfoDB{}, err
	}

	return userInfoDB, nil

}

// !
// SIGN UP
func (x *Repository) signup(userInfo *UserBasicInfo) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO users(name,surname,email,password) VALUES($1, $2, $3, $4)`
	_, err := conn.Exec(ctx, query, userInfo.Name, userInfo.Surname, userInfo.Email, userInfo.Password)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return errors.New("email is already in use, please try another email")
			}
		}
		return err
	}
	return nil
}

// !
// EDIT
func (x *Repository) edit(userID uuid.UUID, newUserInfo UserBasicInfo) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `UPDATE users
			SET 
				name=$1, surname=$2, email=$3
			WHERE 
				user_id=$4`
	_, err := conn.Exec(ctx, query,
		newUserInfo.Name,
		newUserInfo.Surname,
		newUserInfo.Email,
		userID,
	)
	if err != nil {
		return errors.New("something went wrong while infos are updating")
	}
	return nil
}

// !
// CHANGE PASSWORD
func (x *Repository) changePassword(newPassword string, userID uuid.UUID) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `UPDATE users
			SET password=$1
			WHERE user_id=$2`
	_, err := conn.Exec(ctx, query,
		newPassword,
		userID,
	)
	if err != nil {
		return errors.New("something went wrong while password is updating")
	}
	return nil
}

// !
// DELETE USER
func (x *Repository) delete(userID uuid.UUID) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `UPDATE users SET name='', surname='', email='', password='' WHERE user_id=$1`
	_, err := conn.Exec(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}

// QUERY METHODS

func (x *Repository) getUserInfoDB(userID uuid.UUID) (UserInfoDB, error) {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//query
	query := `SELECT * FROM users WHERE user_id=$1`
	var userInfo UserInfoDB
	err := conn.QueryRow(ctx, query, userID).Scan(
		&userInfo.ID,
		&userInfo.Name,
		&userInfo.Surname,
		&userInfo.Email,
		&userInfo.Password,
		&userInfo.IsAdmin,
	)
	if err != nil {
		return userInfo, err
	}
	return userInfo, nil
}

// check e-mail to edit user email
func (x *Repository) checkEmailIsUnique(email string, userID uuid.UUID) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var count int
	var dbUserID uuid.UUID
	query := `SELECT user_id, count(*) over () as total_count
		FROM users WHERE email=$1`
	err := conn.QueryRow(ctx, query, email).Scan(&dbUserID, &count)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	if dbUserID != userID {
		return errors.New("email in in use")
	}
	return nil
}

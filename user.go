package main

import "log"

type User struct {
	ID    int
	Email string
}

func InsertUser(user *User) error {
	log.Println("InsertUser called ...")
	var id int
	err := db.QueryRow(`
		INSERT INTO users(email)
		VALUES ($1)
		RETURNING id
	`, user.Email).Scan(&id)
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func GetUserByID(id int) (*User, error) {
	log.Println("GetUserByID called ...")
	var email string
	err := db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    id,
		Email: email,
	}, nil
}

func RemoveUserByID(id int) error {
	log.Println("RemoveUserByID called ...")
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

func Follow(followerID, followeeID int) error {
	log.Println("Follow called ...")
	_, err := db.Exec(`
		INSERT INTO followers(follower_id, followee_id)
		VALUES ($1, $2)
	`, followerID, followeeID)
	return err
}

func Unfollow(followerID, followeeID int) error {
	log.Println("Unfollow called ...")
	_, err := db.Exec(`
		DELETE FROM followers
		WHERE follower_id=$1
		AND followee_id=$2
	`, followerID, followeeID)
	return err
}

func GetFollowerByIDAndUser(id int, userID int) (*User, error) {
	log.Println("GetFollowerByIDAndUser called ...")
	var email string
	err := db.QueryRow(`
		SELECT u.email
		FROM users AS u, followers AS f
		WHERE u.id = f.follower_id
		AND f.follower_id=$1
		AND f.followee_id=$2
		LIMIT 1
	`, id, userID).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    id,
		Email: email,
	}, nil
}

func GetFollowersForUser(id int) ([]*User, error) {
	log.Println("GetFollowersForUser called ...")
	rows, err := db.Query(`
		SELECT u.id, u.email
		FROM users AS u, followers AS f
		WHERE u.id=f.follower_id
		AND f.followee_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		users = []*User{}
		uid   int
		email string
	)
	for rows.Next() {
		if err = rows.Scan(&uid, &email); err != nil {
			return nil, err
		}
		users = append(users, &User{ID: id, Email: email})
	}
	return users, nil
}

func GetFolloweeByIDAndUser(id int, userID int) (*User, error) {
	log.Println("GetFolloweeByIDAndUser called ...")
	var email string
	err := db.QueryRow(`
		SELECT u.email
		FROM users AS u, followers AS f
		WHERE u.id = f.followee_id
		AND f.followee_id=$1
		AND f.follower_id=$2
		LIMIT 1
	`, id, userID).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    id,
		Email: email,
	}, nil
}

func GetFolloweesForUser(id int) ([]*User, error) {
	log.Println("GetFolloweesForUser called ...")
	rows, err := db.Query(`
		SELECT u.id, u.email
		FROM users AS u, followers AS f
		WHERE u.id=f.follower_id
		AND f.follower_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		users = []*User{}
		uid   int
		email string
	)
	for rows.Next() {
		if err = rows.Scan(&uid, &email); err != nil {
			return nil, err
		}
		users = append(users, &User{ID: id, Email: email})
	}
	return users, nil
}

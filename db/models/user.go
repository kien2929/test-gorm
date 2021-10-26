package models

type User struct {
	Model
	username     string
	email        string
	token        string
	userId       string
	password     string
	avatarId     int
	refreshToken string
}

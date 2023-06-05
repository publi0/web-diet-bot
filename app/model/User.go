package model

type User struct {
	Id        string
	LastPhoto string
	Auth      WebDietAuth
}

type WebDietAuth struct {
	N string
	P string
}

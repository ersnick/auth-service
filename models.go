package main

type User struct {
	ID    string // GUID
	Email string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

type AuthDetails struct {
	UserId string
	Ip     string
}

package main

type Jwt struct {
	Route      string      `default:"/auth/login"`
	Credential Credentials `json:"credential"`
}

type Credentials struct {
	User     string
	Password string
}

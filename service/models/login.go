package models

type LoginRequest struct {
	Type       string
	Identifier string
	Password   string
}

type LoginResponse struct {
}

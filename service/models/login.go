package models

type LoginRequest struct {
	Group      string `bson:"group"`
	Type       string `bson:"type"`
	Identifier string `bson:"identifier"`
	Password   string `bson:"password"`
}

type LoginResponse struct {
	TwoFactor    bool   `bson:"twoFactor"`
	Token        string `bson:"token"`
	JsonWebToken string `bson:"jsonWebToken"`
}

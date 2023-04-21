package endpoints

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"github.com/go-uniform/uniform/common/nosql"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"service/service/_base"
	"service/service/info"
	"service/service/models"
	"time"
)

func init() {
	_base.Subscribe("endpoints.post.auth.login", authLogin)
}

func authLogin(r uniform.IRequest, p diary.IPage) {
	var model models.LoginRequest
	r.Read(&model)

	// todo: get database & collection information based on type (logic service call)
	database := "test"
	collection := "administrators"
	var account models.Account

	db := nosql.Connector(r.Conn(), p, "")
	db.FindOne(r.Remainder(), database, collection, "", 0, bson.D{}, &account)

	if account.TwoFactor {
		if r.CanReply() {
			if err := r.Reply(uniform.Request{
				Model: models.LoginResponse{
					TwoFactor: true,
					Token:     "xyz123",
				},
			}); err != nil {
				p.Error("reply", err.Error(), diary.M{
					"err": err,
				})
			}
		}
		return
	}

	var out struct {
		TwoFactor  bool                   `json:"two-factor"`
		Issuer     string                 `json:"issuer"`
		Audience   string                 `json:"audience"`
		ExpiresAt  time.Time              `json:"expires-at"`
		ActivateAt *time.Time             `json:"activate-at"`
		Inverted   bool                   `json:"inverted"`
		Tags       []string               `json:"tags"`
		Links      map[string][]string    `json:"links"`
		Meta       map[string]interface{} `json:"meta"`
	}

	claims := jwt.MapClaims{
		"id":                   account.Id,
		"group":                model.Group,
		"type":                 model.Type,
		"permissions.inverted": out.Inverted,
		"permissions.tags":     out.Tags,
		"links":                out.Links,
		"meta":                 out.Meta,
		"aud":                  out.Audience,
		"exp":                  out.ExpiresAt.Unix(),

		// subtract 30 min from issue time to handle potentially unsynced server times
		"iat": time.Now().Add(time.Minute * -30).Unix(),
	}

	if out.ActivateAt != nil {
		claims["nbf"] = out.ActivateAt.Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(info.Get("jwt.key"))
	if err != nil {
		panic(err)
	}

	signed, err := token.SignedString(rsaPrivateKey)
	if err != nil {
		panic(err)
	}

	if r.CanReply() {
		if err := r.Reply(uniform.Request{
			Model: models.LoginResponse{
				TwoFactor:    false,
				Token:        "",
				JsonWebToken: signed,
			},
		}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}

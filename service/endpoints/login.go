package endpoints

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
	"service/service/models"
)

func init() {
	_base.Subscribe("endpoints.post.auth.login", authLogin)
}

func authLogin(r uniform.IRequest, p diary.IPage) {
	var model models.LoginRequest
	r.Read(&model)

	// todo: handle login routine

	if r.CanReply() {
		if err := r.Reply(uniform.Request{
			Model: models.LoginResponse{},
		}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}

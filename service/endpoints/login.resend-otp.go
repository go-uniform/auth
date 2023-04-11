package endpoints

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
)

func init() {
	_base.Subscribe("endpoints.post.auth.login.resend-otp", authLoginResendOtp)
}

func authLoginResendOtp(r uniform.IRequest, p diary.IPage) {
	var model interface{}
	r.Read(&model)

	// todo: handle login routine

	if r.CanReply() {
		if err := r.Reply(uniform.Request{}); err != nil {
			p.Error("reply", err.Error(), diary.M{
				"err": err,
			})
		}
	}
}

package service

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"net/http"
	"service/service/_base"
	"service/service/info"
	"sync"
	"time"
)

func RunBefore(shutdown chan bool, group *sync.WaitGroup, p diary.IPage) {
	var endpoints = []struct {
		Timeout time.Duration
		Path    string
		Method  string
		Topic   string
	}{
		{
			Timeout: time.Second * 5,
			Path:    "/auth/login",
			Method:  http.MethodPost,
			Topic:   "endpoints.post.auth.login",
		},
		{
			Timeout: time.Second * 5,
			Path:    "/auth/logout",
			Method:  http.MethodPost,
			Topic:   "endpoints.post.auth.logout",
		},
		{
			Timeout: time.Second * 5,
			Path:    "/auth/login/resend-otp",
			Method:  http.MethodPost,
			Topic:   "endpoints.post.auth.login.resend-otp",
		},
		{
			Timeout: time.Second * 5,
			Path:    "/auth/login/otp",
			Method:  http.MethodPost,
			Topic:   "endpoints.post.auth.login.otp",
		},
	}
	for _, endpoint := range endpoints {
		if err := info.Conn.Request(p, _base.TargetAction("api", "bind"), time.Second*10, uniform.Request{
			Model: endpoint,
		}, func(r uniform.IRequest, p diary.IPage) {
			if r.HasError() {
				panic(r.Error())
			}
			if r.CanReply() {
				if err := r.Reply(uniform.Request{}); err != nil {
					p.Error("reply", err.Error(), diary.M{
						"err": err,
					})
				}
			}
		}); err != nil {
			panic(err)
		}
	}
}

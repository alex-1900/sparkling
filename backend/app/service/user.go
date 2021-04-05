package service

import "sync"

type Auth struct {
}

var serviceAuth *Auth
var onceAuth sync.Once

func GetAuth() *Auth {
	onceAuth.Do(func() {
		serviceAuth = new(Auth)
	})
	return serviceAuth
}

func (s *Auth) Register() string {
	return "hello"
}

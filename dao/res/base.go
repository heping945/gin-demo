package res

import "gin-demo/dao"

type LoginResponse struct {
	User      dao.User `json:"user"`
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expiresAt"`
}

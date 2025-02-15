package typestore

import "context"

type User struct {
	ID       string `json:"userId"   required:"true"`
	Name     string `json:"name"     required:"min=3,max=15"`
	Username string `json:"username" required:"min=3,max=15"`
	Password string `json:"-"        required:"min=3,max=15"`
}

type UserStore interface {
	GetUserByID(ctx context.Context, id string) (*User, error)
}

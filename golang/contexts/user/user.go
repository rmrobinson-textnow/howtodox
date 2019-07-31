package user

import "golang.org/x/net/context"

type ctxKey int

const (
	userKey ctxKey = iota
)

type User struct {
	Username string
	Name     string
	Age      int
}

func NewContext(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func FromContext(ctx context.Context) (*User, bool) {
	user, ok := ctx.Value(userKey).(*User)
	return user, ok
}

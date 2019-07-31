package connection

import "golang.org/x/net/context"

type ctxKey int

const (
	connKey ctxKey = iota
)

type Connection struct {
	IPAddress string
	UserAgent string
}

func NewContext(ctx context.Context, conn Connection) context.Context {
	return context.WithValue(ctx, connKey, conn)
}

func FromContext(ctx context.Context) (Connection, bool) {
	user, ok := ctx.Value(connKey).(Connection)
	return user, ok
}

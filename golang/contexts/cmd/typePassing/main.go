package main

import (
	"fmt"

	"github.com/rmrobinson-textnow/howtodox/golang/contexts/connection"
	"github.com/rmrobinson-textnow/howtodox/golang/contexts/user"
	"golang.org/x/net/context"
)

func requestHandler(ctx context.Context, body string) {
	conn, ok := connection.FromContext(ctx)

	if ok {
		fmt.Printf("Connection from %+v\n", conn)
	} else {
		fmt.Printf("Unknown connnection source, but not fatal\n")
	}

	u, ok := user.FromContext(ctx)

	if !ok {
		fmt.Printf("User not found, cannot process '%s'!\n", body)
		return
	}

	fmt.Printf("User %+v made request '%s'\n", u, body)

	u.Age++

	if u.Age == 32 {
		requestHandler(user.NewContext(ctx, u), "new request")
	}

	fmt.Printf("User %+v\n", u)
}

func userMiddlewearHandler(ctx context.Context, username string) context.Context {
	// Do a lookup somewhere to retrieve the user.
	// Hardcoded here as an example
	u := &user.User{
		Username: username,
		Name:     "John Doe",
		Age:      31,
	}

	return user.NewContext(ctx, u)
}

func connMiddlewearHandler(ctx context.Context, ip string, ua string) context.Context {
	conn := connection.Connection{
		IPAddress: ip,
		UserAgent: ua,
	}

	return connection.NewContext(ctx, conn)
}

func main() {
	reqCtx := context.Background()
	reqCtx = connMiddlewearHandler(reqCtx, "192.168.10.10", "IE6, because why not?")
	reqCtx = userMiddlewearHandler(reqCtx, "testUsername")

	fmt.Printf("Okay request:\n")
	requestHandler(reqCtx, "request body")
	fmt.Printf("\n")

	fmt.Printf("Invalid request:\n")
	requestHandler(context.Background(), "second request body")
	fmt.Printf("\n")
}

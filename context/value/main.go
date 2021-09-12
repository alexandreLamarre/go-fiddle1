package main

import (
	"context"
	"fmt"
)

type database map[string]bool
type userIdKey string

var db database = database{
	"jane": true,
}

func processRequest(ctx context.Context, userid string) {
	vctx := context.WithValue(ctx, userIdKey("userIDKey"), "jane")
	ch := checkMembership(vctx)
	status := <-ch
	fmt.Printf("membership status of user id : %v : %v\n", userid, status)
}

func checkMembership(ctx context.Context) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		userid := ctx.Value(userIdKey("userIDKey")).(string)
		status := db[userid]
		ch <- status
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	processRequest(ctx, "jane")
}

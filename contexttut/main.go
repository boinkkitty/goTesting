package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	exampleTimeout(ctx)
	exampleWithValues(ctx)
}

func exampleTimeout(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 3)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Call api")
	case <-ctxWithTimeout.Done():
		fmt.Println("Timeout expired", ctxWithTimeout.Err())
	}
}

func exampleWithValues(ctx context.Context) {
	type key int
	const UserKey = 0

	ctxWithValue := context.WithValue(ctx, UserKey, "123")

	if userID, ok := ctxWithValue.Value(UserKey).(string); ok {
		fmt.Println("this is userId", userID)
	} else {
		fmt.Println("this is aprotected route - no userID foud")
	}
}

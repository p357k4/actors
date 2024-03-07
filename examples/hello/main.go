package main

import (
	"actors/pkg"
	"context"
	"log/slog"
	"time"
)

// myActor is a concrete implementation of the Receiver interface
type myActor[T any] struct{}

// NewMyReceiver creates a new instance of myActor
func NewMyReceiver[T any]() pkg.Receiver[T] {
	return &myActor[T]{}
}

// Receive implements the Receiver interface and receives messages
func (a *myActor[T]) Receive(ctx context.Context, msg T) error {
	slog.InfoContext(ctx, "Custom logic", slog.Any("msg", msg))
	return nil
}

func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//create system
	system := pkg.NewSystem()

	// create actors
	actorStringRef := pkg.Start(system, ctx, func() pkg.Receiver[string] {
		return NewMyReceiver[string]()
	})
	intActorRef := pkg.Start(system, ctx, func() pkg.Receiver[int] {
		return NewMyReceiver[int]()
	})

	// Send string messages with the custom logic
	actorStringRef.Send(ctx, "Hello")
	actorStringRef.Send(ctx, "World")

	// Send integer messages without custom logic (uses default logic)
	intActorRef.Send(ctx, 42)
	intActorRef.Send(ctx, 100)

	time.Sleep(2 * time.Second)
	cancel()

	system.Wait()
}

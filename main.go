package main

import (
	"fmt"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

type Hello struct {
	Who string
}

type GreeterActor struct{}

// Receive is called when the actor receives a message.
func (state *GreeterActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *Hello:
		fmt.Printf("Hello, %s\n", msg.Who)
	}
}

func main() {
	// create a new actor system
	system := actor.NewActorSystem()

	// create a new actor props
	props := actor.PropsFromProducer(
		func() actor.Actor {
			return &GreeterActor{}
		},
	)

	// create a new actor
	pid := system.Root.Spawn(props)

	// send a message to the actor
	system.Root.Send(pid, &Hello{Who: "World"})

	// wait for a while so the actor has time to process the message
	time.Sleep(100 * time.Millisecond)
}

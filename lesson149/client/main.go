package main

import "github.com/nats-io/nats.go"

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)

	defer nc.Close()

	nc.Publish("foo", []byte("Hello world"))

}

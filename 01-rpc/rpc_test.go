package main

import (
	"testing"

	"github.com/GuilhermeVendramini/microservices-in-go/01-rpc/client"
	"github.com/GuilhermeVendramini/microservices-in-go/01-rpc/server"
)

func BenchmarkDial(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c := client.CreateClient()
		c.Close()
	}
}

// go test -v -run="none" -bench=. -benchtime="5s"
func BenchmarkHelloWorldHandler(b *testing.B) {
	b.ResetTimer()

	c := client.CreateClient()

	for i := 0; i < b.N; i++ {
		client.PerformRequest(c)
	}

	c.Close()

}

func init() {
	// start the server
	go server.StartServer()
}
package main

import "github.com/dyara/dyara/pkg/server"

func main() {
	s, err := server.New()
	if err != nil {
		panic(err)
	}

	err = s.Run()
	if err != nil {
		panic(err)
	}
}

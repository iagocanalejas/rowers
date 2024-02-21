package main

import (
	"rowers/pkg/service"
)

func main() {
	service := service.NewServer()

	err := service.ListenAndServe()
	if err != nil {
		panic("cannot start service")
	}
}

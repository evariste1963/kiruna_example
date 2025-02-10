package main

import "kiruna_example/internal/platform"

func main() {
	err := platform.Kiruna.Build()
	if err != nil {
		panic(err)
	}
}

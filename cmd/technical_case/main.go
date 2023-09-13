package main

import (
	"jump/technical-case/internal/bootstrap"
)

func main() {
	r := bootstrap.Router()
	if err := r.Run("0.0.0.0:80"); err != nil {
		panic("Unable to start HTTP server")
	}
}

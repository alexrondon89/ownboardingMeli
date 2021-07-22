package main

import (
	"ownboardingMeli/internal/server"
	"ownboardingMeli/internal/server/Initial"
)

func main () {
	init := Initial.NewInit()
	router := server.GetMeliServer(init)
	router.Run(":8080")
}

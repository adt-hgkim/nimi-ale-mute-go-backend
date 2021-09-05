package main

import (
	"github.com/adt-hgkim/nimi-ale-mute/database"
	"github.com/adt-hgkim/nimi-ale-mute/plugin"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	plugin.InitRouter(router)
	database.Migrate(&database.User{}, &database.Word{})

	if err := router.Run(); err != nil {
		panic(err)
	}
}

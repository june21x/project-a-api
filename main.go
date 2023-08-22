package main

import (
	"github.com/june21x/project-a-api/api/route"
	initialization "github.com/june21x/project-a-api/init"
	"github.com/june21x/project-a-api/util"
)

func main() {
	// Load config
	util.LoadConfig()

	init := initialization.Initialize()

	app := route.Init(init)

	app.Run() // listen and serve on 0.0.0.0:8080

}

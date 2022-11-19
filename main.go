package main

import "github.com/dhiemaz/alien-invasion/cmd"

func main() {
	app := cmd.NewCommand()
	app.Start()
}

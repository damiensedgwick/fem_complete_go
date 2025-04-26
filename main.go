package main

import "fem_complete_go/internal/app"

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("Hello World")
}

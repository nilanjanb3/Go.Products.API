package main

func main() {
	app := App{}
	app.Initialize()
	logPanic(app.Run("localhost:3000"))
}

package main

func main() {
	app := App{}
	app.Initialize()
	HandleError(app.Run("localhost:3000"))
}

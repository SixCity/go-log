package main

func main() {

	e := NewRouter()

	e.Logger.Fatal(e.Start(":1222"))
}

package main

func main() {
	router := httprouter.New()
	router.GET("/", IndexHandler)
}

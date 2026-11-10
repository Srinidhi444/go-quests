package go_routine

func SendRequest(input string) string {
	// TODO: implement
	// Read README.md for the instructions
	ch := make(chan string)
	go Server(input, ch)
	return <-ch

}

func Server(request string, response chan string) {
	// TODO: implement
	// Read README.md for the instructions
	res := "processed: " + request

	response <- res

}

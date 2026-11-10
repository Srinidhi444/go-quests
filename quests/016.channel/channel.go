package channel

func ReadFromBoth(ch1 chan string, ch2 chan string) string {
	val1 := <-ch1
	val2 := <-ch2

	return "read: " + val1 + " & " + val2
}

func WriteToBoth(ch1 chan string, ch2 chan string, msg string) {
	go func() {
		ch1 <- "write: " + msg
	}()

	go func() {
		ch2 <- "write: " + msg
	}()
}

func ReadThenWrite(input chan string, output chan string) {
	val := <-input

	output <- "transform: " + val
}

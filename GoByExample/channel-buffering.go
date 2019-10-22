package main

import "fmt"

/**
By default channels are unbuffered,
meaning that they will only accept sends (chan <-)
if there is a corresponding receive (<- chan) ready to receive the sent value.
Buffered channels accept a limited number of values without
a corresponding receiver for those values.
 */

func main()  {
	// make a channel of string buffering up to 2 values
	message := make(chan string, 2)
	message <- "aaa"
	message <- "bbb"
	//message <- "ccc"

	fmt.Println(<-message)
	fmt.Println(<-message)
	//fmt.Println(<-message)
}
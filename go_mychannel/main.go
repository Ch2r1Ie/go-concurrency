package gomychannel

import "fmt"

func Channels() {

	myChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	msg := <-myChannel

	fmt.Println(msg)
}

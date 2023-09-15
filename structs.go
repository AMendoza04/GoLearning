package main

import "fmt"

type messageToSend struct{
	phoneNumber int
	message string
}

func test(m messageToSend){
	fmt.Printf("Sending message: %s to %v\n", m.message, m.phoneNumber)
	fmt.Println("==================================")
}

func main(){
	test(messageToSend{ phoneNumber: 1122334455, message: "Calling you back later",})
}

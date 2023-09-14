package main

import "fmt"

func main(){
//functions initial
		//sendsSoFar := 430
		//const sendsToAdd = 25
		//sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
		//fmt.Println("you ve sent ", sendsSoFar, "messages")

//ignore
		//firstName, _/*lastName*/ := getNames()
		//fmt.Println("Welcome to Go dear ", firstName)


		
}

func getNames()(string, string){
	return "Daniel", "Mendoza"
}

func incrementSends(sendsSoFar, sendsToAdd int) int{

	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}

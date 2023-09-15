package main

import( 
	"fmt"
)

func sendMesagge(msg message){
	fmt.Println(msg.getMessage())
}
type message interface{
	getMessage() string
}
type birthdayMessage struct{
	birthdayTime string
	recipientName string
}
func (bm birthdayMessage) getMessage() string{
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime)
}
type sendingReport struct{
	reportName string
	numberOfSeconds int
}
func (sr sendingReport) getMessage() string{
	return fmt.Sprintf("Your %s report is ready. Youve sent %v messages", sr.reportName, sr.numberOfSeconds)
}

func test(m message){
	sendMesagge(m)
	fmt.Println("==============================================")
}

func main(){
	test(sendingReport{
		reportName: "first report",
		numberOfSeconds: 10,
	})
	test(birthdayMessage{
		recipientName: "Jhon Doe",
		birthdayTime: "1994/ 03/ 21",
	})
	test(sendingReport{
		reportName: "first report",
		numberOfSeconds: 10,
	})
	
}

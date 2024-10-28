//struct similar to objects in javascript

//key value

package main

import (
	"fmt"
)

type messageToSend struct {
	message string
	phoneNumber int
}

func testi(m messageToSend){  //messageToSend here is the datatype for the structs
	fmt.Printf("sending message %s, to  contact %d\n ", m.message, m.phoneNumber)
	fmt.Println("==============================================")
}



type mrRobot struct {
	character string
	monologue string
	partners int
}


func test(r mrRobot){
	fmt.Printf("%s said '%s' while he has %d balls", r.character, r.monologue, r.partners)
}

func main1(){
   test(mrRobot{
	monologue: "fuck society" ,
	character: "Elliot",
	partners: 0,
   })
}


//methods on struct 


type rect struct {
	height int
	width int 	
}

func(r rect) area ()int {
	return r.height * r.width
   }






type basicAuthentication struct {
	username string
	password int
}
func auth (b basicAuthentication) string {
    return fmt.Sprintf("%s with the password of %d has been authorized", b.username, b.password)
}


func main(){
    result := auth(basicAuthentication{
		username : "Shivansh",
		password : 320320,
	})
	fmt.Printf(result)
}



package main

import (
	"fmt"
)

// three primitive datatypes in any lang gotta be string, number and boolean.

// similarly in go, we go

// 1. String
// 2. Bool
// 3. Numbers (this is where things start to get a little complicated)

// so numbers are
// int, int8, int16, int32, int64  (Whole Numbers)
// uint, uint8, uint16, uint32, uint64, uintptr (Positive Whole Numbers)

// byte and rune are just fancy names for uint8 and int32, idk what they did to deserve that fancy but we'll see

// float32, float64
//these are just for fractional and decimal values

// complex64, complex128
//and these are for represeting imaginary numbers, think of a usecase where u might need them?

//task1
func unmain0(){

	var smsSendingLimit int
	var costPerSMS float64 //u fucked here but cool not that it wa
	var hasPermission bool
	var username string 


	fmt.Printf(
		"%v %f %v %q\n", //these weird looking operands are for string formatting
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
	)
}

// unlike javascript we don't have to assign our variables a value we can just initialize em and print or use wtv


// Theory-101
//yay, now we'll shorten the syntax of our assigning a varible (:=)

// func main2(){
// to assign a string 
	//empty:= "" //this is same as,
	//var empty string //lmao it is same thats why you'll get an compiler error herel;
//to assign a number 
    // numBitches := 10

// to assign a float 
//    numInterest := 0.0;

//to assign a boolean
//   isFunny := true
// }

// := this called short assignment operator

//task2

func unmain1(){
	numCar := 10 
    fmt.Printf("Integer : %T", numCar)

}  //looks like we got our first bug in go 


//Damn seems like it only prints the function named main, try figuring out why that is though

// We ran into shit tons of errors i should have just documented but anyway so seems seems like printf is for printing strings only, cool


// When dealing with numbers most of our lifting will be done through;

// int
// uint
// float64
// complex128 (not this LoL)

//conversion of datatypes
func unmain2(){
	temp := 21.323
    
	tempNew := int(temp)
	fmt.Println(tempNew) //u need to use different print operators for printing diff datatypes
	fmt.Printf("type of : %T", tempNew)
}

// yes so now we'll discuss about const 
//just like in javascript we got const varible which doesn't mutate

func unmain3(){
	const firstName = "shivansh"
	const lastName = "srivastav"

	const fullName = firstName + " " + lastName
	fmt.Printf(fullName)
}


//now we are getting close to conditionals and believe me they're very similar to js

func main(){
	messagelength := 11
	messageMaxLength := 20
	fmt.Println("trying to send the message ", messagelength) //u can't concat strings w "+" but with a comma

	if messagelength <= messageMaxLength {
		fmt.Println("message sent")
	} else {
		fmt.Println("message to long")
	}
	if age:= getAge(person); age > 18 { //another cool feature about golang which i just got to know is about conditionals


		\\\Println("minor")
	}
};


package main

import "fmt"

//function signature

//so just like while defining a varible in go we need to define its datatype aswell, in functions also we need to provide
// its datatype for its input, parameters and outputs

//so main function shouldn't have any argument or return value

func add(x string , y string) string {
	return x + y
}



// so as far as i've coded i have noticed that only main function gets printed 




//Multiple Parameters 

// so another syntactic sugar provided by our golang in this we can declare many parameter in just one function for example






//ran into an hour of syntax error but anyway u understand whys it good to have multiple parameter


//pass by value

func incrementSends(sendSoFar, sendToAdd int) int {
	sendSoFar = sendSoFar + sendToAdd
	return sendSoFar 
}
//u failed miserably here, do it again tommarrow


//ignored return values 
func getNames()(string, string) {
   return "shiv" ,"sri"
}


// func yearsUntilEvent() (yearsUntilAdult int,
// 	 yearsUntilDrink int, 
// 	 yearsUntilHome int) {
// 	yearsUntilAdult = 18 - age
// 	if yearsUntilAdult < 0 {
// 		yearsUntilAdult = 0
// 	}
// 	yearsUntilDrink = 21 - age 
// 	if yearsUntilDrink < 0 {
// 		yearsUntilDrink = 0
// 	}
// 	yearsUntilHome = 25 - age 
// 	if yearsUntilHome < 0 {
// 		yearsUntilHome = 0
// 	}
// 	return
// }



func main(){
	fmt.Println("fucjk todays class dawg")
}
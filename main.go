//Name: Vincent G.
//Date: 4/13/2020
//Description: Tip calculator



package main

import "fmt"

//make a function that calculates the tip amount of a bill givena percentage.
func Tip(a, b int)(int){
  var tip = (b*a)/100
 return tip
}

func main() {
  //Declare variable for a and b as int
  var a int
  var b int 
//ask user to enter in bill amount and scan for a
 fmt.Println("Enter in the bill amount")
 fmt.Scanln(&a)
 //ask for tip percentage and scan for b
 fmt.Println("Enter in the percent you'd like to tip")
 fmt.Scanln(&b)
 //call Tip(a, b)
fmt.Println("A tip of",b,"% of",a,"is",Tip(a, b))
}

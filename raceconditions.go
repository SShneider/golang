package main; 

import (
	"fmt"
)

var Trigger = 0

func main() {
  go minus()
  go plus()
  var Counter = 0
  for true{
    if Trigger != 0{
      fmt.Printf("Final Counter: %d\n", Counter)
      return
    } else {
     Counter++
     if Counter>1000000{
       break
     }
    }
  }
  fmt.Printf("Final Counter: %d\n", Counter)
  return
}

func minus(){
  Trigger--
  fmt.Printf("Minus Counter: %d\n", Trigger)
}
func plus(){
  Trigger++
  fmt.Printf("Plus Counter: %d\n", Trigger)
}

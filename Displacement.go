package main

import (
  "fmt"
 "strconv"
 )
func main() {
  var a float64
  var v float64
  var s float64
  var t float64
  var aP string = "Please enter the acceleration:\n"
  var vP string = "\nPlease enter the initial velocity:\n"
  var sP string = "\nPlease enter the initial displacement:\n"
  var tP string = "\nPlease enter the time:\n"
  var varArr = []*float64 {&a, &v, &s, &t}
  var promptArr = []string {aP, vP, sP, tP}
  var input string
  for idx, _ := range varArr{
    fmt.Println(promptArr[idx])
    fmt.Scanln(&input)
    procNum, err := strconv.ParseFloat(input, 64)
    if err != nil {
      fmt.Println(err)
    }
    *varArr[idx] = procNum
  }
  fn := GenDisplaceFn(a, v, s)
  fmt.Println("The result is:")
  fmt.Println(fn(t))
}

func GenDisplaceFn(a, v, s float64) func(t float64) float64{
  return func(t float64) float64{
    return a/2*t*t+v*t+s
  }
}

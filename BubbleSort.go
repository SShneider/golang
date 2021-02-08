package main

import (
  "fmt"
 "bufio"
 "strconv"
 "os"
 )
func main() {
  var arrNums = make([]int, 3)
  fmt.Printf("Please enter your nums or 'X' to exit\n")
  in := bufio.NewScanner(os.Stdin)
  idx := 0
  for in.Scan(){
    input:= in.Text()
    if input == "X" || input =="x"{
      break
    }
    newNum, err := strconv.Atoi(input)
    if err != nil {
      fmt.Println(err)
    }
    if(idx<len(arrNums)){
      arrNums[idx] = newNum
    }else{
      var extraSpace = make([]int, len(arrNums))
      arrNums = append(arrNums, extraSpace...)
      arrNums[idx] = newNum
    }
    idx++
  }
  var sliceToSort = arrNums[:idx]
  //fmt.Println(sliceToSort)
  //fmt.Println(arrNums)
  fmt.Println(bubbleSort(sliceToSort))
}

func bubbleSort(arrIn []int) []int{
  var isBeingSorted bool = true
  var i int = 0
  for isBeingSorted{

    isBeingSorted = false
    i = 0 
    for i<len(arrIn)-1{
      if arrIn[i]>arrIn[i+1]{
        arrIn[i], arrIn[i+1] = arrIn[i+1], arrIn[i]
        isBeingSorted = true
      }
      i++
    }
  }
  return arrIn
}

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
  fmt.Println(RadixSort(sliceToSort))
}

func digitCount(num int) int{
  return len(strconv.Itoa(num))
}

func mostDigits(numArr []int) int{
  var max int = 0
  for _, val := range numArr{
    if digitCount(val)>max{
      max = digitCount(val)
    }
  }
  return max
}

func RadixSort(numArr []int) []int{
  var buckets = make([][]int, 10)
  var remainder int = 10
  var maxDigits int = mostDigits(numArr)
  divider := 1
  j := 0
  for j<maxDigits {
    buckets = make([][]int, 10)
    for _, val := range numArr{
      var lastDigit int = (val/divider)%remainder
      buckets[lastDigit] = append(buckets[lastDigit], val)
    }
    numArr = make([]int, 0, cap(numArr))
    for _, subarray := range buckets{
      numArr = append(numArr, subarray...)
    }
    divider*=10
    j++
  }
  //fmt.Println(buckets)
  return numArr
}

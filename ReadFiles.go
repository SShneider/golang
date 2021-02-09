package main

import (
  "fmt"
  "os"
 "bufio" 
"strings"
"log"
 )


type Data struct{
  fname string
  lname string
}

var arrOfData = make([]Data, 5)

func main() {
 scanner := bufio.NewReader(os.Stdin)
 fmt.Println("Please input file name\n")
 name, _ := scanner.ReadString('\n')
 name = strings.TrimRight(name, "\r\n")
 //Hey you, you probably need to edit how the path is generated based on where youre running and your operating system.
 i:= 0
 file, err := os.Open(name)
 if err!= nil{
   log.Fatal(err)
 }
 defer file.Close()
 fileScanner := bufio.NewScanner(file)
 for fileScanner.Scan(){
    if(i==len(arrOfData)){
      var extraSpace = make([]Data, len(arrOfData))
      arrOfData = append(arrOfData, extraSpace...)
    }
    currentLine:= fileScanner.Text()
    names := strings.Split(currentLine, " ")
    newStruct := Data{names[0], names[1]}
    arrOfData[i] = newStruct
    i++
  }
  j:= 0
  for j<i{
    fmt.Printf("First Name: %s Last Name: %s\n", arrOfData[j].fname, arrOfData[j].lname)
    j++
  }
}

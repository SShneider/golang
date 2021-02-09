package main

import (
  "fmt"
  "os"
 "bufio" 
"encoding/json"
"strings"
 )

var data = make(map[string]string)



func main() {
 scanner := bufio.NewReader(os.Stdin)
 fmt.Println("Please input name\n")
 name, _ := scanner.ReadString('\n')
 data["Name"] = strings.TrimRight(name, "\r\n")
 fmt.Println("Please input address\n")
 address, _ := scanner.ReadString('\n')
 data["Address"] = strings.TrimRight(address, "\r\n")
 userData, _ := json.Marshal(data)
os.Stdout.Write(userData)
}

package main

import (
  "fmt"
  "os"
 "bufio" 
 "strings"
 )
type Animal struct{
  eat, move, speak string
}
func (a Animal) Eat(){
   fmt.Println(a.eat)
}
func (a Animal) Move(){
   fmt.Println(a.move)
}
func (a Animal) Speak(){
   fmt.Println(a.speak)
}

func main() {
 cow := Animal{"grass", "walk", "moo"}
 bird := Animal {"worms", "fly", "peep"}
 snake := Animal {"mice", "slither", "hsss"}
 //var looping bool = true
 var input string
 scanner := bufio.NewScanner(os.Stdin)
 fmt.Println("Please in put your query in the format of Animal Action\nEnter 'x' to exit")
 for scanner.Scan(){
   input = scanner.Text()
   var myAnimal Animal
   if input == "x" {
     break
   }else{
     commandArr := strings.Split(input, " ")
     switch(commandArr[0]){
      case "cow":
        myAnimal = cow
      case "bird":
        myAnimal = bird
      case "snake":
        myAnimal = snake
      default:
        fmt.Println("Input not understood please try again")
      }
      switch(commandArr[1]){
        case "eat":
          myAnimal.Eat()
        case "move":
          myAnimal.Move()
        case "speak":
          myAnimal.Speak()
        default:
          fmt.Println("Input not understood please try again")
      }
    }
  }
}

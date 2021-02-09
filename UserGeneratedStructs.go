package main

import (
  "fmt"
  "os"
 "bufio" 
 "strings"
 )
type Animal struct{
  name, eat, move, speak string
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
var animalGenerator = make(map[string][]string)
var myAnimals = make(map[string]Animal)



func main() {
 animalGenerator["cow"] = []string{"temp", "grass", "walk", "moo"}
 animalGenerator["bird"] = []string{"temp", "worms", "fly", "peep"}
 animalGenerator["snake"] = []string{"temp", "mice", "slither", "hsss"}
 var input string
 scanner := bufio.NewScanner(os.Stdin)
 fmt.Println("Please in put your query in the format of query Animal Action\nOr newanimal Animalname Animaltype\nEnter 'x' to exit")
 for scanner.Scan(){
   input = scanner.Text()
   if input == "x" {
     break
   }else{
     commandArr := strings.Split(input, " ")
     if len(commandArr)!= 3{
       fmt.Println("Input not understood please try again")
     }else if commandArr[0]=="query"{
       queryAnimal(commandArr[1], commandArr[2])
     }else if commandArr[0]=="newanimal"{
       var err bool
       var tempAnimal Animal
       tempAnimal, err = createAnimal(commandArr[2], commandArr[1])
       if err{
        myAnimals[commandArr[1]] = tempAnimal
        fmt.Println("Created it!")
       } else{
         fmt.Println("Error creating new animal")
       }
     }else{
       fmt.Println("Input not understood please try again")
     }
    
  
    }
  }
}
func createAnimal(typeIn, nameIn string) (Animal, bool){
   var ok bool
   var myAnimal Animal
   if _, ok = animalGenerator[typeIn]; !ok{
     fmt.Println("invalid animal type")
   }else{
    myAnimal  = Animal {
    nameIn, 
    animalGenerator[typeIn][1],
    animalGenerator[typeIn][2],
    animalGenerator[typeIn][3]}
   }
   return myAnimal, ok
}
func queryAnimal(nameIn, commandIn string){
      if _, ok := myAnimals[nameIn]; !ok {
        fmt.Println("no such animal exists")
      }
      switch(commandIn){
        case "eat":
          myAnimals[nameIn].Eat()
        case "move":
          myAnimals[nameIn].Move()
        case "speak":
          myAnimals[nameIn].Speak()
        default:
          fmt.Println("Input not understood please try again")
      }
}

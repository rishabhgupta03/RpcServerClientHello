package main

import(
    "person"
    "fmt"
)

func main(){
    p := person.Person{};
    p.Name = "Rishabh";
    p.Age = 20;
    fmt.Println("Name = ", p.GetName())
    fmt.Println("Age = ", p.GetAge())
}

 // /usr/local/go/src/Golang
package main

import (
    "design.patterns/CreationalPatterns/SimpleFactory/simplefactory"
    "fmt"
)

func main()  {
    api := simplefactory.NewAPI("cn")
    server := api.Say("管仲才")
    fmt.Println(server)

    api2 := simplefactory.NewAPI("en")
    server2 := api2.Say("GZC")
    fmt.Println(server2)
}
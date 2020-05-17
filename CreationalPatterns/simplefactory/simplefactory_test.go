package simplefactory

import (
    "fmt"
    "testing"
)

func TestSay(t *testing.T)  {
    api := NewAPI("cn")
    server := api.Say("科比")
    fmt.Println(server)

    api2 := NewAPI("en")
    server2 := api2.Say("kobe")
    fmt.Println(server2)
}
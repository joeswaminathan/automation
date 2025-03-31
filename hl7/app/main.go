package main

import (
    "atian.com/hl7"
    "os"
    "fmt"
)

func main() {
    def := &hl7.Definitions{}

    file := "hl7.json"
    if len(os.Args) > 1 {
        fmt.Println("User specified File = ", file)
        file = os.Args[1]
    } else {
        fmt.Println("Default File = ", file)
    }
    def.Load(file)
    def.MakeTypes()
}

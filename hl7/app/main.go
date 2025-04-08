package main

import (
    "atian.com/hl7"
    "fmt"
)

func main() {
    def := &hl7.Definitions{}

    def.Load("hl7.json")
    fmt.Println("package hl7\n")
    def.MakeTypes()
}

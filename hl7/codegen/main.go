package main

import (
    "joeswaminathan.com/Edge_Ble/hl7"
)

func main() {
    def := &hl7.Definitions{}

    def.Load("hl7.json")
    def.Initialize()
    def.MakeTables()
    def.MakeTypes()
}

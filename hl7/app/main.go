package main

import (
	"fmt"

	"atian.com/hl7"
)

func main() {
	def := &hl7.Definitions{}

	def.Load("hl7_update.json")
	fmt.Println("package hl7\n")
	def.MakeTypes()
}

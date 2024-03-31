// gob2.go
package main

import (
	"encoding/gob"
	"log"
	"os"
	"bufio"
	"fmt"
)

type Address struct {
	Type             string
	City             string
	Country          string
}

type VCard struct {
	FirstName	string
	LastName	string
	Addresses	[]*Address
	Remark		string
}

var content	string

func main() {
	var vc VCard
	file, _ := os.Open("vcard.gob")
	defer file.Close()
	inReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inReader)
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
	fmt.Println(vc)
}
package main

import (
	"fmt"
	"strings"
	"unicode"
)

type student struct {
	name string
}

type Chiper interface {
	Encode() string
	Decode() string
}

const KEY string = "slckytvhrmpdojgefiuxwnqzba"


func (s *student) Encode() string {
	var nameEncode = ""

	// your code here

	for _, v := range s.name {
		if unicode.IsUpper(v){
			v = unicode.ToLower(v)
			nameEncode += strings.ToUpper(string(KEY[(v-97)]))
			continue
		}
		if ok := strings.IndexRune(KEY,v); ok == -1 {
			nameEncode += string(v)
			continue
		}

		nameEncode += string(KEY[(v-97)])
	}

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

  // your code here
	
	for _, v := range s.name {
		if unicode.IsUpper(v) {
			v = unicode.ToLower(v)
			runePlain := rune(strings.IndexRune(KEY,v) + 97)
			nameDecode += strings.ToUpper(string(runePlain))
			continue
		}
		if ok := strings.IndexRune(KEY,v); ok == -1 {
			nameDecode += string(v)
			continue
		}
		runePlain := rune(strings.IndexRune(KEY,v) + 97)
		nameDecode += string(runePlain)

	}

  return nameDecode
}

func main() {
	var menu int
  var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)
	
	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student’s name " + a.name + " is : " + c.Decode())
	}
}
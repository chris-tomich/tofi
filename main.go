package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

var grammar map[string]bool = map[string]bool{
	"function": true,
	"if": true,
	"this": true,
	"undefined": true,
}

var operators map[string]bool = map[string]bool{
	"==": true,
	"===": true,
	"!=": true,
	"!==": true,
	"<": true,
	">": true,
	"<=": true,
	">=": true,
}

var separators map[rune]bool = map[rune]bool{
	' ': true,
	'=': true,
	'[': true,
	']': true,
	'(': true,
	')': true,
	'\'': true,
	'"': true,
	'/': true,
	'.': true,
	'{': true,
	'}': true,
}

func main() {
	jsFile := flag.String("js", "", "the name of a javascript file")
	flag.Parse()

	if *jsFile == "" {
		fmt.Println("please provide a file using -js <file>")
	}

	buf, err := ioutil.ReadFile(*jsFile)

	if err != nil {
		panic(err)
	}

	contents := []rune(string(buf))
	
	tokens := make([][]rune, 0, len(contents))
	
	token := make([]rune, 0, 10)

	start := time.Now()
	for _, c := range contents {
		token = append(token, c)

		if grammar[string(token)] {
			tokens = append(tokens, token)
			token = make([]rune, 0, 10)
			continue
		}

		if separators[c] {
			token = make([]rune, 0, 10)
			continue
		}
	}
	fmt.Println(time.Now().Sub(start))

	for _, token := range tokens {
		fmt.Println(string(token))
	}
}

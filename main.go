package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

var keywords map[string]bool = map[string]bool{
	"break":      true,
	"do":         true,
	"instanceof": true,
	"typeof":     true,
	"case":       true,
	"else":       true,
	"new":        true,
	"var":        true,
	"catch":      true,
	"finally":    true,
	"return":     true,
	"void":       true,
	"continue":   true,
	"for":        true,
	"switch":     true,
	"while":      true,
	"debugger":   true,
	"function":   true,
	"this":       true,
	"with":       true,
	"default":    true,
	"if":         true,
	"throw":      true,
	"delete":     true,
	"in":         true,
	"try":        true,
}

var punctuators map[string]func(rune)(string, bool) = map[string]func(rune)(string, bool){
	"{":    SingleRunePunctuator,
	"}":    SingleRunePunctuator,
	"(":    SingleRunePunctuator,
	")":    SingleRunePunctuator,
	"[":    SingleRunePunctuator,
	"]":    SingleRunePunctuator,
	".":    SingleRunePunctuator,
	";":    SingleRunePunctuator,
	",":    SingleRunePunctuator,
	"<":    SingleRunePunctuator,
	">":    SingleRunePunctuator,
	"<=":   true,
	">=":   true,
	"==":   true,
	"!=":   true,
	"===":  true,
	"!==":  true,
	"+":    true,
	"-":    true,
	"*":    true,
	"%":    true,
	"++":   true,
	"--":   true,
	"<<":   true,
	">>":   true,
	">>>":  true,
	"&":    true,
	"|":    true,
	"^":    true,
	"!":    true,
	"~":    true,
	"&&":   true,
	"||":   true,
	"?":    true,
	":":    true,
	"=":    true,
	"+=":   true,
	"-=":   true,
	"*=":   true,
	"%=":   true,
	"<<=":  true,
	">>=":  true,
	">>>=": true,
	"&=":   true,
	"|=":   true,
	"^=":   true,
}

var whitespace map[rune]bool = map[rune]bool{
	'\u0009': true,
	'\u000B': true,
	'\u000C': true,
	'\u0020': true,
	'\u00A0': true,
	'\u1680': true,
	'\u2000': true,
	'\u2001': true,
	'\u2002': true,
	'\u2003': true,
	'\u2004': true,
	'\u2005': true,
	'\u2006': true,
	'\u2007': true,
	'\u2008': true,
	'\u2009': true,
	'\u200A': true,
	'\u202F': true,
	'\u205F': true,
	'\u3000': true,
	'\uFEFF': true,
}

var lineterminators map[rune]bool = map[rune]bool{
	'\u000A': true,
	'\u000D': true,
	'\u2028': true,
	'\u2029': true,
}

var stringliterals map[rune]bool = map[rune]bool{
	'\'': true,
	'"': true,
}

func SingleRunePunctuator(r rune) (string, bool) {
	return string(r), true
}

func MultiRunePunctuator(complete string) func(r rune)(string, bool) {
	handler := func(r rune)(string, bool) {
		return "", false
	}

	return handler
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
	
	var token []rune

	var cminus1 rune
	var inString bool

	start := time.Now()
	for _, c := range contents {
		if stringliterals[c] && !inString {
			cminus1 = 0
			inString = true
			continue
		}

		if inString {
			if stringliterals[c] {
				inString = false
				tokens = append(tokens, token)
				token = make([]rune, 0, 10)
			} else {
				token = append(token, c)
			}
			continue
		}

		if cminus1 == 0 {
			if !whitespace[c] && !lineterminators[c] && punctuators[string(c)] {
				cminus1 = c
			}
			continue
		}

		if whitespace[c] || lineterminators[c] || punctuators[string(c)] {
			if len(token) > 0 {
				tokens = append(tokens, token)
			}

			token = make([]rune, 0, 10)
			cminus1 = 0
			continue
		}

		token = append(token, c)
		cminus1 = c
	}
	fmt.Println(time.Now().Sub(start))

	for _, token := range tokens {
		fmt.Println(string(token))
	}
}

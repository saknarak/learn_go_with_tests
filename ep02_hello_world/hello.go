package main

import "fmt"

const langEN = "en"
const langTH = "th"
const langJP = "jp"
const prefixEN = "Hello, "
const prefixTH = "Sawasdee, "
const prefixJP = "Ohiyo, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return Greeting(lang) + name
}

func Greeting(lang string) string {
	prefix := prefixEN
	switch lang {
	case langTH:
		prefix = prefixTH
	case langJP:
		prefix = prefixJP
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Somchai", "English"))
}

package main

import (
	"io"
	"fmt"
	"net/http"
	"os"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "hello there"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

type logWriter struct{}

func main() {
	// eb := englishBot{}
	// sb := spanishBot{}

	// printGreeting(eb)
	// printGreeting(sb)
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// bs := make([]byte, 9999)

	// resp.Body.Read(bs)  

	// fmt.Println(string(bs))

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int , error ) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this byte slice:", len(bs))

	return len(bs), nil
}
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//An example showcasing working with stdLib interfaces

//Create a custom Writer
type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//Initialize custom writer
	logWriter := logWriter{}

	//Copy func requires a type of Writer interface & Reader interface
	io.Copy(logWriter, resp.Body)

	//stdLib way of going about printing the body object below.

	//io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) { //Make our custom writer to implement the Writer interface
	fmt.Println(string(bs))
	fmt.Println("Just read this much bytes:", len(bs))
	return len(bs), nil
}

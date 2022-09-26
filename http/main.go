package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}
type lwTest map[string]string
type testing []string

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Creating an empty byte slice that will be sent to the read function with resp
	// It will then read resp into the byte slice
	//bs := make([]byte, 99999)

	// resp.Body implements the reader interface based on documentation
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	//Alternative method
	// Copy takes in a function that implements a writer interface and a func that implements a reader interface
	//io.Copy(os.Stdout, resp.Body)

	//Create a Custom writer
	// The three composite types are used essentially so that they can be associated to an interface.
	// You can associated in-built types int,float,string, so you have to define composite types to create interfaces.
	lw := lwTest{}
	t := testing{}
	tw := logWriter{}
	io.Copy(lw, resp.Body)
	io.Copy(t, resp.Body)
	io.Copy(tw, resp.Body)
}

// Create a Custom writer
// Writer Definition in docs
//
//	type Writer interface {
//		Write(p []byte) (n int, err error)
//	}
//
// If you implement Write function defined in the Writer interface,
// logWriter is now also associated to Writer type, hence has a writer
// logwriter is now a value that implements a writer interface because it has a function called Write
func (lwTest) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func (testing) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

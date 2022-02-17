package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	/*
		bs := make([]byte, 99999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))

	*/

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
	//io.Copy(os.Stdout, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote this many bytes %v", len(bs))
	return len(bs), nil
}

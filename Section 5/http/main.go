package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")

	if err != nil {
		errMessage := fmt.Errorf("Error: %v", err)
		log.Println(errMessage)
		os.Exit(1)

	}
	// bytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	errMessage := fmt.Errorf("Error: %v", err)
	// 	log.Println(errMessage)
	// 	os.Exit(1)
	// }
	// resp.Body.Close()
	// fmt.Println(string(bytes))
	logWriter := logWriter{}
	io.Copy(logWriter, resp.Body)
}

func (logWriter) Write(bytes []byte) (int, error) {
	fmt.Println(string(bytes))
	fmt.Println("Just wrote this many bytes:", len(bytes))
	return len(bytes), nil
}

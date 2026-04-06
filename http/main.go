package main

import (
	"fmt"
	"net/http"
	"os"
)

type logWriter struct {
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("wrote this many bytes: ", len(bs))
	return len(bs), nil
}
//since
func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// fmt.Println(resp)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// instead of writing 3 lines of code above, we can write

	// io.Copy(os.Stdout, resp.Body)
	// this does the same
	// pipes info from source to destinatn.
	// Copy 1st argument > smthing that uses Write, 2nd argument > smthing that uses
	// resp.Body is the source from where data is read
	// os.Stdout is the destination. It represents the standard output stream (usually the console/terminal).
	// os.Stdout implements the io.Writer interface.

}

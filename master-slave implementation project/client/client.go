package main

import (
	"fmt"
	// "io/ioutil"
	"net/http"
	// "os"
)

func main() {
	
	// call the master (to get the file)
	resp, err := http.Get("http://localhost:9090")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	
	fmt.Println("the file downloaded successfully !")
}

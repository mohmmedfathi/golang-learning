package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Read the file to be split
	data, err := ioutil.ReadFile("/home/mohmmedfathi/Desktop/golang/final/file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println()
	
	fmt.Println("slave 3 waiting to split his part ...")
	
	// Split the file into three parts
	partSize := len(data) / 3
	part3 := data[2*partSize:]
	
	// fmt.Println("----------------------------------------------")
	

	// fmt.Println(string(part3))

	http.HandleFunc("/part3", func(w http.ResponseWriter, r *http.Request) {
		// Handle the request
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		w.Write(part3)
		fmt.Println("part 3 ok.")
	})

	log.Fatal(http.ListenAndServe(":8083", nil))

}

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
	fmt.Println("slave 1 waiting to split his part ...")
	
	// Split the file into three parts
	partSize := len(data) / 3
	part1 := data[:partSize]
	
	// fmt.Println("----------------------------------------------")
	

	// fmt.Println(string(part1))

	// Start the file sharing server
	http.HandleFunc("/part1", func(w http.ResponseWriter, r *http.Request) {

		// Handle the request
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		w.Write(part1)
		fmt.Println("part 1 ok.")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

	
}

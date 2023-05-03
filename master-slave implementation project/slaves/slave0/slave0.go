package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
)

func main() {
	// Set the IP address of the device
	ipAddress := "localhost"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

			fmt.Println("slave 0 working on merge parts .")

			client := &http.Client{}

			result := ""
			// Create a GET request to download the first part of the file
			req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8081/part1", ipAddress), nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			// Send the request and get the response
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(body))
			
			result += string(body)
			
			// Read the first part of the file
			
			// part1, err := ioutil.ReadAll(resp.Body)
			// if err != nil {
			// 	fmt.Println("Error reading response:", err)
			// 	return
			// }

			// // Write the part1 file to disk
			// err = ioutil.WriteFile("/home/mohmmedfathi/Desktop/golang/final/client/part1.txt", part1, os.ModePerm)
			// if err != nil {
			// 	fmt.Println("Error writing file:", err)
			// 	return
			// }

			fmt.Println("File downloaded and part1 successfully.")

			// Create a GET request to download the second part of the file
			req, err = http.NewRequest("GET", fmt.Sprintf("http://%s:8082/part2", ipAddress), nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			// Send the request and get the response
			resp, err = client.Do(req)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}

			defer resp.Body.Close()

			body2, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(body2))
			
			result += string(body2)

			// Read the second part of the file
			// part2, err := ioutil.ReadAll(resp.Body)
			// if err != nil {
			// 	fmt.Println("Error reading response:", err)
			// 	return
			// }

			// // Write the part2 file to disk
			// err = ioutil.WriteFile("/home/mohmmedfathi/Desktop/golang/final/client/part2.txt", part2, os.ModePerm)
			// if err != nil {
			// 	fmt.Println("Error writing file:", err)
			// 	return
			// }

			fmt.Println("File downloaded and part2 successfully.")

			// Create a GET request to download the third part of the file
			req, err = http.NewRequest("GET", fmt.Sprintf("http://%s:8083/part3", ipAddress), nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			// Send the request and get the response
			resp, err = client.Do(req)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}

			defer resp.Body.Close()
			
			body3, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(body3))
			
			result += string(body3)
			// Read the third part of the file
			// part3, err := ioutil.ReadAll(resp.Body)
			// if err != nil {
			// 	fmt.Println("Error reading response:", err)
			// 	return
			// }

			// Write the part3 file to disk
			// err = ioutil.WriteFile("/home/mohmmedfathi/Desktop/golang/final/client/part3.txt", part3, os.ModePerm)
			// if err != nil {
			// 	fmt.Println("Error writing file:", err)
			// 	return
			// }

			fmt.Println("File downloaded and part3 successfully.")
			
			fmt.Println(result)

			err_combined := ioutil.WriteFile("/home/mohmmedfathi/Desktop/golang/final/client/output_merged.txt", []byte(result), 0644)
			if err_combined != nil {
				log.Fatal(err)
			}
			
			
		}
	})

	fmt.Println()
	fmt.Println("slave0 waiting the master command (get request) to run ....")
	fmt.Println()

	log.Fatal(http.ListenAndServe(":8080", nil))

	
}

# Master-Slave implementation in Golang

This is a simple implementation of a master-slave architecture in Golang. <br>
The architecture consists of a client, a master, and 4 slaves.<br>
The client sends a request to the master, which then handles the request and forwards it to the slave0 (who will collect and merge the file parts from other slaves).<br>
The slaves perform the requested task(each one will splitting his part) and send the result back to slave0, which then sends it back to the client. 

![go](https://user-images.githubusercontent.com/64088888/236265897-edef4399-1ed9-4bdb-bed0-974bd48fabed.png)

## Built With

* ![Go](https://img.shields.io/badge/Go-Golang%20v1.20-blue?style=for-the-badge&logo=appveyor)

https://user-images.githubusercontent.com/64088888/235845709-5b8f45d1-87ea-43e6-b29a-2494c802e6dd.mp4



## Getting started

* 1 - Install Golang on your machine.
* 2 - Clone the repository.
* 3 - Navigate to the root directory of the repository.
* 4 - Run go get to install the necessary dependencies.

### Project Layout 

```tree
├── server
│   ├── server.go
│── client
|   ├── clinet.go
|── slaves
|    └── slave0
|        └── slave0.go
|    └── slave1
|        └── slave1.go
|    └── slave2
|        └── slave2.go
|    └── slave3
|        └── slave3.go
```
A brief description of the layout:
* `server.go` which is the master, will handle client request
* `client.go` request(GET) the file from master
* `slave0.go` will merge the file parts(1,2,3) from other slaves
* `slave1.go` split the first part from file and write it in the  response
* `slave2.go` split the second part from file and write it in the  response
* `slave3.go` split the third part from file and write it in the  response

## Usage
1 - Start the slaves by running the `slave0.go`, `slave1.go`, `slave2.go`, and `slave3.go` files. These will start HTTP servers on ports 8080, 8081, 8082, and 8083, respectively. 

2 - Start the master by running the `server.go` file. This will start an HTTP server on port 9090.

3 - Start the client by running the `client.go` file. This will make a request to the master.

4 - The slaves will perform the requested task and send the result back to the `slave0`, which will send it back to the client.

### notes :

* The client sends a request to the master using the HTTP protocol.

* The master uses the Go standard library's net/http/httputil package to create a reverse proxy that forwards the request to the appropriate slave.

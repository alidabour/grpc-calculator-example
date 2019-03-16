# Coding Test
### Creation of a RPC server and client
Server application, which can calculate simple math tasks. 
And a client which gets input via STDIN and send it to the server. After the server response, show the result on STDOUT.
### System requirements
- Make sure you install go 1.11 or higher
```$ go version```

### Installation instructions
1. Clone this repo in directory outside of **$GOPATH** or if you work in **GOPATH** then run ```export GO111MODULE=on```
2. Navigate to server directory and run main.go with any port (default 8000)
3. ```$ go run main.go --port=5000```
4. Open another shell and naviagate to client directory and run main.go with the same port
5. ```$ go run main.go --port=5000```
6. Start using calculator! (from the client shell) 
```
1+1
2
3*3*3
27
```

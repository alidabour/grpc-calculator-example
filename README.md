# Coding Test
### Creation of a RPC server and client
Server application, which can calculate simple math tasks. 
And a client which gets input via STDIN and send it to the server. After the server response, show the result on STDOUT.
### System requirements
- Make sure you install go 1.11 or higher
```$ go version```

### Installation instructions
1. Clone this repo in directory outside of **$GOPATH** or if you work in **GOPATH** then run ```export GOMODULLE111=on```
2. Navigate to server directory and run main.go with any port (default 8000)
3. (Optional) ```$ go test ./...```
4. (Optional) ```$ go test all``` Run the tests for your module plus the tests for all direct and indirect dependencies to check for incompatibilities
5. ```$ go run main.go --port=5000```
6. Open another shell and naviagate to client directory and run main.go with the same port
7. ```$ go run main.go --port=5000```
8. Start using calculator! (from the client shell) 
```
1+1
2
3*3*3
27
```

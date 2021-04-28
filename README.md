# Polyline CLI
This is a simple `Golang` CLI that receives a file.txt and based the locations coords it found in this file generates and print an encoded polyline.

## Requirements
As we know, this is a cli developed in go, so the first thing that you need to do is install go, please go to [this page](https://golang.org/doc/install) to obtain more info, download Go 1.16.x or higher preferably. 
Besides this we are using an external go library to encoded the polyline, this library is `github.com/twpayne/go-polyline` to implementing it only need to run the next command `go get github.com/twpayne/go-polyline`

### Running Polyline CLI
You just run the next command `go run main.go` in your terminal to running the polyline cli
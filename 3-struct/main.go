package main

import (
	"3-struct/api"
	"3-struct/bins"
	"3-struct/file"
	"3-struct/storage"
)

func main() {
	api.Hello()
	bins.Bins()
	file.File()
	storage.Storage()
}

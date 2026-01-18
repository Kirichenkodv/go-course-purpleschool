package main

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

var BinList = []Bin{}

func main() {

}

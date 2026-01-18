package main

import (
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

var BinList = []Bin{}
var lastNumber uint64

func NextNumericId() string {
	n := atomic.AddUint64(&lastNumber, 1)
	return fmt.Sprintf("%d", n)
}

type CreateBinInput struct {
	Name    string
	Private bool
}

func CreateBin(input CreateBinInput) (Bin, error) {
	name := strings.TrimSpace(input.Name)
	if name == "" {
		return Bin{}, errors.New("Поле Name не может быть пустым")
	}
	id := NextNumericId()
	b := Bin{
		Id:        id,
		Private:   input.Private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	BinList = append(BinList, b)
	return b, nil
}

func main() {

}

package bins

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

func nextNumericId() string {
	n := atomic.AddUint64(&lastNumber, 1)
	return fmt.Sprintf("%d", n)
}

type CreateBinInput struct {
	Name    string
	Private bool
}

// Service — реализация работы со списком bin с собственным состоянием (список + счётчик id).
// Используется для DI: main передаёт BinsStore, при загрузке из storage вызывает SetList.
type Service struct {
	list      []Bin
	lastNum   uint64
}

// CreateBin создаёт новый bin и добавляет его в список этого сервиса.
func (s *Service) CreateBin(input CreateBinInput) (Bin, error) {
	name := strings.TrimSpace(input.Name)
	if name == "" {
		return Bin{}, errors.New("Поле Name не может быть пустым")
	}
	id := fmt.Sprintf("%d", atomic.AddUint64(&s.lastNum, 1))
	b := Bin{
		Id:        id,
		Private:   input.Private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	s.list = append(s.list, b)
	return b, nil
}

// List возвращает текущий список bin.
func (s *Service) List() []Bin {
	return s.list
}

// SetList задаёт список bin (например после загрузки из storage).
// Счётчик id обновляется по длине списка, чтобы новые bin получали уникальные id.
func (s *Service) SetList(list []Bin) {
	s.list = list
	s.lastNum = uint64(len(list))
}

// Пакетные функции для обратной совместимости (работают с глобальным BinList).
func CreateBin(input CreateBinInput) (Bin, error) {
	name := strings.TrimSpace(input.Name)
	if name == "" {
		return Bin{}, errors.New("Поле Name не может быть пустым")
	}
	id := nextNumericId()
	b := Bin{
		Id:        id,
		Private:   input.Private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	BinList = append(BinList, b)
	return b, nil
}

func Bins() {}

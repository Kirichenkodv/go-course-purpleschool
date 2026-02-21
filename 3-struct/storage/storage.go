package storage

import (
	"3-struct/bins"
	"encoding/json"
	"os"
)

// Service — конкретная реализация хранилища (JSON-файл).
// Имеет методы Save и Load, поэтому автоматически удовлетворяет интерфейсу BinStorage из main.
type Service struct{}

// Save записывает список bin в файл path в формате JSON.
func (s *Service) Save(path string, list []bins.Bin) error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// Load читает список bin из JSON-файла path.
func (s *Service) Load(path string) ([]bins.Bin, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var list []bins.Bin
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	if list == nil {
		list = []bins.Bin{}
	}
	return list, nil
}

// Сохранённые функции-обёртки для обратной совместимости (можно вызывать storage.Save/Load без создания Service).
func Save(path string, list []bins.Bin) error {
	return (&Service{}).Save(path, list)
}

func Load(path string) ([]bins.Bin, error) {
	return (&Service{}).Load(path)
}

func Storage() {}

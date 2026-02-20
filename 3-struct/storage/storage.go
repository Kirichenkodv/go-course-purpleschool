package storage

import (
	"3-struct/bins"
	"encoding/json"
	"os"
)

// Save записывает список bin в файл path в формате JSON.
// Если файл не существует — он создаётся, если существует — перезаписывается.
func Save(path string, list []bins.Bin) error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// Load читает список bin из JSON-файла path.
// Если файл не существует или пустой — возвращает пустой срез и ошибку (или nil, если файл пустой и валидный "[]").
// При ошибке чтения или неверном JSON возвращает nil и ошибку.
func Load(path string) ([]bins.Bin, error) {
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

func Storage() {}

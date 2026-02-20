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

func Storage() {}

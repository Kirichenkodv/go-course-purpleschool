package file

import (
	"os"
	"strings"
)

// IsJSON проверяет, что имя файла имеет расширение .json (без учёта регистра).
// Примеры: "data.json" -> true, "backup.JSON" -> true, "readme.txt" -> false.
func IsJSON(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".json")
}

// Read читает содержимое файла по пути path и возвращает байты и ошибку.
// Путь может быть относительным или абсолютным.
func Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func File() {}

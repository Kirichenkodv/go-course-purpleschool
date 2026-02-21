package file

import (
	"os"
	"strings"
)

// Service — реализация чтения файлов (реальная файловая система).
// Методы Read и IsJSON позволяют использовать *Service как FileReader из main.
type Service struct{}

// Read читает содержимое файла по пути path и возвращает байты и ошибку.
func (s *Service) Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// IsJSON проверяет, что имя файла имеет расширение .json (без учёта регистра).
func (s *Service) IsJSON(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".json")
}

// Функции-обёртки для обратной совместимости.
func Read(path string) ([]byte, error) {
	return (&Service{}).Read(path)
}

func IsJSON(filename string) bool {
	return (&Service{}).IsJSON(filename)
}

func File() {}

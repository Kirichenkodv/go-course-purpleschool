package main

import (
	"3-struct/api"
	"3-struct/bins"
	"3-struct/file"
	"3-struct/storage"
	"log"
	"os"
)

// BinStorage — интерфейс хранилища списка bin (контракт для DI).
// main зависит только от этого контракта: "кто-то, кто умеет Save и Load".
// Конкретная реализация (файл, БД, мок в тестах) подставляется снаружи.
type BinStorage interface {
	Save(path string, list []bins.Bin) error
	Load(path string) ([]bins.Bin, error)
}

// FileReader — интерфейс для чтения файлов и проверки расширения (контракт для DI).
// main зависит от контракта: "кто-то, кто умеет Read и IsJSON".
type FileReader interface {
	Read(path string) ([]byte, error)
	IsJSON(filename string) bool
}

// BinsStore — интерфейс для работы со списком bin (контракт для DI).
// main зависит от контракта: создание bin, получение списка и установка списка (после загрузки из storage).
type BinsStore interface {
	CreateBin(input bins.CreateBinInput) (bins.Bin, error)
	List() []bins.Bin
	SetList(list []bins.Bin)
}

// Проверка на этапе компиляции: *storage.Service реализует BinStorage.
var _ BinStorage = (*storage.Service)(nil)

// Проверка на этапе компиляции: *file.Service реализует FileReader.
var _ FileReader = (*file.Service)(nil)

// Проверка на этапе компиляции: *bins.Service реализует BinsStore.
var _ BinsStore = (*bins.Service)(nil)

// Путь к файлу с сохранённым списком bin.
const storagePath = "bins.json"

// run собирает приложение воедино: принимает зависимости через интерфейсы (DI)
// и выполняет логику — загрузка списка, приветствие, сохранение списка.
func run(binsStore BinsStore, fileReader FileReader, binStorage BinStorage) {
	// Загрузка списка bin из файла при старте (если файл есть).
	if fileReader.IsJSON(storagePath) {
		list, err := binStorage.Load(storagePath)
		if err != nil {
			if !os.IsNotExist(err) {
				log.Printf("не удалось загрузить список: %v", err)
			}
			// иначе файла нет — начинаем с пустого списка
		} else {
			binsStore.SetList(list)
		}
	}

	api.Hello()

	// Сохранение списка в файл перед выходом.
	if err := binStorage.Save(storagePath, binsStore.List()); err != nil {
		log.Printf("не удалось сохранить список: %v", err)
	}
}

func main() {
	// Создаём конкретные реализации (один раз при старте).
	binsStore := &bins.Service{}
	fileReader := &file.Service{}
	binStorage := &storage.Service{}

	// Передаём их в run как интерфейсы — main зависит только от контрактов.
	run(binsStore, fileReader, binStorage)
}

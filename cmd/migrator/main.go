package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storageURL, migrationsTable, migrationsPath string

	// Получаем необходимые значения из флагов запуска
	// URL к БД
	flag.StringVar(&storageURL, "storage-path", "", "url storage")
	// Путь до папки с миграциями
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse() // Выполняем парсинг флагов

	// Валидация параметров
	if storageURL == "" {
		// Простейший способ обработки ошибки :)
		// При необходимости, можете выбрать более подходящий вариант.
		// Меня паника пока устраивает, поскольку это вспомогательная утилита.
		panic("storage-path is required")
	}
	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		"postgres://pmp:pmp1226@localhost:5432/firstcode_auth?sslmode=disable")

	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("o migrations to apply")

			return
		}

		panic(err)
	}
}

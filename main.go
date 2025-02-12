package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const userName = "illia"
const userPass = "2204illia"
const dbName = "english_db"
const englishTableName = "EnglishWords"

func main() {
	connString := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", userName, userPass, dbName)

	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatal("Не удалось подключиться к БД:", err)
	}

	defer conn.Close(context.Background())

	var version string
	err = conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal("Ошибка выполнения запроса:", err)
	}

	fmt.Println("PostgreSQL версия:", version, "\n")

	CreateTable(conn)

	InsertData(conn, "intention", "намерение")

	GetData(conn)
	//DeleteTable(conn)
}

func InsertData(conn *pgx.Conn, word string, translation string) {
	// Проверка, существует ли пользователь с таким именем
	var exists bool
	//"SELECT EXISTS(SELECT 1 FROM users WHERE name=$1 AND age=$2)", name, age
	existQuery := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM "%s" WHERE word=$1 AND translation=$2)`, englishTableName)
	err := conn.QueryRow(context.Background(), existQuery, word, translation).Scan(&exists)
	if err != nil {
		log.Fatal("Ошибка выполнения запроса:", err)
	}

	if exists {
		fmt.Printf("Cлово \"%s\" уже существует.\n", word)
	} else {
		// Вставка нового пользователя
		insertQuery := fmt.Sprintf(`INSERT INTO "%s" (word, translation) VALUES ($1,$2)`, englishTableName)
		_, err = conn.Exec(context.Background(), insertQuery, word, translation)
		if err != nil {
			log.Fatal("Ошибка вставки данных:", err)
		}
		fmt.Printf("Слово %s успешно добавлено!\n", word)
	}
}

func GetData(conn *pgx.Conn) {
	// Выполняем запрос
	getQuery := fmt.Sprintf(`SELECT id, word, translation FROM "%s"`, englishTableName)
	rows, err := conn.Query(context.Background(), getQuery)
	if err != nil {
		log.Fatal("Ошибка запроса:", err)
	}
	defer rows.Close()

	// Обрабатываем строки результата
	for rows.Next() {
		var id int
		var word string
		var translation string

		err := rows.Scan(&id, &word, &translation)
		if err != nil {
			log.Fatal("Ошибка чтения строки:", err)
		}

		fmt.Printf("ID: %d | Word: %s | Translation: %s\n", id, word, translation)
	}

	// Проверяем на ошибки после завершения итерации
	if err := rows.Err(); err != nil {
		log.Fatal("Ошибка при чтении данных:", err)
	}
}

func CreateTable(conn *pgx.Conn) {
	// Создание таблицы
	createQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s" (
		id SERIAL PRIMARY KEY,
		word TEXT NOT NULL,
		translation TEXT,
	    example TEXT
	)`, englishTableName)
	_, err := conn.Exec(context.Background(), createQuery)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}

	fmt.Println("Таблица EnglishWords создана!")
}

func DeleteTable(conn *pgx.Conn) {
	// Создание таблицы
	deleteQuery := fmt.Sprintf(`TRUNCATE TABLE "%s" RESTART IDENTITY`, englishTableName)
	_, err := conn.Exec(context.Background(), deleteQuery)
	if err != nil {
		log.Fatal("Ошибка удаления таблицы:", err)
	}

	fmt.Println("Таблица EnglishWords успешно удалена!")
}

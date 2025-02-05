package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://user:password@localhost:5432/mydb")

	if err != nil {
		log.Fatal("Не удалось подключиться к БД:", err)
	}
	defer conn.Close(context.Background())

	var version string
	err = conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal("Ошибка выполнения запроса:", err)
	}

	fmt.Println("PostgreSQL версия:", version)

	CreateTable(conn)
	
	InsertData(conn, "Illia", 25)
	for i := range 10 {
		InsertData(conn, "kwk", i*3)
	}
	GetData(conn)
}

func InsertData(conn *pgx.Conn, name string, age int) {
	// Проверка, существует ли пользователь с таким именем
	var exists bool
	err := conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE name=$1 AND age=$2)", name, age).Scan(&exists)
	if err != nil {
		log.Fatal("Ошибка выполнения запроса:", err)
	}

	if exists {
		fmt.Println("Пользователь с таким именем уже существует.")
	} else {
		// Вставка нового пользователя
		_, err = conn.Exec(context.Background(), "INSERT INTO users (name, age) VALUES ($1, $2)", name, age)
		if err != nil {
			log.Fatal("Ошибка вставки данных:", err)
		}
		fmt.Println("Пользователь успешно добавлен!")
	}
}

func GetData(conn *pgx.Conn) {
	 // Выполняем запрос
	rows, err := conn.Query(context.Background(), "SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal("Ошибка запроса:", err)
	}
	defer rows.Close()

	

	// Обрабатываем строки результата
	for rows.Next() {
		var id int
		var name string
		var age int

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("Ошибка чтения строки:", err)
		}

		fmt.Printf("ID: %d | Name: %s | Age: %d\n", id, name, age)
	}

	// Проверяем на ошибки после завершения итерации
	if err := rows.Err(); err != nil {
		log.Fatal("Ошибка при чтении данных:", err)
	}
}

func CreateTable(conn *pgx.Conn) {
	  // Создание таблицы
	_, err := conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT
	)`)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}

	fmt.Println("Таблица users создана и данные добавлены!")
}

Здесь пример запуска контейнера с бд postgres, подключение к бд, просмотр и добавление данных, проверкаой существует ли данные уже в таблице

GO TO THE PROJECT FOLDER in CMD

EXECUTE COMMANDS: `go main init myproject`; `go mod tidy`.

RUN COMMAND:
`docker run --name my_postgres \
  -p 5432:5432 \
  -v "$(pwd)/postgres_db:/var/lib/postgresql/data" \
  -d postgres`


RUN `main.go` file.


`-v ~/Desktop/postgres_db:/var/lib/postgresql/data \` – volumes for desktop

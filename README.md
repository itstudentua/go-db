Здесь пример запуска контейнера с бд postgres, подключение к бд, просмотр и добавление данных, проверкаой существует ли данные уже в таблице

GO TO THE PROJECT FOLDER in CMD

EXECUTE COMMANDS: `go mod init myproject`; `go get "github.com/jackc/pgx/v5"`; `go mod tidy`.

RUN COMMAND:
`docker run --name my_postgres \
  -e POSTGRES_USER=illia \ 
  -e POSTGRES_PASSWORD=2204illia \ 
  -e POSTGRES_DB=english_db \
  -p 5432:5432 \
  -v "$(pwd)/postgres_db:/var/lib/postgresql/data" \
  -d postgres`


RUN `main.go` file.

Файлы БД будут храниться в папке `postgres_db`, что указана в volumes.

`-v ~/Desktop/postgres_db:/var/lib/postgresql/data \` – volumes for desktop

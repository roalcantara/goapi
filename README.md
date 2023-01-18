# Go API

Simple Golang API with Gorm and Postgres

[![MIT license](https://img.shields.io/badge/License-MIT-brightgreen.svg)](LICENSE) [![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.0-4baaaa.svg)][2] [![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)][5] [![Editor Config](https://img.shields.io/badge/Editor%20Config-1.0.1-crimson.svg)][4] [![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)][3]

## Install

  ```sh
  git clone https://github.com/roalcantara/goapi.git
  ```

### Dependencies

- [Git][6]
- [Docker][8] [Compose][9]

## Usage

### [Golang][7]

- Run

  ```sh
  go run main.go
  ```

### [Docker][8] [Compose][9]

- Build

  ```sh
  docker compose build --remove-orphans
  ```

- Run

  ```sh
  docker compose up --remove-orphans
  ```

- Build and Run

  ```sh
  docker compose up --build --remove-orphans
  ```

- Shutdown

  ```sh
  docker compose down --remove-orphans
  ```

## Recipe

1. `docker run -it --rm -v $PWD:/app --workdir="/app" golang:1.19 sh`
    
    - `go mod init github.com/roalcantara/api`
    - `go get github.com/githubnemo/CompileDaemon`
    - `go install github.com/githubnemo/CompileDaemon`
    - `go get github.com/joho/godotenv`
    - `go get -u github.com/gin-gonic/gin`
    - `go get -u gorm.io/gorm`
    - `gorm.io/driver/postgres`

2. `touch .env`

    ```env
    PORT=3000
    POSTGRES_HOST=db
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=postgres
    POSTGRES_DB=postgres
    POSTGRES_PORT=5432
    ```

3. `touch main.go`

    ```go
    // main.go
    package main

    import (
      "log"

      "github.com/gin-gonic/gin"
      "github.com/joho/godotenv"
    )

    func init() {
      err := godotenv.Load()
      if err != nil {
        log.Fatal("Error loading .env file")
      }
    }

    func main() {
      r := gin.Default()
      r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
          "message": "Hello World!",
        })
      })
      r.Run()
    }
    ```

4. `touch Dockerfile`

    ```dockerfile
    # Dockerfile
    FROM golang:1.19-alpine as build
    WORKDIR /app
    COPY . .
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api

    FROM gcr.io/distroless/static:latest as prod
    WORKDIR /app
    COPY --from=build /app/api /bin
    ENTRYPOINT [ "/bin" ]

    FROM golang:1.19-alpine as dev
    WORKDIR /app
    COPY . .
    RUN go install -mod=mod github.com/githubnemo/CompileDaemon
    ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main"

    FROM golang:1.19-alpine as test
    COPY --from=build /app/api /bin
    RUN go test    
    ```

5. `touch docker-compose.yml`

    ```yml
    version: '3.9'

    services:
      db:
        image: postgres:alpine
        restart: always
        env_file: .env
        ports:
          - 5432:5432
        volumes:
          - postgresqldb:/var/lib/postgresql/data
      api:
        depends_on:
          - db
        build:
          context: .
          target: dev
        restart: unless-stopped
        env_file: .env
        environment:
          - DB_HOST=db
          - DB_USER=$POSTGRES_USER
          - DB_PASSWORD=$POSTGRES_PASSWORD
          - DB_NAME=$POSTGRES_DB
          - DB_PORT=$POSTGRES_PORT
        expose:
          - 3001
        ports:
          - 3001:3001
        volumes:
          - .:/app
        stdin_open: true
        tty: true
    volumes:
      postgresqldb:
        driver: local
    ```

6. `docker compose up --build`
7. `curl http://localhost:3001`

## Acknowledgements

- [Standard Readme][5]
- [Conventional Commits][3]
- [Creating a JSON CRUD API in Go (Gin/GORM)][12]

## Contributing

- Bug reports and pull requests are welcome on [GitHub][0]
- Do follow [Editor Config][4] rules.
- Do follow [Git lint][10] rules.
- Everyone interacting in the project’s codebases, issue trackers, chat rooms and mailing lists is expected to follow the [Contributor Covenant][2] code of conduct.

## License

The project is available as open source under the terms of the [MIT][1] [License](LICENSE)

[0]: https://github.com/roalcantara/goreact
[1]: https://opensource.org/licenses/MIT "Open Source Initiative"
[2]: https://contributor-covenant.org "A Code of Conduct for Open Source Communities"
[3]: https://conventionalcommits.org "Conventional Commits"
[4]: https://editorconfig.org "EditorConfig"
[5]: https://github.com/RichardLitt/standard-readme "Standard Readme"
[6]: https://git-scm.com "Git"
[7]: https://go.dev "Go: An open-source programming language supported by Google"
[8]: https://docker.com "Docker: The most-loved Tool in Stack Overflow’s 2022 Developer Survey"
[9]: https://docs.docker.com/compose "Docker Compose: Defining and running multi-container Docker applications"
[10]: https://jorisroovers.com/gitlint "git commit message linter"
[11]: https://pre-commit.com "A framework for managing and maintaining multi-language pre-commit hooks"
[12]: https://youtu.be/lf_kiH_NPvM "Creating a JSON CRUD API in Go (Gin/GORM)"

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

FROM golang:alpine

RUN apk update && apk add git

ENV MYSQL_HOST=host.docker.internal
ENV MYSQL_PORT=3306
ENV MYSQL_USER="root"
ENV MYSQL_PASSWORD=""
ENV MYSQL_DBNAME="go_todo"

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o gotodo

ENTRYPOINT [ "./gotodo" ]
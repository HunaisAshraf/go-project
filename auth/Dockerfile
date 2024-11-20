FROM golang:alpine3.20

WORKDIR /app

COPY . .

WORKDIR /app/cmd/server

EXPOSE 3001

CMD [ "go","run","server.go" ]
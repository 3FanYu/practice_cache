FROM golang:1.22

WORKDIR /app

COPY . .

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1


RUN go mod download

FROM golang:1.20

ENV TZ=Europe/Moscow

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server
RUN go build -o mail ./cmd/mail

EXPOSE 8080

CMD ["sh", "-c", "./server & ./mail"]

FROM golang:1.23-alpine
LABEL authors="IX"

WORKDIR /usr/local/go/src/SongsLibrary
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum /usr/local/go/src/SongsLibrary/
RUN go mod download

COPY app /usr/local/go/src/SongsLibrary/app
COPY air.toml .
#COPY .env .env

EXPOSE 4000:8080

CMD ["air", "-c", "air.toml"]
FROM golang:1.23-alpine
LABEL authors="IX"

WORKDIR /SongsLibrary
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum /SongsLibrary/
RUN go mod download

COPY app/* /SongsLibrary
COPY air.toml .

EXPOSE 4000:8080

CMD ["air", "-c", "air.toml"]
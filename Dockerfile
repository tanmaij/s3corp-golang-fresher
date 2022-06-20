#Dockerfile

FROM golang:1.18.3

WORKDIR /app

COPY . .

RUN go mod download && go mod tidy

RUN go build -o app

EXPOSE 3333

ENTRYPOINT ["./app"]
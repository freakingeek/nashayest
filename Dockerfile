FROM golang:latest

ENV PORT 8080

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /nashayest

EXPOSE 8080

ENTRYPOINT [ "/nashayest" ]
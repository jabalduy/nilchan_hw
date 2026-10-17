FROM golang:1.27-bookworm

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["make", "service-run"]
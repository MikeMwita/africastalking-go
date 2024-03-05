FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM scratch

COPY --from=builder /app/main /app/main

EXPOSE 5556

CMD ["/app/main"]
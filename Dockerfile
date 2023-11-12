FROM golang:alpine as builder

WORKDIR /go/src

COPY main.go .
COPY go.mod .
COPY go.sum .

RUN go build -o /go/main .

RUN ls /go

FROM scratch

COPY --from=builder /go/main /app/main

ENTRYPOINT ["/app/main"]

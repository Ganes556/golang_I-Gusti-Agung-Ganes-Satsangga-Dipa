FROM golang:1.20-alpine as builder

WORKDIR /usr/src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./app

FROM alpine:3.14

RUN apk --no-cache add ca-certificates

WORKDIR /usr/src

COPY --from=builder /usr/src/app .

EXPOSE 8000

CMD [ "./app" ]
FROM golang:1.20-alpine as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o app .

# Copy the binary to the production image from the builder stage.

FROM golang:1.20-alpine

ARG env_file_path=./.env.production

ENV $(cat $env_file_path | xargs)

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/app .

EXPOSE 8000

CMD [ "./app" ]
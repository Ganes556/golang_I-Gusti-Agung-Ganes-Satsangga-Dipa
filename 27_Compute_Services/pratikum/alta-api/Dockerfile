FROM golang:1.20-alpine as builder

WORKDIR /usr/src/alta_api

COPY go.mod go.sum ./

RUN go mod download

COPY .  .

RUN go build -v -o alta_api_app .


FROM golang:1.20-alpine

ARG env_file_path=./.env.production

ENV $(cat $env_file_path | xargs)

WORKDIR /usr/src/alta_api

COPY --from=builder /usr/src/alta_api/alta_api_app .

EXPOSE 8000

CMD [ "./alta_api_app" ]


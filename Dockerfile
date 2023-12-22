FROM golang:1.21.5-alpine3.19

WORKDIR /app

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .

RUN go build -o /main

CMD [ "/main" ]
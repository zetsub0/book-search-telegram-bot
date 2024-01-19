FROM golang:1.21-alpine3.19 as builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash gcc gettext musl-dev

COPY . .
RUN go mod download
RUN go build -o /usr/local/bin/bot cmd/bot/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/bin/bot /usr/local/bin/

WORKDIR /usr/local/bin

COPY main.yaml .

CMD ["bot"]


FROM golang:1.21 AS gobuild

WORKDIR /build

COPY ./ ./
RUN go mod download && \
    CGO_ENABLED=0 go build cmd/bot/main.go

CMD [ "./main" ]


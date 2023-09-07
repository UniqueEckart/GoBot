FROM golang:1.15-alpine as build

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o bot bot.go

#-------------------------------#

FROM alpine:3 AS final
WORKDIR /app

COPY --from=build /build/bot ./bot
ENTRYPOINT ["./bot"]
CMD ["-c", "config.json"]
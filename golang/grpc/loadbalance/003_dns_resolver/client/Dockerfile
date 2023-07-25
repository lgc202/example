FROM golang:1.19-alpine AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o client .

FROM alpine

RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/client /app/client
CMD ["./client"]

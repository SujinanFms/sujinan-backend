FROM golang:1.24-alpine

# ติดตั้ง netcat (nc) สำหรับ wait-for.sh
RUN apk add --no-cache netcat-openbsd

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

COPY wait-for.sh /wait-for.sh
RUN chmod +x /wait-for.sh

RUN ls -l /
RUN ls -l /wait-for.sh
RUN ls -l /app


EXPOSE 8080

CMD ["./main"]


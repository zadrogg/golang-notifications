# компилируем приложение
FROM golang:alpine as builder
WORKDIR /notifications
RUN apk --no-cache add ca-certificates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./notifications

# запускаем собранное приложение
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /notifications/notifications /notifications
ENTRYPOINT ["/notifications"]
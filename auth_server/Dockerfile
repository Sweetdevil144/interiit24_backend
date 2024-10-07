FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o auth_server .
FROM alpine:3.17
COPY --from=build /app/auth_server /app/auth_server
EXPOSE 6969
CMD ["/app/auth_server"]

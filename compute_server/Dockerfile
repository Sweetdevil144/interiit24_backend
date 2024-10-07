FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o compute_server .
FROM alpine:3.17
COPY --from=build /app/compute_server /app/compute_server
EXPOSE 6942
CMD ["/app/compute_server"]
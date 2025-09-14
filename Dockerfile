# Build the binary
FROM golang:1.25.1-alpine3.22 AS build
RUN apk update
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/gotodo

# Get the binary from the build stage and
# Execute in alpine container
FROM alpine:3.17
RUN apk add --no-cache ca-certificates &&\
  rm -rf /var/lib/cache
WORKDIR /app
COPY --from=build /app/.env .
COPY --from=build /app/bin/gotodo .
EXPOSE 9000
ENTRYPOINT [ "./gotodo" ]

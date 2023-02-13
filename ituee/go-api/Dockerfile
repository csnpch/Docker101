FROM golang:1.19-alpine3.16 AS builder

# Set workdir create any path/name
WORKDIR /app
# copy current ./ or . to path /app
COPY ./go-api/. /app


# RUN go build -o main main.go
RUN go build main.go



# Build small image
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY ./go-api/.env /app

# FOR Volume uploads directory  -p is mean auto create directory if some path is lose eg: /usr/ituee -> /usr/src/ituee
RUN mkdir -p /usr/src/ituee/uploads 
# PORT
EXPOSE 3000

# Command run api but if not have line 22 don't thinks to bad
CMD ["/app/main"]


# docker run --name goapi -d -p 3000:3000 goapi
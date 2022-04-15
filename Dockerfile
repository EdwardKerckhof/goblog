# BUILD
FROM golang:1.18-alpine3.15 AS build

WORKDIR /apps/blog/api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build

# RUN
FROM alpine:3.15
WORKDIR /apps/blog/api

COPY --from=build /apps/blog/api/api .

EXPOSE 3000
CMD ["./api"]

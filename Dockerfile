FROM golang:1.25.3-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/ginmongo

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/ginmongo .

CMD [ "./learnginmongo" ]
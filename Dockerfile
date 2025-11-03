FROM golang:1.25.3-alpine AS BUILD

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/ginmongo

FROM alpine:latest

WORKDIR /app

COPY --from=BUILD /app/ginmongo .

CMD [ "./learnginmongo" ]
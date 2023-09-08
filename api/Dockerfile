FROM golang:alpine as builder

RUN mkdir /build

WORKDIR /build
ADD . .

RUN go build -o main .

#stage 2

FROM alpine

RUN adduser -S -D -H -h /app appuser
USER appuser

COPY . /app
COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000
CMD ["./main"]
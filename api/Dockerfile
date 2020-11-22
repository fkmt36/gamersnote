FROM golang:1.15-alpine AS dev

WORKDIR /usr/src/api

COPY ./go.mod .

RUN go get github.com/codegangsta/gin

FROM golang:1.15-alpine AS build

WORKDIR /usr/src/api

COPY . .

RUN go build -o a.out

FROM alpine

WORKDIR /usr/src/api

COPY --from=build /usr/src/api/a.out .

CMD [ "./a.out" ]

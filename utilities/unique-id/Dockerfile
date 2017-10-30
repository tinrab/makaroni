FROM golang:1.9.2-alpine3.6 AS build
RUN apk --no-cache add gcc g++ make

WORKDIR /go/src/app
COPY . .
RUN go build -a -o /go/bin/app .

FROM alpine:3.6
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 3000
CMD [ "app" ]

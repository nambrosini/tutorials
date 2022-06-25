FROM golang:1.18.3 AS build
WORKDIR /go/src/github.com/nambrosini/tutorials
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/nambrosini/tutorials/app .
CMD ["./app"]
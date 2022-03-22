FROM golang:1.18-buster as build

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR go/src/lambda

RUN go env -w GOPROXY=direct

ADD go.mod go.sum ./
RUN go mod download

ADD .. .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main

FROM scratch
COPY --from=build /main /bin/main
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT [ "/bin/main" ]

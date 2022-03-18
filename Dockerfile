FROM alpine as build

RUN apk add go git
RUN go env -w GOPROXY=direct

ADD go.mod go.sum ./
RUN go mod download

ADD . .
RUN go build -o /main

FROM scratch
COPY --from=build /main /main
ENTRYPOINT [ "/main" ]

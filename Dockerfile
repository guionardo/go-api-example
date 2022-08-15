FROM golang:1.19 AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go env | grep 'CACHE'
RUN go build -o /go-api-example
RUN ls -la /
#14 0.524 GOCACHE="/root/.cache/go-build"
#14 0.524 GOMODCACHE="/go/pkg/mod"

FROM alpine:3.16 AS run
WORKDIR /
RUN adduser -D nonroot
USER nonroot
COPY --from=build /go-api-example /go-api-example
EXPOSE 8080
RUN ls -la /go-api-example
ENTRYPOINT ["/go-api-example"]

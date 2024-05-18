FROM golang:1.18 as build
WORKDIR /
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux GOPROXY="https://goproxy.cn" go build .

FROM alpine:3.14
RUN apk --no-cache add ca-certificates
COPY --from=build /cicd-apiserver /
ENTRYPOINT ["/cicd-apiserver"]
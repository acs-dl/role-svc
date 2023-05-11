FROM golang:1.19-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/acs-dl/role-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/role-svc /go/src/github.com/acs-dl/role-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/role-svc /usr/local/bin/role-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["role-svc"]

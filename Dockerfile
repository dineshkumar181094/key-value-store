FROM golang:1.15-alpine3.13 as builder
WORKDIR /srv/gocode
COPY . /srv/gocode
RUN ls
RUN go build -o bin/apiserver ./exec/api && go build -o bin/kvstore ./exec/cli


FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /srv/gocode/bin/* ./
CMD ["/root/apiserver"]

FROM golang:1.22.3-alpine3.20 as builder

WORKDIR /workdir
COPY . /workdir

ENV CGO_ENABLED=0
RUN go build -o http-dump .

FROM scratch

WORKDIR /
EXPOSE 8080

COPY --from=builder /workdir/http-dump .

USER 65532:65532

ENTRYPOINT ["/http-dump"]

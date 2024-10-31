FROM golang:1.23.2-alpine3.20 as builder

WORKDIR /src

RUN apk --update --no-cache add git make

ENV CGO_ENABLED=0

COPY go.mod go.mod
COPY Makefile Makefile

RUN go mod download

COPY *.go ./

RUN make build

FROM alpine:3.20

RUN apk --update --no-cache add ca-certificates

COPY --from=builder /src/sops-compliance-checker /sops-compliance-checker

USER nobody

ENTRYPOINT ["/sops-compliance-checker"]

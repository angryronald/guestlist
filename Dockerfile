FROM golang:1.16-alpine as builder

RUN apk add --no-cache ca-certificates git

WORKDIR /guestlist
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go install ./cmd/guestlist

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin /bin
USER nobody:nobody
ENTRYPOINT ["/bin/guestlist"]
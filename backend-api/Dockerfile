FROM golang:latest as builder

ENV GO111MODULE=on
WORKDIR /anotherdirectory

COPY go.mod /anotherdirectory
COPY go.sum /anotherdirectory
RUN go mod download

COPY . /anotherdirectory
RUN cd /anotherdirectory/cmd/thing && CGO_ENABLED=0 GOOS=linux go build -o /go/bin/thing .

FROM alpine:latest
COPY --from=builder /go/bin/thing /go/bin/thing
ENTRYPOINT ["/go/bin/thing"]

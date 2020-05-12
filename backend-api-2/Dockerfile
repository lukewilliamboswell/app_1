FROM golang:latest as builder

ENV GO111MODULE=on
WORKDIR /anotherdirectory

COPY go.mod /anotherdirectory
COPY go.sum /anotherdirectory
RUN go mod download

COPY . /anotherdirectory
RUN cd /anotherdirectory/cmd/calc && CGO_ENABLED=0 GOOS=linux go build -o /go/bin/calc .

FROM alpine:latest
COPY --from=builder /go/bin/calc /go/bin/calc
ENTRYPOINT ["/go/bin/calc"]

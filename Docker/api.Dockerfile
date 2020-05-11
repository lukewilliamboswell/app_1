# 
#  This file should build a goa API service and starts a dev server
#  I think it should be a multi stage Dockerfile following builder pattern
# 
FROM golang:alpine

WORKDIR /src/go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]

FROM golang

WORKDIR /go/src/client
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["client"]

FROM golang

WORKDIR /go/src/manganato-cli
COPY . .

RUN go get

CMD ["go","run","."]
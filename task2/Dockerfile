FROM golang:1.18-alpine
ADD . /go/modules/flaconi-challenge
WORKDIR /go/modules/flaconi-challenge
RUN go install -v
CMD ["/go/bin/flaconi-challenge"]
EXPOSE 8080
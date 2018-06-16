FROM golang:1.10.3-alpine3.7
EXPOSE 3000

COPY ./main.go /go/src/github.com/morix1500/study-k8s/main.go

WORKDIR /go/src/github.com/morix1500/study-k8s/

RUN go build -o app

CMD ./app

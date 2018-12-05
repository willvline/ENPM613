# from node:alpine

# WORKDIR /web-app/web

# RUN npm install -g live-server

# COPY ./web /web-app/web

# EXPOSE 8080

# CMD live-server

FROM golang:1.10.1

RUN go version

WORKDIR /go/src/github.com/Johnlovescoding/ENPM613/HOLMS/

COPY . /go/src/github.com/Johnlovescoding/ENPM613/HOLMS

RUN cd ./cmd && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

EXPOSE 8000

CMD ["./cmd/main"]



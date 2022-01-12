FROM golang:1.17.6-bullseye
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app
EXPOSE 8080
CMD ["./note-app"]

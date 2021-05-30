FROM golang:latest as build-env

RUN mkdir /events-watcher
WORKDIR /events-watcher
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/events-watcher

FROM scratch
COPY --from=build-env /go/bin/events-watcher /go/bin/events-watcher
ENTRYPOINT ["/go/bin/events-watcher"]



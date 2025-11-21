FROM golang:1.22.0
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /newtestapp

EXPOSE 8090

CMD ["/newtestapp"]

FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-chi

EXPOSE 3000

CMD [ "/go-chi" ]
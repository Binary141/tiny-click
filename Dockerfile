from golang:1.22

workdir /usr/src/app

copy go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]

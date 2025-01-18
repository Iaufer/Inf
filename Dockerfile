FROM golang:1.22.6

RUN go version

WORKDIR /app

COPY . .

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download   
RUN go build -o apiserver cmd/apiserver/main.go

CMD [ "./apiserver" ]
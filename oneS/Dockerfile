FROM golang:1.16.13-alpine3.14

RUN mkdir /servb

WORKDIR /servb

COPY . /servb

RUN go mod download

RUN go build -o binOneS

CMD [ "./binOneS" ]

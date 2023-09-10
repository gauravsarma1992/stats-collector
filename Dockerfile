FROM golang:1.19

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY config/config.json /config.json
COPY statscollector ./statscollector
COPY cmd ./cmd
COPY Makefile ./

RUN ls
RUN make build

EXPOSE 8050

CMD ["bin/statscollector"]
FROM golang:1.17 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

RUN go get github.com/markbates/pkger/cmd/pkger
COPY go.mod go.sum ./
RUN go mod download

ADD . /app/
RUN pkger && \
    go build -a -installsuffix cgo -o gemini-forum

FROM scratch
COPY --from=builder /etc/ssl /etc/ssl

COPY --from=builder /app/gemini-forum /gemini-forum
CMD ["/gemini-forum"]
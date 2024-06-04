FROM golang:1.22

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/rubenv/sql-migrate/...@latest

CMD ["air"]
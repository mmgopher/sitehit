FROM golang:1.12

WORKDIR /app
COPY . .

RUN go build -o /app/sitehit
CMD /app/sitehit
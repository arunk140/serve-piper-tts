FROM golang:1.20

WORKDIR /go/src/app
COPY . .

RUN chmod +x ./download-piper.sh
RUN ./download-piper.sh

RUN chmod +x ./download-voices.sh
RUN ./download-voices.sh

RUN go build .

EXPOSE 8080

CMD ["./serve-piper-go"]
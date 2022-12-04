FROM golang:1.19

WORKDIR /work/

COPY go.mod go.sum /work/
RUN go mod download && go mod verify 

COPY . . /work/

RUN go build -o imbored

ENTRYPOINT ["./imbored"]
CMD [""]
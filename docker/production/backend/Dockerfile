FROM golang:1.23

WORKDIR /app

ENV PORT 80
ENV USER hakushi

RUN useradd -m $USER && \
    chown -R $USER:$USER /app  && \
    chown -R $USER:$USER /go

COPY go.mod go.sum ./
RUN go mod download

RUN mkdir -p /data && chown -R $USER:$USER /data

COPY . .

RUN go build -o /app/main .

EXPOSE $PORT

USER $USER

CMD ["/app/main"]

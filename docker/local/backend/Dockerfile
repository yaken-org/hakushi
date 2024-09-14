FROM golang:1.23

WORKDIR /app

ENV AIR_CONFIG_PATH .air.toml
ENV PORT 8080
ENV USER hakushi

RUN useradd -m $USER && \
    chown -R $USER:$USER /app  && \
    chown -R $USER:$USER /go

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

RUN mkdir -p /data && chown -R $USER:$USER /data

EXPOSE $PORT

USER $USER
ENTRYPOINT ["air", "-c", ".air.toml"]

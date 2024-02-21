FROM golang:1.22.0-bookworm

WORKDIR /app

RUN GOBIN=/usr/local/bin/ go install github.com/gocraft/work/cmd/workwebui@v0.5.1

# ENTRYPOINT [ "workwebui" ]

EXPOSE 5040

CMD workwebui -redis=$REDIS_HOST -ns=$NAMESPACE -listen=$APP_LISTEN_HOST

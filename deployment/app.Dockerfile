FROM golang:alpine as builder

ARG SERVICE_NAME

WORKDIR /

COPY go.mod ./
COPY go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/$SERVICE_NAME ./cmd/$SERVICE_NAME/main.go

# deploy
FROM alpine as production

ARG SERVICE_NAME
ARG SERVICE_PORT

WORKDIR /

COPY --from=builder /bin/${SERVICE_NAME} /${SERVICE_NAME}

# COPY config.yaml /config.yaml

# COPY static /static

COPY .env .env
COPY ./deployment/scripts/entrypoint.sh entrypoint.sh

RUN chmod +x entrypoint.sh

EXPOSE ${SERVICE_PORT}

ENTRYPOINT ["/entrypoint.sh"]

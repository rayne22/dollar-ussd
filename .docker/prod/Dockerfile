FROM golang:1.20.0-alpine AS base
WORKDIR /app

ARG DOTENV_KEY
ARG DEBUG=false
ENV DOTENV_KEY=${DOTENV_KEY}
ENV DEBUG=${DEBUG}

# builder
FROM base AS builder
ENV GOOS linux
ENV GOARCH amd64

RUN apk --no-cache add bash git openssh

# modules: utlize build cache
COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

RUN go build -o main .

# runner
FROM base AS runner
RUN apk add --no-cache libc6-compat tini
# Tini is now available at /sbin/tini

COPY --from=builder /app/main /app/main
COPY --from=builder /app/.env.vault /app/.env.vault
EXPOSE 8880

ENTRYPOINT [ "/sbin/tini", "--" ]
CMD [ "/app/main" ]


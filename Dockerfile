FROM golang:1.20.0-alpine AS base
WORKDIR /app

ARG DEBUG=true
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

# RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# runner
FROM base AS runner

ARG DOTENV_KEY="dotenv://:key_fdf7141f6d0c6f8e21cabc9719d095076c28cd97ecad12231f2d50289eaab936@dotenv.org/vault/.env.vault?environment=staging"
ENV DOTENV_KEY=${DOTENV_KEY}

RUN apk add --no-cache libc6-compat tini
# Tini is now available at /sbin/tini

COPY --from=builder /app .
#COPY --from=builder /app/main /app/main
#COPY --from=builder /app/simulator /app/simulator
COPY --from=builder /app/.env.vault /app/.env.vault
EXPOSE 8980

CMD ["go", "run", "main.go"]
# CMD ["./main"]

#ENTRYPOINT [ "/sbin/tini", "--" ]
#CMD [ "/app/main" ]


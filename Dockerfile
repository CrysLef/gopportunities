FROM golang:1.21.1-alpine AS builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd cmd/api/ && CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/cmd/api/main .
ARG PORT
ENV PORT=$PORT
EXPOSE $PORT
CMD [ "./main" ]


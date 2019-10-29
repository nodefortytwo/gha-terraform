FROM golang:1.13 as builder
WORKDIR /code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint main.go

FROM hashicorp/terraform:0.12.12
RUN apk add tree
COPY --from=builder entrypoint /entrypoint

ENTRYPOINT /entrypoint
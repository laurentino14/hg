

FROM golang:1.21.5 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY vendor ./vendor
COPY . .
RUN go mod tidy
RUN go mod vendor
RUN CGO_ENABLED=0 go build ./cmd/sirius
RUN ls -la

FROM scratch
WORKDIR /app
COPY --from=builder /app/sirius ./
EXPOSE 3000
ENTRYPOINT ["./sirius"]

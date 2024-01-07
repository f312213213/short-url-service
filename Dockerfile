FROM golang:1.20-alpine3.18 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV GOPATH=/short-url
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /short-url


FROM scratch

COPY --from=builder /short-url/short-url /short-url

EXPOSE 8080


CMD ["./short-url"]
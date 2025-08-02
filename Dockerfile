FROM golang:alpine AS builder
WORKDIR /build

COPY . .
RUN go build -ldflags "-s -w" -o dockhome

FROM scratch
WORKDIR /dockhome

COPY --from=builder /build/dockhome .
ENTRYPOINT ["./dockhome"]

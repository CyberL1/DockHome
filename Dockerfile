FROM node:alpine AS web-builder

WORKDIR /build
COPY web .

RUN npm i -D
RUN npm run build

FROM golang:alpine AS builder
WORKDIR /build

COPY . .
COPY --from=web-builder /build/build web/build

RUN go build -ldflags "-s -w" -o dockhome

FROM scratch
WORKDIR /dockhome

COPY --from=builder /build/dockhome .
ENTRYPOINT ["./dockhome"]

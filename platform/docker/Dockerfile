FROM golang:1.24 AS build
ENV CGO_ENABLED=1

WORKDIR /app
COPY . .
RUN go build -o api-binary ./cmd/app

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /app/api-binary /api-binary
COPY --from=build /app/configs/ /configs/
COPY --from=build /app/keys/ /keys/

WORKDIR /
CMD ["./api-binary"]

LABEL org.opencontainers.image.title="api" \
      org.opencontainers.image.authors="Ibrahim Konsowa <ibrahim@konsowa.com>"

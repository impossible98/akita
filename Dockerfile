# one
FROM golang:1.17 as one
WORKDIR /src/
ENV CGO_ENABLED=1
COPY . .
RUN apt-get update; \
    apt-get install --yes --no-install-recommends make; \
    make install; \
    make build
# two
FROM debian:10
WORKDIR /app
COPY --from=one /src/dist/ ./
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/app/akita", "serve"]

FROM golang:1.24-alpine AS builds
RUN mkdir -p /isa/builds/{restaurant,binaries}
WORKDIR /isa/builds/restaurant
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /isa/builds/binaries/restaurant /isa/builds/restaurant

FROM gcr.io/distroless/base-debian12
COPY --from=builds /isa/builds/binaries/restaurant /bin/restaurant
ENTRYPOINT ["/bin/restaurant"]
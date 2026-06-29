FROM golang:1.24-alpine AS builds
RUN mkdir -p /isa/builds/{api,binaries}
WORKDIR /isa/builds/api
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /isa/builds/binaries/api /isa/builds/api

FROM gcr.io/distroless/base-debian12
COPY --from=builds /isa/builds/binaries/api /bin/api
ENTRYPOINT ["/bin/api"]
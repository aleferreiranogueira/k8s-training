## Build stage
FROM golang:1.20.3-bullseye as build-stage

WORKDIR /src

# Download modules
COPY go.mod go.sum ./
RUN go mod download

# Copy app source code
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /k8s-training-app

## Release image
FROM gcr.io/distroless/base-debian11 AS release-stage

WORKDIR /

COPY --from=build-stage /k8s-training-app /k8s-training-app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/k8s-training-app"]

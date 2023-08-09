# Start by building the application.
FROM golang:1.20 as build

WORKDIR /
COPY main.go main.go
COPY go.mod go.mod

RUN go mod download
RUN CGO_ENABLED=0 go build -o app

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11
COPY --from=build /app /
CMD ["/app"]
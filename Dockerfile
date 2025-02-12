FROM golang:1.24 AS build
RUN go install github.com/google/go-licenses@latest
WORKDIR /app
COPY . .
RUN make build

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/bin/file-server /app
EXPOSE 8080
CMD ["/app/file-server"]
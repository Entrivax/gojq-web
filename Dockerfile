FROM golang:1.24 AS build
RUN go install github.com/google/go-licenses@latest
WORKDIR /app
COPY . .
RUN make build-gojq

FROM nginx:stable
COPY --from=build /app/front/ /usr/share/nginx/html/
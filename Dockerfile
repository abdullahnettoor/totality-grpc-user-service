FROM golang:alpine3.20 AS build
WORKDIR /app
COPY ./ /app
RUN mkdir -p /app/build
RUN go mod tidy
RUN go build -v -o /app/build/user_service ./cmd/main.go

FROM gcr.io/distroless/static-debian11
COPY --from=build /app/build/user_service /
COPY --from=build /app/pkg/config/dev.env /pkg/config/
EXPOSE 9000
ENTRYPOINT ["/user_service"]
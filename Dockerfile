FROM golang:alpine3.20 AS build
WORKDIR /app
COPY ./ /app
RUN mkdir -p /app/build
RUN go mod download
RUN go build -v -o /app/build/user_service ./cmd/main.go

FROM build AS test
RUN go test ./...

FROM gcr.io/distroless/static-debian11 AS release
COPY --from=build /app/build/user_service /
COPY --from=build /app/dev.env /
EXPOSE 9000
ENTRYPOINT ["/user_service"]
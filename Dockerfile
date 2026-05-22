FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/certcheck ./...

FROM gcr.io/distroless/static:nonroot
COPY --from=build /out/certcheck /certcheck
ENTRYPOINT ["/certcheck"]

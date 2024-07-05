FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o exporter cmd/main.go

FROM scratch

COPY --from=build /app/exporter /exporter

COPY public /public

USER nobody

ENTRYPOINT ["/exporter"]

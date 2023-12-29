FROM golang as build

WORKDIR /app
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go/bin/app

FROM scratch
COPY --from=build /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]

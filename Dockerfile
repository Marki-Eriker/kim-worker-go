FROM golang:latest as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /worker cmd/app/*

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /worker /worker
COPY ./.env ./

EXPOSE 5432
EXPOSE 3030

CMD ["/worker"]

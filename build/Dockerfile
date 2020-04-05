FROM golang:1.14 as build

ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN make build_local

FROM gcr.io/distroless/base
COPY --from=build /app/compagnyhelper /
# COPY --from=build /app/api/openapi.yml /api/openapi.yml
WORKDIR /app/data
CMD ["/compagnyhelper"]

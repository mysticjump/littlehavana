FROM golang:1.16-alpine as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine
COPY --from=build /app/main /serve
COPY static static
CMD ["/serve"]
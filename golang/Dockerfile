ARG BUILD_IMAGE=golang
ARG BUILD_TAG=1.22-alpine

FROM ${BUILD_IMAGE}:${BUILD_TAG} as builder

ARG COMPONENT

WORKDIR /app

COPY $COMPONENT/go.mod $COMPONENT/go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o main $COMPONENT/main/main.go

FROM alpine:latest

ENV GIN_MODE=release

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=builder /app/main .

CMD ["./main"]
FROM public.ecr.aws/docker/library/golang:1.21.4 as base
WORKDIR /app/
ENV GOOS="linux"
ENV CGO_ENABLED=0

FROM base as dev
WORKDIR /app/
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
ENTRYPOINT ["air"]

FROM base as builder
ADD . /app/
WORKDIR /app/
RUN CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o simplebank .

FROM public.ecr.aws/docker/library/debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /app/
COPY --from=builder /app/simplebank .
ENTRYPOINT ["/app/simplebank"]
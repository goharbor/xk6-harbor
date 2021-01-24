ARG GOLANG
FROM ${GOLANG} as builder

WORKDIR /go/src/xk6-harbor
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-w' -i -o k6 ./cmd/k6/main.go


FROM photon:3.0

RUN mkdir -p /var/lib/harbor

ENV HARBOR_ROOT /var/lib/harbor
ENV K6_NO_USAGE_REPORT=true

COPY --from=builder /go/src/xk6-harbor/k6 /usr/local/bin

VOLUME /var/lib/harbor

EXPOSE 6565

ENTRYPOINT ["k6"]

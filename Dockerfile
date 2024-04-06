FROM golang:1.22.1-alpine AS build

ADD db /app/db
ADD handlers /app/handlers
ADD model /app/model
ADD oauth2 /app/oauth2
ADD security /app/security
ADD main.go go.mod go.sum /app/


WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o fyp main.go

FROM scratch

COPY --from=build /app/fyp /app/fyp
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/

CMD ["/app/fyp"]
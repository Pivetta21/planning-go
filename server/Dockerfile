FROM golang:1.21.0-alpine AS builder    

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/planning-go ./cmd/api

FROM scratch AS runner

COPY --from=builder /usr/local/bin/planning-go /usr/local/bin/planning-go

EXPOSE 9000

CMD ["planning-go"]
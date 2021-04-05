# create enviroment
FROM golang:1.15-alpine as enviroment

RUN apk add --update --no-cache ca-certificates git\
    && go env -w GOPROXY=https://goproxy.cn,direct\
    && go env -w GO111MODULE=on\
    && go env -w GOSUMDB=off

WORKDIR /sparkling

COPY go.mod .
COPY go.sum .

RUN go mod download

# builder frame
FROM enviroment AS builder

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /sparkling/sparkling

# prod frame
FROM scratch AS prod

COPY --from=builder /sparkling/sparkling /sparkling/sparkling
COPY --from=builder /sparkling/.env /sparkling/.env

WORKDIR /sparkling

ENTRYPOINT ["/sparkling/sparkling"]

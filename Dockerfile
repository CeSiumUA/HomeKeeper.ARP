FROM golang:alpine
ARG apiname=homekeeperapi

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./

RUN go build -o /hkarp

CMD [ "/hkarp", apiname ]
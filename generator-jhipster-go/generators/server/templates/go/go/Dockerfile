FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

EXPOSE <%= serverPort %>

CMD ["go","run","."]       
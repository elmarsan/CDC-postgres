FROM golang:1.17.2
WORKDIR /app

COPY . ./
COPY ../go.mod ./
COPY ../go.sum ./
RUN go mod download
COPY ../*.go ./

CMD ["go", "test", "./test", "-v"]
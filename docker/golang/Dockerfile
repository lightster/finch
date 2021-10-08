FROM golang:1.17

WORKDIR /app

RUN go get github.com/githubnemo/CompileDaemon
RUN go get golang.org/x/tools/...

COPY ./go.mod /app/go.mod
RUN go mod download

CMD ["CompileDaemon", \
  "-exclude-dir=.git", "-exclude-dir=node_modules", "-exclude-dir=data", \
  "--build=go build -o /tmp/app .", \
  "--include=*.gohtml", \
  "--command=/tmp/app"]

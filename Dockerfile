FROM golang:1.25.1

WORKDIR /app

COPY ./server ./
RUN go mod download
RUN go build -o /application

EXPOSE 8080

ENTRYPOINT [ "/application" ]
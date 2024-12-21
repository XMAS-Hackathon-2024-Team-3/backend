FROM golang:1.22.3
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY datamodel ./datamodel/
COPY ml ./ml/
COPY server ./server/
COPY storage ./storage/
COPY notificator ./notificator/
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend
RUN chmod +x /backend

EXPOSE 8080

CMD ["/backend"]
FROM golang:latest
WORKDIR /api

# pre-copy deps to download & verify checksums
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# copy the rest of the app code
COPY . .

# expose the port to publish
EXPOSE 8080

# run the app
CMD ["go","run","/api/cmd/web/main.go"]
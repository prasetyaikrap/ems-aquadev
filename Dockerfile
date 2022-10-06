# Get Alpine Golang Image
FROM golang:1.18-alpine

# Use app directory as Workdir
WORKDIR /app

# Copy Go module
COPY go.mod ./
COPY go.sum ./

# Download module
RUN go mod download
# Copy Project file into workdir
COPY . ./
# build Go
RUN go build -o /ems-aquadev

EXPOSE 1323

#Execute Go Binary
CMD [ "/ems-aquadev" ]
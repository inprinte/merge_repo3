# # syntax=docker/dockerfile:1

# FROM golang:1.17-alpine

# ENV GO111MODULE=on

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# COPY *.go ./

# # RUN go build -o /inprinte-backoffice

# EXPOSE 8080

# CMD ["echo", "Hello World"]

# CMD [ "/app/inprinteBackoffice" ]


## We specify the base image we need for our
## go application
FROM golang:1.17.0-alpine
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
RUN go mod download
RUN go build -o main .
EXPOSE 80
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]
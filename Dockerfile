#Fetch base layer image
FROM golang:1.18-alpine as dependencies

#Create and set working directory
WORKDIR /app

#Initial setup for downloading dependencies
COPY go.mod go.sum ./
RUN go mod download

#Build application
FROM dependencies as build
COPY . ./
RUN CGO_ENABLED=0 go build -o /main -ldflags="-w -s" ./cmd

#Copy application from base env to docker env
FROM golang:1.18-alpine
COPY --from=build /main /main

#Executable
EXPOSE 50051
CMD [ "/main" ]
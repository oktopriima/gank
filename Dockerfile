FROM golang:1.17.3-buster

# install migration from rubenv
#RUN go get -v github.com/rubenv/sql-migrate/...

# Create a directory for the app
RUN mkdir /app

# Set working directory
WORKDIR /app

 # cache and install dependencies, only when they change
COPY go.mod go.sum ./
RUN go mod download

# Copy all files from the current directory to the app directory
COPY . /app

RUN apt-get update && apt-get install -y \
    imagemagick libmagickwand-dev --no-install-recommends

RUN mv /etc/localtime localtime.backup
RUN ln -s /usr/share/zoneinfo/Asia/Singapore /etc/localtime

RUN go build -o gank .

ENTRYPOINT ["./gank"]
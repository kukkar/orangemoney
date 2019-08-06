# Use ubuntu LTS as Base image
FROM ubuntu:16.04
MAINTAINER glitter.sankalp@gmail.com

# Install required packages
RUN apt-get update && apt-get install -y wget vim git

# Checkout and install go 1.10
RUN wget https://dl.google.com/go/go1.10.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.10.linux-amd64.tar.gz

# Set required environment vars
ENV PATH=$PATH:/usr/local/go/bin:/gospace/bin

# Make gopath space
RUN mkdir -p /gospace
RUN chmod 766 /gospace
ENV GOPATH="/gospace"

# Install reflorest cli
RUN go get -u github.com/BoutiqaatREPO/getsetgo/reflorest


# Using minideb:bullseye for small size and functionality
FROM bitnami/minideb:bullseye

# Update and upgrade
RUN apt-get update --fix-missing && apt-get upgrade -y
RUN apt-get install -y wget git

# Install Go
ENV GOLANG_VERSION 1.22.5
RUN wget -q https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz
RUN rm go${GOLANG_VERSION}.linux-amd64.tar.gz

# Add Go ENV variables
ENV PATH ${PATH}:/usr/local/go/bin
ENV GO111MODULE=on

# Clone project
RUN git clone https://github.com/Auwate/Go_mail_sender.git

# Specify your username
ENV AWS_DIR "C:/Users/auwat/.aws"
# Mount AWS credentials
VOLUME ${AWS_DIR}

# Change directory to where mod file is
WORKDIR /Go_mail_sender

# Download dependencies
RUN go mod tidy
RUN go mod download

EXPOSE 8080

CMD ["go", "run", "./src/main.go"]
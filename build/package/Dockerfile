# GENERATE GO BINARY
FROM  golang:latest AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/muharik19/boiler-plate-grpc
COPY . ./

# caches
COPY go.mod go.sum ./

# get dependencies
RUN go mod tidy
# build golang binary
RUN CGO_ENABLED=0 go build -o /muharik19/boiler-plate-grpc .

FROM debian:bullseye

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
    apt-get -y -q install \
    libreoffice \
    tzdata \
    telnet \
    libreoffice-writer \
    ure \
    libreoffice-java-common \
    libreoffice-core \
    libreoffice-common \
    openjdk-11-jre \
    fonts-opensymbol \
    hyphen-fr \
    hyphen-de \
    hyphen-en-us \
    hyphen-it \
    hyphen-ru \
    fonts-dejavu \
    fonts-dejavu-core \
    fonts-dejavu-extra \
    fonts-droid-fallback \
    fonts-dustin \
    fonts-f500 \
    fonts-fanwood \
    fonts-freefont-ttf \
    fonts-liberation \
    fonts-lmodern \
    fonts-lyx \
    fonts-sil-gentium \
    fonts-texgyre \
    librdkafka-dev \
    fonts-tlwg-purisa && \
    apt-get -y -q remove libreoffice-gnome && \
    apt -y autoremove && \
    rm -rf /var/lib/apt/lists/*

# copy env from the host & copy go binary from the compiler
COPY --from=builder /muharik19/boiler-plate-grpc ./
COPY ./docs ./docs
COPY ./.env ./.env

ENV TZ=Asia/Jakarta

ENTRYPOINT ["/boiler-plate-grpc"]
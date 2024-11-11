ARG GOLANG_VERSION
ARG GOTENBERG_VERSION

FROM golang:$GOLANG_VERSION AS builder

ENV CGO_ENABLED=0

# Define the working directory outside of $GOPATH (we're using go modules).
WORKDIR /home

# Install module dependencies.
COPY go.mod go.sum ./

RUN go mod download &&\
    go mod verify

# Copy the source code.
COPY cmd ./cmd
COPY pkg ./pkg

# Build the binary.
ARG GOTENBERG_VERSION
ARG APP_NAME
ARG APP_VERSION

RUN go build -o gotenberg -ldflags "-X 'github.com/gotenberg/gotenberg/v8/cmd.Version=$GOTENBERG_VERSION-$APP_NAME-$APP_VERSION'" cmd/app/main.go

FROM gotenberg/gotenberg:$GOTENBERG_VERSION

ARG GOTENBERG_VERSION
ARG APP_AUTHOR
ARG APP_REPOSITORY
ARG APP_NAME
ARG APP_VERSION

LABEL org.opencontainers.image.title="Gotenberg $APP_NAME" \
      org.opencontainers.image.description="A Docker-powered stateless API for PDF files." \
      org.opencontainers.image.version="$GOTENBERG_VERSION-$APP_NAME-$APP_VERSION" \
      org.opencontainers.image.authors="$APP_AUTHOR" \
      org.opencontainers.image.source="$APP_REPOSITORY"

USER root

# TODO: add your instructions.

# RUN apt-get update -qq &&\
#     DEBIAN_FRONTEND=noninteractive apt-get install -y -qq --no-install-recommends -t bookworm-backports libreoffice-l10n-de &&\
#     sed -i '/de_DE.UTF-8/s/^# //g' /etc/locale.gen &&\
#     locale-gen &&\
#     # Cleanup.
#     # Note: the Debian image does automatically a clean after each install thanks to a hook.
#     # Therefore, there is no need for apt-get clean.
#     # See https://stackoverflow.com/a/24417119/3248473.
#     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# ENV LANG de_DE.UTF-8
# ENV LANGUAGE de_DE:de
# ENV LC_ALL de_DE.UTF-8

# Copy the Gotenberg binary from the builder stage.
COPY --from=builder /home/gotenberg /usr/bin/

USER gotenberg
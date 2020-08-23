ARG IMAGE="docker.pkg.github.com/nolte/vscode-devcontainers/devops"
ARG VERSION="latest"
FROM ${IMAGE}:${VERSION}

USER root

ARG HUGO_VERSION=0.74.3

#https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_extended_${HUGO_VERSION}_Linux-64bit.tar.gz

RUN curl -LO https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_extended_${HUGO_VERSION}_Linux-64bit.tar.gz \
    && tar xvzf hugo_extended_${HUGO_VERSION}_Linux-64bit.tar.gz -C /usr/local/bin hugo \
    && rm hugo_extended_${HUGO_VERSION}_Linux-64bit.tar.gz \
    && chmod +x /usr/local/bin/hugo

USER ${USERNAME}

RUN npm install -g markdownlint-cli

FROM golang:1.11 as build_env

ARG SSH_KEY
ARG SSH_KEY_PUB

RUN mkdir -p /{{cookiecutter.svc_name}}
RUN apt-get install ca-certificates
COPY . /{{cookiecutter.svc_name}}
COPY gitconfig /root/.gitconfig
RUN cd /{{cookiecutter.svc_name}}
RUN mkdir -p ~/.ssh && \
    echo "$SSH_KEY" > ~/.ssh/id_rsa && \
    echo "$SSH_KEY_PUB" > ~/.ssh/id_rsa.pub && \
    chmod 0700 ~/.ssh/id_rsa && \
    eval `ssh-agent` && \
    ssh-add ~/.ssh/id_rsa && \
    ssh-keyscan gitlab.internal.unity3d.com >> ~/.ssh/known_hosts && \
    echo "Host gitlab.internal.unity3d.com\n\tStrictHostKeyChecking no\n" >> ~/.ssh/config
RUN cd /{{cookiecutter.svc_name}} && CGO_ENABLED=0 GOOS=linux go build -o /{{cookiecutter.svc_name}}-server ./cmd/server/main.go
# Remove ssh key
RUN rm ~/.ssh/id_rsa ~/.ssh/id_rsa.pub && unset SSH_KEY && unset SSH_KEY_PUB

FROM scratch
LABEL maintainer="{{cookiecutter.author}} <{{cookiecutter.email}}>"

ADD ./gcp /etc/gcp
COPY --from=build_env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY ./src/config/*.json /etc/
COPY --from=build_env /{{cookiecutter.svc_name}}-server .

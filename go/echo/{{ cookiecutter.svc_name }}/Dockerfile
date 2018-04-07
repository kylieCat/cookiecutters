FROM golang:1.9 as build_env
MAINTAINER {{ cookiecutter.author }} <{{ cookiecutter.email }}>

RUN mkdir -p /go/src/ggithub.com/{{ cookiecutter.github_username }}/{{ cookiecutter.svc_name }}
RUN apt-get install ca-certificates
COPY . /go/src/github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.svc_name }}
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /{{ cookiecutter.svc_name }} /go/src/github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.svc_name }}/main.go


FROM scratch
MAINTAINER {{ cookiecutter.author }} <{{ cookiecutter.email }}>

COPY --from=build_env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY ./config/*.json /etc/
COPY --from=build_env /{{ cookiecutter.svc_name }} .
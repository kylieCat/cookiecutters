FROM golang:1.9 as build_env
MAINTAINER {{ cookiecutter.author }} <{{ cookiecutter.email }}>

RUN mkdir -p /go/src/gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}
RUN apt-get install ca-certificates
COPY ../../../ /go/src/gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /{{ cookiecutter.svc_name }} /go/src/gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/main.go


FROM scratch
MAINTAINER {{ cookiecutter.author }} <{{ cookiecutter.email }}>

COPY --from=build_env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build_env /{{ cookiecutter.svc_name }} .
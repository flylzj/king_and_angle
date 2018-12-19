FROM golang:1.11

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo 'Asia/Shanghai' >/etc/timezone
WORKDIR /app


CMD ["main"]
FROM alpine

COPY jesi /usr/bin/jesi

RUN ["/usr/bin/jesi"]
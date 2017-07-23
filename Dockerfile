FROM alpine

COPY jesi /usr/bin/jesi

CMD ["/usr/bin/jesi"]
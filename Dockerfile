FROM go:latest

RUN make build

ENTRYPOINT [ "sh", "-c", "" ]
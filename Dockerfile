FROM alpine:latest

COPY bin/entrypoint.sh entrypoint.sh
RUN chmod a+x entrypoint.sh

ENTRYPOINT [ "entrypoint.sh" ]


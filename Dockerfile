FROM hashicorp/terraform:0.12.12

RUN apk add tree

ADD ./bin/entrypoint.sh /entrypoint.sh
RUN chmod a+x /entrypoint.sh

ENTRYPOINT /entrypoint.sh


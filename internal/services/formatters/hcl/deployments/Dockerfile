# fork aledj02/horusx
# base: python:3.12-alpine — tools: tfsec 1.28.14 + checkov (última)
FROM python:3.12-alpine

RUN apk add --no-cache wget

RUN wget -O /usr/bin/tfsec https://github.com/aquasecurity/tfsec/releases/download/v1.28.14/tfsec-linux-amd64 && chmod +x /usr/bin/tfsec

RUN apk add --no-cache --virtual .build_deps build-base libffi-dev \
 && pip install --no-cache-dir -U checkov \
 && apk del .build_deps

ENTRYPOINT []
CMD ["/bin/sh"]

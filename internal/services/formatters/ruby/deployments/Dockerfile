# fork aledj02/horusx
# base: ruby:3.3-alpine — tools: brakeman 8.1.0 + bundler-audit 0.9.3
FROM ruby:3.3-alpine

RUN apk add --no-cache git

RUN gem install brakeman -v 8.1.0
RUN gem install bundler-audit -v 0.9.3

ENTRYPOINT []
CMD ["/bin/sh"]

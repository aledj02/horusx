# fork aledj02/horusx
# base: php:8.3-alpine — tools: php_codesniffer + phpcs-security-audit
FROM php:8.3-alpine

RUN apk add --no-cache curl

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

RUN composer global config bin-dir /usr/local/bin

RUN composer global require "squizlabs/php_codesniffer=*"

RUN composer require --dev pheromone/phpcs-security-audit

RUN phpcs --config-set installed_paths /vendor/pheromone/phpcs-security-audit/Security

ENTRYPOINT []
CMD ["/bin/sh"]

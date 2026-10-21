# fork aledj02/horusx
# base: python:3.12-alpine + jlink JDK 21 — tools: semgrep 1.182.0, dependency-check 12.1.0, trivy 0.76.1
FROM azul/zulu-openjdk-alpine:21 AS jlink

RUN "$JAVA_HOME/bin/jlink" --compress=2 \
    --module-path /opt/java/openjdk/jmods \
    --add-modules java.base,java.compiler,java.datatransfer,jdk.crypto.ec,java.desktop,java.instrument,java.logging,java.management,java.naming,java.rmi,java.scripting,java.security.sasl,java.sql,java.transaction.xa,java.xml,jdk.unsupported \
    --output /jlinked

FROM python:3.12-alpine

RUN apk add --no-cache build-base curl

RUN pip install --no-cache-dir semgrep==1.182.0

COPY --from=jlink /jlinked /opt/jdk/

ENV JAVA_HOME=/opt/jdk

RUN curl -fsSL -o /tmp/dc.zip https://github.com/jeremylong/DependencyCheck/releases/download/v12.1.0/dependency-check-12.1.0-release.zip \
 && unzip /tmp/dc.zip -d /opt \
 && ln -s /opt/dependency-check/dependency-check.sh /usr/local/bin/dependency-check \
 && rm /tmp/dc.zip

RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.76.1

ENTRYPOINT []
CMD ["/bin/sh"]

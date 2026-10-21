# fork aledj02/horusx
# base: python:3.12-slim (glibc) + jlink JDK 21 — tools: opengrep 1.31.0 (fork OSS de semgrep, sin login), dependency-check 12.1.0, trivy 0.76.1
# multi-arch: linux/amd64 + linux/arm64
FROM azul/zulu-openjdk-alpine:21 AS jlink

RUN "$JAVA_HOME/bin/jlink" --compress=2 \
    --module-path /opt/java/openjdk/jmods \
    --add-modules java.base,java.compiler,java.datatransfer,jdk.crypto.ec,java.desktop,java.instrument,java.logging,java.management,java.naming,java.rmi,java.scripting,java.security.sasl,java.sql,java.transaction.xa,java.xml,jdk.unsupported \
    --output /jlinked

FROM python:3.12-slim

RUN apt-get update && apt-get install -y --no-install-recommends curl ca-certificates unzip \
    && rm -rf /var/lib/apt/lists/*

# OpenGrep v1.31.0 — binarios estáticos manylinux (glibc), multi-arch
ARG TARGETARCH
RUN set -e; \
    case "${TARGETARCH}" in \
        amd64)  OG_ARCH=x86 ;; \
        arm64)  OG_ARCH=aarch64 ;; \
        *) echo "Unsupported architecture: ${TARGETARCH}" >&2; exit 1 ;; \
    esac; \
    curl -fsSL -o /usr/local/bin/opengrep \
        "https://github.com/opengrep/opengrep/releases/download/v1.31.0/opengrep_manylinux_${OG_ARCH}"; \
    chmod +x /usr/local/bin/opengrep; \
    /usr/local/bin/opengrep --version

COPY --from=jlink /jlinked /opt/jdk/

ENV JAVA_HOME=/opt/jdk

RUN curl -fsSL -o /tmp/dc.zip https://github.com/jeremylong/DependencyCheck/releases/download/v12.1.0/dependency-check-12.1.0-release.zip \
    && unzip /tmp/dc.zip -d /opt \
    && ln -s /opt/dependency-check/dependency-check.sh /usr/local/bin/dependency-check \
    && rm /tmp/dc.zip

RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.76.1

ENTRYPOINT []
CMD ["/bin/sh"]

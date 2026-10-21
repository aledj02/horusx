# fork aledj02/horusx
# base: dotnet sdk 8.0 (LTS) — tools: security-scan (última) + jq
FROM mcr.microsoft.com/dotnet/sdk:8.0-alpine

RUN apk add --no-cache jq

RUN dotnet tool install --global security-scan

ENV PATH="$PATH:/root/.dotnet/tools"
ENTRYPOINT []
CMD ["/bin/sh"]

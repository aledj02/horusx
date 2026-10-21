# Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# Fork horusx — FASE 5 (2026-10-21)
#   - golang:1.24.5-alpine (upstream: 1.17.8)
#   - nancy v2.1.0 (upstream: v1.0.33) — binary directo (no tar.gz)
#   - gosec v2.29.0 (upstream: v2.11.0) — tar.gz, el install script pone en BINDIR
#   - multi-arch via TARGETARCH (buildx lo setea)
#
# Uso del parser (debe seguir corriendo):
#   - gosec:   `gosec -quiet -fmt=json -out=$FILE_PATH ./...`
#   - nancy:   `go list -json -m all | nancy sleuth -o json`

FROM golang:1.24.5-alpine

RUN apk add --no-cache curl tar

ARG TARGETARCH

# nancy v2.1.0 (upstream: v1.0.33)
# Asset: nancy-v2.1.0-linux-{amd64,arm64} (binary directo, no gzip)
RUN curl -fsSL \
      "https://github.com/sonatype-nexus-community/nancy/releases/download/v2.1.0/nancy-v2.1.0-linux-${TARGETARCH}" \
      -o /usr/local/bin/nancy \
    && chmod +x /usr/local/bin/nancy \
    && nancy sleuth --help >/dev/null

# gosec v2.29.0 (upstream: v2.11.0)
# Asset: gosec_2.29.0_linux_{amd64,arm64}.tar.gz (el install.sh de securego extrae)
RUN echo "gosec install for TARGETARCH=${TARGETARCH}" \
    && curl -fsSL https://raw.githubusercontent.com/securego/gosec/master/install.sh \
       | BINDIR=/usr/local/bin sh -s v2.29.0 \
    && gosec -version

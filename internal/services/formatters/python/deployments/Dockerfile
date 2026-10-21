# Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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
#   - python 3.12.7-alpine3.20
#   - bandit 1.9.4 (upstream: 1.7.4)
#   - safety 1.10.3 (MANTENIDO — parser config.go usa el JSON de safety 1.x;
#     safety 2+/3 es breaking: requiere reescribir config.go + formatter.go)
#   - jq

FROM python:3.12.7-alpine3.20

RUN apk add --no-cache jq curl

# Opción A firmada (D2026-10-21): bandit 1.9.4 + safety 1.10.3 (compat parser)
RUN pip install --no-cache-dir \
      bandit==1.9.4 \
      safety==1.10.3 \
      && python -m bandit --version \
      && safety --version

# shell helper que sube el upstream (no modifica el parser)
COPY ./internal/services/formatters/python/deployments/horusx-file-ignore.sh /usr/local/bin/horusx-file-ignore.sh
RUN chmod +x /usr/local/bin/horusx-file-ignore.sh

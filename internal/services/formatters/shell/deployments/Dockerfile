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
#   - alpine 3.21 (upstream: 3.16.0)
#   - shellcheck 0.10.0 (upstream: >0.7.2) — pin del repo community de alpine 3.21
#   - multi-arch (buildx)
#
# Uso del parser (debe seguir corriendo):
#   - shellcheck --format=json <file>

FROM alpine:3.21

RUN apk add --no-cache shellcheck=0.10.0-r1 \
    && shellcheck --version

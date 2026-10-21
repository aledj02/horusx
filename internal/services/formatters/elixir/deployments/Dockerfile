# fork aledj02/horusx
# base: elixir:1.17-alpine — tools: mix_audit 2.1.5 + sobelow (última)
FROM elixir:1.17-alpine

RUN apk add --no-cache git

RUN mix local.hex --force
RUN mix local.rebar --force

RUN mix escript.install hex mix_audit 2.1.5 --force
RUN mix archive.install hex sobelow --force

RUN mv ./root/.mix/escripts/mix_audit ./bin/

ENTRYPOINT []
CMD ["/bin/sh"]

FROM caddy:2.6.4-builder-alpine as builder
RUN xcaddy build \
  --with github.com/caddyserver/cache-handler

FROM caddy:2.6.4-alpine
COPY --from=builder /usr/bin/caddy /usr/bin/caddy

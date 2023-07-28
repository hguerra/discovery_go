#################################################
# STAGE-1
# Base Kong image
FROM kong:3.3.0-ubuntu AS kong-base

USER root

RUN mkdir -p /opt/kong
COPY kong.yaml /opt/kong/kong.yaml
RUN chown -R kong:kong /opt/kong

ENV KONG_NGINX_PROXY_GZIP="on"
ENV KONG_NGINX_PROXY_GZIP_DISABLE="msie6"
ENV KONG_NGINX_PROXY_GZIP_VARY="on"
ENV KONG_NGINX_PROXY_GZIP_PROXIED="any"
ENV KONG_NGINX_PROXY_GZIP_COMP_LEVEL="6"
ENV KONG_NGINX_PROXY_GZIP_BUFFERS="16 8k"
ENV KONG_NGINX_PROXY_GZIP_HTTP_VERSION="1.1"
ENV KONG_NGINX_PROXY_GZIP_MIN_LENGTH="256"
ENV KONG_NGINX_PROXY_GZIP_TYPES="application/atom+xml application/geo+json application/javascript application/x-javascript application/json application/ld+json application/manifest+json application/rdf+xml application/rss+xml application/xhtml+xml application/xml font/eot font/otf font/ttf image/svg+xml text/css text/javascript text/plain text/xml"

ENV KONG_DATABASE="off"
ENV KONG_PROXY_LISTEN="0.0.0.0:8080"
ENV KONG_ADMIN_LISTEN="0.0.0.0:8001"
ENV KONG_STATUS_LISTEN="0.0.0.0:8100"
ENV KONG_LOG_LEVEL="error"
ENV KONG_ADMIN_ACCESS_LOG="off"
ENV KONG_PROXY_ACCESS_LOG="off"
ENV KONG_PROXY_ERROR_LOG="/dev/stderr"
ENV KONG_ADMIN_ERROR_LOG="/dev/stderr"
ENV KONG_DECLARATIVE_CONFIG="/opt/kong/kong.yaml"


#################################################
# STAGE-2
# Build the kong plugin
FROM golang:1.18.10-bullseye AS plugin-builder

WORKDIR /builder

COPY . .

RUN cd plugins/header-validation \
    && GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /tmp/header-validation main.go


#################################################
# STAGE-3
# Build kong including the plugin that already build in previous stage
FROM kong-base

USER root

COPY --from=plugin-builder /tmp/header-validation /usr/local/bin/header-validation
RUN chmod +x /usr/local/bin/header-validation

ENV KONG_PLUGINS="bundled,header-validation"
ENV KONG_PLUGINSERVER_NAMES="header-validation"
ENV KONG_PLUGINSERVER_HEADER_VALIDATION_START_CMD="/usr/local/bin/header-validation"
ENV KONG_PLUGINSERVER_HEADER_VALIDATION_QUERY_CMD="/usr/local/bin/header-validation -dump"

USER kong
ENTRYPOINT ["/docker-entrypoint.sh"]
EXPOSE 8080 8443 8001 8444 8100
STOPSIGNAL SIGQUIT
HEALTHCHECK --interval=10s --timeout=10s --retries=10 CMD kong health
CMD ["kong", "docker-start"]

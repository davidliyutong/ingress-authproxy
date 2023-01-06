
# shellcheck disable=SC2188
echo set \$AUTHPROXY_BACKEND_URL "$AUTHPROXY_BACKEND_URL"; > /etc/nginx/conf.d/env.variable

nginx

sh
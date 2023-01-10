#!/bin/bash

# shellcheck disable=SC2016
echo "set \$AUTHPROXY_BACKEND_URL $AUTHPROXY_BACKEND_URL ;" > /etc/nginx/conf.d/env.variable
echo "set \$AUTHPROXY_OIDC_SERVICE_URL $AUTHPROXY_OIDC_SERVICE_URL ;" >> /etc/nginx/conf.d/env.variable

nginx -g 'daemon off;'
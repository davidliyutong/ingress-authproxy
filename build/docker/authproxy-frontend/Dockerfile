FROM node:16.18.1 AS builder

WORKDIR /opt/build

COPY frontend /opt/build
RUN yarn install && yarn build

FROM nginx:1.23.3

COPY --from=builder /opt/build/dist /var/www/html/public
#COPY frontend/dist /var/www/html/public

COPY manifests/authproxy-frontend/nginx/nginx.conf /etc/nginx/nginx.conf

COPY manifests/authproxy-frontend/entrypoint/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

EXPOSE 80

SHELL ["/bin/bash", "-c"]
ENTRYPOINT ["./entrypoint.sh"]


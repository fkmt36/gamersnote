#!/bin/sh

envsubst '$$API_HOST $$API_PORT $$CLIENT_HOST $$CLIENT_PORT' < \
  /etc/nginx/conf.d/default.conf.template > \
  /etc/nginx/conf.d/default.conf

nginx -g 'daemon off;'

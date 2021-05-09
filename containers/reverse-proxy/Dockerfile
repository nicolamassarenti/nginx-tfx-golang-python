FROM nginx:1.19.0-alpine

ARG AUTHPROXY="src/authproxy"
ENV SRC="${AUTHPROXY}"

# Copy the main config
COPY ${SRC}/nginx.conf /etc/nginx/nginx.conf

# Copy because this image substitutes all env variables in *.templates and outputs a * file
COPY ${SRC}/routes.conf.template /etc/nginx/templates/routes.conf.template

# Copy the static content
COPY ${SRC}/html/ /usr/share/nginx/html/

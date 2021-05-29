FROM nginx:1.20.0-alpine


# Copy the main config
COPY ./mime.types /etc/nginx/mime.types
COPY ./proxy.conf /etc/nginx/proxy.conf
COPY ./nginx.conf /etc/nginx/nginx.conf

FROM nginx:1.15.0-alpine
EXPOSE 80

COPY ./conf/default.conf /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]

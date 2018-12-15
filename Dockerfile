#=== NodeJS
#==============================================================================================
FROM node:8.4.0-alpine as nodejs

ENV TIMEZONE Asia/Tokyo

#  Setup
#-----------------------------------------------
RUN set -x \
  && apk add --no-cache -U tzdata \
  && ln -sf /usr/share/zoneinfo/$TIMEZONE /etc/localtime

RUN apk add --no-cache -U yarn

#  Build
#-----------------------------------------------
WORKDIR /app

COPY src/ src/
COPY public/ public/
COPY webpack/ webpack/
COPY package.json .
COPY yarn.lock .

RUN yarn
RUN yarn build
RUN mv public/assets/index.html public/index.html

#=== Nginx
#==============================================================================================
FROM nginx:1.15.7-alpine

# envsubst
RUN set -x \
  && apk --no-cache add libintl \
  && apk --no-cache add --virtual .gettext gettext \
  && cp /usr/bin/envsubst /usr/local/bin/envsubst \
  && apk del .gettext

COPY --from=nodejs /app/public/ /usr/share/nginx/html/
COPY nginx/default.conf.tmpl /etc/nginx/conf.d/
COPY script/server /usr/local/bin/

CMD ["server"]

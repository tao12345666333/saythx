FROM node:10.13 as builder

WORKDIR /app
COPY . /app

RUN yarn install \
        && yarn build


FROM nginx:1.15

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist /usr/share/nginx/html/

EXPOSE 80

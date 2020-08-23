FROM klakegg/hugo:0.74.3-ext-nodejs AS hugo

RUN npm install -D --save autoprefixer@9.8.6 && npm install -D --save postcss-cli@7.1.1

FROM nginx:alpine
COPY --from=hugo /target /usr/share/nginx/html

FROM klakegg/hugo:0.74.3-ext-nodejs AS hugo

RUN npm install -D --save autoprefixer && npm install -D --save postcss-cli

FROM nginx:alpine
COPY --from=hugo /target /usr/share/nginx/html

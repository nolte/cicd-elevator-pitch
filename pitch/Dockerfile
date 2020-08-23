FROM klakegg/hugo:0.74.3-ext-alpine-onbuild AS hugo

FROM nginx:alpine
COPY --from=hugo /target /usr/share/nginx/html

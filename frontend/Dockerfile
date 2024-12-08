FROM node:slim as builder
COPY package.json .
RUN npm i
COPY . .
RUN npm run build

FROM nginx:stable
COPY --from=builder /dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]

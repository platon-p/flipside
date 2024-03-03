FROM node as builder
WORKDIR /app
COPY package.json .
RUN npm i --only=dev
COPY . .
RUN npm run build

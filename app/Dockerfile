FROM node:14.12.0-alpine AS dev

WORKDIR /usr/src/app

# ================================= #

FROM node:14.12.0-alpine

WORKDIR /usr/src/app

COPY ./package*.json ./

RUN npm install

COPY . .

ARG S3_URL=""

ARG APP_URL=""

ENV HOST=0.0.0.0 PORT=3000

RUN npm run build

CMD [ "npm", "run", "start" ]

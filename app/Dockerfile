FROM node:14.12.0-alpine AS dev

WORKDIR /usr/src/app

# ================================= #

FROM node:14.12.0-alpine

WORKDIR /usr/src/app

COPY ./package*.json ./

RUN npm install

COPY . .

RUN npm run build

ENV HOST=0.0.0.0 PORT=3000

CMD [ "npm", "run", "start" ]

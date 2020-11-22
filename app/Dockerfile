FROM node:14.12.0-alpine AS dev

WORKDIR /usr/src/app

# ================================= #

FROM node:14.12.0-alpine

WORKDIR /usr/src/app

COPY ./package*.json ./

RUN npm install

COPY . .

RUN npm run build

# ENV NUXT_HOST=0.0.0.0 NUXT_PORT=3000 API_URL=http://nginx API_URL_BROWSER=http://localhost:3000

CMD [ "npm", "run", "start" ]

FROM node:20-alpine as build

COPY . /app

WORKDIR /app

RUN npm ci --omit=dev

FROM node:20-alpine

COPY --from=build /app /app

WORKDIR /app

CMD ["node", "app.js"]
FROM node:22.8-bookworm

WORKDIR /app

ENV PORT 80

COPY . ./
RUN rm -rf node_modules

RUN npm install
RUN npm run build

EXPOSE $PORT

CMD ["npm", "start"]

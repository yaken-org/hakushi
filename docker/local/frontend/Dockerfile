FROM node:22.8-alpine

WORKDIR /app

ENV PORT 3000

COPY package.json package-lock.json ./

RUN npm ci

COPY . .

EXPOSE $PORT

CMD ["npm", "run", "dev"]

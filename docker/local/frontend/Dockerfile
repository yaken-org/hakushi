FROM node:22.8-bookworm

WORKDIR /app

ENV PORT 3000

COPY package.json package-lock.json ./

RUN npm install

COPY . .

EXPOSE $PORT

CMD ["npm", "run", "dev"]

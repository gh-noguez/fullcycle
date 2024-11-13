FROM node:latest

WORKDIR /usr/src/app

RUN npm init -y && npm install express && npm install mysql --save
    
COPY . .

CMD [ "node", "index.js" ]
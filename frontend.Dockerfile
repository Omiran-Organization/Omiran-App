FROM node:12
WORKDIR /home/node/app
COPY frontend /home/node/app
RUN npm install
CMD npm run app
EXPOSE 3000
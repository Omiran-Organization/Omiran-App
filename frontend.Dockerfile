FROM node:12
WORKDIR /home/node/app
COPY frontend /home/node/app
RUN npm install
CMD npm run dev
EXPOSE 3000
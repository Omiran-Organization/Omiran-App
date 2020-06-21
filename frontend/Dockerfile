

FROM node:13.12.0 as build

WORKDIR /usr/src/app/
ENV TZ America/New_York
ENV PATH /usr/src/app/node_modules/.bin:$PATH
COPY package.json . /usr/src/app/
RUN npm install
COPY . ./
EXPOSE 3000
# RUN npm run build
# CMD npm run start
CMD npm run dev
=======
FROM node
COPY . /app/
WORKDIR /app/
RUN npm install
CMD npm run dev
ENTRYPOINT [ "/app" ]
EXPOSE 3000

